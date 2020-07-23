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
    monitor "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
)

type GetSiteMonitorRequest struct {

    core.JDCloudRequest

    /*  (Optional) */
    Name *string `json:"name"`

    /*  (Optional) */
    Id []string `json:"id"`

    /*  (Optional) */
    Type *string `json:"type"`

    /*  (Optional) */
    PageNumber *int `json:"pageNumber"`

    /*  (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    WithStats *string `json:"withStats"`

    /* 1:包含删除对象，默认：0 (Optional) */
    WithDeleted *string `json:"withDeleted"`

    /* 查询的可用率、响应时间的时间间隔，最大不超过30天，支持分钟级别,小时级别，天级别，例如：1m、1h、1d (Optional) */
    TimeInterval *string `json:"timeInterval"`

    /* name为'id' - 站点监控id (Optional) */
    Filters []monitor.Filter `json:"filters"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetSiteMonitorRequest(
) *GetSiteMonitorRequest {

	return &GetSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/siteMonitor",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param name:  (Optional)
 * param id:  (Optional)
 * param type_:  (Optional)
 * param pageNumber:  (Optional)
 * param pageSize:  (Optional)
 * param withStats:  (Optional)
 * param withDeleted: 1:包含删除对象，默认：0 (Optional)
 * param timeInterval: 查询的可用率、响应时间的时间间隔，最大不超过30天，支持分钟级别,小时级别，天级别，例如：1m、1h、1d (Optional)
 * param filters: name为'id' - 站点监控id (Optional)
 */
func NewGetSiteMonitorRequestWithAllParams(
    name *string,
    id []string,
    type_ *string,
    pageNumber *int,
    pageSize *int,
    withStats *string,
    withDeleted *string,
    timeInterval *string,
    filters []monitor.Filter,
) *GetSiteMonitorRequest {

    return &GetSiteMonitorRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Name: name,
        Id: id,
        Type: type_,
        PageNumber: pageNumber,
        PageSize: pageSize,
        WithStats: withStats,
        WithDeleted: withDeleted,
        TimeInterval: timeInterval,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetSiteMonitorRequestWithoutParam() *GetSiteMonitorRequest {

    return &GetSiteMonitorRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/siteMonitor",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param name: (Optional) */
func (r *GetSiteMonitorRequest) SetName(name string) {
    r.Name = &name
}

/* param id: (Optional) */
func (r *GetSiteMonitorRequest) SetId(id []string) {
    r.Id = id
}

/* param type_: (Optional) */
func (r *GetSiteMonitorRequest) SetType(type_ string) {
    r.Type = &type_
}

/* param pageNumber: (Optional) */
func (r *GetSiteMonitorRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: (Optional) */
func (r *GetSiteMonitorRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param withStats: (Optional) */
func (r *GetSiteMonitorRequest) SetWithStats(withStats string) {
    r.WithStats = &withStats
}

/* param withDeleted: 1:包含删除对象，默认：0(Optional) */
func (r *GetSiteMonitorRequest) SetWithDeleted(withDeleted string) {
    r.WithDeleted = &withDeleted
}

/* param timeInterval: 查询的可用率、响应时间的时间间隔，最大不超过30天，支持分钟级别,小时级别，天级别，例如：1m、1h、1d(Optional) */
func (r *GetSiteMonitorRequest) SetTimeInterval(timeInterval string) {
    r.TimeInterval = &timeInterval
}

/* param filters: name为'id' - 站点监控id(Optional) */
func (r *GetSiteMonitorRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetSiteMonitorRequest) GetRegionId() string {
    return ""
}

type GetSiteMonitorResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetSiteMonitorResult `json:"result"`
}

type GetSiteMonitorResult struct {
    List []monitor.SiteMonitor `json:"list"`
    Total int64 `json:"total"`
}