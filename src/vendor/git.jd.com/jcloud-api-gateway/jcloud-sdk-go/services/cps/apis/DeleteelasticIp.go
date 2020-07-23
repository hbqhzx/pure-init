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

type DeleteelasticIpRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 弹性公网IPID  */
    ElasticIpId string `json:"elasticIpId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param elasticIpId: 弹性公网IPID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteelasticIpRequest(
    regionId string,
    elasticIpId string,
) *DeleteelasticIpRequest {

	return &DeleteelasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ElasticIpId: elasticIpId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param elasticIpId: 弹性公网IPID (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 */
func NewDeleteelasticIpRequestWithAllParams(
    regionId string,
    elasticIpId string,
    clientToken *string,
) *DeleteelasticIpRequest {

    return &DeleteelasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ElasticIpId: elasticIpId,
        ClientToken: clientToken,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteelasticIpRequestWithoutParam() *DeleteelasticIpRequest {

    return &DeleteelasticIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DeleteelasticIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param elasticIpId: 弹性公网IPID(Required) */
func (r *DeleteelasticIpRequest) SetElasticIpId(elasticIpId string) {
    r.ElasticIpId = elasticIpId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *DeleteelasticIpRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteelasticIpRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteelasticIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteelasticIpResult `json:"result"`
}

type DeleteelasticIpResult struct {
    Success bool `json:"success"`
}