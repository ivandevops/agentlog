package test

//
//import (
//	"context"
//	"fmt"
//	"github.com/coreos/etcd/clientv3"
//	"time"
//)
//
//
//// watch demo
//
//func main() {
//	cli, err := clientv3.New(clientv3.Config{
//		Endpoints:   []string{"192.168.1.136:2379"},
//		DialTimeout: 5 * time.Second,
//	})
//	if err != nil {
//		fmt.Printf("connect to etcd failed, err:%v\n", err)
//		return
//	}
//	fmt.Println("connect to etcd success")
//	defer cli.Close()
//	// watch key:q1mi change
//	resp, err := cli.Get(context.Background(), "q1mi")
//	for _, ev := range resp.Kvs {
//		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
//	}
//	rch := cli.Watch(context.Background(), "q1mi") // <-chan WatchResponse
//	for wresp := range rch {
//		for _, ev := range wresp.Events {
//			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
//		}
//	}
//}
