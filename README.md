Для запуска используйте команду
docker-compose up --build

1)Сначала соберется образ приложения 

2)Запустит контейнеры для приложения и базы данных

3)Приложение открывается по адресу: http://localhost:3000

База данных PostgreSQL будет доступна на:
host: db

port: 5432

user: postgres

password: 12345

database: todo_db

Остановка проекта : docker-compose down

API-эндпоинты:
1) GET
   ![alt text](./screenshots/{7B1FC019-6B61-4FB6-9E46-0B239BFDAE8C}.png)
![alt text](./screenshots/{A91BFD3D-61E0-41AF-957C-5BA3D92FEDBE}.png)

2) POST
![alt text](./screenshots/{AECC9B65-22E5-49F8-AB01-C5F1E08E9021}.png)
![alt text](./screenshots/{BA1C12CE-5017-41B2-90CC-EB1316B5954A}.png)


3) DELETE (удалил 1 и 2 таску)
![alt text](./screenshots/{9CE95FCF-DC04-4B5E-9B9A-64C6EE94C83D}.png)
![alt text](./screenshots/{15053751-272F-47F5-94EA-B4FC60BAD31A}.png)

4) PUT
![alt text](./screenshots/{7AD62184-6D62-4F68-B58A-2F138912638D}.png)
![alt text](./screenshots/{4D7CBBA0-CB5E-4F43-B0CD-8204920B9583}.png)

![alt text](./screenshots/{293E63DA-CEA9-48F6-A0B9-88CDB6CC78FD}.png)

В файле docker-compose.yml можно указать
services:
app:
build: .
ports:
- "3000:3000"
env_file:
- .env
depends_on:
- db
Так как файл уже явно объявлен в проекте