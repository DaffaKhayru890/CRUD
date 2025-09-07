package controllers

import (
	"github.com/DaffaKhayru890/src/configs"
	"github.com/DaffaKhayru890/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetContacts(ctx *fiber.Ctx) error {
	var Contacts []models.Contacts

	if err := configs.DB().Find(&Contacts).Error; err != nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Can not get contacts"})
	}

	return ctx.Status(fiber.StatusOK).JSON(Contacts)
}

func GetContact(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var Contact models.Contacts

	if err := configs.DB().Where("id = ?", id).First(&Contact).Error; err != nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Failed to get contact"})
	}

	return ctx.Status(fiber.StatusOK).JSON(Contact)
}

func PostContact(ctx *fiber.Ctx) error {
	var bodyReq models.Contacts

	if err := ctx.BodyParser(&bodyReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body request"})
	}

	newContact := models.Contacts{
		FirstName: bodyReq.FirstName,
		LastName: bodyReq.LastName,
		Email: bodyReq.Email,
		Phone: bodyReq.Phone,
	}

	if err := configs.DB().Create(&newContact).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to post contact"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Post contact successfully"})
}

func UpdateContacts(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "hello"})
}

func DeleteContacts(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := configs.DB().Where("id = ?", id).Delete(&models.Contacts{}).Error; err != nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"error": "Failed to delete contacts"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Delete contact successfully"})
}