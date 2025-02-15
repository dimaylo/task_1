package handlers

import (
	"REST_API/internal/userService"
	"REST_API/internal/web/users"
	"golang.org/x/net/context"
)

type UserHandler struct {
	UserService *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (u *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.UserService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var response users.GetUsers200JSONResponse
	for _, usr := range allUsers {
		response = append(response, users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		})
	}
	return response, nil
}

func (u *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	newUser := userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}
	createdUser, err := u.UserService.CreateUser(newUser)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (u *UserHandler) PatchUsersUserId(_ context.Context, request users.PatchUsersUserIdRequestObject) (users.PatchUsersUserIdResponseObject, error) {
	updatedData := userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}
	userID := request.UserId
	updatedUser, err := u.UserService.UpdateUserByID(userID, updatedData)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsersUserId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil
}

func (u *UserHandler) DeleteUsersUserId(_ context.Context, request users.DeleteUsersUserIdRequestObject) (users.DeleteUsersUserIdResponseObject, error) {
	userID := request.UserId
	if err := u.UserService.DeleteUserByID(userID); err != nil {
		return nil, err
	}
	response := users.DeleteUsersUserId204Response{}
	return response, nil
}
