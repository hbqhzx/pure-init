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
    kmscap "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/kmscap/apis"
    "encoding/json"
    "errors"
)

type KmscapClient struct {
    core.JDCloudClient
}

func NewKmscapClient(credential *core.Credential) *KmscapClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("kmscap.jdcloud-api.com")

    return &KmscapClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "kmscap",
            Revision:    "0.1.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *KmscapClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *KmscapClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 获取可信应用列表 */
func (c *KmscapClient) DescribeTrustedAppList(request *kmscap.DescribeTrustedAppListRequest) (*kmscap.DescribeTrustedAppListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kmscap.DescribeTrustedAppListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询申请状态 */
func (c *KmscapClient) DescribeApplyStatus(request *kmscap.DescribeApplyStatusRequest) (*kmscap.DescribeApplyStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kmscap.DescribeApplyStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取用户列表 */
func (c *KmscapClient) DescribeCustomerList(request *kmscap.DescribeCustomerListRequest) (*kmscap.DescribeCustomerListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kmscap.DescribeCustomerListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 新建可信应用 */
func (c *KmscapClient) CreateTrustedApp(request *kmscap.CreateTrustedAppRequest) (*kmscap.CreateTrustedAppResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kmscap.CreateTrustedAppResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改用户配额 */
func (c *KmscapClient) ModifyQuota(request *kmscap.ModifyQuotaRequest) (*kmscap.ModifyQuotaResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kmscap.ModifyQuotaResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 申请使用密钥管理服务（KMS） */
func (c *KmscapClient) Apply(request *kmscap.ApplyRequest) (*kmscap.ApplyResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kmscap.ApplyResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除可信应用 */
func (c *KmscapClient) DeleteTrustedApp(request *kmscap.DeleteTrustedAppRequest) (*kmscap.DeleteTrustedAppResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kmscap.DeleteTrustedAppResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取报表信息 */
func (c *KmscapClient) QueryReportInfo(request *kmscap.QueryReportInfoRequest) (*kmscap.QueryReportInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &kmscap.QueryReportInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}
