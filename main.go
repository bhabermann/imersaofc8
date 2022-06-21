package main

import (
	"fmt" 
	myroute "github.com/bhabermann/imersaofc8-simulator/app/route"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	route := myroute.Route{
		ID: "1",
		ClientID: "1",
	}
	route.LoadPosition()
	json, _ := route.ExportJsonPositions()
	fmt.Println(json[1])
}