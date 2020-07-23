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

package apis

import (
    "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/core"
)

type CreateTokenRequest struct {

    core.JDCloudRequest

    /* 产品名称 (Optional) */
    Product *string `json:"product"`

    /* 服务名称  */
    ServiceName string `json:"serviceName"`

    /* 分组名称 (Optional) */
    Group *string `json:"group"`

    /* IP地址  */
    Ip string `json:"ip"`

    /* 随机数 (Optional) */
    Nonce *string `json:"nonce"`

    /* token的失效时间（秒级时间戳），不传则默认token有效时间为1年 (Optional) */
    Deadline *int64 `json:"deadline"`

    /* erp (Optional) */
    Erp *string `json:"erp"`
}

/*
 * param serviceName: 服务名称 (Required)
 * param ip: IP地址 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTokenRequest(
    serviceName string,
    ip string,
) *CreateTokenRequest {

	return &CreateTokenRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/tokens",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        ServiceName: serviceName,
        Ip: ip,
	}
}

/*
 * param product: 产品名称 (Optional)
 * param serviceName: 服务名称 (Required)
 * param group: 分组名称 (Optional)
 * param ip: IP地址 (Required)
 * param nonce: 随机数 (Optional)
 * param deadline: token的失效时间（秒级时间戳），不传则默认token有效时间为1年 (Optional)
 * param erp: erp (Optional)
 */
func NewCreateTokenRequestWithAllParams(
    product *string,
    serviceName string,
    group *string,
    ip string,
    nonce *string,
    deadline *int64,
    erp *string,
) *CreateTokenRequest {

    return &CreateTokenRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/tokens",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Product: product,
        ServiceName: serviceName,
        Group: group,
        Ip: ip,
        Nonce: nonce,
        Deadline: deadline,
        Erp: erp,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTokenRequestWithoutParam() *CreateTokenRequest {

    return &CreateTokenRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/tokens",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param product: 产品名称(Optional) */
func (r *CreateTokenRequest) SetProduct(product string) {
    r.Product = &product
}

/* param serviceName: 服务名称(Required) */
func (r *CreateTokenRequest) SetServiceName(serviceName string) {
    r.ServiceName = serviceName
}

/* param group: 分组名称(Optional) */
func (r *CreateTokenRequest) SetGroup(group string) {
    r.Group = &group
}

/* param ip: IP地址(Required) */
func (r *CreateTokenRequest) SetIp(ip string) {
    r.Ip = ip
}

/* param nonce: 随机数(Optional) */
func (r *CreateTokenRequest) SetNonce(nonce string) {
    r.Nonce = &nonce
}

/* param deadline: token的失效时间（秒级时间戳），不传则默认token有效时间为1年(Optional) */
func (r *CreateTokenRequest) SetDeadline(deadline int64) {
    r.Deadline = &deadline
}

/* param erp: erp(Optional) */
func (r *CreateTokenRequest) SetErp(erp string) {
    r.Erp = &erp
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTokenRequest) GetRegionId() string {
    return ""
}

type CreateTokenResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTokenResult `json:"result"`
}

type CreateTokenResult struct {
    Token string `json:"token"`
}