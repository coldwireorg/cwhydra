package cwhydra

type Client struct {
	Api Api
}

type ClientInterface interface {
	List() ([]Oauth2Client, error)
	Create(client Oauth2Client) (Oauth2Client, error)
	Delete(id string) error
}

func ClientManager(a Api) Client {
	return Client{
		Api: a,
	}
}

func (c Client) Create(client Oauth2Client) (Oauth2Client, error) {
	var res Oauth2Client
	err := c.Api.Post("/clients", client, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c Client) Delete(id string) error {
	return c.Api.Del("/clients/"+id, nil, nil)
}

func (c Client) List() ([]Oauth2Client, error) {
	var res []Oauth2Client
	err := c.Api.Get("/clients", &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c Client) Get(id string) (Oauth2Client, error) {
	var res Oauth2Client
	err := c.Api.Get("/clients/"+id, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
