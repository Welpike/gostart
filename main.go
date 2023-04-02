package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
 * Run : ..\..\go1.20.2\go\bin\go run main.go
 * Tutorial : https://youtu.be/e8zL5PD8164?t=2508
 */

type Todo struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main () {
	ch := make(chan string)
	errCh := make(chan error)
	go func(){
		title, err := fetchTodoTitle()
		if err != nil{
			errCh <- err
		}else{
			ch <- title
		}
	}()
	fmt.Println("hello")
	defer close(ch)
	defer close(errCh)
	select{
	case err := <- errCh:
		panic(err)
	case title := <- ch:
		fmt.Printf("todo title : %s", title)
	}
}

func fetchTodoTitle () (string, error){
	r, err := http.Get("https://jsonplaceholder.typicode.com/todos?_limit=3")
	if err != nil {
		return "", fmt.Errorf("server connection error : %s", err)
	}
	defer r.Body.Close()
	var todolist []Todo
	err = json.NewDecoder(r.Body).Decode(&todolist)
	if err != nil {
		return "", fmt.Errorf("error when trying to parse server response : %s", err.Error())
	}
	if len(todolist)>0{
		return todolist[0].Title, nil
	}
	return "", fmt.Errorf("server connection error : %s", err)
}
