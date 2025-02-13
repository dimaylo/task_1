package handlers

import (
	"context"
	"strconv"

	"REST_API/internal/taskService"
	"REST_API/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

// Метод для GET /tasks
func (h *Handler) GetTasks(
	ctx context.Context,
	_ tasks.GetTasksRequestObject,
) (tasks.GetTasksResponseObject, error) {

	all, err := h.Service.GetAllTasks()
	if err != nil {
		// В strict-схеме нет "дефолтного" ответа под ошибку,
		// поэтому если хотим кинуть ошибку, можно вернуть nil, err
		return nil, err
	}

	resp := make(tasks.GetTasks200JSONResponse, 0, len(all))
	for _, t := range all {
		id := uint(t.ID)
		taskVal := t.Task
		isDone := t.IsDone
		resp = append(resp, tasks.Task{
			Id:     &id,
			Task:   &taskVal,
			IsDone: &isDone,
		})
	}
	return resp, nil
}

// Метод для POST /tasks
func (h *Handler) PostTasks(
	ctx context.Context,
	req tasks.PostTasksRequestObject,
) (tasks.PostTasksResponseObject, error) {

	body := req.Body
	if body == nil {
		return nil, nil // или вернуть ошибку
	}

	newTask := taskService.Task{
		Task:   ptrToStr(body.Task),
		IsDone: ptrToBool(body.IsDone),
	}
	created, err := h.Service.CreateTask(newTask)
	if err != nil {
		return nil, err
	}
	id := uint(created.ID)
	taskVal := created.Task
	isDoneVal := created.IsDone
	return tasks.PostTasks201JSONResponse{
		Id:     &id,
		Task:   &taskVal,
		IsDone: &isDoneVal,
	}, nil
}

// Метод для DELETE /tasks/{id} (видите, в api.gen.go он называется DeleteTasksId)
func (h *Handler) DeleteTasksId(
	ctx context.Context,
	req tasks.DeleteTasksIdRequestObject,
) (tasks.DeleteTasksIdResponseObject, error) {

	idStr := req.Id
	idUint, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		// Т.к. в strict-схеме не описан 400 или 500, можно вернуть nil, err
		return nil, err
	}
	if err := h.Service.DeleteTask(uint(idUint)); err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

// Метод для PATCH /tasks/{id} (в api.gen.go: PatchTasksId)
func (h *Handler) PatchTasksId(
	ctx context.Context,
	req tasks.PatchTasksIdRequestObject,
) (tasks.PatchTasksIdResponseObject, error) {

	idStr := req.Id
	idUint, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return nil, err
	}
	body := req.Body
	if body == nil {
		return nil, nil // или вернуть err
	}
	updated := taskService.Task{
		Task:   ptrToStr(body.Task),
		IsDone: ptrToBool(body.IsDone),
	}
	res, err := h.Service.UpdateTask(uint(idUint), updated)
	if err != nil {
		return nil, err
	}
	id := uint(res.ID)
	taskVal := res.Task
	isDoneVal := res.IsDone
	return tasks.PatchTasksId200JSONResponse{
		Id:     &id,
		Task:   &taskVal,
		IsDone: &isDoneVal,
	}, nil
}

func ptrToStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ptrToBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
