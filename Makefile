rundev:
	go run main.go
mysql:
	docker run --name todo-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql
rundb:
	docker start todo-mysql
