package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/olebedev/config"
	conf "lib/config"
	neturl "net/url"
	"regexp"
	"strings"
	"time"
)

const (
	DefaultExpireTime = 3600
)

var Oss *config.Config

func LoadOss() {
	if Oss == nil {
		Oss, _ = conf.Config.Get("oss")
	}
}

func SignUrl(url string, expire ...time.Time) string {
	LoadOss()
	accessKeyId := Oss.UString("access_key_id")
	accessKeySecret := Oss.UString("access_key_secret")
	expireAt := time.Now().Add(time.Second * DefaultExpireTime)
	if len(expire) > 0 {
		expireAt = expire[0]
	}

	u, err := neturl.Parse(url)
	if err != nil {
		return url
	}
	bucket := ""
	path := ""
	//bucket名字在前面的情况
	if len(strings.Split(u.RequestURI(), "/")) == 2 || strings.Contains(u.Host, "confcenter") {
		bucket = regexp.MustCompile("(http(s?):\\/\\/)?([a-zA-Z0-9\\-]*)").FindStringSubmatch(url)[3]
		path = regexp.MustCompile("(http(s?):\\/\\/)?[^\\/]*\\/(.*)$").FindStringSubmatch(url)[3]
	} else {
		bucket = regexp.MustCompile("(http(s?):\\/\\/)?[^/]*/([^/]*)").FindStringSubmatch(url)[3]
		path = regexp.MustCompile("(http(s?):\\/\\/)?[^/]*/[^/]*/(.*)$").FindStringSubmatch(url)[3]
	}
	toSign := fmt.Sprintf("GET\n\n\n%d\n/%s/%s", expireAt.Unix(), bucket, path)
	hash := hmac.New(sha1.New, []byte(accessKeySecret))
	hash.Write([]byte(toSign))
	sum := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	urlSigned := fmt.Sprintf("%s?Expires=%d&AccessKey=%s&Signature=%s", url, expireAt.Unix(), accessKeyId, sum)
	return urlSigned
}
