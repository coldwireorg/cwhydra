package cwhydra

type Client struct {
	Api Api
}

type ClientInterface interface {
	List() ([]OAuth2Client, error)
	Create(client OAuth2Client) (OAuth2Client, error)
	Delete(id string) error
}

func ClientManager(a Api) Client {
	return Client{
		Api: a,
	}
}

func (c Client) Create(client OAuth2Client) (OAuth2Client, error) {
	var res OAuth2Client
	err := c.Api.Post("/clients", client, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c Client) Delete(id string) error {
	return c.Api.Del("/clients/"+id, nil, nil)
}

func (c Client) List() ([]OAuth2Client, error) {
	var res []OAuth2Client
	err := c.Api.Get("/clients", &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c Client) Get(id string) (OAuth2Client, error) {
	var res OAuth2Client
	err := c.Api.Get("/clients/"+id, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
