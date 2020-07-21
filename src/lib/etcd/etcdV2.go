/* etcd.go - 模仿redis的用法来使用etcd*/

/*
modification history
--------------------
2017/4/14, by lovejoy, create
2017/4/14 by lovejoy, modify,
*/
/*
DESCRIPTION
etcd:
Usage:
    import ""

*/
package etcd

import (
	"context"
	"github.com/cihub/seelog"
	"github.com/coreos/etcd/client"
	"lib/config"
	"time"
	"errors"
)

const (
	CONFIG_ROOT = "etcd.server"
)

type EndPoints []string

//根据IDC获得配置中的etcd的服务器地址列表
func GetEndPointsByIdc(idc string) (EndPoints, error) {
	if err := config.LoadCofig(); err != nil {
		return nil, err
	}
	//fmt.Printf("%#v",config.Config)
	server, err := config.Config.Get(CONFIG_ROOT)
	if err != nil {
		return nil, err
	}
	serverList := server.UList(idc)
	if len(serverList) == 0 {
		return nil, errors.New("没有"+ idc +"的ETCD配置")
	}
	endPoints := make([]string, len(serverList))
	for k, v := range serverList {
		endPoints[k] = v.(string)
	}
	return endPoints, nil
}

func GetKeysApi(endPoints EndPoints) (client.KeysAPI, error) {
	cfg := client.Config{
		Endpoints: endPoints,
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	kapi := client.NewKeysAPI(c)
	return kapi, nil

}

func (e *EndPoints) GET(key string) (string, error) {
	kapi, kerr := GetKeysApi(*e)
	if kerr != nil {
		seelog.Error(kerr)
		return "", kerr
	}
	resp, err := kapi.Get(context.Background(), key, &client.GetOptions{
		Recursive: false,
		Sort:      false,
		Quorum:    true,
	})
	if err != nil {
		return "", err
	}
	return resp.Node.Value, err

}

//类似于redis的set，expire其实只用第一个
func (e *EndPoints) SET(key, value string, expire ...int) error {
	kapi, kerr := GetKeysApi(*e)
	if kerr != nil {
		seelog.Error(kerr)
		return kerr
	}
	ttl := 0
	if len(expire) > 0 {
		ttl = expire[0]
	} else {
		ttl = -1
	}
	kapi.Set(context.Background(), key, value, &client.SetOptions{
		PrevValue:        "",
		PrevIndex:        0,
		TTL:              time.Duration(ttl) * time.Second,
		Refresh:          false,
		Dir:              false,
		NoValueOnSuccess: false})
	return nil
}
func (e *EndPoints) SETEX(key, value string, expire ...int) error {
	return e.SET(key, value, expire...)
}

//类似于redis的set，expire其实只用第一个
func (e *EndPoints) SETNX(key, value string, expire ...int) error {
	kapi, kerr := GetKeysApi(*e)
	if kerr != nil {
		seelog.Error(kerr)
		return kerr
	}
	ttl := 0
	if len(expire) > 0 {
		ttl = expire[0]
	} else {
		ttl = -1
	}
	_, err := kapi.Set(context.Background(), key, value, &client.SetOptions{
		PrevValue:        "",
		PrevIndex:        0,
		PrevExist:        client.PrevExistType("false"), //false 类似于setnx
		TTL:              time.Duration(ttl) * time.Second,
		Refresh:          false,
		Dir:              false,
		NoValueOnSuccess: false})
	return err
}

//类似于redis的set，expire其实只用第一个

func (e *EndPoints) TTL(key string) (int64, error) {
	kapi, kerr := GetKeysApi(*e)
	if kerr != nil {
		seelog.Error(kerr)
		return -1, kerr
	}
	resp, err := kapi.Get(context.Background(), key, &client.GetOptions{
		Recursive: false,
		Sort:      false,
		Quorum:    true,
	})
	return resp.Node.TTL, err

}
func (e *EndPoints) EXPIRE(key string, expire ...int) error {
	kapi, kerr := GetKeysApi(*e)
	if kerr != nil {
		seelog.Error(kerr)
		return kerr
	}
	ttl := 0
	if len(expire) > 0 {
		ttl = expire[0]
	} else {
		ttl = -1
	}
	kapi.Set(context.Background(), key, "", &client.SetOptions{
		PrevValue:        "",
		PrevIndex:        0,
		PrevExist:        client.PrevExistType("true"),
		TTL:              time.Duration(ttl) * time.Second,
		Refresh:          true,
		Dir:              false,
		NoValueOnSuccess: false})
	return nil
}

func (e *EndPoints) DEL(key string) error {
	kapi, kerr := GetKeysApi(*e)
	if kerr != nil {
		seelog.Error(kerr)

	}
	_, err := kapi.Delete(context.Background(), key, &client.DeleteOptions{
		PrevValue: "", //无脑删
		PrevIndex: 0,
		Recursive: false,
		Dir:       false})
	if err != nil {
		return err
	}
	return nil
}
