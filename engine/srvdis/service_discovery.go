package srvdis

import (
	"time"

	"context"

	"github.com/coreos/etcd/clientv3"
	"github.com/xiaonanln/goworld/engine/gwlog"
	"github.com/xiaonanln/goworld/engine/gwutils"
)

var (
	srvdisNamespace string
)

type ServiceDelegate interface {
	ServiceType() string
	ServiceId() string
}

func Startup(ctx context.Context, etcdEndPoints []string, namespace string, delegate ServiceDelegate) {
	srvdisNamespace = namespace

	cfg := clientv3.Config{
		Endpoints:   etcdEndPoints,
		DialTimeout: time.Second,
		//Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		//HeaderTimeoutPerRequest: time.Second,
		Context: ctx,
	}

	cli, err := clientv3.New(cfg)
	if err != nil {
		gwlog.Panic(err)
	}

	go gwutils.RepeatUntilPanicless(func() {
		registerRoutine(ctx, cli, delegate)
	})
	//kv := clientv3.NewKV(cli)
	//// set "/foo" key with "bar" value
	//gwlog.Infof("Setting '/foo' key with 'bar' value")
	//resp, err := kv.Put(context.Background(), "/foo", "bar")
	//if err != nil {
	//	gwlog.Panic(err)
	//} else {
	//	// print common key info
	//	gwlog.Infof("Set is done. Metadata is %q\n", resp)
	//}
	//// get "/foo" key's value
	//gwlog.Infof("Getting '/foo' key value")
	//getresp, err := kv.Get(context.Background(), "/foo")
	//if err != nil {
	//	gwlog.Panic(err)
	//} else {
	//	// print common key info
	//	gwlog.Infof("Get is done. Metadata is %q\n", getresp)
	//	// print value
	//	gwlog.Infof("%q key has %q value\n", getresp.Kvs[0].Key, getresp.Kvs[0].Value)
	//}
}