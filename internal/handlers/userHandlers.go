package handlers

import (
	"context"
	"project/internal/userService"
	"project/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) GetUsersIdTasks(ctx context.Context, request users.GetUsersIdTasksRequestObject) (users.GetUsersIdTasksResponseObject, error) {
	userID := request.Id
	tasksForUser, err := h.Service.GetTasksForUser(userID)
	if err != nil {
		return nil, err
	}

	response := users.GetUsersIdTasks200JSONResponse{}
	for _, tsk := range tasksForUser {
		task := users.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)

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

func (h *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	var userToUpdate userService.User

	if request.Body.Email != nil {
		userToUpdate.Email = *request.Body.Email
	}
	if request.Body.Password != nil {
		userToUpdate.Password = *request.Body.Password
	}

	updatedUser, err := h.Service.UpdateUserByID(request.Id, userToUpdate)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return response, nil
}
func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.Service.DeleteUserByID(request.Id)
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
