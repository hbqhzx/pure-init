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
    iotgw "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/iotgw/apis"
    "encoding/json"
    "errors"
)

type IotgwClient struct {
    core.JDCloudClient
}

func NewIotgwClient(credential *core.Credential) *IotgwClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("iotgw.jdcloud-api.com")

    return &IotgwClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "iotgw",
            Revision:    "1.3.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *IotgwClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *IotgwClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询iotgw实例列表 */
func (c *IotgwClient) DescribeInstances(request *iotgw.DescribeInstancesRequest) (*iotgw.DescribeInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.DescribeInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一个指定配置的iotgw实例 */
func (c *IotgwClient) CreateInstance(request *iotgw.CreateInstanceRequest) (*iotgw.CreateInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.CreateInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 欠禁止对instance访问 */
func (c *IotgwClient) SuspendInstance(request *iotgw.SuspendInstanceRequest) (*iotgw.SuspendInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.SuspendInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询iotgw实例详情 */
func (c *IotgwClient) DescribeInstance(request *iotgw.DescribeInstanceRequest) (*iotgw.DescribeInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.DescribeInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除单个iotgw实例 */
func (c *IotgwClient) DeleteInstance(request *iotgw.DeleteInstanceRequest) (*iotgw.DeleteInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.DeleteInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 下发设备控制指令 */
func (c *IotgwClient) DeviceControl(request *iotgw.DeviceControlRequest) (*iotgw.DeviceControlResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.DeviceControlResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 停恢复对instance的访问 */
func (c *IotgwClient) RecoverInstance(request *iotgw.RecoverInstanceRequest) (*iotgw.RecoverInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.RecoverInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询instance绑定的ExposedDomain */
func (c *IotgwClient) QueryInstanceExposeDomain(request *iotgw.QueryInstanceExposeDomainRequest) (*iotgw.QueryInstanceExposeDomainResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.QueryInstanceExposeDomainResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询instance当前状态 */
func (c *IotgwClient) QueryInstanceStatus(request *iotgw.QueryInstanceStatusRequest) (*iotgw.QueryInstanceStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.QueryInstanceStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 更新iotgw实例服务配置 */
func (c *IotgwClient) ServiceConfig(request *iotgw.ServiceConfigRequest) (*iotgw.ServiceConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &iotgw.ServiceConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}
