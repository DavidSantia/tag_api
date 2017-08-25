# tag_api
API using Go-lang struct tags to load SQL data, and implement JSON endpoints

* This project includes a database container with sample Image and User data
* The images come from https://clients3.google.com/cast/chromecast/home

You can clone the project with
```sh
$ go get "github.com/DavidSantia/tag_api"
```

It also uses govvv to provide the Github version string in the code.
```sh
$ go get "github.com/ahmetb/govvv"
```

## Database Setup

Build the database container as follows
```sh
$ docker build -t tagdemo ./data
```

Start the MySQL container as follows:
```sh
$ docker run --name tag_api_db --rm -p 6603:3306 tagdemo
```
As shown above, we are mapping the MySQL default port 3306 from the container, to 6603 on localhost.  This was chosen so as to not conflict in case you have locally installed a MySQL server using the default port.

The database will be ready after you see the message:
```
[Entrypoint] MySQL init process done. Ready for start up.

[Entrypoint] Starting MySQL 5.7.19-1.1.0
```

## API Setup

Build the API server as follows
```sh
$ cd api
$ govvv build
```

You can get command-line help as follows:
```sh
$ ./api -help
Usage of ./api:
  -debug
    	Debug logging
  -log string
    	Specify logging filename
```

## How it works

Use the *-debug* flag to see the SQL queries that are being auto-generated from the struct tags.

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
	Media        string  `json:"media" db:"media"`
}

const ImageQuery = "FROM images i " +
	"WHERE i.media IS NOT NULL"
```
Tags shown above are as follows:
* **json**: field name returned in API
* **db**: field name in SQL
* **sql**: optional SQL for SELECT

The **sql** tag is useful when
* using joined statements with otherwise ambiguous field names
* you want to insert an IFNULL or other logic


### func (data *ApiData) MakeQuery
```go
func (data *ApiData) MakeQuery(dt interface{}, query string, v ...interface{}) (finalq string)
```
This takes two inputs:
* **dt**: the struct you are loading data into
* **query**: the FROM and WHERE part of a query

The query can contain optional format 'verbs'; optional **v** parameters replace these via fmt.Sprintf

It returns one output:
The final query, a combination of the auto-generated SELECT statement, and the rest of the query.

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
