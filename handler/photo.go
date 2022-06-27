package handler

import (
	"encoding/json"
	"final-project/entity"
	"final-project/helper"
	"final-project/middleware"
	"final-project/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PhotoHandler struct {
	photoService service.PhotoServiceInterface
}

func NewPhotoHandler(photoService service.PhotoServiceInterface) *PhotoHandler {
	return &PhotoHandler{photoService: photoService}
}

func (ph PhotoHandler) CreatePhoto(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	user := middleware.ForContext(ctx)
	id := strconv.Itoa(user.ID)

	var photo entity.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)

		return
	}

	newPhoto, errCreate := ph.photoService.CreatePhoto(ctx, id, photo)

	if errCreate != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Created Photo", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := map[string]interface{}{
		"id":         newPhoto.ID,
		"title":      newPhoto.Title,
		"caption":    newPhoto.Caption,
		"photo_url":  newPhoto.PhotoUrl,
		"user_id":    newPhoto.UserID,
		"created_at": newPhoto.CreatedAt,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Created Photo", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}

}

func (ph PhotoHandler) GetPhotos(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	photos, errGetPhotos := ph.photoService.GetPhotos(ctx)

	if errGetPhotos != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Get All Photos", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	var photoRequest []*map[string]interface{}

	for _, val := range photos {
		formatter := map[string]interface{}{
			"id":         val.ID,
			"title":      val.Title,
			"caption":    val.Caption,
			"photo_url":  val.PhotoUrl,
			"user_id":    val.UserID,
			"created_at": val.CreatedAt,
			"user": map[string]interface{}{
				"email":    val.User.Email,
				"username": val.User.Username,
			},
		}
		photoRequest = append(photoRequest, &formatter)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Created Photo", http.StatusOK, true, photoRequest)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}

}

func (ph PhotoHandler) GetPhotoById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	photoId, err := ph.photoService.GetPhotoById(ctx, id)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Get Photos By Id", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := map[string]interface{}{
		"id":         photoId.ID,
		"title":      photoId.Title,
		"caption":    photoId.Caption,
		"photo_url":  photoId.PhotoUrl,
		"user_id":    photoId.UserID,
		"created_at": photoId.CreatedAt,
		"updated_at": photoId.UpdatedAt,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Created Photo", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}

}

func (ph PhotoHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	var photo entity.Photo
	errDecode := json.NewDecoder(r.Body).Decode(&photo)

	if errDecode != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)

		return
	}

	updatePhoto, err := ph.photoService.UpdatePhoto(ctx, id, photo)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Update Photo", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := map[string]interface{}{
		"id":         updatePhoto.ID,
		"title":      updatePhoto.Title,
		"caption":    updatePhoto.Caption,
		"photo_url":  updatePhoto.PhotoUrl,
		"user_id":    updatePhoto.UserID,
		"updated_at": updatePhoto.UpdatedAt,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Created Photo", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}

func (ph PhotoHandler) DeletePhoto(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	params := mux.Vars(r)
	id := params["id"]

	err := ph.photoService.DeletePhoto(r.Context(), id)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Deleted Photo", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponseSuccessWithoutData("Your photo has been successfully deleted")
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}

}
