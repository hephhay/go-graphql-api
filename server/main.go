package main

import (
	"github.com/hephhay/go-graphql/app/container"
)

func main() {
	err := container.initApp()

	if err != nil {
		panic(err)
	}
}
