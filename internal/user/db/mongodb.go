package db

import (
	"context"
	"errors"
	"fmt"
	"rest-api-server/internal/user"
	"rest-api-server/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

// Create implements user.Storage.
func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("создание пользователя")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("Ошибка создания пользователя: %v", err)
	}
	d.logger.Debug("конвертация в оид")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("ошибка конвертации объекта(юзера) в хекс: %s", oid)

}

// FindOne implements user.Storage.
func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("ошибка перевода хекса в оid: %s", id)
	}
	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO ErrEntityNotFound
			return u, fmt.Errorf("not found")
		}

		return u, fmt.Errorf("ошибка поиска юзера id:%s. ошибка:%v", id, err)
	}

	if err := result.Decode(&u); err != nil {
		return u, fmt.Errorf("ошибка декода юзера id:%s. ошибка:%v", id, err)
	}
	return u, nil
}

// Update implements user.Storage.
func (d *db) Update(ctx context.Context, user user.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("ошибка конвертации user ID в objectID. ID=%s", user.ID)
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("ошибка маршала user.error: %v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("ошибка анмаршала user.error: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{"$set": updateUserObj}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса на обновление пользователя: ", err)
	}

	if result.MatchedCount == 0 {
		//TODO ErrEntityNotFound
		fmt.Errorf("пользователь не найден")
	}
	d.logger.Tracef("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil
}

// Delete implements user.Storage.
func (d *db) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("ошибка конвертации user ID в objectID. ID=%s", id)
	}

	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("ошибка удаления пользователя", err)
	}
	if result.DeletedCount == 0 {
		//TODO ErrEntityNotFound
		fmt.Errorf("пользователь не найден")
	}
	d.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil
}

func NewStorage(database *mongo.Database, colletion string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(colletion),
		logger:     logger,
	}
}
