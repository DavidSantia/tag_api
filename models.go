package tag_api

// User data

type User struct {
	Id         int64  `json:"id" db:"id"`
	GroupId    int64  `json:"group_id" db:"group_id"`
	Guid       string `json:"guid" db:"guid"`
	FirstName  string `json:"first_name" db:"first_name"`
	MiddleInit string `json:"middle_init" db:"middle_init"`
	LastName   string `json:"last_name" db:"last_name"`
	Email      string `json:"email" db:"email"`
	Addr       string `json:"addr" db:"addr"`
	City       string `json:"city" db:"city"`
	State      string `json:"state" db:"state"`
	Zip        string `json:"zip" db:"zip"`
	Gender     string `json:"gender" db:"gender"`
	Status     bool   `json:"status" db:"status"`
}

const UserQuery = "FROM users u " +
	"WHERE u.status IS NOT NULL"

// Image data

// Pointers to int/string to allow for 'null' value in JSON
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

// Group data

type Group struct {
	Id              int64           `json:"id" db:"id"`
	Name            string          `json:"name" db:"name"`
	ImagesGroupsMap ImagesGroupsMap `json:"-"`
}

const GroupQuery = "FROM groups g"

// ImagesGroups data

type ImagesGroups struct {
	GroupId int64 `json:"group_id" db:"group_id"`
	ImageId int64 `json:"image_id" db:"image_id"`
}

const ImagesGroupsQuery = "FROM images_groups ig"
