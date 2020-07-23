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
    rds "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/rds/models"
)

type DescribeActiveQueryPerformanceRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 需要查询的数据库名，多个数据库名之间用英文逗号分隔，默认所有数据库 (Optional) */
    Db *string `json:"db"`

    /* 返回执行时间大于等于threshold的记录，默认10，单位秒 (Optional) */
    Threshold *int `json:"threshold"`

    /* 显示数据的页码，默认为1，取值范围：[-1,1000)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页。 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 每页显示的数据条数，默认为50，取值范围：[1,100]，只能为10的倍数 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeActiveQueryPerformanceRequest(
    regionId string,
    instanceId string,
) *DescribeActiveQueryPerformanceRequest {

	return &DescribeActiveQueryPerformanceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/performance:describeActiveQueryPerformance",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param db: 需要查询的数据库名，多个数据库名之间用英文逗号分隔，默认所有数据库 (Optional)
 * param threshold: 返回执行时间大于等于threshold的记录，默认10，单位秒 (Optional)
 * param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,1000)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页。 (Optional)
 * param pageSize: 每页显示的数据条数，默认为50，取值范围：[1,100]，只能为10的倍数 (Optional)
 */
func NewDescribeActiveQueryPerformanceRequestWithAllParams(
    regionId string,
    instanceId string,
    db *string,
    threshold *int,
    pageNumber *int,
    pageSize *int,
) *DescribeActiveQueryPerformanceRequest {

    return &DescribeActiveQueryPerformanceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/performance:describeActiveQueryPerformance",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Db: db,
        Threshold: threshold,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeActiveQueryPerformanceRequestWithoutParam() *DescribeActiveQueryPerformanceRequest {

    return &DescribeActiveQueryPerformanceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/performance:describeActiveQueryPerformance",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribeActiveQueryPerformanceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *DescribeActiveQueryPerformanceRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param db: 需要查询的数据库名，多个数据库名之间用英文逗号分隔，默认所有数据库(Optional) */
func (r *DescribeActiveQueryPerformanceRequest) SetDb(db string) {
    r.Db = &db
}

/* param threshold: 返回执行时间大于等于threshold的记录，默认10，单位秒(Optional) */
func (r *DescribeActiveQueryPerformanceRequest) SetThreshold(threshold int) {
    r.Threshold = &threshold
}

/* param pageNumber: 显示数据的页码，默认为1，取值范围：[-1,1000)。pageNumber为-1时，返回所有数据页码；超过总页数时，显示最后一页。(Optional) */
func (r *DescribeActiveQueryPerformanceRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 每页显示的数据条数，默认为50，取值范围：[1,100]，只能为10的倍数(Optional) */
func (r *DescribeActiveQueryPerformanceRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeActiveQueryPerformanceRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeActiveQueryPerformanceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeActiveQueryPerformanceResult `json:"result"`
}

type DescribeActiveQueryPerformanceResult struct {
    ActiveQueryPerformanceResult []rds.ActiveQueryPerformanceResult `json:"activeQueryPerformanceResult"`
    TotalCount int `json:"totalCount"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
}