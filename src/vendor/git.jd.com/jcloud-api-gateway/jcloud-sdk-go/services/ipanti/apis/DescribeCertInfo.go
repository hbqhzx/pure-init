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
    ipanti "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ipanti/models"
)

type DescribeCertInfoRequest struct {

    core.JDCloudRequest

    /* 区域 Id  */
    RegionId string `json:"regionId"`

    /* 高防实例 Id  */
    InstanceId int `json:"instanceId"`

    /* 网站规则 Id  */
    WebRuleId int `json:"webRuleId"`

    /* 查询证书预览请求参数  */
    CertInfoDescribeSpec *ipanti.CertInfoDescribeSpec `json:"certInfoDescribeSpec"`
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param certInfoDescribeSpec: 查询证书预览请求参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeCertInfoRequest(
    regionId string,
    instanceId int,
    webRuleId int,
    certInfoDescribeSpec *ipanti.CertInfoDescribeSpec,
) *DescribeCertInfoRequest {

	return &DescribeCertInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}:describeCertInfo",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        CertInfoDescribeSpec: certInfoDescribeSpec,
	}
}

/*
 * param regionId: 区域 Id (Required)
 * param instanceId: 高防实例 Id (Required)
 * param webRuleId: 网站规则 Id (Required)
 * param certInfoDescribeSpec: 查询证书预览请求参数 (Required)
 */
func NewDescribeCertInfoRequestWithAllParams(
    regionId string,
    instanceId int,
    webRuleId int,
    certInfoDescribeSpec *ipanti.CertInfoDescribeSpec,
) *DescribeCertInfoRequest {

    return &DescribeCertInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}:describeCertInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        WebRuleId: webRuleId,
        CertInfoDescribeSpec: certInfoDescribeSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeCertInfoRequestWithoutParam() *DescribeCertInfoRequest {

    return &DescribeCertInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/webRules/{webRuleId}:describeCertInfo",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域 Id(Required) */
func (r *DescribeCertInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 高防实例 Id(Required) */
func (r *DescribeCertInfoRequest) SetInstanceId(instanceId int) {
    r.InstanceId = instanceId
}

/* param webRuleId: 网站规则 Id(Required) */
func (r *DescribeCertInfoRequest) SetWebRuleId(webRuleId int) {
    r.WebRuleId = webRuleId
}

/* param certInfoDescribeSpec: 查询证书预览请求参数(Required) */
func (r *DescribeCertInfoRequest) SetCertInfoDescribeSpec(certInfoDescribeSpec *ipanti.CertInfoDescribeSpec) {
    r.CertInfoDescribeSpec = certInfoDescribeSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeCertInfoRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeCertInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeCertInfoResult `json:"result"`
}

type DescribeCertInfoResult struct {
    Data ipanti.CertInfo `json:"data"`
}