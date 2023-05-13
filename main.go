/*
Copyright © 2023 tochiman developer@tochiman.com. All rights reserved.
*/
package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tochiman/DriveManegement/cmd"
)

func main() {
	LoadEnv()
	cmd.Execute()
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}
