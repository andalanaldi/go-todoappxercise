// 2nd part : API with DB (MongoDB)
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //int, if value is false, do not add to response, what's gonna happen if dont add omitempty field, the null id return null or 0000, if there is omitempty field we got default id generated by mongodb
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

// bson = binary json
var collection *mongo.Collection

func main() {
	fmt.Println("Hello universe")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil) //:=
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB ATLAS")

	collection = client.Database("golang_db").Collection("todos")
	// (*mongo.collection)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{}) // trying to pass our filter, in case this is no filter, we ant all todos in collection, if u pass filter, should intro object, this case we want to pass all collection todos, assign var with cursor, mongo db return cursor, pointer to the result set, write over the doc return query

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())
	// suround function is completed run the right handside ?

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
} // we have created connection, fetching, we want to this fucntion to be done

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	// {id:0,completed:false,body:""}

	if err := c.BodyParser(todo); err != nil { // bind req body, pass into struct
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	} //then we want to insert it to MongoDB right?

	// if we pass all of the if checks so that
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id") // string to object ready
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}

// 1st Part : API without DB

// package main
// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/joho/godotenv"
// )

// type Todo struct {
// 	ID 		  int 	 `json:"id"`
// 	Completed bool 	 `json:"completed"`
// 	Body 	  string `json:"body"`

// }

// func main()  {
// 	fmt.Println("Hello Universe")
// 	// var myName string = "John Doe"
// 	// const mySecondName string = "Jane Doe"

// 	// myThirdName := "Bob Doe"
// 	// fmt.Println(myName)
// 	// fmt.Println(mySecondName)
// 	// fmt.Println(myThirdName)
// 	app := fiber.New()

// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	PORT := os.Getenv("PORT")

// 	todos := []Todo{}

// 	// Get ToDos
// 	app.Get("/api/todos", func(c *fiber.Ctx) error {
// 		return c.Status((200)).JSON(todos) // formerly fiber.Map{"msg": "hello universe"}
// 	})

// 	// Create a ToDo
// 	app.Post("/api/todos", func(c *fiber.Ctx) error {
// 		todo := &Todo{} // {id:0,completed:false,body:""} ; this object comes from app.Post function and is assined into todo

// 		if err := c.BodyParser(todo); err != nil{
// 			return err
// 		}
// 		// BodyParser binds body to request

// 		if todo.Body == "" {
// 			return c.Status(400).JSON(fiber.Map{"error":"Todo body is required"})
// 		}

// 		todo.ID = len(todos) + 1
// 		todos = append(todos, *todo)

// 		var x int = 5 // 0x00001
// 		var p *int = &x // 0x00001
// 		fmt.Println(p) // 0x00001
// 		fmt.Println(p) // 5

// 		return c.Status(201).JSON(todo)

// 	})

// 	// Update a Todo
// 	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")

// 		for i, todo := range todos {
// 			if fmt.Sprint(todo.ID) == id {
// 				todos[i].Completed = true // todos[i].Completed = !todos[i].Completed // true is gonna be false, false is gonna be true
// 				return c.Status(200).JSON(todos[i])
// 			}
// 		}

// 		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
// 	})

// 	// Delete a Todo
// 	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")

// 		for i, todo := range todos {
// 			if fmt.Sprint(todo.ID) == id {
// 				todos = append(todos[:i],todos[i+1:]...)
// 				return c.Status(200).JSON(fiber.Map{"success": true})
// 			}
// 		}

// 		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
// 		// we have 5 values, each of them are 1 2 3 4 5, we delete one which is the 3rd, so it becomes:
// 		// 1 2 4 5
// 		// ... == periodic operator === spread operator in js to unpack the values
// 	})

// 	log.Fatal(app.Listen(":"+PORT)) // formerly ":4000" wo PORT err
// }
