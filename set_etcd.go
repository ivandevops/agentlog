package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)



const (
	EtcdKey = "/oldboy/backend/logagent/config/192.168.1.237"
)


//模拟在etcd中更改配置
func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{												//生成一个etcd客户端
		Endpoints:   []string{"192.168.1.136:2379"},										//etcd是由多个，所以是个集群的模式，有多个地址，建议写域名
		DialTimeout: 5 * time.Second,														//超时控制
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

		//
		//var logConfArr []tailf.CollectConf
		//logConfArr = append(
		//	logConfArr,
		//	tailf.CollectConf{
		//		LogPath: "E:/log/bbb.txt",
		//		Topic:   "bbb_nginx_log",
		//	},
		//)
		//logConfArr = append(
		//	logConfArr,
		//	tailf.CollectConf{
		//		LogPath: "E:/log/sss.txt",
		//		Topic:   "sss_nginx_log_err",
		//	},
		//)
		//
		//data, err := json.Marshal(logConfArr)
		//if err != nil {
		//	fmt.Println("json failed, ", err)
		//	return
		//}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		cli.Delete(ctx, EtcdKey)
		return
		//_, err = cli.Put(ctx, EtcdKey, string(data))
		cancel()
		//if err != nil {
		//	fmt.Println("put failed, err:", err)
		//	return
		//}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, EtcdKey)
	//cancel()
	//if err != nil {
	//	fmt.Println("get failed, err:", err)
	//	return
	//}
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	//}
}

func main() {
	SetLogConfToEtcd()
}
