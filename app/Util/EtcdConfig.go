package Util

import (
	"api-skeleton/app/Global"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
	"strings"
	"time"
)

//GetEtcdCli 获取etcd客户端
func GetEtcdCli() (cli *clientv3.Client, err error) {
	var endPoints []string
	nodes := strings.Split(Global.Configs.Etcd.Node, ",")
	for _, v := range nodes {
		endPoints = append(endPoints, fmt.Sprintf("%s:%s", Global.Configs.Etcd.Host, v))
	}

	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   endPoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		// handle error!
		logrus.WithFields(logrus.Fields{
			"node": endPoints,
		}).Error(fmt.Sprintf("connect to etcd failed, err:%v\n", err))
		panic(fmt.Sprintf("connect to etcd failed, err:%v\n", err))
	}

	return
}

//GetEtcdConf 获取etcd key->value值
func GetEtcdConf(key string) (string, error) {
	cli, _ := GetEtcdCli()
	defer cli.Close()
	// get
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key": key,
		}).Error(fmt.Sprintf("get from etcd failed, err:%v\n", err))
		return "", err
	}
	val := ""
	for _, ev := range resp.Kvs {
		val = string(ev.Value)
	}
	fmt.Println(val)
	return val, nil
}

//SetEtcdConf 设置etcd值
func SetEtcdConf(key string, value string) (err error) {
	cli, _ := GetEtcdCli()
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err = cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key":   key,
			"value": value,
		}).Error(fmt.Sprintf("put to etcd failed, err:%v\n", err))
		return err
	}
	return nil
}
