package tag_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

// DB loaders

func (data *ApiData) LoadImages() {
	var err error
	var query string
	var i Image
	var rows *sqlx.Rows

	// Load images
	query = data.MakeQuery(i, ImageQuery)
	logging.Debug.Printf("ImageQuery: %s\n", query)
	rows, err = data.Db.Queryx(query)
	if err != nil {
		logging.Error.Printf("Load Images: %v\n", err)
		return
	}
	for rows.Next() {
		err = rows.StructScan(&i)
		if err != nil {
			logging.Error.Printf("Load Images: %v\n", err)
			continue
		}
		data.ImageMap[i.Id] = i
	}
	logging.Info.Printf("Load Images: %d entries total\n", len(data.ImageMap))
}

// HTTP Handlers

func HandleAllImages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var i Image
	var is []Image

	group_id, err := GetGroupIdFromSession(r)
	if err != nil {
		HandleError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// See if group valid
	g, ok := d.GroupMap[group_id]
	if !ok {
		err = fmt.Errorf("GroupId %s not valid", group_id)
		HandleError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Make array of images for group
	for image_id := range g.ImagesGroupsMap {
		i, err = d.FindImage(image_id, group_id)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, err.Error())
			return
		}
		is = append(is, i)
	}

	b, err := json.Marshal(is)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	HandleReply(w, http.StatusOK, string(b))
}

func HandleImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var image_id int64
	var i Image

	group_id, err := GetGroupIdFromSession(r)
	if err != nil {
		HandleError(w, http.StatusUnauthorized, err.Error())
		return
	}

	iid := ps.ByName("Id")
	image_id, err = strconv.ParseInt(iid, 10, 64)
	if err != nil {
		HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	i, err = d.FindImage(image_id, group_id)
	if err != nil {
		HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	var b []byte
	b, err = json.Marshal(i)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, err.Error())
		return
	}

	HandleReply(w, http.StatusOK, string(b))
}

// Data Interfaces

func (data *ApiData) FindImage(image_id, group_id int64) (i Image, err error) {
	var g Group
	var ok bool

	// See if image exists
	_, ok = data.ImageMap[image_id]
	if !ok {
		err = fmt.Errorf("ImageId %d not valid", image_id)
		return
	}

	// See if group valid
	g, ok = data.GroupMap[group_id]
	if !ok {
		err = fmt.Errorf("GroupId %d not valid", group_id)
		return
	}

	// See if group has this image
	if !g.ImagesGroupsMap[image_id] {
		err = fmt.Errorf("ImageId %d not valid for GroupId %d", image_id, group_id)
	}
	return
}

func (data *ApiData) CreateImage(i Image) (err error) {
	var reply interface{}

	// Store in local Map
	data.ImageMap[i.Id] = i
	tag := fmt.Sprintf("m:%d", i.Id)

	// Save JSON blob to Redis
	b, err := json.Marshal(i)
	reply, err = data.Redis.Do("SET", tag, string(b))
	if err != nil {
		return
	}

	logging.Debug.Printf("SET %s %s\n", tag, reply)
	return
}

func (data *ApiData) DeleteImage(id int64) {

	// Store in local Map
	delete(data.ImageMap, id)
	tag := fmt.Sprintf("m:%d", id)

	// Delete in Redis
	reply, err := data.Redis.Do("DEL", tag)
	if err != nil {
		return
	}
	if reply.(int) != 1 {
		err = fmt.Errorf("DEL %s failed: %s", tag, reply)
		return
	}

	logging.Debug.Printf("DEL %s %s\n", tag, reply)
	return
}
