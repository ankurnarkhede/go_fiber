package controllers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	// "github.com/RohitKuwar/go_fiber/models"
)

type Goal struct {
  Id        int    `json:"id"`
  Title     string `json:"title"`
  Status    bool   `json:"status"`
}

var goals = []*Goal{
	{
		Id:        1,
		Title:     "Read about Promises",
		Status:    true,
	},
	{
		Id:        2,
		Title:     "Read about Closures",
		Status:    false,
	},
}

func GetGoals(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"goals": goals,
		},
	})
}

func GetGoal(c *fiber.Ctx) error {
	// get parameter value
	paramID := c.Params("id")

	// convert parameter value string to int
	id, err := strconv.Atoi(paramID)

	// if error in parsing string to int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse Id",
			"error":   err,
		})
	}

	// find goal and return
	for _, goal := range goals {
		if goal.Id == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"goal": goal,
				},
			})
		}
	}

	// if goal not available
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Goal not found",
	})
}

func CreateGoal(c *fiber.Ctx) error {
	type Request struct {
		Title string `json:"title"`
	}

	var body Request

	err := c.BodyParser(&body)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// create a goal variable
	goal := &Goal{
		Id:        len(goals) + 1,
		Title:     body.Title,
		Status: false,
	}

	// append in goal
	goals = append(goals, goal)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"goal": goal,
		},
	})
}


func DeleteGoal(c *fiber.Ctx) error {
	// get param
	paramID := c.Params("id")

	// convert param string to int
	id, err := strconv.Atoi(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
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
		"success": false,
		"message": "Goal not found",
	})
}

func UpdateGoal(c *fiber.Ctx) error {
	// find parameter
	paramID := c.Params("id")

	// convert parameter string to int
	id, err := strconv.Atoi(paramID)

	// if parameter cannot parse
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	// request structure
	type Request struct {
		Title     *string `json:"title"`
		Status *bool   `json:"status"`
	}

	var body Request
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	var goal *Goal

	for _, t := range goals {
		if t.Id == id {
			goal = t
			break
		}
	}

	if goal.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not found",
		})
	}

	if body.Title != nil {
		goal.Title = *body.Title
	}

	if body.Status != nil {
		goal.Status = *body.Status
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"goal": goal,
		},
	})
}
