package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"todocli/delivery/deliveryparam"
	"todocli/repository/memorystore"
	"todocli/service/category"
	"todocli/service/task"
)


func main() {

	const (
		network = "tcp"
		address = "127.0.0.1:9986"
	)

	//create new listener
	listner, err := net.Listen(network, address)

	if err != nil {
		log.Fatalln("can't listen on given address", address, err)

	}
	fmt.Println("listening on : ", listner.Addr())

	taskMemoryRepo:= memorystore.NewTaskStore()
	taskService := task.NewService(taskMemoryRepo)

	categoryMemoryRepo:= memorystore.NewCategoryStore()
	catgoryService := category.NewService(categoryMemoryRepo)

	for {
		//listen for new connection
		connection, aErr := listner.Accept()
		if aErr != nil {
			log.Println("can't listen to new connection", aErr)
		}

		//process request
		var rawRequest = make([]byte,1024)
		numberOfReadBytes, rErr := connection.Read(rawRequest)
		if rErr!=nil{
			log.Println("can't read data from connection",rErr)

			continue
		}
		
		fmt.Printf("client addres: %s, numberOfReadBytes: %d, data: %s\n",
		connection.RemoteAddr(),numberOfReadBytes,string(rawRequest))

		req := &deliveryparam.Request{}
		if uErr:= json.Unmarshal(rawRequest[:numberOfReadBytes],req); uErr!=nil{
			log.Println("bad request",uErr)

			continue
		}

		switch req.Command{
		case "create-task":
			response,cErr:= taskService.Create(task.CreateRequest{
				Title: req.CreateTaskRequest.Title,
				DueDate: req.CreateTaskRequest.DueDate,
				CategoryID: req.CreateTaskRequest.CategoryID,
				AuthenticatedUserID: 0,
			})
			if cErr!=nil{
				_,wErr:=connection.Write([]byte(cErr.Error()))
				if wErr!=nil{
					log.Println("can't write data to connection",wErr)

					continue
				}
			}

			data,mErr:=json.Marshal(response)
			if mErr!=nil{
				_,wErr:=connection.Write([]byte(mErr.Error()))
				if wErr!=nil{
					log.Println("can't marshal response",wErr)

					continue
				}
				continue
			}

			_,wErr:=connection.Write(data)
			if wErr!=nil{
				log.Println("can't write data to connection",wErr)

				continue
			}
		case "create-category":
			response,cErr := catgoryService.Create(category.CreateRequest{
				Title: req.CreateCategoryRequest.Title,
				Color: req.CreateCategoryRequest.Color,
				AuthenticatedUserID: 0,
			})
			
			if cErr!=nil{
				_,wErr:=connection.Write([]byte(cErr.Error()))
				if wErr!=nil{
					log.Println("can't write data to connection",wErr)

					continue
				}
			}

			data,mErr:=json.Marshal(response)
			if mErr!=nil{
				_,wErr:=connection.Write([]byte(mErr.Error()))
				if wErr!=nil{
					log.Println("can't marshal response",wErr)

					continue
				}
				continue
			}

			_,wErr:=connection.Write(data)
			if wErr!=nil{
				log.Println("can't write data to connection",wErr)

				continue
			}


		}

		connection.Close()
	}

}
