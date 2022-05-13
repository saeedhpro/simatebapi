package main

import (
	"github.com/saeedhpro/apisimateb/controller"
	"github.com/saeedhpro/apisimateb/repository"
)

func main() {
	repository.Init()
	http.Run("8000")
}
