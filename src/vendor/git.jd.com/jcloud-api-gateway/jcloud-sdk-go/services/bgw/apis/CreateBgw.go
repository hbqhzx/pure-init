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

type CreateBgwRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 边界网关的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。  */
    BgwName string `json:"bgwName"`

    /* 边界网关的描述，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional) */
    Description *string `json:"description"`
}

/*
 * param regionId: Region ID (Required)
 * param bgwName: 边界网关的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateBgwRequest(
    regionId string,
    bgwName string,
) *CreateBgwRequest {

	return &CreateBgwRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bgws",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BgwName: bgwName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param bgwName: 边界网关的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。 (Required)
 * param description: 边界网关的描述，允许输入UTF-8编码下的全部字符，不超过256字符。 (Optional)
 */
func NewCreateBgwRequestWithAllParams(
    regionId string,
    bgwName string,
    description *string,
) *CreateBgwRequest {

    return &CreateBgwRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bgws",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BgwName: bgwName,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateBgwRequestWithoutParam() *CreateBgwRequest {

    return &CreateBgwRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bgws",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateBgwRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bgwName: 边界网关的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符。(Required) */
func (r *CreateBgwRequest) SetBgwName(bgwName string) {
    r.BgwName = bgwName
}

/* param description: 边界网关的描述，允许输入UTF-8编码下的全部字符，不超过256字符。(Optional) */
func (r *CreateBgwRequest) SetDescription(description string) {
    r.Description = &description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateBgwRequest) GetRegionId() string {
    return r.RegionId
}

type CreateBgwResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateBgwResult `json:"result"`
}

type CreateBgwResult struct {
    BgwId string `json:"bgwId"`
}