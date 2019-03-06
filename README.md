# RIL

A Go web app for storing articles to read later

## Local setup 

Build the app

```
go build -o ril cmd/web/*
```

Set up a local mysql database in docker. This will map port `3306` to the container and populate the db with the
[`.sql/init.sql`](https://github.com/edwlarkey/ril/tree/master/.sql/init.sql) script

```
docker-compose up -d
```

Run the app

```
./ril
```

Visit http://localhost:4000/

and sign up a user at http://localhost:4000/user/signup

## Usage

```
./ril -help
```
