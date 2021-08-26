package handlers

import (
	"github.com/ajiepangestu/go-rest-api-example/database"
	"github.com/ajiepangestu/go-rest-api-example/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	insert, err := database.DB.Query("INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?)",
		c.Params("first_name"), c.Params("last_name"), c.Params("email"),
	)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	defer insert.Close()
	return err
}

func GetUserList(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT first_name, last_name, email FROM users")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	defer rows.Close()
	result := models.Users{}
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.FirstName, &user.LastName, &user.Email); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result.Users = append(result.Users, user)
	}
	if err := c.JSON(&fiber.Map{
		"success": true,
		"data":    result,
		"message": "All user data returned successfully",
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return err
}

func GetUser(c *fiber.Ctx) error {
	row, err := database.DB.Query("SELECT first_name, last_name, email FROM users WHERE id = ?", c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	defer row.Close()
	user := models.User{}
	for row.Next() {
		if err := row.Scan(&user.FirstName, &user.LastName, &user.Email); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
	}
	if c.JSON(&fiber.Map{
		"success": true,
		"data":    user,
		"message": "User data returned successfully",
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return err
}

func UpdateUser(c *fiber.Ctx) error {
	updateRow, err := database.DB.Query("UPDATE users SET first_name=?, last_name=?, email=?  WHERE id = ?",
		c.Params("first_name"), c.Params("last_name"), c.Params("email"), c.Params("id"),
	)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	defer updateRow.Close()
	return err
}

func DeleteUser(c *fiber.Ctx) error {
	deleteRow, err := database.DB.Query("DELETE FROM users WHERE email = ?", c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	defer deleteRow.Close()
	return err
}
