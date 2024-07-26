package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// в  этом файле мы описываем подключение в базе данных (клинт) и вазвращаем базу данных
func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (*mongo.Database, error) {
	var mongoDBURL string
	var isAuth bool
	if username == "" && password == "" {
		//это url для не авторизованого пользователя и он выглядит как
		//"mongodb://host:port"
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		//это url для авторизованого пользователя и он выглядит как
		//"mongodb://username:password@host:port"
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
		isAuth = true
	}
	clientOptions := options.Client().ApplyURI(mongoDBURL)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOptions.SetAuth(options.Credential{
			AuthSource: "",
			Username:   username,
			Password:   password,
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Ошибка конекта. Ошибка: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Ошибка пинга. Ошибка: %v", err)
	}
	return client.Database(database), nil
}
