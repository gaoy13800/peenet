package http

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"strings"
	"io"
	"encoding/base64"
	"peenet/common"
	"log"
	"peenet/memery"
	"time"
	"encoding/json"
)


/**
	http 接口
 */

const (

	User_Increment_id = "user:id"

	User_Detail = "user:detail:"

	Device_Detail = "device:detail:"

)


func Login			(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")

	if auth == "" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	auths := strings.SplitN(auth, " ", 2)

	if len(auths) != 2 {
		fmt.Println("error")
		return
	}
	authMethod := auths[0]
	authB64 := auths[1]

	if authMethod != "Basic"{
		io.WriteString(w, "Auth type invalid")
		return
	}

	authStr, err := base64.StdEncoding.DecodeString(authB64)

	if err != nil {
		io.WriteString(w, "Unauthorized!\n")
		return
	}

	userInfo := strings.SplitN(string(authStr), ":", 2)

	username := userInfo[0]
	password := userInfo[1]

	if password != common.LOGIN_COMMON_PASSWD {
		log.Printf("用户 %s 登录失败！ 密码错误", username)

		io.WriteString(w, "Unauthorized!\n")

		return
	}

	userMemory := memery.SelectMemory(common.Common_Library) //选取user存储库

	var uid string

	userId, _ := userMemory.Get(User_Increment_id) //获取当前用户要绑定的userId

	if userId == nil {
		uid = "1"
		userMemory.Set(User_Increment_id, "1", common.LongTimeMemory)
	}else {
		uid = userId.(string)
	}

	userMemory.Set(User_Detail + uid  , username + "|" + time.Now().String(), common.LongTimeMemory) //存储详情

	userMemory.Increment(User_Increment_id, 1) //新增下一个userId

	fmt.Printf("username: %s userpwd:%s\n",username, password)

	io.WriteString(w, "hello, world!\n")

	fmt.Fprint(w, "Welcome!" + username + "\n")
}

func Logout			(w http.ResponseWriter, r *http.Request){
	userId := r.FormValue("userId")


	fmt.Println(userId)
	//todo remove all current user data
}

func SaveDeviceData (w http.ResponseWriter, r *http.Request) {

	protoType := r.FormValue("protoType")
	deviceId  := r.FormValue("deviceId")
	userId    := r.FormValue("userId")

	resultMap := make(map[string]string)

	if deviceId == "" || userId == "" {
		resultMap["result"] = "error"
		resultMap["msg"] = "invalid param"
		result, _ := json.Marshal(resultMap)

		fmt.Fprint(w, result)
		return
	}

	userMemory := memery.SelectMemory(common.Common_Library)

	userMemory.Set(Device_Detail + deviceId, protoType+"|"+userId, common.LongTimeMemory)

	resultMap["result"] = "ok"
	resultMap["msg"] = "success"
	result, _ := json.Marshal(&resultMap)

	fmt.Fprint(w, string(result))
	return
}



func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %s\n", todoId)
}
