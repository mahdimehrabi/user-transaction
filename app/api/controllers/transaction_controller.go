package controller

import (
	"bbdk/app/api/dto"
	"bbdk/app/api/response"
	transactionRepo "bbdk/domain/repository/transaction"
	"bbdk/domain/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransactionController struct {
	service service.TransactionService
}

func NewTransactionController(service service.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (tc *TransactionController) CreateTransaction(c *gin.Context) {
	var req dto.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Response(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	transaction := req.ToEntity()
	if err := tc.service.CreateTransaction(transaction); err != nil {
		if errors.Is(err, transactionRepo.ErrAlreadyExist) {
			response.Response(c, nil, http.StatusBadRequest, err.Error())
			return
		}
		response.InternalServerError(c)
		return
	}

	response.Response(c, nil, http.StatusCreated, "")
}

func (tc *TransactionController) GetTransactionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := tc.service.GetTransactionByID(uint(id))
	if err != nil {
		if errors.Is(err, transactionRepo.ErrNotFound) {
			response.NotFound(c)
		} else {
			response.InternalServerError(c)
		}
		return
	}
	transactionResponse := &dto.TransactionResponse{}
	transactionResponse.FromEntity(transaction)
	response.Response(c, transactionResponse, http.StatusOK, "")
}

func (tc *TransactionController) UpdateTransaction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Response(c, nil, http.StatusBadRequest, "invalid id")
		return
	}

	var req dto.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Response(c, nil, http.StatusBadRequest, err.Error())
		return
	}
	transaction := req.ToEntity()
	transaction.ID = uint(id)

	if err := tc.service.UpdateTransaction(transaction); err != nil {
		if errors.Is(err, transactionRepo.ErrNotFound) {
			response.NotFound(c)
		} else if errors.Is(err, transactionRepo.ErrAlreadyExist) {
			response.Response(c, nil, http.StatusBadRequest, err.Error())
			return
		} else {
			response.InternalServerError(c)
		}
		return
	}

	response.Response(c, nil, http.StatusOK, "")
}

func (tc *TransactionController) DeleteTransaction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Response(c, nil, http.StatusBadRequest, "invalid id")
		return
	}

	if err := tc.service.DeleteTransaction(uint(id)); err != nil {
		if errors.Is(err, transactionRepo.ErrNotFound) {
			response.NotFound(c)
		} else {
			response.InternalServerError(c)
		}
		return
	}

	response.Response(c, nil, http.StatusOK, "")
}

func (tc *TransactionController) GetAllTransactions(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	transactions, err := tc.service.GetAllTransactions(page, pageSize)
	if err != nil {
		response.InternalServerError(c)
		return
	}

	lnTransactions := len(transactions)
	responses := make([]dto.TransactionResponse, lnTransactions)
	for i := 0; i < lnTransactions; i++ {
		transactionResponse := dto.TransactionResponse{}
		transactionResponse.FromEntity(transactions[i])
		responses[i] = transactionResponse
	}

	response.Response(c, responses, http.StatusOK, "")
}
