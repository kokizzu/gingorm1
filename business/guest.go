package business

type Guest_LoginIn struct {
	// CommonRequest
	Username string
	Password string
}

type Guest_LoginOut struct {
	// CommonResponse
	Success bool
}

const Guest_LoginRoute = `/guest/login`

func Guest_Login(in *Guest_LoginIn) (out Guest_LoginOut) {
	if in.Username == in.Password {
		out.Success = true
	}
	return
}
