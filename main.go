package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// 2. Creando el struct del TODO
type Todo struct {
	ID        int    `json: id`
	Completed bool   `json: "completed"`
	Body      string `json: "body"`
}

// 1
func main() {
	fmt.Println("Hola mundo")
	//Creando nueva instancia de fiber
	app := fiber.New()

	//5.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	PORT := os.Getenv("PORT")

	//3. Creando el array del todo
	todos := []Todo{}

	//agregando la primera ruta
	//La funcion handler que recibe un contexto 'c' de fiber
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		// devuelve una respuesta http y un objeto json
		return c.Status(200).JSON(todos)
	})

	// esta ruta maneja la creacion de las nuevas tareas
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
		// parsea el cuerpo de la solicitud http a la estructura 'Todo'. Si hay un error, lo devuelve.
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		// validacion del body de la respuesta
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "El cuerpo de la tarea es requerido"})
		}

		// asigna un id unico al nuevo todo
		todo.ID = len(todos) + 1
		// agrega el nuevo elemento al array todos
		todos = append(todos, *todo)

		// devuelve el nuevo todo
		return c.Status(201).JSON(todo)
	})

	// actualizando un todo
	// patch se usa para modificaciones parciales.
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		// se extrae el parametro id de la url
		id := c.Params("id")

		for i, todo := range todos {
			// se compara el id del parametro de la url con el id de la tarea actual
			if fmt.Sprint(todo.ID) == id {
				// se actualiza el estado completed a true
				todos[i].Completed = true
				// devuelve la tarea actualizada
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Tarea no encontrada"})
	})

	// eliminar una tarea
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				// crea una nueva lista que ignora el elemento actual (borra)
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Tarea no encontrada"})
	})
	log.Fatal(app.Listen(":" + PORT))
}
