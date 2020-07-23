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
    antipro "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/antipro/apis"
    "encoding/json"
    "errors"
)

type AntiproClient struct {
    core.JDCloudClient
}

func NewAntiproClient(credential *core.Credential) *AntiproClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("antipro.jdcloud-api.com")

    return &AntiproClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "antipro",
            Revision:    "1.0.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *AntiproClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *AntiproClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询防护包实例列表(不区分区域) */
func (c *AntiproClient) DescribeAntiproInstances(request *antipro.DescribeAntiproInstancesRequest) (*antipro.DescribeAntiproInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeAntiproInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询防护包实例 */
func (c *AntiproClient) DescribeInstance(request *antipro.DescribeInstanceRequest) (*antipro.DescribeInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加防护包防护 IP */
func (c *AntiproClient) AddProtectedIp(request *antipro.AddProtectedIpRequest) (*antipro.AddProtectedIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.AddProtectedIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 试用防护包转为正式防护包 */
func (c *AntiproClient) BuyTrialInstance(request *antipro.BuyTrialInstanceRequest) (*antipro.BuyTrialInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.BuyTrialInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询云物理服务器公网 IP 安全信息 */
func (c *AntiproClient) DescribeCpsIpResources(request *antipro.DescribeCpsIpResourcesRequest) (*antipro.DescribeCpsIpResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeCpsIpResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询攻击来源 */
func (c *AntiproClient) DescribeAttackSource(request *antipro.DescribeAttackSourceRequest) (*antipro.DescribeAttackSourceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeAttackSourceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取防护包实例或 IP 的防护规则 */
func (c *AntiproClient) DescribeProtectionRule(request *antipro.DescribeProtectionRuleRequest) (*antipro.DescribeProtectionRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeProtectionRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改独享 IP 类型实例的防护对象和 IP */
func (c *AntiproClient) ModifyProtectedObjectIp(request *antipro.ModifyProtectedObjectIpRequest) (*antipro.ModifyProtectedObjectIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.ModifyProtectedObjectIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询防护包实例列表 */
func (c *AntiproClient) DescribeInstances(request *antipro.DescribeInstancesRequest) (*antipro.DescribeInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询攻击记录, 参数 ip 优先级大于 instanceId
  - 指定 ip 参数时, 忽略 instanceId 参数, 查询 ip 相关攻击记录
  - 未指定 ip 时, 查询 instanceId 指定实例相关攻击记录
  - ip 和 instanceId 均未指定时, 查询用户所有攻击记录
 */
func (c *AntiproClient) DescribeAttackLogs(request *antipro.DescribeAttackLogsRequest) (*antipro.DescribeAttackLogsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeAttackLogsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加防护包实例到购物车, 当前支持区域, cn-north-1: 华北-北京, cn-east-2: 华东-上海 */
func (c *AntiproClient) AddProductToCart(request *antipro.AddProductToCartRequest) (*antipro.AddProductToCartResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.AddProductToCartResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询公网 IP 基本信息 */
func (c *AntiproClient) DescribeIpResourceInfo(request *antipro.DescribeIpResourceInfoRequest) (*antipro.DescribeIpResourceInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeIpResourceInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 添加区域内所有公网ip到防护包, 防护包仅能防护所在区域的公网ip, 本接口只能在防护包防护 ip 数量为无限时, 添加防护包所在区域的所有公网ip(包括弹性公网ip和与物理服务器公网ip)到防护包, 已经添加到其他防护包的ip也会被更改到当前防护包 */
func (c *AntiproClient) AddAllProtectedIp(request *antipro.AddAllProtectedIpRequest) (*antipro.AddAllProtectedIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.AddAllProtectedIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询弹性公网 IP 安全信息 */
func (c *AntiproClient) DescribeElasticIpResources(request *antipro.DescribeElasticIpResourcesRequest) (*antipro.DescribeElasticIpResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeElasticIpResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询操作日志 */
func (c *AntiproClient) DescribeOperationRecords(request *antipro.DescribeOperationRecordsRequest) (*antipro.DescribeOperationRecordsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeOperationRecordsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除防护包防护 IP */
func (c *AntiproClient) DeleteProtectedIp(request *antipro.DeleteProtectedIpRequest) (*antipro.DeleteProtectedIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DeleteProtectedIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询公网 IP 的监控流量 */
func (c *AntiproClient) DescribeIpMonitorFlow(request *antipro.DescribeIpMonitorFlowRequest) (*antipro.DescribeIpMonitorFlowResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeIpMonitorFlowResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 检测是否试用过防护包 */
func (c *AntiproClient) CheckTriedAntiPro(request *antipro.CheckTriedAntiProRequest) (*antipro.CheckTriedAntiProResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.CheckTriedAntiProResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询防护规则 Geo 拦截可设置区域 */
func (c *AntiproClient) DescribeGeoAreas(request *antipro.DescribeGeoAreasRequest) (*antipro.DescribeGeoAreasResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeGeoAreasResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 攻击记录统计, 参数 ip 优先级大于 instanceId
  - 指定 ip 参数时, 忽略 instanceId 参数, 查询 ip 相关攻击记录
  - 未指定 ip 时, 查询 instanceId 指定实例相关攻击记录
  - ip 和 instanceId 均未指定时, 查询用户所有攻击记录
 */
func (c *AntiproClient) DescribeAttackStatistics(request *antipro.DescribeAttackStatisticsRequest) (*antipro.DescribeAttackStatisticsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeAttackStatisticsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 防护信息概要 */
func (c *AntiproClient) DescribeProtectionOutline(request *antipro.DescribeProtectionOutlineRequest) (*antipro.DescribeProtectionOutlineResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeProtectionOutlineResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改防护包防护对象 */
func (c *AntiproClient) ModifyProtectedObjects(request *antipro.ModifyProtectedObjectsRequest) (*antipro.ModifyProtectedObjectsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.ModifyProtectedObjectsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改防护包实例名称 */
func (c *AntiproClient) ModifyInstanceName(request *antipro.ModifyInstanceNameRequest) (*antipro.ModifyInstanceNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.ModifyInstanceNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询已防护 IP 列表 */
func (c *AntiproClient) DescribeProtectedIpList(request *antipro.DescribeProtectedIpListRequest) (*antipro.DescribeProtectedIpListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeProtectedIpListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 检测是否开通防护包 */
func (c *AntiproClient) CheckAntiPro(request *antipro.CheckAntiProRequest) (*antipro.CheckAntiProResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.CheckAntiProResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 试用防护包, 支持区域, cn-north-1: 华北-北京, cn-east-2: 华东-上海 */
func (c *AntiproClient) TryAntiPro(request *antipro.TryAntiProRequest) (*antipro.TryAntiProResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.TryAntiProResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 检测实例名称是否可用, 同一用户的防护包实例名称不可重复 */
func (c *AntiproClient) CheckInstanceName(request *antipro.CheckInstanceNameRequest) (*antipro.CheckInstanceNameResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.CheckInstanceNameResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询各类型攻击次数, 参数 ip 优先级大于 instanceId
  - 指定 ip 参数时, 忽略 instanceId 参数, 查询 ip 相关攻击记录
  - 未指定 ip 时, 查询 instanceId 指定实例相关攻击记录
  - ip 和 instanceId 均未指定时, 查询用户所有攻击记录
 */
func (c *AntiproClient) DescribeAttackTypeCount(request *antipro.DescribeAttackTypeCountRequest) (*antipro.DescribeAttackTypeCountResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.DescribeAttackTypeCountResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改防护包实例或 IP 的防护规则 */
func (c *AntiproClient) ModifyProtectionRule(request *antipro.ModifyProtectionRuleRequest) (*antipro.ModifyProtectionRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &antipro.ModifyProtectionRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

