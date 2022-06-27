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

type SocialMediaHandler struct {
	smService service.SocialMediaServiceInterface
}

func NewSocialMediaHandler(smService service.SocialMediaServiceInterface) *SocialMediaHandler {
	return &SocialMediaHandler{smService: smService}
}

func (s SocialMediaHandler) CreateSocialMedia(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	user := middleware.ForContext(ctx)
	id := strconv.Itoa(user.ID)

	var input entity.SocialMediaInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)

		return
	}

	newSm, errCreate := s.smService.CreateSocialMedia(ctx, id, input)

	if errCreate != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Created Social Media", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := map[string]interface{}{
		"id":               newSm.ID,
		"name":             newSm.Name,
		"social_media_url": newSm.SocialMediaUrl,
		"user_id":          newSm.UserID,
		"created_at":       newSm.CreatedAt,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Created Social Media", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}

func (s SocialMediaHandler) GetAllSocialMedia(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	socialMedias, errSocialMedias := s.smService.GetAllSocialMedia(ctx)

	if errSocialMedias != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Get All Social Medias", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	var socialMediasResponse []map[string]interface{}

	for _, val := range socialMedias {
		formatter := map[string]interface{}{
			"id":               val.ID,
			"name":             val.Name,
			"social_media_url": val.SocialMediaUrl,
			"user_id":          val.UserID,
			"created_at":       val.CreatedAt,
			"updated_at":       val.UpdatedAt,
			"user": map[string]interface{}{
				"id":       val.User.ID,
				"username": val.User.Username,
			},
		}
		socialMediasResponse = append(socialMediasResponse, formatter)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Get All Social Medias", http.StatusOK, true, socialMediasResponse)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}

func (s SocialMediaHandler) GetByIdSocialMedia(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	getById, errById := s.smService.GetSocialMediaById(ctx, id)

	if errById != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Get By Id Social Media", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := entity.SocialMedia{
		ID:             getById.ID,
		Name:           getById.Name,
		SocialMediaUrl: getById.SocialMediaUrl,
		UserID:         getById.UserID,
		CreatedAt:      getById.CreatedAt,
		UpdatedAt:      getById.UpdatedAt,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Get By Id Social Media", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}

func (s SocialMediaHandler) UpdateSocialMedia(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	var input entity.SocialMediaInput
	errDecode := json.NewDecoder(r.Body).Decode(&input)

	if errDecode != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)

		return
	}

	updateSm, errUpdateComment := s.smService.UpdateSocialMedia(ctx, id, input)

	if errUpdateComment != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Updated Social Media", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := entity.SocialMedia{
		ID:             updateSm.ID,
		Name:           updateSm.Name,
		SocialMediaUrl: updateSm.SocialMediaUrl,
		UserID:         updateSm.UserID,
		UpdatedAt:      updateSm.UpdatedAt,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Updated Social Media", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}

func (s SocialMediaHandler) DeleteSocialMedia(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	err := s.smService.DeleteSocialMedia(ctx, id)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Deleted Social Media", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponseSuccessWithoutData("Your social media has been successfully deleted")
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}
