package main
import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID 		  int 	 `json:"id"`
	Completed bool 	 `json:"completed"`
	Body 	  string `json:"body"`

}

func main()  {
	fmt.Println("Hello Universe")
	// var myName string = "John Doe"
	// const mySecondName string = "Jane Doe"

	// myThirdName := "Bob Doe"
	// fmt.Println(myName)
	// fmt.Println(mySecondName)
	// fmt.Println(myThirdName)
	app := fiber.New()

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status((200)).JSON(todos) // formerly fiber.Map{"msg": "hello universe"}
	})

	// Create a ToDo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} // {id:0,completed:false,body:""} ; this object comes from app.Post function and is assined into todo
		
		if err := c.BodyParser(todo); err != nil{
			return err	
		}
		// BodyParser binds body to request

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error":"Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		var x int = 5 // 0x00001
		var p *int = &x // 0x00001
		fmt.Println(p) // 0x00001
		fmt.Println(p) // 5

		return c.Status(201).JSON(todo)

	})

	// Update a Todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true // todos[i].Completed = !todos[i].Completed // true is gonna be false, false is gonna be true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	// Delete a Todo
	// app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {

	// })

	log.Fatal(app.Listen(":4000"))
}