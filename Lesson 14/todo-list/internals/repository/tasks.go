package repository

import "todo/internals/models"

var TaskID uint

func SaveTask(task models.Task) {
	TaskID++
	Tasks[TaskID] = task
}

func SaveTaskChanges(ID uint, changedTask models.Task) {
	Tasks[ID] = changedTask
}

func GetAllTasks() map[uint]models.Task {
	return Tasks
}

func GetTaskByID(ID uint) (askedTask models.Task) {
	askedTask = Tasks[ID]
	return
}

func DeleteTaskByID(ID uint){
	delete(Tasks, ID)
}
