package controllers

import "github.com/gofiber/fiber/v2"

func Signup(ctx *fiber.Ctx) error {
	

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "hello"})
}

func Login(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "hello"})
}