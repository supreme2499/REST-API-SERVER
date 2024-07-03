# REST-API-SERVER
это личный репозиторий в котором я учуть работать с апи

rest-api-serv

logic:

GET /user -- list of users -- 200, 404, 500
   этот запрос даёт нам информацию о всех юзерах

GET /user/:id -- user by id -- 200, 404,   500 этот запрос возвращает нам информацию о пользователе по его id 

POST /user/:id --create user -- 204, 4xx, hendler location: url
этот запрос создаёт нового пользователя и возвращает ссылку на него

PUT /user/:id -- fully update user -- 204, 4xx, 500
Этот запрос полностью обновляет данные о пользователе

PATH /user/:id -- partailly update user -- 204, 4xx, 500
Этот запрос частично обновляет данные о пользователе

DELETE /user/:id -- delete user by id -- 204, 4xx, 500
Этот запрос удаляет пользователя по id

