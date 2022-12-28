package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task1 := Tasks{
		ID:         "1",
		TaskName:   "1st Projects",
		TaskDetail: "Finish the project",
		Date:       "20-22",
	}
	tasks = append(tasks, task1)
	task2 := Tasks{
		ID:         "2",
		TaskName:   "New projects",
		TaskDetail: "Finish the project",
		Date:       "20-22",
	}
	tasks = append(tasks, task2)
	task3 := Tasks{
		ID:         "3",
		TaskName:   "New projects",
		TaskDetail: "Finish the project",
		Date:       "20-22",
	}
	tasks = append(tasks, task3)
	fmt.Println("Your tasks are", tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home page")
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func getTask(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "task not found"})
	}
}
func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Tasks
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(1000))
	currentTime := time.Now().Format("01-02-2006")
	task.Date = currentTime
	// fmt.Println(task)
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//loop , hit id, perform remove (index,index+1)

	for index, task := range tasks {

		if task.ID == params["id"] {

			tasks = append(tasks[:index], tasks[index+1:]...)
			json.NewEncoder(w).Encode("Item Deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Item not Found")

}
func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//grab the id from params
	params := mux.Vars(r)
	//loop , get id , remove ,add with my ID

	for index, task := range tasks {
		if task.ID == params["id"] {
			var firstList []Tasks
			var secondList []Tasks
			firstList = tasks[:index]
			secondList = tasks[index+1:]
			var task Tasks
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			tasks = append(firstList, task)
			tasks = append(tasks, secondList...)
			json.NewEncoder(w).Encode(task)
			return
		}
	}
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":4000", router))

}

func main() {
	fmt.Println("Welcome to Task Management Backend")

	allTasks()
	handleRoutes()
}
