package task

import (
	"fmt"
	"todocli/entity"
)
type ServiceRepository interface {
	//DoesThisUserHaveThisCategoryID(userID, categoryID int) bool
	CreateNewTask(t entity.Task) (entity.Task,error)
	ListUserTask(userID int) ([]entity.Task ,error)
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
	Title      string
	DueDate    string
	CategoryID int
	AuthenticatedUserID     int
}

type CreateResponse struct{
	Task entity.Task
	MetaData string
}

func (s Service) Create(req CreateRequest) (CreateResponse,error){

	
/*	if !s.repository.DoesThisUserHaveThisCategoryID(
		req.AuthenticatedUserID,req.CategoryID){
		return CreateResponse{}, fmt.Errorf("user does not have this category: %d ",req.CategoryID)
	}
*/
	
	createdTask,cErr:=s.repository.CreateNewTask(entity.Task{
		Title:      req.Title,
		DueDate:    req.DueDate,
		CategoryID: req.CategoryID,
		IsDone:     false,
		UserID:     req.AuthenticatedUserID,
	})
	if cErr!=nil{
		return CreateResponse{},fmt.Errorf("can't create new task %v",cErr)
	}
	 
	return CreateResponse{Task: createdTask},nil
}

type ListRequest struct{
	UserID int
}

type ListResponse struct{
	Tasks []entity.Task
}

func (s Service) List(req ListRequest) (ListResponse,error){
	tasks,err := s.repository.ListUserTask(req.UserID)
	if err!=nil{
		return ListResponse{},fmt.Errorf("can't list user tasks: %v",err)
	}
	return ListResponse{Tasks: tasks},nil
}