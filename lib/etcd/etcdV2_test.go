/* etcd.go -   */

/*
modification history
--------------------
2017/4/17, by lovejoy, create
2017/4/17 by lovejoy, modify,
*/
/*
DESCRIPTION
etcd:
Usage:
    import ""

*/
package etcd

import (
	"fmt"
	"testing"
)

const (
	testKey     = "/unittest/etcd/key"
	notExitsKey = "/unittest/etcd/key_not_exits"
)

func TestGetServer(t *testing.T) {
	s, err := GetEndPointsByIdc("yf")
	fmt.Printf("%#v", s)
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	e := &EndPoints{"http://192.168.168.183:2379"}

	r, err := e.GET(testKey)
	if err != nil {
		t.Error(err)
	}

	r, err = e.GET(notExitsKey) //这里看get 会不会异常不用显式error

	//test set
	e.SET(testKey, "set", 1024)
	r, _ = e.GET(testKey)
	if r != "set" {
		t.Error("SET FAIL")
	}
	//test set ttl
	i, _ := e.TTL(testKey)
	if i != 1024 {
		t.Error("SET TTL FAIL")
	}
	//test setnx
	e.SETNX(notExitsKey, "setnx", 1024)
	r, _ = e.GET(notExitsKey)
	if r != "setnx" {
		t.Error("SETNX FAIL")
	}
	//test setnx ttl
	i, _ = e.TTL(testKey)
	if i != 1024 {
		t.Error("SET TTL FAIL")
	}
	//test expire
	e.EXPIRE(testKey, 4096)
	i, _ = e.TTL(testKey)
	if i != 4096 {
		t.Error("EXPIRE TTL FAIL")
	}
}
