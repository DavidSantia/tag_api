package tag_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"strconv"
)

// HTTP Handlers

func HandleAllImages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var i Image
	var is []Image

	group_id, err := GetGroupIdFromSession(r)
	if err != nil {
		HandleError(w, http.StatusUnauthorized, r.RequestURI, err)
		return
	}

	// See if group valid
	g, ok := d.GroupMap[group_id]
	if !ok {
		err = fmt.Errorf("GroupId %s not valid", group_id)
		HandleError(w, http.StatusUnauthorized, r.RequestURI, err)
		return
	}

	// Make array of images for group
	for image_id := range g.ImagesGroupsMap {
		i, err = d.FindImage(image_id, group_id)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
			return
		}
		is = append(is, i)
	}

	b, err := json.Marshal(is)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
		return
	}

	HandleReply(w, http.StatusOK, string(b))
}

func HandleImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var image_id int64
	var i Image

	group_id, err := GetGroupIdFromSession(r)
	if err != nil {
		HandleError(w, http.StatusUnauthorized, r.RequestURI, err)
		return
	}

	iid := ps.ByName("Id")
	image_id, err = strconv.ParseInt(iid, 10, 64)
	if err != nil {
		HandleError(w, http.StatusBadRequest, r.RequestURI, err)
		return
	}

	i, err = d.FindImage(image_id, group_id)
	if err != nil {
		HandleError(w, http.StatusBadRequest, r.RequestURI, err)
		return
	}

	var b []byte
	b, err = json.Marshal(i)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
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
		return
	}

	i, _ = data.ImageMap[image_id]
	return
}
