package cmd

import (
	"e-wallet-transaction/external"
	"e-wallet-transaction/helpers"
	"e-wallet-transaction/internal/api"
	"e-wallet-transaction/internal/interfaces"
	"e-wallet-transaction/internal/repository"
	"e-wallet-transaction/internal/services"
	"github.com/gin-gonic/gin"
	"log"
)

func ServeHTTP() {
	d := dependencyInject()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal("Failed to set trusted proxies", err)
	}

	transactionV1 := r.Group("/transaction/v1")
	transactionV1.POST("/create", d.MiddlewareValidateToken, d.CreateTransactionAPI.CreateTransaction)
	transactionV1.PUT("/update-status/:reference", d.MiddlewareValidateToken, d.CreateTransactionAPI.UpdateTransactionStatus)
	transactionV1.GET("/", d.MiddlewareValidateToken, d.CreateTransactionAPI.GetTransaction)
	transactionV1.GET("/:reference", d.MiddlewareValidateToken, d.CreateTransactionAPI.GetTransactionDetail)
	transactionV1.POST("/refund", d.MiddlewareValidateToken, d.CreateTransactionAPI.RefundTransaction)

	err = r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
	log.Println("Server started")
}

type Dependency struct {
	External             interfaces.IExternal
	CreateTransactionAPI interfaces.ITransactionHandler
}

func dependencyInject() *Dependency {
	ext := &external.External{}
	trxRepo := &repository.TransactionRepository{DB: helpers.DB}

	createTrxSvc := &services.TransactionService{
		TransactionRepo: trxRepo,
		External:        ext,
	}
	createTrxApi := &api.TransactionHandler{TransactionService: createTrxSvc}

	return &Dependency{
		External:             ext,
		CreateTransactionAPI: createTrxApi,
	}
}
