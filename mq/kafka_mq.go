package mq

type KafkaClient struct {
}

func GetKafkaClient() MqRepository {
	return new(KafkaClient)
}

func (client *KafkaClient) init(mqconfig MqConfig) error {
	return nil
}
