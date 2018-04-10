package cas

import (
	"github.com/golang_cas/client"
	"github.com/golang_cas/service"
)

func NewClient(server, username, password string) client.CasClientConfig {
	return client.New(server, username, password)
}

func NewService(server, hostService string) service.CasServiceConfig {
	return service.New(server, hostService)
}
