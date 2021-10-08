package controllers

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/RohitKuwar/go_fiber/models"
	"github.com/gofiber/fiber/v2"

	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var goals = []models.Goal{}

var firestoreCredentialsLocation = "../serviceAccountKey.json"

type goalsModel struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func GetGoals(c *fiber.Ctx) error {

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	var newGoals []models.Goal
	iter := client.Collection("goals").Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		fmt.Print(doc.Data())

		var tempGoals models.Goal
		if err := doc.DataTo(&tempGoals); err != nil {
			break
		}
		newGoals = append(newGoals, tempGoals)
	}

	return c.Status(fiber.StatusOK).JSON(newGoals)
}

func GetGoal(c *fiber.Ctx) error {
	// get parameter value
	paramID := c.Params("id")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	dsnap, err := client.Collection("goals").Doc(paramID).Get(ctx)
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Goal not found",
		})
	}
	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)

	if m == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Goal not found",
		})
	}

	return c.Status(fiber.StatusNotFound).JSON(m)

}

func CreateGoal(c *fiber.Ctx) error {
	type Request struct {
		Id     string `json:"id"`
		Title  string `json:"title"`
		Status string `json:"status"`
	}

	var body Request

	err := c.BodyParser(&body)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// create a goal variable
	goal := &models.Goal{
		Title:  body.Title,
		Status: body.Status,
	}

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	ref := client.Collection("goals").NewDoc()

	goal.Id = ref.ID

	_, err = ref.Set(ctx, map[string]interface{}{
		"id":     goal.Id,
		"title":  goal.Title,
		"status": goal.Status,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "data added successfully",
		"goal":    goal,
	})
}

func UpdateGoal(c *fiber.Ctx) error {
	type request struct {
		Title  string `json:"title"`
		Status string `json:"status"`
	}
	var body request

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	paramID := c.Params("id")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	fmt.Print(paramID, body.Status, body.Title)

	_, err = client.Collection("goals").Doc(paramID).Set(ctx, map[string]interface{}{
		"title":  body.Title,
		"status": body.Status,
	}, firestore.MergeAll)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data updated successfully",
	})

	// return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
}

func DeleteGoal(c *fiber.Ctx) error {
	// get param
	paramID := c.Params("id")

	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile(firestoreCredentialsLocation)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	dsnap, err := client.Collection("goals").Doc(paramID).Get(ctx)
	if err != nil {
		fmt.Print(err)
		// if goal not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Goal with id " + paramID + " Not found",
		})
	}

	// Test Print line
	m := dsnap.Data()
	fmt.Printf("Document data: %#v\n", m)

	_, err = client.Collection("goals").Doc(paramID).Delete(ctx)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Goal with id" + paramID + "Deleted Successfully",
	})
}
