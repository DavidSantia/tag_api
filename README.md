# tag_api
API using Go-lang struct tags to load SQL data and implement JSON endpoints.

## Goal
Demonstrate how simple it is to prototype and modify an API in Go.

By simply adjusting or adding a field to a Go struct, you automatically update both how the server loads from the database, as well as what it outputs for the API endpoint.

* This project builds a MySQL database Docker image, initialized with sample data
* The sample data uses photos from https://clients3.google.com/cast/chromecast/home
* It also builds an api-server application in Go, which it installs on a Docker image

It uses govvv to provide the Github version string in the application.  Install as follows:
```sh
go get "github.com/ahmetb/govvv"
```

Then clone the `tag_api` project:
```sh
go get "github.com/DavidSantia/tag_api"
cd $GOPATH/src/github.com/DavidSantia/tag_api
```

## Building and Running the System
First install govvv, clone this project, and start in the `tag_api` directory as shown above.

Next, run the `build.sh` script to compiles the apps and images.
```sh
cd docker
./build.sh
```

Then type `docker images`, and you should see the following images:
```
REPOSITORY                      TAG                 IMAGE ID            CREATED             SIZE
tagdemo/api-server              latest              0fc6d6e09ab2        About an hour ago   8.98MB
tagdemo/mysql                   latest              9d795daac22c        About an hour ago   255MB
```

There is also a `clean.sh` script to remove containers and images from your previous builds.
```sh
./clean.sh
```

Finally, start the database, NATS server, and api-server as follows:
```sh
docker-compose up
```
The NATS server handles communication between the services.  In this Basic example, the API server server subscribes to NATS, but
it doesn't do anything further.

### Accessing the API
Once you have the API server up and running, use your browser to authenticate and access data. The following endpoints are defined:
```go
router.Handle("GET", "/authenticate", handleAuthTestpage)
router.Handle("POST", "/authenticate", makeHandleAuthenticate(cs))
router.Handle("GET", "/keepalive", makeHandleAuthKeepAlive(cs))
router.Handle("GET", "/image", makeHandleAllImages(cs))
router.Handle("GET", "/image/:Id", makeHandleImage(cs))
router.Handle("GET", "/user", makeHandleUser(cs))
```

