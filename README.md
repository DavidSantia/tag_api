# tag_api
API using Go-lang struct tags to load SQL data, and implement JSON endpoints

* This project builds a Docker database container with sample Image and User data
* The sample data images come from https://clients3.google.com/cast/chromecast/home

You can clone the project with
```sh
go get "github.com/DavidSantia/tag_api"
```

It also uses govvv to provide the Github version string in the code.
```sh
go get "github.com/ahmetb/govvv"
```

## Database Setup

Build the database container as follows
```sh
docker build -t tagdemo ./data
```

Start the MySQL container as follows:
```sh
docker run --name tag_api_db --rm -p 6603:3306 tagdemo
```
As shown above, we are mapping the MySQL default port 3306 from the container, to 6603 on localhost.

* Port 6603 does not conflict with any local installations of MySQL server on the default port
* If you want to specify a different port on the docker run command, also edit **DbPort** in [config.go](https://github.com/DavidSantia/tag_api/blob/master/config.go)

The database will be ready after you see the message:
```
[Entrypoint] MySQL init process done. Ready for start up.
```

If you need to stop the MySQL container, use
```sh
docker kill tag_api_db
```

## API Server Setup

Build the API server as follows
```sh
cd api
govvv build
```

Use `./api -help` to get command-line usage
```
Usage of ./api:
  -debug
    	Debug logging
  -log string
    	Specify logging filename
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

const ImageQuery = "FROM images i " +
	"WHERE i.media IS NOT NULL"
```
Tags shown above are interpreted as follows:
* **json**: field name returned in API
* **db**: field name in SQL
* **sql**: optional SQL for SELECT

The **sql** tag is useful when
* using joined statements with otherwise ambiguous field names
* you want to insert an IFNULL or other logic

Use `./api -debug` to debug the SQL queries that are being auto-generated from the struct tags.
```
DEBUG: 2017/08/26 14:07:03 Select "id" for field: Id [int64]
DEBUG: 2017/08/26 14:07:03 Select "width" for field: Width [int64]
DEBUG: 2017/08/26 14:07:03 Select "height" for field: Height [int64]
DEBUG: 2017/08/26 14:07:03 Select "url" for field: Url [string]
DEBUG: 2017/08/26 14:07:03 Select "title" for field: Title [*string]
DEBUG: 2017/08/26 14:07:03 Select "artist" for field: Artist [*string]
DEBUG: 2017/08/26 14:07:03 Select "gallery" for field: Gallery [*string]
DEBUG: 2017/08/26 14:07:03 Select "organization" for field: Organization [*string]
DEBUG: 2017/08/26 14:07:03 Select "media" for field: Media [string]
DEBUG: 2017/08/26 14:07:03 ImageQuery: SELECT id, width, height, url, title, artist, gallery, organization, media FROM images i WHERE i.media IS NOT NULL
```

### func (data *ApiData) MakeQuery
```go
func (data *ApiData) MakeQuery(dt interface{}, query string, v ...interface{}) (finalq string)
```
This takes two inputs:
* **dt**: the struct you are loading data into
* **query**: the FROM and WHERE part of a query

It can also take optional **v** parameters.  If using these, include a format 'verb' (see the Go [fmt](https://golang.org/pkg/fmt/#hdr-Printing) package) in your query for each parameter.

MakeQuery returns one output, the final query. This will be a combination of the auto-generated SELECT statement, and the rest of the query.

### Example Code
```go
var i Image

// Load images
query := data.MakeQuery(i, ImageQuery)
rows, err := data.Db.Queryx(query)
if err != nil {
	fmt.Printf("Load Images: %v\n", err)
	return
}
```

Notice we have automatically assembled the query as follows:
```sql
SELECT id, width, height, url, title, artist, gallery, organization, media
  FROM images i
  WHERE i.media IS NOT NULL
```

Because we are using the sqlx package, we also load each struct in one step with **rows.StructScan()** as shown
```go
for rows.Next() {
	err = rows.StructScan(&i)
	if err != nil {
		fmt.Printf("Load Images: %v\n", err)
		continue
	}
	data.ImageMap[i.Id] = i
}
```
### Accessing the API
Once you have the API server up and running, you can use your browser to authenticate, and access data.

The following endpoints are defined:
```go
router.Handle("GET", "/authenticate", HandleAuthTester)
router.Handle("POST", "/authenticate", HandleAuthenticate)
router.Handle("GET", "/image", HandleAllImages)
router.Handle("GET", "/image/:Id", HandleImage)
router.Handle("GET", "/user", HandleUser)
```

By browsing to [localhost:8080/authenticate](http://localhost:8080/authenticate), you will see a test framework with two buttons.  Each button authenticates you as a particular user from the sample database, either in the Basic or Premium group.

Once authenticated, your browser will have a Session cookie to allow you to continue using the API.

You can then browse to

* [localhost:8080/image](http://localhost:8080/image) to see images you have access to
* [localhost:8080/image/4](http://localhost:8080/image/4) to see image id 4
* [localhost:8080/user](http://localhost:8080/user) to see your user profile data
