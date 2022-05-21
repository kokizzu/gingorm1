package business

type CommonResponse struct {
	ErrorCode    int
	ErrorMsg     string
	SetAuthToken string
}

func (o *CommonResponse) SetError(code int, msg string) {
	o.ErrorCode = code
	o.ErrorMsg = msg
	if o.ErrorCode >= 500 {
		// kirim ke telegram/slack
	}
}

type CommonRequest struct {
	AuthToken string
}
