package handler

import (
	"encoding/json"
	"final-project/entity"
	"final-project/helper"
	"final-project/middleware"
	"final-project/service"
	"fmt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) Register(writer http.ResponseWriter, req *http.Request) {
	var user entity.User
	jsonDecoder := json.NewDecoder(req.Body)
	err := jsonDecoder.Decode(&user)

	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		helper.APIResponseFailed("Failed to Decode User", http.StatusInternalServerError, false)
		return
	}

	newUser, err := uh.userService.Register(req.Context(), user)

	if err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		dataErr := helper.APIResponseFailed("Failed to Created User", http.StatusBadRequest, false)
		jsonData, _ := json.Marshal(&dataErr)
		_, errWrite := writer.Write(jsonData)
		if errWrite != nil {
			return
		}
		return
	}

	formatter := map[string]interface{}{
		"id":       newUser.ID,
		"username": newUser.Username,
		"email":    newUser.Email,
		"password": newUser.Password,
		"age":      newUser.Age,
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	data := helper.APIResponse("Success to Created User", http.StatusOK, true, formatter)
	jsonData, _ := json.Marshal(&data)
	_, errWrite := writer.Write(jsonData)
	if errWrite != nil {
		return
	}

}

func (uh *UserHandler) Login(writer http.ResponseWriter, req *http.Request) {
	var input entity.UserLogin
	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		response, _ := json.Marshal(helper.APIResponseFailed(err.Error(), http.StatusBadRequest, false))
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(response)
		return
	}

	errValidation := helper.CheckEmpty(input.Email, input.Password)

	if errValidation != nil {
		response, _ := json.Marshal(helper.APIResponseFailed(errValidation.Error(), http.StatusBadRequest, false))
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(response)
		return
	}

	user, errLogin := uh.userService.Login(req.Context(), input)

	if errLogin != nil {
		response, _ := json.Marshal(helper.APIResponseFailed("Login Failed", http.StatusBadRequest, false))
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(response)
		return
	}

	id := strconv.Itoa(user.ID)

	token, errToken := middleware.GenerateToken(id)

	if errToken != nil {
		response, _ := json.Marshal(helper.APIResponseFailed("Failed to Generate Token", http.StatusBadRequest, false))
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(response)
		return
	}

	formatter := map[string]interface{}{
		"token": token,
	}

	response, _ := json.Marshal(helper.APIResponse("Success to Generate Token", http.StatusOK, true, formatter))
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (uh *UserHandler) Update(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user := middleware.ForContext(ctx)

	var input entity.UserLogin
	err := json.NewDecoder(req.Body).Decode(&input)

	if err != nil {
		fmt.Println("errDecode", err.Error())
		return
	}

	id := strconv.Itoa(user.ID)

	userUpdate, errUpdate := uh.userService.Update(ctx, id, input)

	if errUpdate != nil {
		response, _ := json.Marshal(helper.APIResponseFailed(errUpdate.Error(), http.StatusInternalServerError, false))

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(response)
	}

	formatter := map[string]interface{}{
		"id":         userUpdate.ID,
		"username":   userUpdate.Username,
		"email":      userUpdate.Email,
		"age":        userUpdate.Age,
		"updated_at": userUpdate.UpdatedAt,
	}

	response, _ := json.Marshal(helper.APIResponse("Success Update User", http.StatusOK, true, formatter))

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func (uh *UserHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user := middleware.ForContext(ctx)
	id := strconv.Itoa(user.ID)

	userId, err := uh.userService.GetById(ctx, id)

	if err != nil {
		response, _ := json.Marshal(helper.APIResponseFailed(err.Error(), http.StatusInternalServerError, false))

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write(response)
	}

	formatter := map[string]interface{}{
		"id":       userId.ID,
		"username": userId.Username,
		"email":    userId.Email,
		"age":      userId.Age,
	}

	response, _ := json.Marshal(helper.APIResponse("Success to toGet User", http.StatusOK, true, formatter))

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
