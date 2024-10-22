package main

import (
	"fmt"
	"net/http"
)

var taskItems = []string{"Görev 1", "Görev 2", "Görev 3"}

func main() {
	fmt.Println("##### İlk uygulama ####")
	http.HandleFunc("/",helloUser)
	http.HandleFunc("/show-tasks",showTasks)
	http.ListenAndServe(":8082",nil)
   
  printTasks(taskItems)
  fmt.Println()
  taskItems = addTask(taskItems, "Go for a run")
  taskItems = addTask(taskItems, "Pratik yapıyoruz Go")
  printTasks(taskItems)
  
}

func showTasks(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer,taskItems)
}

func helloUser(writer http.ResponseWriter, request *http.Request){
	var greeting = "Hello user. Welcome to our Todolist App!"
	fmt.Fprintln(writer,greeting)
}

func printTasks(taskItems []string){
  for index, task := range taskItems{
    //fmt.Println(task)
    fmt.Printf("%d. %s\n",index+1,task)
  }
}

func addTask(taskItems []string, newTask string) []string {

  var updatedTaskItems = append(taskItems,newTask)
  return updatedTaskItems
}
