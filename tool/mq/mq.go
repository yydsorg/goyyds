package mq

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

var _client *mqtt.Client

func init() {
	log.Println("111111111")
	_client = New("47.104.20.105", 1883)
	c = *_client
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func New(host string, p int) *mqtt.Client {
	o := mqtt.NewClientOptions().SetClientID("goyyds").AddBroker(fmt.Sprintf("tcp://%s:%d", host, p))
	o.Username = "emqx"
	o.Password = "123456"
	client := mqtt.NewClient(o)
	return &client
}

func GetClient() *mqtt.Client {
	return _client
}
