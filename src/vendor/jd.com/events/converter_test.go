package events

import (
	"encoding/json"
	"testing"
)

var deploySample = map[string]interface{}{
	"app":         "b2b-m-web-pre",
	"cfg_version": []string{"2020"},
	"corp":        "solution-service",
	"deployment":  "balloon",
	"end_time":    1527039300,
	"env":         []string{"[{\"key\":\"SKYWING_LOG_ENABLE\",\"value\":\"true\"},{\"key\":\"JAVA_OPS\",\"value\":\"-Djava.library.path=/usr/local/lib -server -Xms2048m -Xmx2048m -XX:MaxPermSize=256m -Djava.awt.headless=true -Dsun.net.client.defaultConnectTimeout=60000 -Dsun.net.client.defaultReadTimeout=60000 -Djmagick.systemclassloader=no -Dnetworkaddress.cache.ttl=300 -Dsun.net.inetaddr.ttl=300 -Dfile.encoding=UTF-8\"},{\"key\":\"LANG\",\"value\":\"zh_CN.utf-8\"}]"},
	"group":       []string{"skywing-pre2"},
	"host_ip":     []string{"10.160.13.98"},
	"id":          "balloon_58396",
	"idc":         []string{"gyyhb"},
	"ins":         []string{"1.b2b-m-web-pre"},
	"operation":   "DeployTypeUpdate",
	"operator":    "chenxin35",
	"pdl":         "b2b",
	"start_time":  1527039277,
	"status":      "DEPLOY_STATUS_FINISH",
	"sys":         "b2b-hb-pre",
	"system":      "balloon",
	"tenant":      "",
	"update_time": 1527039300,
}

func Test_ConvertDeploy(t *testing.T) {
	c := NewConvert(map[string]string{"pdl": "product"}, []string{"pdl", "group", "env", "deployment", "cfg_version"})
	e, err := c.Convert(deploySample)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(e)
	}

	bytes, _ := json.Marshal(e)
	t.Log(string(bytes))

}
