package handlers

import (
	"context"
	"fmt"
	"project/internal/taskService"
	"project/internal/web/tasks"
)

type TaskHandler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
			UserId: &tsk.UserID,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	if taskRequest.UserId == nil {
		return nil, fmt.Errorf("user_id is required")
	}

	createdTask, err := h.Service.CreateTask(taskToCreate, *taskRequest.UserId)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserID,
	}
	return response, nil
}

func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	var taskToUpdate taskService.Task

	if request.Body.Task != nil {
		taskToUpdate.Task = *request.Body.Task
	}
	if request.Body.IsDone != nil {
		taskToUpdate.IsDone = *request.Body.IsDone
	}

	updatedTask, err := h.Service.UpdateTaskByID(request.Id, taskToUpdate)
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTaskByID(request.Id)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}
