package handler

import (
	"net/http"

	"github.com/stwile/go_todo_app/entity"
)

type ListTask struct {
	Service ListTasksService
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHttp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Service.ListTasks(ctx)
	if err != nil {
		RespondJson(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJson(ctx, w, rsp, http.StatusOK)
}
