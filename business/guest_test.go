package business

import (
	"errors"
	"testing"

	"gingorm1/model"

	"github.com/hexops/autogold"
	"golang.org/x/crypto/bcrypt"
)

func TestGuestLogin(t *testing.T) {

	guest := GuestDeps{}

	t.Run(`userNotFound`, func(t *testing.T) {
		guest.GetUserByEmail = func(email string) (user *model.User, err error) {
			return nil, errors.New(`user not found`)
		}
		in := &Guest_LoginIn{
			Email:    "test",
			Password: "test1",
		}
		out := guest.Guest_Login(in)

		want := autogold.Want(`userNotFound1`, Guest_LoginOut{CommonResponse: CommonResponse{
			ErrorCode: 500,
			ErrorMsg:  "user not found",
		}})
		want.Equal(t, out)
	})

	t.Run(`wrongPasswordCase`, func(t *testing.T) {
		guest.GetUserByEmail = func(email string) (user *model.User, err error) {
			user = &model.User{
				Email:    "test",
				Password: ``,
			}
			return
		}
		in := &Guest_LoginIn{
			Email:    "test",
			Password: "test1",
		}
		out := guest.Guest_Login(in)

		want := autogold.Want(`wrongPasswordCase1`, Guest_LoginOut{CommonResponse: CommonResponse{
			ErrorCode: 400,
			ErrorMsg:  "username or password not match",
		}})
		want.Equal(t, out)
	})

	t.Run(`correctPasswordCase`, func(t *testing.T) {
		guest.GetUserByEmail = func(email string) (user *model.User, err error) {
			var pass []byte
			pass, err = bcrypt.GenerateFromPassword([]byte(`test1`), bcrypt.DefaultCost)
			user = &model.User{
				Email:    "test",
				Password: string(pass),
			}
			return
		}
		in := &Guest_LoginIn{
			Email:    "test",
			Password: "test1",
		}
		out := guest.Guest_Login(in)

		want := autogold.Want(`correctPasswordCase1`, Guest_LoginOut{})
		want.Equal(t, out)
	})
}

func TestGuestRegister(t *testing.T) {

	guest := GuestDeps{}

	t.Run(`registerEmailMustValid`, func(t *testing.T) {
		guest.InsertUser = func(email, password string) (err error) {
			return
		}
		in := &Guest_RegisterIn{}
		out := guest.Guest_Register(in)

		want := autogold.Want(`registerEmailMustValid1`, Guest_RegisterOut{CommonResponse: CommonResponse{
			ErrorCode: 400,
			ErrorMsg:  "email too short",
		}})
		want.Equal(t, out)
	})

	t.Run(`registerPasswordValidation`, func(t *testing.T) {
		guest.InsertUser = func(email, password string) (err error) {
			return
		}
		in := &Guest_RegisterIn{
			Email: `foo@bar.com`,
		}
		out := guest.Guest_Register(in)

		want := autogold.Want(`registerPasswordValidation1`, Guest_RegisterOut{CommonResponse: CommonResponse{
			ErrorCode: 400,
			ErrorMsg:  "password too short",
		}})
		want.Equal(t, out)
	})

	t.Run(`registerMustSuccess`, func(t *testing.T) {
		guest.InsertUser = func(email, password string) (err error) {
			return
		}
		in := &Guest_RegisterIn{
			Email:    `foo@bar.com`,
			Password: `123`,
		}
		out := guest.Guest_Register(in)

		want := autogold.Want(`registerMustSuccess1`, Guest_RegisterOut{})
		want.Equal(t, out)
	})
}
