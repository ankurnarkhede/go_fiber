package controllers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/RohitKuwar/go_fiber/models"
)

var goals = []models.Goal{
	{
		Id:        1,
		Title:     "Read about Promises",
		Status:    "completed",
	},
	{
		Id:        2,
		Title:     "Read about Closures",
		Status:    "active",
	},
}

func GetGoals(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(goals)
}

func GetGoal(c *fiber.Ctx) error {
	// get parameter value
	paramID := c.Params("id")

	// convert parameter value string to int
	id, err := strconv.Atoi(paramID)

	// if error in parsing string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse Id",
			"error":   err,
		})
	}

	// find goal and return
	for _, goal := range goals {
		if goal.Id == id {
			return c.Status(fiber.StatusOK).JSON(goal)
		}
	}

	// if goal not available
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Goal not found",
	})
}

func CreateGoal(c *fiber.Ctx) error {
	type Request struct {
		Title 	string 	`json:"title"`
		Status  string  `json:"status"`
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
		Id:        len(goals) + 1,
		Title:     body.Title,
		Status: 	 body.Status,
	}

	// append in goal
	goals = append(goals, *goal)

	return c.Status(fiber.StatusCreated).JSON(goal)
}

func DeleteGoal(c *fiber.Ctx) error {
	// get param
	paramID := c.Params("id")

	// convert param string to int
	id, err := strconv.Atoi(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	// find and delete goal
	for i, goal := range goals {
		if goal.Id == id {
			goals = append(goals[:i], goals[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	// if goal not found
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Goal not found",
	})
}

func UpdateGoal(c *fiber.Ctx) error {
	type request struct {
		Title      string `json:"title"`
		Status string   `json:"status"`
	}
	var body request

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid id.",
		})
	}

	for i, goal := range goals {
		if goal.Id == id {
			goals[i] = models.Goal{
				Id:       id,
				Title:    body.Title,
				Status: 	body.Status,
			}
			return c.Status(fiber.StatusOK).JSON(goals[i])
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Record not found"})
}

