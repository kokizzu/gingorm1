package main

import (
	"gingorm1/business"
	"gingorm1/config"
	"gingorm1/model"
	"gingorm1/presentation"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.ConnectDB()
	db.AutoMigrate(&model.User{})
	userRepo := model.UserConn{DB: db}

	guest := &business.GuestDeps{
		GetUserByEmail: userRepo.GetUserByEmail,
		InsertUser:     userRepo.InsertUser,
	}

	r.Use(func(context *gin.Context) {
		// if route segment 1 bukan guest, harus punya cookie
		// parse dari cookie/header
		// context.Set(`authToken`, `whatever`)
		context.Next()
	})

	// TODO: refactor common serialization/deserialization to presentation/rest.go
	r.POST(business.Guest_RegisterRoute, func(c *gin.Context) {
		//authToken := c.Get(`authToken`)
		in := business.Guest_RegisterIn{
			//CommonRequest{authToken},
		}
		presentation.ReadRestInput[business.Guest_RegisterIn](c, &in, &in.CommonRequest)
		out := guest.Guest_Register(&in)
		presentation.WriteRestOutput[business.Guest_RegisterOut](c, out, &out.CommonResponse)
	})

	r.POST(business.Guest_LoginRoute, func(c *gin.Context) {
		in := business.Guest_LoginIn{}
		presentation.ReadRestInput[business.Guest_LoginIn](c, &in, &in.CommonRequest)
		out := guest.Guest_Login(&in)
		presentation.WriteRestOutput[business.Guest_LoginOut](c, out, &out.CommonResponse)
	})

	r.Run(config.ListenAddress)
}
