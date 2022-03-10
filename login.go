package cwhydra

type Login struct {
	Api Api
}

type LoginInterface interface {
	Get(challenge string) (LogoutResponse, error)
	Accept(challenge string) (string, error)
	Reject(challenge string, request RejectRequest) (string, error)
}

func LoginManager(a Api) Login {
	return Login{
		Api: a,
	}
}

func (c Login) Get(challenge string) (LoginRequest, error) {
	var res LoginRequest
	err := c.Api.Get("/oauth2/auth/requests/login?login_challenge="+challenge, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (c Login) Accept(challenge string, request AcceptLoginRequest) (AcceptLoginRequest, error) {
	var res AcceptLoginRequest
	err := c.Api.Put("/oauth2/auth/requests/login/accept?login_challenge="+challenge, request, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (c Login) Reject(challenge string, request RejectRequest) (string, error) {
	var res struct {
		RedirectTo string `json:"redirect_to"`
	}

	err := c.Api.Put("/oauth2/auth/requests/login/reject?login_challenge="+challenge, request, &res)
	if err != nil {
		return res.RedirectTo, err
	}

	return res.RedirectTo, err
}
