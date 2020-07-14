![Docker Compose Actions Workflow](https://github.com/AlexanderOkhrimenko/go_redis/workflows/Docker%20Compose%20Actions%20Workflow/badge.svg?branch=master)

## Go + Redis + Docker 

The project that joins NoSQL Redis and Go, for connecting to database.  The Project is packaged in Docker and consists of two services: redis and the api .The Go code connects to the database, creates and reads a fixed structure.


Build: `bash ./build.sh`

Purge: `bash ./purge.sh`

Test: `curl "http://localhost:8080/v1/url.insert"`

***

## Go + Redis + Docker 

Проект соединяющий NoSQL Redis и код  на языке Go для подключения к ней, Проект упакован в Docker, состоит из двух сервисов : redis ,  api .Gо код соединяется с базой, создает и читает некую фиксированную структуру.


Build: `bash ./build.sh`

Purge: `bash ./purge.sh`

Test: `curl "http://localhost:8080/v1/url.insert"`
