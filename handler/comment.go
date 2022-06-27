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

type CommentHandler struct {
	commentService service.CommentServiceInterface
}

func NewCommentHandler(commentService service.CommentServiceInterface) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (c CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	user := middleware.ForContext(ctx)
	id := strconv.Itoa(user.ID)

	var input entity.CommentInput
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)

		return
	}

	newComment, errCreate := c.commentService.CreateComment(ctx, id, input)

	if errCreate != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Created Comment", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := map[string]interface{}{
		"id":         newComment.ID,
		"message":    newComment.Message,
		"photo_id":   newComment.PhotoID,
		"user_id":    newComment.UserID,
		"created_at": newComment.CreatedAt,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Created Comment", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}

func (c CommentHandler) GetAllComment(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	comments, errGetComments := c.commentService.GetAllComment(ctx)

	if errGetComments != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Get All Comments", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	var commentResponse []map[string]interface{}

	for _, val := range comments {
		formatter := map[string]interface{}{
			"id":         val.ID,
			"message":    val.Message,
			"photo_id":   val.PhotoID,
			"user_id":    val.UserID,
			"created_at": val.CreatedAt,
			"updated_at": val.UpdatedAt,
			"user": map[string]interface{}{
				"id":       val.User.ID,
				"email":    val.User.Email,
				"username": val.User.Username,
			},
			"photo": map[string]interface{}{
				"id":        val.Photo.ID,
				"title":     val.Photo.Title,
				"caption":   val.Photo.Caption,
				"photo_url": val.Photo.PhotoUrl,
				"user_id":   val.Photo.UserID,
			},
		}
		commentResponse = append(commentResponse, formatter)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Get All Comments", http.StatusOK, true, commentResponse)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}

}

func (c CommentHandler) GetCommentById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	getById, errById := c.commentService.GetCommentById(ctx, id)

	if errById != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Get By Id Comments", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := map[string]interface{}{
		"id":         getById.ID,
		"message":    getById.Message,
		"photo_id":   getById.PhotoID,
		"user_id":    getById.UserID,
		"created_at": getById.CreatedAt,
		"updated_at": getById.UpdatedAt,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Get By Id Comment", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}

func (c CommentHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	var input entity.CommentInput
	errDecode := json.NewDecoder(r.Body).Decode(&input)

	if errDecode != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		helper.APIResponseFailed("Failed to Decode", http.StatusInternalServerError, false)

		return
	}

	updateComment, errUpdateComment := c.commentService.UpdateComment(ctx, id, input)

	if errUpdateComment != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Updated Comments", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := map[string]interface{}{
		"id":         updateComment.ID,
		"message":    updateComment.Message,
		"photo_id":   updateComment.PhotoID,
		"user_id":    updateComment.UserID,
		"updated_at": updateComment.UpdatedAt,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Updated Comment", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}

func (c CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	err := c.commentService.DeleteComment(ctx, id)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Deleted Comment", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := w.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data := helper.APIResponseSuccessWithoutData("Your comment has been successfully deleted")
	jsonData, _ := json.Marshal(&data)
	_, errWrite := w.Write(jsonData)
	if errWrite != nil {
		return
	}
}
