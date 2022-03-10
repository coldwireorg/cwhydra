package cwhydra

import (
	"log"
)

var (
	AdminApi *Api
)

func Init(config Config) bool {
	api := New(config)

	for {
		ready, err := InfosManager(api).Ready()

		if ready && err == nil {
			log.Println("Connected to hydra!")
			AdminApi = &api
			return true
		} else {
			log.Println(err.Error())
		}
	}
}
