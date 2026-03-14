## INIT PROJECT


## Starting a local docker postgress for this application

`docker run --name postgres-container -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres`

connect though a sql client with credentials below

username: postgres
database: postgres
password: secret

create the database we will be working with

```
    CREATE DATABASE golang_sql;
```

## Init golang
https://go.dev/doc/tutorial/getting-started#install

`
    go mod init golang-sql-tutorial
`

## Install sqlx and logrus

https://github.com/jmoiron/sqlx

```
    go get github.com/jmoiron/sqlx
    go get github.com/lib/pq
    go get github.com/sirupsen/logrus
```