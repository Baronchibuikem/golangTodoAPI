package handler

import (
	"todoapp/database"
	"todoapp/model"
	"context"
    "fmt"
	"os"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)



func GetTodos(c *fiber.Ctx) error {
	todoCollection := database.Database.DB.Collection(os.Getenv("TODO_COLLECTION"))
	cursor, err := todoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer cursor.Close(context.Background())

	var todos []model.Todo
	if err := cursor.All(context.Background(), &todos); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo model.Todo
	todoCollection := database.Database.DB.Collection(os.Getenv("TODO_COLLECTION"))
	err := todoCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		return c.Status(404).SendString("Todo not found")
	}
	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	fmt.Println("Starting post method...")
	var todo model.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	todo.ID = ""
	todoCollection := database.Database.DB.Collection(os.Getenv("TODO_COLLECTION"))
	result, err := todoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Convert ObjectID to string
	todo.ID = result.InsertedID.(primitive.ObjectID).Hex()
    fmt.Println("Created")
	return c.JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	var updatedTodo model.Todo
	if err := c.BodyParser(&updatedTodo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	todoCollection := database.Database.DB.Collection(os.Getenv("TODO_COLLECTION"))
	_, err := todoCollection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": updatedTodo})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	updatedTodo.ID = id
	return c.JSON(updatedTodo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todoCollection := database.Database.DB.Collection(os.Getenv("TODO_COLLECTION"))
	_, err := todoCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
