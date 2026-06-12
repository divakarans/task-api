package handlers

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
	"task-api/models"
	"task-api/store"
	"time"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(store.Tasks)
	return

}

func CreateTasksHandler(w http.ResponseWriter, r *http.Request) {

	var tsk models.Task

	err := json.NewDecoder(r.Body).Decode(&tsk)
	if err != nil {
		http.Error(w, "Not found", http.StatusBadRequest)
		return
	}

	if tsk.Title == "" {
		http.Error(w, "Input required", http.StatusBadRequest)
		return

	} else if tsk.Status == "" {
		tsk.Status = "pending"

	} else if tsk.Status != "pending" && tsk.Status != "done" && tsk.Status != "in_progress" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tsk.ID = store.NextID
	tsk.CreatedAt = time.Now()
	store.NextID++
	store.Tasks = append(store.Tasks, tsk)
	json.NewEncoder(w).Encode(tsk)
	return
}

func GetTasksByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	Idnum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(store.Tasks); i++ {
		if Idnum == store.Tasks[i].ID {
			json.NewEncoder(w).Encode(store.Tasks[i])
			return
		}
	}
	http.Error(w, "no task id found", http.StatusNotFound)
}

func DeleteTasksByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	Idnum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid Id", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(store.Tasks); i++ {
		if Idnum == store.Tasks[i].ID {
			store.Tasks = slices.Delete(store.Tasks, i, i+1)

			w.WriteHeader(http.StatusOK)
			return

		}
	}
	http.Error(w, "no task id found", http.StatusNotFound)
}

func UpdateTasksByIdHandler(w http.ResponseWriter, r *http.Request) {

	var tsk models.Task
	err := json.NewDecoder(r.Body).Decode(&tsk)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	id := r.PathValue("id")
	Idnum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(store.Tasks); i++ {
		if Idnum == store.Tasks[i].ID {

			if tsk.Title != "" {
				store.Tasks[i].Title = tsk.Title
			}
			if tsk.Status != "" {

				if tsk.Status != "pending" && tsk.Status != "in_progress" && tsk.Status != "done" {

					http.Error(w, "Invalid status", http.StatusBadRequest)
					return
				}

				store.Tasks[i].Status = tsk.Status
			}
			json.NewEncoder(w).Encode(store.Tasks[i])
			return
		}
	}
	http.Error(w, "no task id found", http.StatusNotFound)

}
