package start

import "fmt"


type Todo struct {
	id int
	title string
	completed bool
}

type Toggleable interface{
	toggle()
}

func (t *Todo) toggle () {
	t.completed = !t.completed
}

func main() {
	// arrays
	a := [3]int{1, 2, 3}
	b := []int{1,2,3}
	b = append(b, 4)

	fmt.Printf("%#v %#v", a, b)
	fmt.Println("")

	// maps
	c := map[string]string{
		"key1": "value1",
		"key2": "value2",
		}
		c["key1"]="value1updated"
		fmt.Printf("%#v", c["key1"])
		fmt.Println("")

		// structs
		d := Todo{id: 1, title: "Learn Golang", completed: false,}
		fmt.Printf("%#v", d)
		fmt.Println("")

		// methods
		d.toggle()
		fmt.Printf("%#v", d)
		fmt.Println("")

		// interfaces
		toggleTodo(&d)
		fmt.Printf("%#v", d)
		fmt.Println("")

		todolist := []Toggleable{
			&d,
			&Todo{id: 2, title: "Learn ruby", completed: false,},
			&Todo{id: 1, title: "Learn kotlin", completed: false,},
			}
			toggleTodolist(todolist)
		fmt.Printf("%#v", todolist)
}

func toggleTodo(t Toggleable){
	t.toggle()
}

func toggleTodolist(t []Toggleable){
	fmt.Printf("toggle todolist")
	fmt.Println("")
}
