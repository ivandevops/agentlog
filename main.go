package main

import (
	"agentlog/conf"
	"agentlog/etcd"
	"agentlog/kafka"
	"agentlog/log"
	"agentlog/server"
	"agentlog/tailf"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {

	//读取本地配置文件
 	filename := "conf/logagent.ini"
	err := conf.LoadConf("ini", filename)
	if err != nil {
		panic("load conf failed")
		return
	}

	//加载日志模块
	err = log.InitLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}
	logs.Debug("load conf succ, config:%v", conf.AppConfig)

	fmt.Println(conf.AppConfig)

	//初始化etcd并且获取配置
	collectConf, err := etcd.InitEtcd(conf.AppConfig.EtcdAddr, conf.AppConfig.EtcdKey, conf.AppConfig.LocalIP)
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}
	logs.Debug("initialize etcd succ")
	fmt.Println(collectConf)



	err = tailf.InitTail(collectConf, conf.AppConfig.ChanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}
	logs.Debug("initialize tailf succ")


	err = kafka.InitKafka(conf.AppConfig.KafkaAddr)			//初始化kafka留下一个kafka客户端
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}
	logs.Debug("initialize all succ")


	err = server.ServerRun()										//调用已经初始化的组件，向kafka里发送数据
	if err != nil {
		logs.Error("serverRUn failed, err:%v", err)
		return
	}
	logs.Info("program exited")

}




