package api

import (
	"e-wallet-transaction/constants"
	"e-wallet-transaction/helpers"
	"e-wallet-transaction/internal/interfaces"
	"e-wallet-transaction/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionHandler struct {
	TransactionService interfaces.ITransactionService
}

func (api *TransactionHandler) CreateTransaction(c *gin.Context) {
	var (
		log = helpers.Logger
		req = &models.Transaction{}
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to bind json", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to bind json",
			nil,
		)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to validate request",
			nil,
		)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token",
			nil,
		)
		return
	}

	tokenData, ok := token.(*models.TokenData)
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token data",
			nil,
		)
		return
	}

	if !constants.MapTransaction[req.TransactionType] {
		log.Error("invalid transaction type")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"invalid transaction type",
			nil,
		)
		return
	}

	req.UserID = int(tokenData.UserID)

	resp, err := api.TransactionService.CreateTransaction(c.Request.Context(), req)
	if err != nil {
		log.Error("failed to create transaction", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to create transaction",
			nil,
		)
		return
	}

	fmt.Println(resp)

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"success",
		resp,
	)
}

func (api *TransactionHandler) UpdateTransactionStatus(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.UpdateStatusTransaction
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to bind json", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to bind json",
			nil,
		)
		return
	}

	req.Reference = c.Param("reference")

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to validate request",
			nil,
		)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token",
			nil,
		)
		return
	}

	tokenData, ok := token.(*models.TokenData)
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token data",
			nil,
		)
		return
	}

	err := api.TransactionService.UpdateStatusTransaction(c.Request.Context(), tokenData, &req)
	if err != nil {
		log.Error("failed to update transaction status", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to update transaction status",
			nil,
		)
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"success",
		nil,
	)
}

func (api *TransactionHandler) GetTransaction(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token",
			nil,
		)
		return
	}

	tokenData, ok := token.(*models.TokenData)
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token data",
			nil,
		)
		return
	}

	resp, err := api.TransactionService.GetTransaction(c.Request.Context(), int(tokenData.UserID))
	if err != nil {
		log.Error("failed to update transaction status", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to update transaction status",
			nil,
		)
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"success",
		resp,
	)
}

func (api *TransactionHandler) GetTransactionDetail(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	reference := c.Param("reference")
	if reference == "" {
		log.Error("reference is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"reference is required",
			nil,
		)
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token",
			nil,
		)
		return
	}

	_, ok = token.(*models.TokenData)
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token data",
			nil,
		)
		return
	}

	resp, err := api.TransactionService.GetTransactionDetail(c.Request.Context(), reference)
	if err != nil {
		log.Error("failed to update transaction status", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to update transaction status",
			nil,
		)
		return
	}

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"success",
		resp,
	)
}

func (api *TransactionHandler) RefundTransaction(c *gin.Context) {
	var (
		log = helpers.Logger
		req = &models.RefundTransaction{}
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to bind json", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to bind json",
			nil,
		)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to validate request",
			nil,
		)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token",
			nil,
		)
		return
	}

	tokenData, ok := token.(*models.TokenData)
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest,
			false,
			"failed to get token data",
			nil,
		)
		return
	}

	resp, err := api.TransactionService.RefundTransaction(c.Request.Context(), tokenData, req)
	if err != nil {
		log.Error("failed to create transaction", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError,
			false,
			"failed to refund transaction",
			nil,
		)
		return
	}

	fmt.Println(resp)

	helpers.SendResponseHTTP(
		c,
		http.StatusOK,
		true,
		"success",
		resp,
	)
}
