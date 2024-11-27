package generator

import (
	"math/rand"
)

type Device interface {
	GetId() string
	GetName() string
	GenerateData()
}

// BaseDevice provides common fields and methods
type BaseDevice struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (d *BaseDevice) GetId() string {
	return d.ID
}

func (d *BaseDevice) GetName() string {
	return d.Name
}

// SmartWatch device implementation
type SmartWatch struct {
	BaseDevice
	Info struct {
		BatteryLevel int     `json:"batteryLevel"`
		HeartRate    int     `json:"heartRate"`
		SleepTime    int     `json:"sleepTime"`
		Steps        int     `json:"steps"`
		Calories     int     `json:"calories"`
		Temperature  float32 `json:"temperature"`
	} `json:"info"`
}

func (s *SmartWatch) GenerateData() {
	s.Info.BatteryLevel = genRandomInt(0, 100)
	s.Info.HeartRate = genRandomInt(60, 100)
	s.Info.SleepTime = genRandomInt(0, 8)
	s.Info.Steps = genRandomInt(0, 10000)
	s.Info.Calories = genRandomInt(0, 500)
	s.Info.Temperature = genRandomFloat32(36.0, 37.5)
}

// YogaPants device implementation
type YogaPants struct {
	BaseDevice
	Info struct {
		BatteryLevel int     `json:"batteryLevel"`
		Temperature  float32 `json:"temperature"`
		Humidity     int     `json:"humidity"`
		Movement     int     `json:"movement"`
	} `json:"info"`
}

func (y *YogaPants) GenerateData() {
	y.Info.BatteryLevel = genRandomInt(0, 100)
	y.Info.Temperature = genRandomFloat32(20.0, 30.0)
	y.Info.Humidity = genRandomInt(40, 60)
	y.Info.Movement = genRandomInt(0, 100)
}

// SmartBelt device implementation
type SmartBelt struct {
	BaseDevice
	Info struct {
		BatteryLevel  int `json:"batteryLevel"`
		HeartRate     int `json:"heartRate"`
		MuscleTension int `json:"muscleTension"`
		Stress        int `json:"stress"`
	} `json:"info"`
}

func (s *SmartBelt) GenerateData() {
	s.Info.BatteryLevel = genRandomInt(0, 100)
	s.Info.HeartRate = genRandomInt(60, 100)
	s.Info.MuscleTension = genRandomInt(0, 100)
	s.Info.Stress = genRandomInt(0, 100)
}

// Helper functions for generating random data
func genRandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func genRandomFloat32(min, max float64) float32 {
	return float32(min + rand.Float64()*(max-min))
}
