package tag_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"strconv"
)

// HTTP Handlers

func makeHandleAllImages(cs ContentService) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var err error
		var image Image
		var imageArr []Image
		var group Group
		var imageId, groupId int64
		var ok bool

		groupId, err = GetGroupIdFromSession(r)
		if err != nil {
			HandleError(w, http.StatusUnauthorized, r.RequestURI, err)
			return
		}

		// See if group valid
		group, ok = cs.GetGroup(groupId)
		if !ok {
			err = fmt.Errorf("GroupId %s not valid", groupId)
			HandleError(w, http.StatusUnauthorized, r.RequestURI, err)
			return
		}

		// Make array of images for group
		for imageId = range group.ImagesGroupsMap {
			image, err = findImage(cs, imageId, groupId)
			if err != nil {
				HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
				return
			}
			imageArr = append(imageArr, image)
		}

		b, err := json.Marshal(imageArr)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
			return
		}

		HandleReply(w, http.StatusOK, string(b))
	}
}

func makeHandleImage(cs ContentService) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var err error
		var imageId, groupId int64
		var image Image

		groupId, err = GetGroupIdFromSession(r)
		if err != nil {
			HandleError(w, http.StatusUnauthorized, r.RequestURI, err)
			return
		}

		imageId, err = strconv.ParseInt(ps.ByName("Id"), 10, 64)
		if err != nil {
			HandleError(w, http.StatusBadRequest, r.RequestURI, err)
			return
		}

		image, err = findImage(cs, imageId, groupId)
		if err != nil {
			HandleError(w, http.StatusBadRequest, r.RequestURI, err)
			return
		}

		var b []byte
		b, err = json.Marshal(image)
		if err != nil {
			HandleError(w, http.StatusInternalServerError, r.RequestURI, err)
			return
		}

		HandleReply(w, http.StatusOK, string(b))
	}
}

// Data Interfaces

func findImage(cs ContentService, imageId, groupId int64) (i Image, err error) {
	var group Group
	var ok bool

	// See if image exists
	_, ok = cs.GetImage(imageId)
	if !ok {
		err = fmt.Errorf("ImageId %d not valid", imageId)
		return
	}

	// See if group valid
	group, ok = cs.GetGroup(groupId)
	if !ok {
		err = fmt.Errorf("GroupId %d not valid", groupId)
		return
	}

	// See if group has this image
	if !group.ImagesGroupsMap[imageId] {
		err = fmt.Errorf("ImageId %d not valid for GroupId %d", imageId, groupId)
		return
	}

	i, _ = cs.GetImage(imageId)
	return
}
