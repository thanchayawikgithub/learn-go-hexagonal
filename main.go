package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanchayawikgithub/learn-go-hexagonal/adapters"
	"github.com/thanchayawikgithub/learn-go-hexagonal/core"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&core.Order{})

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/orders", orderHandler.CreateOrder)

	app.Listen(":8000")
}
