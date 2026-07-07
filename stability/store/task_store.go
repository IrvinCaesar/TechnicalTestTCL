package store

import "stability-test-task-api/models"

var Tasks = []models.Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build API", Done: false},
}
var nextID = len(Tasks) + 1

func GetAllTasks() []models.Task {
	return Tasks
}

func GetTaskByID(id int) *models.Task {
	for i := range Tasks {
		if Tasks[i].ID == id {
			return &Tasks[i]
		}
	}
	return nil
}

func AddTask(task models.Task) models.Task {
	task.ID = nextID
	nextID++
	Tasks = append(Tasks, task)
	return task
}
func UpdateTask(id int, updatedTask models.Task) *models.Task {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks[i] = updatedTask
			return &Tasks[i]
		}
	}
	return nil
}

func DeleteTask(id int) {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
		}
	}
}
