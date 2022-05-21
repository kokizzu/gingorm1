package business

import (
	"testing"

	"github.com/hexops/autogold"
)

func TestGuestLogin(t *testing.T) {

	t.Run(`wrongPasswordCase`, func(t *testing.T) {
		in := &Guest_LoginIn{
			Username: "test",
			Password: "test1",
		}
		out := Guest_Login(in)

		want := autogold.Want(`wrongPasswordCase1`, Guest_LoginOut{})
		want.Equal(t, out)
	})

	t.Run(`correctPasswordCase`, func(t *testing.T) {
		in := &Guest_LoginIn{
			Username: "test",
			Password: "test",
		}
		out := Guest_Login(in)

		want := autogold.Want(`correctPasswordCase1`, Guest_LoginOut{Success: true})
		want.Equal(t, out)
	})

}
