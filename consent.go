package cwhydra

type Consent struct {
	Api Api
}

type ConsentInterface interface {
	Get(challenge string) (ConsentRequest, error)
	Accept(challenge string, request AcceptConsentRequest) (AcceptConsentRequest, error)
	Reject(challenge string, request RejectRequest) (string, error)
}

func ConsentManager(a Api) Consent {
	return Consent{
		Api: a,
	}
}

func (c Consent) Get(challenge string) (ConsentRequest, error) {
	var res ConsentRequest
	err := c.Api.Get("/oauth2/auth/requests/consent?consent_challenge="+challenge, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (c Consent) Accept(challenge string, request AcceptConsentRequest) (AcceptConsentRequest, error) {
	var res AcceptConsentRequest
	err := c.Api.Put("/oauth2/auth/requests/consent/accept?consent_challenge="+challenge, request, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (c Consent) Reject(challenge string, request RejectRequest) (string, error) {
	var res struct {
		RedirectTo string `json:"redirect_to"`
	}

	err := c.Api.Put("/oauth2/auth/requests/consent/reject?consent_challenge="+challenge, request, &res)
	if err != nil {
		return res.RedirectTo, err
	}

	return res.RedirectTo, err
}
