package usecase

import (
	"golang-simple-api/internal/entity"
	"golang-simple-api/internal/model"
	"golang-simple-api/internal/repository"
)

type SimpleUsecase struct {
	SimpleRepository repository.SimpleRepositoryInterface
}

func NewSimpleUsecase(simpleRepository repository.SimpleRepositoryInterface) *SimpleUsecase {
	return &SimpleUsecase{
		SimpleRepository: simpleRepository,
	}
}

type SimpleUsecaseInterface interface {
	GetAll() ([]entity.Person, error)
	Create(request model.PersonRequest) error
	GetByID(id int) (entity.Person, error)
}

func (s *SimpleUsecase) GetAll() ([]entity.Person, error) {
	return s.SimpleRepository.GetAll()
}

func (s *SimpleUsecase) Create(request model.PersonRequest) error {
	return s.SimpleRepository.Create(model.ToEntity(request))
}

func (s *SimpleUsecase) GetByID(id int) (entity.Person, error) {
	return s.SimpleRepository.GetByID(id)
}
