package cwhydra

type Infos struct {
	Api Api
}

type InfosInterface interface {
	Ready() (bool, error)
	Version() (string, error)
}

func InfosManager(a Api) Infos {
	return Infos{
		Api: a,
	}
}

func (c Infos) Ready() (bool, error) {
	var res struct {
		Status string `json:"status"`
	}

	err := c.Api.Get("/health/ready", &res)
	if err != nil {
		return false, err
	}

	if res.Status != "ok" {
		return false, err
	}

	return true, nil
}

func (c Infos) Version() (string, error) {
	var res struct {
		Version string `json:"version"`
	}

	err := c.Api.Get("/version", &res)
	if err != nil {
		return res.Version, err
	}

	return res.Version, nil
}
