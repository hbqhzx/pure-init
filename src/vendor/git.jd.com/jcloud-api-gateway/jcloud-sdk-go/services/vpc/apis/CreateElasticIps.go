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
    vpc "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/vpc/models"
)

type CreateElasticIpsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 购买弹性ip数量；取值范围：[1,100]  */
    MaxCount int `json:"maxCount"`

    /* 指定弹性ip地址进行创建，当申请创建多个弹性ip时，必须为空 (Optional) */
    ElasticIpAddress *string `json:"elasticIpAddress"`

    /* 弹性ip规格  */
    ElasticIpSpec *vpc.ElasticIpSpec `json:"elasticIpSpec"`
}

/*
 * param regionId: Region ID (Required)
 * param maxCount: 购买弹性ip数量；取值范围：[1,100] (Required)
 * param elasticIpSpec: 弹性ip规格 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateElasticIpsRequest(
    regionId string,
    maxCount int,
    elasticIpSpec *vpc.ElasticIpSpec,
) *CreateElasticIpsRequest {

	return &CreateElasticIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIps/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        MaxCount: maxCount,
        ElasticIpSpec: elasticIpSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param maxCount: 购买弹性ip数量；取值范围：[1,100] (Required)
 * param elasticIpAddress: 指定弹性ip地址进行创建，当申请创建多个弹性ip时，必须为空 (Optional)
 * param elasticIpSpec: 弹性ip规格 (Required)
 */
func NewCreateElasticIpsRequestWithAllParams(
    regionId string,
    maxCount int,
    elasticIpAddress *string,
    elasticIpSpec *vpc.ElasticIpSpec,
) *CreateElasticIpsRequest {

    return &CreateElasticIpsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        MaxCount: maxCount,
        ElasticIpAddress: elasticIpAddress,
        ElasticIpSpec: elasticIpSpec,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateElasticIpsRequestWithoutParam() *CreateElasticIpsRequest {

    return &CreateElasticIpsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/elasticIps/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateElasticIpsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param maxCount: 购买弹性ip数量；取值范围：[1,100](Required) */
func (r *CreateElasticIpsRequest) SetMaxCount(maxCount int) {
    r.MaxCount = maxCount
}

/* param elasticIpAddress: 指定弹性ip地址进行创建，当申请创建多个弹性ip时，必须为空(Optional) */
func (r *CreateElasticIpsRequest) SetElasticIpAddress(elasticIpAddress string) {
    r.ElasticIpAddress = &elasticIpAddress
}

/* param elasticIpSpec: 弹性ip规格(Required) */
func (r *CreateElasticIpsRequest) SetElasticIpSpec(elasticIpSpec *vpc.ElasticIpSpec) {
    r.ElasticIpSpec = elasticIpSpec
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateElasticIpsRequest) GetRegionId() string {
    return r.RegionId
}

type CreateElasticIpsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateElasticIpsResult `json:"result"`
}

type CreateElasticIpsResult struct {
    ElasticIpIds []string `json:"elasticIpIds"`
    RequestId string `json:"requestId"`
}