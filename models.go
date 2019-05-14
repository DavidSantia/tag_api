package tag_api

import (
	"github.com/newrelic/go-agent"
)

// Image data

// Pointers to int/string to allow for 'null' value in JSON
type Image struct {
	Id           int64   `json:"id" db:"id"`
	Width        int     `json:"width" db:"width"`
	Height       int     `json:"height" db:"height"`
	Url          string  `json:"url" db:"url"`
	Title        *string `json:"title" db:"title"`
	Artist       *string `json:"artist" db:"artist"`
	Gallery      *string `json:"gallery" db:"gallery"`
	Organization *string `json:"organization" db:"organization"`
}

var ImageSegment = newrelic.DatastoreSegment{
	Collection:         "images",
	Operation:          "SELECT",
	ParameterizedQuery: makeQuery(Image{}, "FROM images i WHERE i.media IS NOT NULL"),
}

// Group data

type Group struct {
	Id              int64           `json:"id" db:"id"`
	Name            string          `json:"name" db:"name"`
	SessionSeconds  *int64          `json:"sess_seconds" db:"sess_seconds"`
	ImagesGroupsMap ImagesGroupsMap `json:"-"`
}

var GroupSegment = newrelic.DatastoreSegment{
	Collection:         "groups",
	Operation:          "SELECT",
	ParameterizedQuery: makeQuery(Group{}, "FROM groups g"),
}

// ImagesGroups data

type ImagesGroups struct {
	GroupId int64 `json:"group_id" db:"group_id"`
	ImageId int64 `json:"image_id" db:"image_id"`
}

var ImageGroupSegment = newrelic.DatastoreSegment{
	Collection:         "images_groups",
	Operation:          "SELECT",
	ParameterizedQuery: makeQuery(ImagesGroups{}, "FROM images_groups ig"),
}

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

var UserSegment = newrelic.DatastoreSegment{
	Collection:         "users",
	Operation:          "SELECT",
	ParameterizedQuery: makeQuery(User{}, "FROM users u WHERE u.id = ?"),
}

type UserMessage struct {
	Command   string `json:"command"`
	Id        int64  `json:"id"`
	GroupId   int64  `json:"group_id"`
	Guid      string `json:"guid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
