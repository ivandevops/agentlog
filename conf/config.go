package conf

import (
	"agentlog/tailf"
	"fmt"
	"github.com/astaxie/beego/config"
)


var (
	AppConfig *Config
)

type Config struct {
	LogLevel string
	LogPath  string

	ChanSize    int
	KafkaAddr   string
	CollectConf []tailf.CollectConf   	//日志的路径和topic

	EtcdAddr string
	EtcdKey  string
	LocalIP  string
}

//func loadCollectConf(conf config.Configer) (err error) {			//加载 收集日志操作 的配置
//
//	var cc tailf.CollectConf
//	cc.LogPath = conf.String("collect::log_path")				//要收集的日志
//	if len(cc.LogPath) == 0 {
//		err = errors.New("invalid collect::log_path")
//		return
//	}
//
//	cc.Topic = conf.String("collect::topic")					//收集的日志发送到那个topic
//	if len(cc.LogPath) == 0 {
//		err = errors.New("invalid collect::topic")
//		return
//	}
//
//	AppConfig.CollectConf = append(AppConfig.CollectConf, cc)		//添加到applog
//	return
//}

func LoadConf(confType, filename string) (err error) {

	conf, err := config.NewConfig(confType, filename)				//读取logagent配置文件
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	AppConfig = &Config{}
	AppConfig.LogLevel = conf.String("logs::log_level")	//本地日志等级
	if len(AppConfig.LogLevel) == 0 {
		AppConfig.LogLevel = "debug"
	}

	AppConfig.LogPath = conf.String("logs::log_path")		//本地日志记录到哪里
	if len(AppConfig.LogPath) == 0 {
		AppConfig.LogPath = "./logs"
	}
	AppConfig.ChanSize, err = conf.Int("kafka::chan_size")	//管道大小
	if err != nil {
		fmt.Println(err)
		AppConfig.ChanSize = 100
	}

	AppConfig.KafkaAddr = conf.String("kafka::server_addr")	//kfka的地址
	if len(AppConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	AppConfig.EtcdAddr = conf.String("etcd::addr")				//etcd的地址
	if len(AppConfig.EtcdAddr) == 0 {
		err = fmt.Errorf("invalid etcd addr")
		return
	}

	AppConfig.EtcdKey = conf.String("etcd::configKey")			//etcd的key
	if len(AppConfig.EtcdKey) == 0 {
		err = fmt.Errorf("invalid etcd key")
		return
	}

	AppConfig.LocalIP = conf.String("etcd::localIP")			//etcd的key
	if len(AppConfig.LocalIP) == 0 {
		err = fmt.Errorf("invalid etcd")
		return
	}

	//err = loadCollectConf(conf)
	//if err != nil {
	//	fmt.Printf("load collect conf failed, err:%v\n", err)
	//	return
	//}
	return
}

