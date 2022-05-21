package controllers

import (
	"BlogAPI/config"
	"BlogAPI/models"
	"github.com/gofiber/fiber/v2"
)

func GetBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var blog models.Blogs
	result := config.DB.Find(&blog, id)
	if result.RowsAffected == 0 {
		return ctx.SendStatus(404)
	}
	return ctx.Status(fiber.StatusOK).JSON(&blog)
}

func CreateBlog(ctx *fiber.Ctx) error {
	blog := new(models.Blogs)
	if err := ctx.BodyParser(blog); err != nil {
		return ctx.Status(503).JSON(err)
	}
	config.DB.Create(&blog)
	return ctx.Status(fiber.StatusOK).JSON(blog)
}

func DeleteBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var blog models.Blogs
	search := config.DB.Delete(&blog, id)
	if search.RowsAffected == 0 {
		return ctx.Status(404).JSON(models.Message{
			Status:      "Error",
			Description: "404: Not Found",
		})
	}
	return ctx.Status(200).JSON(models.Message{
		Status:      "Success",
		Description: "The post was successfully deleted",
	})
}

func UpdateBlog(ctx *fiber.Ctx) error {
	blog := new(models.Blogs)
	id := ctx.Params("id")
	if err := ctx.BodyParser(blog); err != nil {
		return ctx.Status(503).JSON(models.Message{
			Status:      "Error",
			Description: "There was an error parsing your request",
		})
	}
	config.DB.Where("id = ?", id).Updates(&blog)
	return ctx.Status(200).JSON(blog)
}
