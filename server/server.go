package server

import (
	"agentlog/kafka"
	"agentlog/tailf"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

func ServerRun() (err error) {
	for {
		msg := tailf.GetOneLine()   //获取一次 tail 信息 （日志和 对应的 topic）
		err = sendToKafka(msg)		//发送到kafka
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	fmt.Printf("read msg:%s, topic:%s\n", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}