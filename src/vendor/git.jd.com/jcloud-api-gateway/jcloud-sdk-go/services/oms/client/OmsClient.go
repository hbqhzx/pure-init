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
    oms "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/oms/apis"
    "encoding/json"
    "errors"
)

type OmsClient struct {
    core.JDCloudClient
}

func NewOmsClient(credential *core.Credential) *OmsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("oms.jdcloud-api.com")

    return &OmsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "oms",
            Revision:    "1.0.1",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *OmsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *OmsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 批量创建bucket
 */
func (c *OmsClient) BatchCreateBuckets(request *oms.BatchCreateBucketsRequest) (*oms.BatchCreateBucketsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &oms.BatchCreateBucketsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据url查询url所属用户pin、ucid、bucket、objectKey
 */
func (c *OmsClient) GetUrlInfo(request *oms.GetUrlInfoRequest) (*oms.GetUrlInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &oms.GetUrlInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据pin、region、bucket、storageClass获取space size
 */
func (c *OmsClient) GetSpace(request *oms.GetSpaceRequest) (*oms.GetSpaceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &oms.GetSpaceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}
