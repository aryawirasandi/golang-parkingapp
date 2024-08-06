# Golang Parking App UseCase

A simple use case parking app written with golang 
this repo still work in progress, so i'll put in the [board](https://github.com/users/aryawirasandi/projects/2) to track this current repo progress 

# Usecase

![usecase parking app](usecase.png)

# Preparation
- install [golang](https://go.dev/)
- install [docker](https://www.docker.com/)

# How to run it ?
- first you need a mysql docker container, all you have to do is compose up the mysql then it will run at port 4444, open your terminal and run
```bash
docker compose up
```
in the root of the project
- copy the .env-example to .env, the env is exactly same as docker mysql environment
- run migration in your machine, if you use docker compose, you just need to run this command 
```bash
migrate -path db/migrations -database "mysql://root:root@tcp(127.0.0.1:4444)/parkingdb?query" up
```
then run 
```bash
 go run main.go
```
thats it have a nice code 😉

# Honorable Mention
 - [Golang Migrate](https://github.com/golang-migrate/migrate)
 - [Echo](https://echo.labstack.com/)
