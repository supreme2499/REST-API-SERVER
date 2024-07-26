package user

import "context"

type Storage interface {
	//метод создание пользователя
	Create(ctx context.Context, user User) (string, error)
	//поиск пользователя по id
	FindOne(ctx context.Context, id string) (User, error)
	//Обновление пользователя по id
	Update(ctx context.Context, user User) error
	//удаление пользователя по id
	Delete(ctx context.Context, id string) error
}
