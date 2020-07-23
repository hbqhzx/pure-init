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
    jdsf "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/jdsf/models"
)

type DescribeAppConfigVersionRequest struct {

    core.JDCloudRequest

    /* 可用区id  */
    RegionId string `json:"regionId"`

    /* 配置id  */
    AppConfigId string `json:"appConfigId"`

    /* 要获取的配置版本 id  */
    AppConfigVersionId int `json:"appConfigVersionId"`
}

/*
 * param regionId: 可用区id (Required)
 * param appConfigId: 配置id (Required)
 * param appConfigVersionId: 要获取的配置版本 id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAppConfigVersionRequest(
    regionId string,
    appConfigId string,
    appConfigVersionId int,
) *DescribeAppConfigVersionRequest {

	return &DescribeAppConfigVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/appconfig/{appConfigId}/versions/{appConfigVersionId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AppConfigId: appConfigId,
        AppConfigVersionId: appConfigVersionId,
	}
}

/*
 * param regionId: 可用区id (Required)
 * param appConfigId: 配置id (Required)
 * param appConfigVersionId: 要获取的配置版本 id (Required)
 */
func NewDescribeAppConfigVersionRequestWithAllParams(
    regionId string,
    appConfigId string,
    appConfigVersionId int,
) *DescribeAppConfigVersionRequest {

    return &DescribeAppConfigVersionRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/appconfig/{appConfigId}/versions/{appConfigVersionId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppConfigId: appConfigId,
        AppConfigVersionId: appConfigVersionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAppConfigVersionRequestWithoutParam() *DescribeAppConfigVersionRequest {

    return &DescribeAppConfigVersionRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/appconfig/{appConfigId}/versions/{appConfigVersionId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 可用区id(Required) */
func (r *DescribeAppConfigVersionRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appConfigId: 配置id(Required) */
func (r *DescribeAppConfigVersionRequest) SetAppConfigId(appConfigId string) {
    r.AppConfigId = appConfigId
}

/* param appConfigVersionId: 要获取的配置版本 id(Required) */
func (r *DescribeAppConfigVersionRequest) SetAppConfigVersionId(appConfigVersionId int) {
    r.AppConfigVersionId = appConfigVersionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAppConfigVersionRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAppConfigVersionResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAppConfigVersionResult `json:"result"`
}

type DescribeAppConfigVersionResult struct {
    AppConfigVersionDetail jdsf.AppConfigVersionDetail `json:"appConfigVersionDetail"`
}