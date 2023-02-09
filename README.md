# TODO APP

# Setup localy
make sure yoou have golang 1.19 or later and docker installed in your machine

- clone repository
  ``` bash
  $ git clone git@github.com:khilmi-aminudin/todo-app.git
  $ cd todo-app
  ```
- create mysql database with docker
  ``` bash
  # create container mysql
  $ make mysql
  # exec and open mysql CLI
  $ make execdb

  # after entering mysql CLI run this command
  $ create database todo_activity;
  ```
- run migration
  ``` bash
  # on todo-app directory
  $ make migrate
  ```

- run the app
  ``` bash
  # on todo-app directory
  $ make rundev
  ```
- build app
  ``` bash
  $ make build 
  ``` 