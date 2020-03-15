package etcd
import (
	"agentlog/tailf"
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"strings"
	"time"
)

type EtcdClient struct {
	client *clientv3.Client
	keys   []string
}



var (
	etcdClient *EtcdClient
)




func InitEtcd(addr string, key string, localIP string) (collectConf []tailf.CollectConf, err error) {
	var IP = []string{localIP}
	//连接etcd服务器
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	etcdClient = &EtcdClient{
		client: cli,
	}

	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	//获取配置
	for _, ip := range IP {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed, err:%v", err)
			continue
		}
		cancel()
		logs.Debug("resp from etcd:%v", resp.Kvs)
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed, err:%v", err)
					continue
				}
				logs.Debug("log config is %v", collectConf)
			}
		}
	}

	//watch key
	initEtcdWatcher(addr)
	return
}

func initEtcdWatcher(addr string) {

	for _, key := range etcdClient.keys {
		go watchKey(addr,key)
	}
}

func watchKey(addr string, key string) {

	//连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return
	}

	logs.Debug("begin watch key:%s", key)
	for {
		rch := cli.Watch(context.Background(), key)
		var collectConf []tailf.CollectConf
		var getConfSucc = true
		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", key)
					continue
				}

				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", collectConf)
				tailf.UpdateConfig(collectConf)
			}
		}

	}
}
