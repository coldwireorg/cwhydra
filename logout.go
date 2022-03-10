package cwhydra

type Logout struct {
	Api Api
}

type LogoutInterface interface {
	Get(challenge string) (LogoutResponse, error)
	Accept(challenge string) (string, error)
	Reject(challenge string, request RejectRequest) (string, error)
}

func LogoutManager(a Api) Logout {
	return Logout{
		Api: a,
	}
}

func (c Logout) Get(challenge string) (LogoutResponse, error) {
	var res LogoutResponse
	err := c.Api.Get("/oauth2/auth/requests/logout?logout_challenge="+challenge, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (c Logout) Accept(challenge string) (string, error) {
	var res struct {
		RedirectTo string `json:"redirect_to"`
	}

	err := c.Api.Put("/oauth2/auth/requests/logout/accept?logout_challenge="+challenge, nil, &res)
	if err != nil {
		return res.RedirectTo, err
	}

	return res.RedirectTo, err
}

func (c Logout) Reject(challenge string, request RejectRequest) (string, error) {
	var res struct {
		RedirectTo string `json:"redirect_to"`
	}

	err := c.Api.Put("/oauth2/auth/requests/logout/reject?logout_challenge="+challenge, request, &res)
	if err != nil {
		return res.RedirectTo, err
	}

	return res.RedirectTo, err
}
