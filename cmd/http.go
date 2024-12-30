package cmd

import (
	"e-wallet-transaction/external"
	"e-wallet-transaction/helpers"
	"e-wallet-transaction/internal/interfaces"
	"github.com/gin-gonic/gin"
	"log"
)

func ServeHTTP() {
	_ = dependencyInject()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("Failed to set trusted proxies", err)
	}

	r.Group("/transaction/v1")

	err = r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Server started")
}

type Dependency struct {
	External interfaces.IExternal
}

func dependencyInject() *Dependency {
	ext := &external.External{}

	return &Dependency{External: ext}
}
