package mq

import "errors"

type MqRepository interface {
	init(mqconfig MqConfig) error
}

type MqStruct struct {
	mqRepository MqRepository
}

type MqConfig struct {
	MqType string
}

func MqInit(mqconfig MqConfig) (error, *MqStruct) {
	mqStruct := new(MqStruct)
	if mqconfig.MqType == "kafka" {

	} else if mqconfig.MqType == "nats" {

	} else {
		return errors.New("no mq to choice"), nil
	}
	return nil, mqStruct
}
