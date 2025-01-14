package main

import (
	"embed"

	"github.com/gin-contrib/static"
	"github.com/rubenhoenle/ci-cd-lecture-project-template/api"
	"github.com/rubenhoenle/ci-cd-lecture-project-template/persistence/inmemory"
)

//go:embed frontendDist/*
var frontendDir embed.FS

func main() {
	var persistence api.Persistence = inmemory.NewInMemoryPersistence()
	router := api.NewRouter(persistence)

	// serve the frontend
	router.Use(static.Serve("", static.EmbedFolder(frontendDir, "frontendDist")))

	router.Run(":8080")
}
