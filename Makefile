rundev:
	go run main.go

mysql:
	docker run --name todo-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql

execdb:
	docker exec todo-mysql mysql -uroot -psecret 

migrate:
	migrate -path app/db/migration/schema -database "mysql://root:secret@tcp(localhost:3306)/todo_activity" -verbose up

rundb:
	docker start todo-mysql

build:
	docker build -t todo-app .
