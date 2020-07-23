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
    logs "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/logs/models"
)

type CreateSystemTemplatesRequest struct {

    core.JDCloudRequest

    /* UID, 对应appName  */
    Uid string `json:"uid"`

    /* 日志类型  */
    Name string `json:"name"`

    /* serviceCode  */
    ServiceCode string `json:"serviceCode"`

    /* log path (Optional) */
    Path *string `json:"path"`

    /* log mode：1-JSON/2-FLATTEN/3-SEPARATE/4-ROW/5-ROWS  */
    LogMode int `json:"logMode"`

    /* tenant：原始日志只有tenantId，没有pin，则传 1，否则传 0  */
    Tenant int `json:"tenant"`

    /* fields,字段列表  */
    Fields []logs.Field `json:"fields"`
}

/*
 * param uid: UID, 对应appName (Required)
 * param name: 日志类型 (Required)
 * param serviceCode: serviceCode (Required)
 * param logMode: log mode：1-JSON/2-FLATTEN/3-SEPARATE/4-ROW/5-ROWS (Required)
 * param tenant: tenant：原始日志只有tenantId，没有pin，则传 1，否则传 0 (Required)
 * param fields: fields,字段列表 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateSystemTemplatesRequest(
    uid string,
    name string,
    serviceCode string,
    logMode int,
    tenant int,
    fields []logs.Field,
) *CreateSystemTemplatesRequest {

	return &CreateSystemTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/teamplatesx",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Uid: uid,
        Name: name,
        ServiceCode: serviceCode,
        LogMode: logMode,
        Tenant: tenant,
        Fields: fields,
	}
}

/*
 * param uid: UID, 对应appName (Required)
 * param name: 日志类型 (Required)
 * param serviceCode: serviceCode (Required)
 * param path: log path (Optional)
 * param logMode: log mode：1-JSON/2-FLATTEN/3-SEPARATE/4-ROW/5-ROWS (Required)
 * param tenant: tenant：原始日志只有tenantId，没有pin，则传 1，否则传 0 (Required)
 * param fields: fields,字段列表 (Required)
 */
func NewCreateSystemTemplatesRequestWithAllParams(
    uid string,
    name string,
    serviceCode string,
    path *string,
    logMode int,
    tenant int,
    fields []logs.Field,
) *CreateSystemTemplatesRequest {

    return &CreateSystemTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/teamplatesx",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Uid: uid,
        Name: name,
        ServiceCode: serviceCode,
        Path: path,
        LogMode: logMode,
        Tenant: tenant,
        Fields: fields,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateSystemTemplatesRequestWithoutParam() *CreateSystemTemplatesRequest {

    return &CreateSystemTemplatesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/teamplatesx",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param uid: UID, 对应appName(Required) */
func (r *CreateSystemTemplatesRequest) SetUid(uid string) {
    r.Uid = uid
}

/* param name: 日志类型(Required) */
func (r *CreateSystemTemplatesRequest) SetName(name string) {
    r.Name = name
}

/* param serviceCode: serviceCode(Required) */
func (r *CreateSystemTemplatesRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param path: log path(Optional) */
func (r *CreateSystemTemplatesRequest) SetPath(path string) {
    r.Path = &path
}

/* param logMode: log mode：1-JSON/2-FLATTEN/3-SEPARATE/4-ROW/5-ROWS(Required) */
func (r *CreateSystemTemplatesRequest) SetLogMode(logMode int) {
    r.LogMode = logMode
}

/* param tenant: tenant：原始日志只有tenantId，没有pin，则传 1，否则传 0(Required) */
func (r *CreateSystemTemplatesRequest) SetTenant(tenant int) {
    r.Tenant = tenant
}

/* param fields: fields,字段列表(Required) */
func (r *CreateSystemTemplatesRequest) SetFields(fields []logs.Field) {
    r.Fields = fields
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSystemTemplatesRequest) GetRegionId() string {
    return ""
}

type CreateSystemTemplatesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateSystemTemplatesResult `json:"result"`
}

type CreateSystemTemplatesResult struct {
}