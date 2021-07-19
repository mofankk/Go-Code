package utils

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var ConsulClient *consulapi.Client

func init() {
	config := consulapi.DefaultConfig()
	config.Address = "192.168.3.237:8500"

	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
}

func RegisterService() {


	reg := &consulapi.AgentServiceRegistration{
		ID: "userService",
		Name: "service_a",
		Address: "192.168.3.153",
		Port: 8080,
		Tags: []string{"primary", "v1"},
	}

	check := &consulapi.AgentServiceCheck{
		Interval: "5s",
		HTTP: "http://192.168.3.153:8080" + "/health",
	}

	reg.Check = check

	err := ConsulClient.Agent().ServiceRegister(reg)
	if err != nil {
		log.Fatal(err)
	}
}

func DeRegisterService() {
	err := ConsulClient.Agent().ServiceDeregister("userService")
	if err != nil {
		log.Fatal(err)
	}
}
