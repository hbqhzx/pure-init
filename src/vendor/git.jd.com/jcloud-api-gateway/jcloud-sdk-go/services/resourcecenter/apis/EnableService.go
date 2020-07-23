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
    resourcecenter "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/resourcecenter/models"
)

type EnableServiceRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*  (Optional) */
    SiteCode *string `json:"siteCode"`

    /*  (Optional) */
    AppCode *string `json:"appCode"`

    /*   */
    ServiceCode string `json:"serviceCode"`
}

/*
 * param regionId:  (Required)
 * param serviceCode:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEnableServiceRequest(
    regionId string,
    serviceCode string,
) *EnableServiceRequest {

	return &EnableServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/serviceEnable",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceCode: serviceCode,
	}
}

/*
 * param regionId:  (Required)
 * param siteCode:  (Optional)
 * param appCode:  (Optional)
 * param serviceCode:  (Required)
 */
func NewEnableServiceRequestWithAllParams(
    regionId string,
    siteCode *string,
    appCode *string,
    serviceCode string,
) *EnableServiceRequest {

    return &EnableServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serviceEnable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SiteCode: siteCode,
        AppCode: appCode,
        ServiceCode: serviceCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEnableServiceRequestWithoutParam() *EnableServiceRequest {

    return &EnableServiceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serviceEnable",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *EnableServiceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param siteCode: (Optional) */
func (r *EnableServiceRequest) SetSiteCode(siteCode string) {
    r.SiteCode = &siteCode
}

/* param appCode: (Optional) */
func (r *EnableServiceRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCode: (Required) */
func (r *EnableServiceRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableServiceRequest) GetRegionId() string {
    return r.RegionId
}

type EnableServiceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EnableServiceResult `json:"result"`
}

type EnableServiceResult struct {
    ServiceEnable resourcecenter.ServiceEnable `json:"serviceEnable"`
}