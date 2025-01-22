package handler

import (
	"golang-simple-api/internal/entity"
	"golang-simple-api/internal/model"
	"golang-simple-api/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SimpleHandler struct {
	SimpleUsecase usecase.SimpleUsecaseInterface
}

func NewSimpleHandler(simpleUsecase usecase.SimpleUsecaseInterface) *SimpleHandler {
	return &SimpleHandler{
		SimpleUsecase: simpleUsecase,
	}
}

func (s *SimpleHandler) GetAll(ctx *gin.Context) {
	persons, err := s.SimpleUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.BaseResponse[[]entity.Person]{
		Data: persons,
		Code: http.StatusOK,
	})
}

func (s *SimpleHandler) Create(ctx *gin.Context) {
	var request model.PersonRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.SimpleUsecase.Create(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Person created"})
}

func (s *SimpleHandler) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person, err := s.SimpleUsecase.GetByID(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.BaseResponse[entity.Person]{
		Code: http.StatusOK,
		Data: person,
	})
}
