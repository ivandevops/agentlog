package test

//
//import (
//	"fmt"
//	"github.com/coreos/etcd/clientv3"
//	"golang.org/x/net/context"
//	"time"
//)
//
//
//
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
//
//
//	// put
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	_, err = cli.Put(ctx, "q1mi", "cheng")
//	cancel()
//	if err != nil {
//		fmt.Printf("put to etcd failed, err:%v\n", err)
//		return
//	}
//
//
//	// get
//	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
//	resp, err := cli.Get(ctx, "q1mi")
//	cancel()
//	if err != nil {
//		fmt.Printf("get from etcd failed, err:%v\n", err)
//		return
//	}
//
//	for _, ev := range resp.Kvs {
//		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
//	}
//}