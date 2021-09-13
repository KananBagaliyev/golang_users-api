package app

import (
	"github.com/kananbagaliyev/golang_users-api/controllers/ping"
	"github.com/kananbagaliyev/golang_users-api/controllers/user"

)

func mapUrls(){
	//Ping controller
	router.GET("/ping", ping.Ping)

	//User controller
	router.POST("/user",user.CreateUser)
	router.GET("/user/:user_id",user.GetUser)
	router.GET("/user/search",user.SearchUser)
}
