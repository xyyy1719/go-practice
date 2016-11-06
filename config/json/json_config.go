package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//create a struct to holding the json configuration values
type configuration struct {
	Enable bool
	Path   string
}

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(config.Path)
}
