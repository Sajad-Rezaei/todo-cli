package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"todocli/delivery/deliveryparam"
)

func main(){
	fmt.Println("command",os.Args[0])
	
	if len(os.Args)<2{
		log.Fatalln("you should set ip address of server")
	}
	serverAddress:= os.Args[1]
	message := "default message"
	if len(os.Args)>2 {
		message = os.Args[2]
	}

	connection,err :=net.Dial("tcp",serverAddress)
	if err!=nil {
		log.Fatalln("can't dial the given address",err)
	}
	defer connection.Close()

	fmt.Println("local address",connection.LocalAddr())

	req:= deliveryparam.Request{Command: message}
	switch req.Command{
	case "create-task":
		req.CreateTaskRequest = deliveryparam.CreateTaskRequest{
			Title: "test",
			DueDate: "test",
			CategoryID: 0,
		}
	case "create-category":
		req.CreateCategoryRequest= deliveryparam.CreateCategoryRequest{
			Title: "test",
			Color: "test",
		}	
	
	}	
	serializedData,mErr := json.Marshal(&req)
	if mErr!=nil{
		log.Fatalln("can't marshal request",mErr)
	
	}
	

	numberOfWrittenBytes,wErr := connection.Write(serializedData)
	if wErr!=nil{
		log.Fatalln("can't write data to connection",wErr)
		
	}
	fmt.Println("numberOfWrittenBytes",numberOfWrittenBytes)


	var data = make([]byte,1024)
	_, rErr := connection.Read(data)
	if rErr!=nil{
		log.Fatalln("can't read data from connection",rErr)

	}
	fmt.Println("server response : ",string(data))

}	