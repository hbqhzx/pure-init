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
    vpcwaf "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vpcwaf/apis"
    "encoding/json"
    "errors"
)

type VpcwafClient struct {
    core.JDCloudClient
}

func NewVpcwafClient(credential *core.Credential) *VpcwafClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("vpcwaf.jdcloud-api.com")

    return &VpcwafClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "vpcwaf",
            Revision:    "0.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *VpcwafClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *VpcwafClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询自定义访问控制规则防御事件分布 */
func (c *VpcwafClient) DescribeACLDataTrend(request *vpcwaf.DescribeACLDataTrendRequest) (*vpcwaf.DescribeACLDataTrendResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeACLDataTrendResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 提交订单 */
func (c *VpcwafClient) SubmitNewOrder(request *vpcwaf.SubmitNewOrderRequest) (*vpcwaf.SubmitNewOrderResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.SubmitNewOrderResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询用户自定义防护配置列表 */
func (c *VpcwafClient) DescribeUserConfigs(request *vpcwaf.DescribeUserConfigsRequest) (*vpcwaf.DescribeUserConfigsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeUserConfigsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询今日攻击事件分布图数据 */
func (c *VpcwafClient) DescribeTodayAttacks(request *vpcwaf.DescribeTodayAttacksRequest) (*vpcwaf.DescribeTodayAttacksResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeTodayAttacksResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询最近攻击事件趋势 */
func (c *VpcwafClient) DescribeAttackTrend(request *vpcwaf.DescribeAttackTrendRequest) (*vpcwaf.DescribeAttackTrendResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeAttackTrendResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询自定义防护规则 */
func (c *VpcwafClient) DescribeUserConfig(request *vpcwaf.DescribeUserConfigRequest) (*vpcwaf.DescribeUserConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeUserConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询WAF实例的防护配置 */
func (c *VpcwafClient) DescribeWafConfig(request *vpcwaf.DescribeWafConfigRequest) (*vpcwaf.DescribeWafConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeWafConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询攻击信息-Web应用攻击(当前仅支持查询7天内的攻击日志) */
func (c *VpcwafClient) DescribeWafVuls(request *vpcwaf.DescribeWafVulsRequest) (*vpcwaf.DescribeWafVulsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeWafVulsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建自定义防护规则 */
func (c *VpcwafClient) CreateUserConfig(request *vpcwaf.CreateUserConfigRequest) (*vpcwaf.CreateUserConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.CreateUserConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改自定义防护规则 */
func (c *VpcwafClient) ModifyUserConfig(request *vpcwaf.ModifyUserConfigRequest) (*vpcwaf.ModifyUserConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.ModifyUserConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询申请状态 */
func (c *VpcwafClient) DescribeApplyStatus(request *vpcwaf.DescribeApplyStatusRequest) (*vpcwaf.DescribeApplyStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeApplyStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询攻击信息-CC攻击(当前仅支持查询7天内的攻击日志) */
func (c *VpcwafClient) DescribeWafCCs(request *vpcwaf.DescribeWafCCsRequest) (*vpcwaf.DescribeWafCCsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeWafCCsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建WAF实例 */
func (c *VpcwafClient) CreateWaf(request *vpcwaf.CreateWafRequest) (*vpcwaf.CreateWafResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.CreateWafResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改WAF实例防护配置 */
func (c *VpcwafClient) ModifyCCEngineConfig(request *vpcwaf.ModifyCCEngineConfigRequest) (*vpcwaf.ModifyCCEngineConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.ModifyCCEngineConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询WAF实例列表 */
func (c *VpcwafClient) DescribeWafs(request *vpcwaf.DescribeWafsRequest) (*vpcwaf.DescribeWafsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeWafsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改WAF工作状态 */
func (c *VpcwafClient) ModifyWafAction(request *vpcwaf.ModifyWafActionRequest) (*vpcwaf.ModifyWafActionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.ModifyWafActionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询CC攻击防护趋势 */
func (c *VpcwafClient) DescribeCCDataTrend(request *vpcwaf.DescribeCCDataTrendRequest) (*vpcwaf.DescribeCCDataTrendResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeCCDataTrendResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改WAF实例基本信息,仅支持实例名称和绑定的albId修改 */
func (c *VpcwafClient) ModifyWaf(request *vpcwaf.ModifyWafRequest) (*vpcwaf.ModifyWafResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.ModifyWafResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询访问安全趋势 */
func (c *VpcwafClient) DescribeQPSDataTrend(request *vpcwaf.DescribeQPSDataTrendRequest) (*vpcwaf.DescribeQPSDataTrendResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeQPSDataTrendResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 申请使用应用安全网关 */
func (c *VpcwafClient) Apply(request *vpcwaf.ApplyRequest) (*vpcwaf.ApplyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.ApplyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除自定义防护规则 */
func (c *VpcwafClient) DeleteUserConfig(request *vpcwaf.DeleteUserConfigRequest) (*vpcwaf.DeleteUserConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DeleteUserConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询还未绑定Waf实例的负载均衡实例Id */
func (c *VpcwafClient) DescribeLBServerIds(request *vpcwaf.DescribeLBServerIdsRequest) (*vpcwaf.DescribeLBServerIdsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeLBServerIdsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启用WAF实例 */
func (c *VpcwafClient) EnableWafInstance(request *vpcwaf.EnableWafInstanceRequest) (*vpcwaf.EnableWafInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.EnableWafInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除WAF实例，支持批量删除，多个instanceId以','隔开 [MFA enabled] */
func (c *VpcwafClient) DeleteWaf(request *vpcwaf.DeleteWafRequest) (*vpcwaf.DeleteWafResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DeleteWafResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询WAF实例价格 */
func (c *VpcwafClient) DescribeWafPrice(request *vpcwaf.DescribeWafPriceRequest) (*vpcwaf.DescribeWafPriceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeWafPriceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 停用WAF实例 */
func (c *VpcwafClient) DisableWafInstance(request *vpcwaf.DisableWafInstanceRequest) (*vpcwaf.DisableWafInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DisableWafInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询攻击信息-访问控制攻击(当前仅支持查询7天内的攻击日志) */
func (c *VpcwafClient) DescribeWafACLs(request *vpcwaf.DescribeWafACLsRequest) (*vpcwaf.DescribeWafACLsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeWafACLsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改反爬虫防护配置 */
func (c *VpcwafClient) ModifyAntisdEngineConfig(request *vpcwaf.ModifyAntisdEngineConfigRequest) (*vpcwaf.ModifyAntisdEngineConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.ModifyAntisdEngineConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改登录防爆破防护配置 */
func (c *VpcwafClient) ModifyAntibfEngineConfig(request *vpcwaf.ModifyAntibfEngineConfigRequest) (*vpcwaf.ModifyAntibfEngineConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.ModifyAntibfEngineConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改Web应用防护配置 */
func (c *VpcwafClient) ModifyWafEngineConfig(request *vpcwaf.ModifyWafEngineConfigRequest) (*vpcwaf.ModifyWafEngineConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.ModifyWafEngineConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改恶意IP封禁防护配置 */
func (c *VpcwafClient) ModifyMaliciousIPEngineConfig(request *vpcwaf.ModifyMaliciousIPEngineConfigRequest) (*vpcwaf.ModifyMaliciousIPEngineConfigResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.ModifyMaliciousIPEngineConfigResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询WAF实例基本信息 */
func (c *VpcwafClient) DescribeWaf(request *vpcwaf.DescribeWafRequest) (*vpcwaf.DescribeWafResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpcwaf.DescribeWafResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}
