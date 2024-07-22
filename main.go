package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 2.
type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //mongodb guarda la info en binaryjson:bson
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

// 3.
var collection *mongo.Collection

// 1.
func main() {
	fmt.Println("hola mundo")

	//4. cargar el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Errir loading .env file:", err)
	}

	//5. uri del db
	MONGODB_URI := os.Getenv("MONGODB_URI")
	//6. Conectandose a mongodb
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	//obtener el cliente y error
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	//12. cerrar la conexion cuando ya no se este usando
	defer client.Disconnect(context.Background())

	//7. Checar conexion con el metodo ping
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB ATLAS")
	//8.
	collection = client.Database("golang_db").Collection("tareas")
	//9. Creando el entorno web con fiber
	app := fiber.New()
	//endpoints
	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodos)
	app.Patch("/api/todos/:id", updateTodos)
	app.Delete("/api/todos/:id", deleteTodos)
	//definidendo puerto y manejando errores
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}

// 10. Creando los handlers para cada endpoint
func getTodos(c *fiber.Ctx) error {
	//instancia del objeto struct todo
	var todos []Todo

	//se pasan filtros a la coleccion. En este caso no hay filtros. Se quieren traer todas las tareas a la coleccion
	//cuando se hace una query en mdb, se retorna un cursor (pointer al resultset)
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	//11. Se crea una funcion que va a cerrar el cursor cuando el handler termine su tarea
	defer cursor.Close(context.Background())

	//se itera el cursor: se crea una variable, se maneja su error.
	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		//no hay error se agrega al objeto
		todos = append(todos, todo)
	}
	//retorna el objeto
	return c.JSON(todos)
}

// 12.
func createTodos(c *fiber.Ctx) error {
	//sera un pointer
	todo := new(Todo)

	//unir el cuerpo del request al struct
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	//se actualiza el valor del id del struct con el valor del insertResult
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

// 13.
func updateTodos(c *fiber.Ctx) error {
	//el id llega tipo string. Se convierte a objetoID.
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	//se actualiza mediante el filtro y una clausula de actualizacion,
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})

}

// 14.
func deleteTodos(c *fiber.Ctx) error {
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
