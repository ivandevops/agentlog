package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll					//确保消息不丢
	config.Producer.Partitioner = sarama.NewRandomPartitioner			//随机分区
	config.Producer.Return.Successes = true								//确保写成功

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka producer failed, err:", err)
		return
	}

	logs.Debug("init kafka succ")
	return
}

func SendToKafka(data, topic string) (err error) {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err:%v data:%v topic:%v", err, data, topic)
		return
	}

	logs.Debug("send succ, pid:%v offset:%v, topic:%v\n", pid, offset, topic)
	return
}
