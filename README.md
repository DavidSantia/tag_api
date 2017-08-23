# tag_api
API using Go-lang struct tags to load SQL data, and implement JSON endpoints

* This project includes a database container with sample Image and User data

## Setup

Build the database container as follows
```sh
docker build -t tagdemo ./data
```

Start the MySQL container as follows:
```sh
docker run --name tag_api_db --rm -p 6603:3306 tagdemo
```
The database will be ready after you see the message:
```
[Entrypoint] MySQL init process done. Ready for start up.

[Entrypoint] Starting MySQL 5.7.19-1.1.0
```
