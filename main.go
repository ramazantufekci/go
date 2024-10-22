package main

import "fmt"

func main() {
  fmt.Println("##### İlk uygulama ####")
  fmt.Println("Todolist App")

  var taskItems = []string{"Görev 1", "Görev 2", "Görev 3"}
  printTasks(taskItems)
  
}

func printTasks(taskItems){
  for index, task := range taskItems{
    //fmt.Println(task)
    fmt.Printf("%d. %s\n",index+1,task)
  }
}
