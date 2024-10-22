package main

import "fmt"
var taskItems = []string{"Görev 1", "Görev 2", "Görev 3"}
func main() {
  fmt.Println("##### İlk uygulama ####")
  fmt.Println("Todolist App")

  
  printTasks()
  
}

func printTasks(){
  for index, task := range taskItems{
    //fmt.Println(task)
    fmt.Printf("%d. %s\n",index+1,task)
  }
}