By browsing to [localhost:8080/authenticate](http://localhost:8080/authenticate), you will see a test page with two buttons.
![Figure 1: Architecture](https://raw.githubusercontent.com/DavidSantia/tag_api/master/README-2buttons.png)

Each button authenticates you as a particular user from the sample database, either in the Basic or Premium group. Once authenticated, your browser will have a Session cookie to allow you to continue using the API.

You can then browse to

* [localhost:8080/image](http://localhost:8080/image) to see images you have access to
* [localhost:8080/image/4](http://localhost:8080/image/4) to see image id 4
* [localhost:8080/user](http://localhost:8080/user) to see your user profile data

## Monitoring
The API server and the MySQL container both have integrations with [New Relic](https://github.com/newrelic).  If you have a license key for this, you can enable monitoring simply by setting a docker environment variable.
```sh
cd docker
cp newrelic.template.env newrelic.env
```
Then edit `newrelic.env` and set the license key on the first line:
```
NEW_RELIC_LICENSE_KEY=YOUR_LICENSE_KEY_HERE
```

Running docker-compose will automatically use this key, and you should see additional lines in the logs like so:
```
api-server      |  INFO: 2019/02/24 01:18:32 New Relic monitor started
```
and
```
tagdemo-mysql   | INFO: New Relic monitor started
```

The `build.sh` script creates an empty stub of `newrelic.env` if you are not using monitoring.

## Developing
For developing, you can run this server locally without Docker.

### Build and Run the Database
Build the database manually as follows:
```sh
cd docker
docker build -t tagdemo/mysql ./mysql
```
You can then start the MySQL container as follows:
```sh
docker run --name tagdemo-mysql --rm -p 3306:3306 tagdemo/mysql
```
As shown above, we are mapping the MySQL default port 3306 from the container, to 3306 on localhost.

* If this conflicts with a local installations of MySQL server, specify a different port
* If you change the MySQL port, also specify `-dbport` in the [api-server/Dockerfile](https://github.com/DavidSantia/tag_api/blob/master/docker/auth-server/Dockerfile) entrypoint.

The database will be ready after you see the message:
```
[Entrypoint] MySQL init process done. Ready for start up.
```

If you need to stop the MySQL container, use
```sh
docker stop tagdemo-mysql
```
### Build and Run the API Server
In a separate terminal, build the API server as follows:
```sh
cd apps/api-server
govvv build
```

Use the help option to get command-line usage
```sh
./api-server -h
Usage of ./api-server:
  -bolt string
    	Specify BoltDB filename (default "./content.db")
  -dbhost string
    	Specify DB host
  -dbport string
    	Specify DB port (default "3306")
  -debug
    	Debug logging
  -host string
    	Specify Api host
  -log string
    	Specify logging filename
  -nhost string
    	Specify NATS host
  -nport string
    	Specify NATS port (default "4222")
  -port string
    	Specify Api port (default "8080")
```

Then run the server as follows:
```sh
./api-server -dbhost 127.0.0.1 -dbload -debug
```

## How it works
The database loader uses the Go [reflect](https://golang.org/pkg/reflect) package to auto-generate SELECT statements from the struct tags.

### Example Struct
```go
type Image struct {
	Id           int64   `json:"id" db:"id"`
	Width        int64   `json:"width" db:"width"`
	Height       int64   `json:"height" db:"height"`
	Url          string  `json:"url" db:"url"`
	Title        *string `json:"title" db:"title"`
	Artist       *string `json:"artist" db:"artist"`
	Gallery      *string `json:"gallery" db:"gallery"`
	Organization *string `json:"organization" db:"organization"`
}
```
Tags shown above are interpreted as follows:
* **json**: field name returned in API
* **db**: field name in SQL
* **sql**: optional SQL for SELECT

The **sql** tag is useful when
* using joined statements with otherwise ambiguous field names
* you want to insert an IFNULL or other logic


### func (data *ApiData) makeQuery
```go
func (data *ApiData) MakeQuery(dt interface{}, query string, v ...interface{}) (finalq string)
```
This takes two inputs:
* **dt**: the struct you are loading data into
* **query**: the FROM and WHERE part of a query

It can also take optional **v** parameters.  If using these, include a format 'verb' (see the Go [fmt](https://golang.org/pkg/fmt/#hdr-Printing) package) in your query for each parameter.

It returns one output, the generated query.

### Example Code
```go
query = makeQuery(Image{},"FROM images i WHERE i.media IS NOT NULL"),
```

This combines the fields in the *Image()* struct, along with the FROM clause provided, as follows:
```sql
SELECT id, width, height, url, title, artist, gallery, organization
FROM images i WHERE i.media IS NOT NULL
```

The **api-server** `-debug` flag reveals the generated SQL queries.
```sh
DEBUG: 2019/05/14 16:31:59 GroupQuery: SELECT id, name, sess_seconds
FROM groups g
 INFO: 2019/05/14 16:31:59 Load Groups: 3 entries total
DEBUG: 2019/05/14 16:31:59 ImageQuery: SELECT id, width, height, url, title, artist, gallery, organization
FROM images i WHERE i.media IS NOT NULL
 INFO: 2019/05/14 16:31:59 Load Images: 12 entries total
```

Using the [sqlx](https://jmoiron.github.io/sqlx) package, the code loads each object with a single *rows.StructScan()* step as shown:
```go
for rows.Next() {
	err = rows.StructScan(&image)
	if err != nil {
		fmt.Printf("Load Images: %v\n", err)
		continue
	}
	bs.ImageMap[image.Id] = image
}
```
Above we see all the parameters from *makeQuery()* are loaded into an image object.

