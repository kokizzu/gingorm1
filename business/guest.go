package business

import (
	"gingorm1/model"
)

type GuestDeps struct {
	GetUserByEmail func(email string) (user *model.User, err error)
	InsertUser     func(email, password string) (err error)
}

type Guest_LoginIn struct {
	CommonRequest
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Guest_LoginOut struct {
	CommonResponse
}

const Guest_LoginRoute = `/guest/login`

func (g *GuestDeps) Guest_Login(in *Guest_LoginIn) (out Guest_LoginOut) {
	user, err := g.GetUserByEmail(in.Email)

	if err != nil {
		// if not found, return 4xx
		// else
		out.SetError(500, err.Error())
		return
	}

	if !user.PasswordMatch(in.Password) {
		out.SetError(400, `username or password not match`)
		return
	}

	return
}

type Guest_RegisterIn struct {
	CommonRequest
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Guest_RegisterOut struct {
	CommonResponse
}

const Guest_RegisterRoute = `/guest/register`

func (g *GuestDeps) Guest_Register(in *Guest_RegisterIn) (out Guest_RegisterOut) {
	if len(in.Email) < 3 {
		out.SetError(400, `email too short`)
		return
	}

	if len(in.Password) < 3 {
		out.SetError(400, `password too short`)
		return
	}

	err := g.InsertUser(in.Email, in.Password)

	if err != nil {
		// if unique insertion error, return 4xx
		// else
		out.SetError(500, err.Error())
	}

	return
}
