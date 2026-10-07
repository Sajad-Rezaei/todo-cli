package category

import (
	"fmt"
	"todocli/entity"
)

type ServiceRepository interface{
	CreateNewCategory(cat entity.Category) (entity.Category,error)
}

type Service struct{
	repository ServiceRepository
}


func NewService(repo ServiceRepository) Service{
	return Service{
		repository: repo,
	}
}


type CreateRequest struct{
	Title  string
	Color  string
	AuthenticatedUserID int
}

type CreateResponse struct{
	Category entity.Category
	MetaData string
}


func (s Service) Create(req CreateRequest) (CreateResponse,error){
	createdCategory,cErr := s.repository.CreateNewCategory(entity.Category{
		Title: req.Title,
		Color: req.Color,
		UserID: req.AuthenticatedUserID,
	})
	if cErr!=nil{
		return CreateResponse{},fmt.Errorf("can't create new category %v",cErr)
	}
	return CreateResponse{Category: createdCategory},nil
}