package main

import "fmt"

import (
	"encoding/json"
)

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func main() {
	type User struct {
		Name string `json:"name"`
	}

	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// json序列化
	user := Response[User]{
		Code: 0,
		Msg:  "成功",
		Data: User{
			Name: "枫枫",
		},
	}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
	userInfo := Response[UserInfo]{
		Code: 0,
		Msg:  "成功",
		Data: UserInfo{
			Name: "枫枫",
			Age:  24,
		},
	}
	byteData, _ = json.Marshal(userInfo)
	fmt.Println(string(byteData))

	//json反序列化
	var userResponse Response[User] // 如果这里不用泛型，那么反序列化以后，对象里面的user和userInfo对象就是一个map，userResponse.Data.Name就会报错
	json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"枫枫"}}`), &userResponse)
	fmt.Println(userResponse.Data.Name)

	var userInfoResponse Response[UserInfo]
	json.Unmarshal([]byte(`{"code":0,"msg":"成功","data":{"name":"枫枫","age":24}}`), &userInfoResponse)
	fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age)
}
