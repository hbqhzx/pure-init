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
    csa "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/csa/models"
)

type QueryAlarmEventsRequest struct {

    core.JDCloudRequest

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为10；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 起始时间, 毫秒时间戳。如果不传，则需要传timeSpan (Optional) */
    TimeBegin *string `json:"timeBegin"`

    /* 结束时间, 毫秒时间戳, 不传为当前时间 (Optional) */
    TimeEnd *string `json:"timeEnd"`

    /* 天数，如果timeBegin，则该参数无效 (Optional) */
    TimeSpan *int `json:"timeSpan"`

    /* 云主机外网IP (Optional) */
    FloatingIp *string `json:"floatingIp"`

    /* 攻击类型 (Optional) */
    JdClass *int `json:"jdClass"`

    /* 云主机内网IP (Optional) */
    FixedIp *string `json:"fixedIp"`

    /* 云主机Id (Optional) */
    ServerId *string `json:"serverId"`

    /* 云主机名称 (Optional) */
    ServerName *string `json:"serverName"`

    /* 事件等级 (Optional) */
    Severity *int `json:"severity"`

    /* 事件等级,支持多个，逗号分隔 (Optional) */
    Severities *string `json:"severities"`

    /* 事件状态, 待处理传0，完成传1,2,3 (Optional) */
    Statuses *string `json:"statuses"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryAlarmEventsRequest(
) *QueryAlarmEventsRequest {

	return &QueryAlarmEventsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/alarmEvents",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为10；取值范围[10, 100] (Optional)
 * param timeBegin: 起始时间, 毫秒时间戳。如果不传，则需要传timeSpan (Optional)
 * param timeEnd: 结束时间, 毫秒时间戳, 不传为当前时间 (Optional)
 * param timeSpan: 天数，如果timeBegin，则该参数无效 (Optional)
 * param floatingIp: 云主机外网IP (Optional)
 * param jdClass: 攻击类型 (Optional)
 * param fixedIp: 云主机内网IP (Optional)
 * param serverId: 云主机Id (Optional)
 * param serverName: 云主机名称 (Optional)
 * param severity: 事件等级 (Optional)
 * param severities: 事件等级,支持多个，逗号分隔 (Optional)
 * param statuses: 事件状态, 待处理传0，完成传1,2,3 (Optional)
 */
func NewQueryAlarmEventsRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    timeBegin *string,
    timeEnd *string,
    timeSpan *int,
    floatingIp *string,
    jdClass *int,
    fixedIp *string,
    serverId *string,
    serverName *string,
    severity *int,
    severities *string,
    statuses *string,
) *QueryAlarmEventsRequest {

    return &QueryAlarmEventsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        TimeBegin: timeBegin,
        TimeEnd: timeEnd,
        TimeSpan: timeSpan,
        FloatingIp: floatingIp,
        JdClass: jdClass,
        FixedIp: fixedIp,
        ServerId: serverId,
        ServerName: serverName,
        Severity: severity,
        Severities: severities,
        Statuses: statuses,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryAlarmEventsRequestWithoutParam() *QueryAlarmEventsRequest {

    return &QueryAlarmEventsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/alarmEvents",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *QueryAlarmEventsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为10；取值范围[10, 100](Optional) */
func (r *QueryAlarmEventsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param timeBegin: 起始时间, 毫秒时间戳。如果不传，则需要传timeSpan(Optional) */
func (r *QueryAlarmEventsRequest) SetTimeBegin(timeBegin string) {
    r.TimeBegin = &timeBegin
}

/* param timeEnd: 结束时间, 毫秒时间戳, 不传为当前时间(Optional) */
func (r *QueryAlarmEventsRequest) SetTimeEnd(timeEnd string) {
    r.TimeEnd = &timeEnd
}

/* param timeSpan: 天数，如果timeBegin，则该参数无效(Optional) */
func (r *QueryAlarmEventsRequest) SetTimeSpan(timeSpan int) {
    r.TimeSpan = &timeSpan
}

/* param floatingIp: 云主机外网IP(Optional) */
func (r *QueryAlarmEventsRequest) SetFloatingIp(floatingIp string) {
    r.FloatingIp = &floatingIp
}

/* param jdClass: 攻击类型(Optional) */
func (r *QueryAlarmEventsRequest) SetJdClass(jdClass int) {
    r.JdClass = &jdClass
}

/* param fixedIp: 云主机内网IP(Optional) */
func (r *QueryAlarmEventsRequest) SetFixedIp(fixedIp string) {
    r.FixedIp = &fixedIp
}

/* param serverId: 云主机Id(Optional) */
func (r *QueryAlarmEventsRequest) SetServerId(serverId string) {
    r.ServerId = &serverId
}

/* param serverName: 云主机名称(Optional) */
func (r *QueryAlarmEventsRequest) SetServerName(serverName string) {
    r.ServerName = &serverName
}

/* param severity: 事件等级(Optional) */
func (r *QueryAlarmEventsRequest) SetSeverity(severity int) {
    r.Severity = &severity
}

/* param severities: 事件等级,支持多个，逗号分隔(Optional) */
func (r *QueryAlarmEventsRequest) SetSeverities(severities string) {
    r.Severities = &severities
}

/* param statuses: 事件状态, 待处理传0，完成传1,2,3(Optional) */
func (r *QueryAlarmEventsRequest) SetStatuses(statuses string) {
    r.Statuses = &statuses
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryAlarmEventsRequest) GetRegionId() string {
    return ""
}

type QueryAlarmEventsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryAlarmEventsResult `json:"result"`
}

type QueryAlarmEventsResult struct {
    Data []csa.AlarmEventDetail `json:"data"`
    Total int `json:"total"`
    Message string `json:"message"`
}