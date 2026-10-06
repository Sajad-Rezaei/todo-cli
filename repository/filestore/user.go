package filestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"todocli/constant"
	"todocli/entity"
	"strconv"
	"strings"
)



type FileStore struct {
	filePath string
	serializationMode string
}

//constructor
func New (path , serializationMode string) FileStore {
	return FileStore{filePath: path,serializationMode: serializationMode}
}

func (f FileStore) Save(u entity.User){
	f.writeUserToFile(u)
}


func (f FileStore) Load() []entity.User{
	var uStore []entity.User
	file, oErr := os.Open(f.filePath)
	if oErr != nil {
		fmt.Println("cant open the file", oErr)

		return nil
	}

	var data = make([]byte, 1024)
	_, rErr := file.Read(data)
	if rErr != nil {
		fmt.Println("can't read from the file", oErr)

		return nil
	}

	var dataStr = string(data)
	// len of "" is 0
	//charecters in the last of dataStr is not "" , it is like "                   " so
	//len of last charecters is not 0 , it is 936 . becuse lenghs of data is 1024 and we have 936 zero value
	// and this zero values are in index [2] of userSlice (by the code in line 283)
	// so we can't use if u=="" in for loop (line 285) to skip index[2] of userSlice
	// we should remove "\n\x00" in the end of your dataStr
	// in this way index[2] will never be created
	dataStr = strings.TrimRight(dataStr, "\n\x00")

	userSlice := strings.Split(dataStr, "\n")

	for _, u := range userSlice {

		//if u=="" {
		//	continue
		//}
		var userStruct = entity.User{}

		switch f.serializationMode{
		case constant.ManDarAvardiSerializationMode:
			var dErr error
			userStruct,dErr = deserilizeFromManDaravardi(u)
			if dErr !=nil{
				fmt.Println("can't deserialize user record to user struct",dErr)

				return nil
			}

		case constant.JsonSerializationMode:
			uErr:=json.Unmarshal([]byte(u), &userStruct)
			if uErr!=nil{
				fmt.Println("can't deserialize user record to user struct with json mode",uErr)

				return nil
			}
		}
		
		uStore=append(uStore, userStruct)

	}	
	return uStore
}



func(f FileStore) writeUserToFile(user entity.User) {
	var file *os.File

	file, err := os.OpenFile(f.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cant create or open file", err)

		return
	}

	defer file.Close()

	var data []byte
	//serialize the user struct/object
	switch f.serializationMode{
	case constant.ManDarAvardiSerializationMode:	
		strdata := fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", user.Id, user.Name, user.Email, user.Password)
		data= []byte(strdata)
	
	case constant.JsonSerializationMode:
		var jErr error
		data,jErr=json.Marshal(user)
		if jErr!=nil {
			fmt.Println("cat't marshal user struct to json",jErr)

			return
		}
		//  ... = spread
		// chon mikhahim yek slice az byte ro bedim behesh (na yek item)
		// bayad az  spread estefadeh konim
		data = append(data, []byte("\n")...)
	
	default:
		fmt.Println("ivalid serialization mode ")
		
		return
	}

	_, wErr := file.Write(data)
	if wErr != nil {
		fmt.Printf("can't write to the file %v\n", wErr)

		return
	}

}


func deserilizeFromManDaravardi (userStr string) (entity.User, error){
		
		if userStr==""{
			return entity.User{},errors.New("user strinng is empty")
		}

		var user = entity.User{}
		userFields := strings.Split(userStr, ",")

		for _, field := range userFields {
			values := strings.Split(field, ": ")
			if len(values) != 2 {
				fmt.Println("field is not valid, skipping ...", len(values))

				continue
			}
			fieldName := strings.ReplaceAll(values[0], " ", "")
			fieldValue := values[1]

			switch fieldName {
			case "id":
				id, err := strconv.Atoi(fieldValue)
				if err != nil {
					return entity.User{},errors.New("strconv error")
				}
				user.Id = id
			case "name":
				user.Name = fieldValue
			case "email":
				user.Email = fieldValue
			case "password":
				user.Password = fieldValue
			}
		}
	return user,nil
}	


