package main

import (
	"iot-mocker/generator"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var Devices []generator.Device

func main() {
	app := fiber.New()

	// Initialize devices
	Devices = initializeDevices()
	generateRandomData()
	go updateDevices()

	app.Get("/devices", getDevices)
	app.Get("/devices/:id", getDeviceData)

	log.Println("Server is running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initializeDevices() []generator.Device {
	return []generator.Device{
		&generator.SmartBelt{BaseDevice: generator.BaseDevice{ID: genUUID(), Name: "Fitbit A112 SmartWatch"}},
		&generator.SmartWatch{BaseDevice: generator.BaseDevice{ID: genUUID(), Name: "Nadi X Yoga Pants"}},
		&generator.YogaPants{BaseDevice: generator.BaseDevice{ID: genUUID(), Name: "Fitbit SmartBelt"}},
	}
}

func updateDevices() {
	for {
		for _, device := range Devices {
			device.GenerateData()
		}

		time.Sleep(3 * time.Second)
	}
}

// Generate random data dynamically
func generateRandomData() {
	for _, device := range Devices {
		device.GenerateData()
	}
}

func getDevices(c *fiber.Ctx) error {
	devices := make([]map[string]string, len(Devices))

	for i, device := range Devices {
		devices[i] = map[string]string{
			"id":   device.GetId(),
			"name": device.GetName(),
		}
	}

	return c.Status(fiber.StatusOK).JSON(devices)
}

func getDeviceData(c *fiber.Ctx) error {
	// Fetch the device ID
	deviceId := c.Params("id")

	// Check if the ID exists
	for _, device := range Devices {
		if device.GetId() == deviceId {
			return c.Status(fiber.StatusOK).JSON(device)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Device not found",
	})
}

func genUUID() string {
	return uuid.New().String()
}
