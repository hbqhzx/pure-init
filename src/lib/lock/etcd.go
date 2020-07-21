package lock

import (
	"github.com/cihub/seelog"
	config2 "github.com/olebedev/config"
	"lib/config"
	"lib/etcd"
	"strconv"
)

type Lock struct {
	LockMap  map[string][]string
	Id       int
	Locked   bool //标记是否加上锁，解锁用这个判断
}

const (
	LockPrefix = "/lock/" //锁名用 /lock/lockname 如/lock/appName-groupName 之类的
)

func (lock *Lock) TryLock() error {
	return TryBatchLock(lock.LockMap, lock.Id)
}

func (lock *Lock) Unlock() error {
	return BatchUnlock(lock.LockMap, lock.Id)
}

func (lock *Lock) Renew() error {
	return BatchRenew(lock.LockMap)
}

func BatchRenew(lockMap map[string][]string) error {
	lockConfig, _ := LoadConfig()
	ttl, terr := lockConfig.Int("ttl")
	if terr != nil {
		seelog.Error(terr)
		return terr
	}

	for idc, lockNames := range lockMap {
		e, err := etcd.GetEndPointsByIdc(idc)
		if err != nil {
			seelog.Error(err)
			return err
		}
		for _, lockName := range lockNames {
			err = e.EXPIRE(LockPrefix+lockName, ttl)
			if err != nil {
				seelog.Error(err)
				return err
			}else {
				seelog.Info("renew the lock",LockPrefix+lockName)
			}
		}
	}
	return nil

}

func TryBatchLock(lockMap map[string][]string, id int) error {
	lockConfig, _ := LoadConfig()
	ttl, terr := lockConfig.Int("ttl")
	if terr != nil {
		seelog.Error(terr)
		return terr
	}

	for idc, lockNames := range lockMap {
		e, err := etcd.GetEndPointsByIdc(idc)
		if err != nil {
			seelog.Error(err)
			return err
		}
		for _, lockName := range lockNames {
			err = e.SETNX(LockPrefix+lockName, strconv.Itoa(id), ttl)
			if err != nil {
				seelog.Error("lock "+lockName+"/"+strconv.Itoa(id)+" on idc ", idc, " fail ", err)
				return err //加锁失败一个就退出
			}
			seelog.Info("lock "+lockName+"/"+strconv.Itoa(id)+" on idc ", idc, " success ")
		}
	}
	return nil
}

func BatchUnlock(lockMap map[string][]string, id int) error {
	for idc, lockNames := range lockMap {
		e, err := etcd.GetEndPointsByIdc(idc)
		if err != nil {
			seelog.Error(err)
			return err
		}
		for _, lockName := range lockNames {
			err = e.DEL(LockPrefix + lockName)
			if err != nil {
				seelog.Error("unlock "+lockName+"/"+strconv.Itoa(id)+" on idc ",idc," fail ", err)
				//这里Log记录下其实就好，解锁失败不管，等TTL结束
			} else {
				seelog.Info("unlock " + lockName + "/" + strconv.Itoa(id) +" on idc ",idc, " success ")
			}
		}
	}
	return nil
}

func LoadConfig() (*config2.Config, error) {
	if err := config.LoadCofig(); err != nil {
		return nil, err
	}
	lockConfig, err := config.Config.Get("lock")
	if err != nil {
		return nil, err
	}
	return lockConfig, nil
}
