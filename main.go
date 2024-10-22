package main

import "fmt"

func main() {
  fmt.Println("##### İlk uygulama ####")
  fmt.Println("Todolist App")
  var taskItems = []string{"Görev 1", "Görev 2", "Görev 3"}
  
  printTasks(taskItems)
  fmt.Println()
  addTask(taskItems, "Go for a run")
  
}

func printTasks(taskItems []string){
  for index, task := range taskItems{
    //fmt.Println(task)
    fmt.Printf("%d. %s\n",index+1,task)
  }
}

func addTask(taskItems []string, newTask string){

  var updatedTaskItems = append(taskItems,newTask)
  printTaks(updatedTaskItems)
}
