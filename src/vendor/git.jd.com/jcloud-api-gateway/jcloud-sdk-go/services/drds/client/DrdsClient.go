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
    drds "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/drds/apis"
    "encoding/json"
    "errors"
)

type DrdsClient struct {
    core.JDCloudClient
}

func NewDrdsClient(credential *core.Credential) *DrdsClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("drds.jdcloud-api.com")

    return &DrdsClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "drds",
            Revision:    "0.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *DrdsClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *DrdsClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 根据实例的的id，获取实例名称 */
func (c *DrdsClient) SelectDetailList(request *drds.SelectDetailListRequest) (*drds.SelectDetailListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &drds.SelectDetailListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除一个DRDS实例 [MFA enabled] */
func (c *DrdsClient) DeleteInstance(request *drds.DeleteInstanceRequest) (*drds.DeleteInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &drds.DeleteInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取当前账号下所有DRDS实例的概要信息，例如实例类型，版本，计费信息等 */
func (c *DrdsClient) DescribeInstances(request *drds.DescribeInstancesRequest) (*drds.DescribeInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &drds.DescribeInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询DRDS实例详细信息 */
func (c *DrdsClient) DescribeInstanceAttributes(request *drds.DescribeInstanceAttributesRequest) (*drds.DescribeInstanceAttributesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &drds.DescribeInstanceAttributesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}
