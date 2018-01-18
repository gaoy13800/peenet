package storage

type userHandler struct {

}

var userList map[int]*userDetail

var incrementId int = 0


type userDetail struct {

	Name string

	CreateTime string

	UserName string

	PassWord string
}

//func SaveUser(data string){
//
//	csvFile, err := os.Create("memory.csv")
//	if err != nil {
//		panic(err)
//	}
//
//	defer csvFile.Close()
//
//	writer := csv.NewWriter(csvFile)
//
//
//
//	writer.Flush()
//
//}


//创建一个用户

func (user *userHandler) CreateUserId(name, userName, password string){

	userId := incrementId + 1

	detail := &userDetail{
		Name:name,
		UserName:userName,
		PassWord:password,
	}

	userList[userId] = detail
}

func (user *userHandler) ValidateUser(userName, password string) bool{

	for _, v := range userList{

		if v.Name == userName {
			return v.PassWord == password
		}
	}

	return false
}

func (user *userHandler) GetUserCount() int{
	return len(userList)
}


func (user *userHandler) GetUserDetail(userId int)map[string]string{

	detail := userList[userId]

	result := make(map[string]string)

	result["Name"] = detail.Name
	result["Password"] = detail.PassWord

	return result
}

func NewUserHandler() *userHandler{


	return &userHandler{}

}

