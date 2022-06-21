package main

import (
	"fmt" 
	myroute "github.com/bhabermann/imersaofc8-simulator/app/route"
)

func main() {
	route := myroute.Route{
		ID: "1",
		ClientID: "1",
	}
	route.LoadPosition()
	json, _ := route.ExportJsonPositions()
	fmt.Println(json[1])
}