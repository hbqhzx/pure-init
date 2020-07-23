// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package client

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
    monitorcm "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitorcm/apis"
    "encoding/json"
    "errors"
)

type MonitorcmClient struct {
    core.JDCloudClient
}

func NewMonitorcmClient(credential *core.Credential) *MonitorcmClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("monitorcm.jdcloud-api.com")

    return &MonitorcmClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "monitorcm",
            Revision:    "1.0.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *MonitorcmClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *MonitorcmClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 上报用户自定义监控数据 */
func (c *MonitorcmClient) PutMetricData(request *monitorcm.PutMetricDataRequest) (*monitorcm.PutMetricDataResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitorcm.PutMetricDataResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

