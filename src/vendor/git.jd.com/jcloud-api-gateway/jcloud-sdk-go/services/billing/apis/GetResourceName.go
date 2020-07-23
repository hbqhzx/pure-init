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

type GetResourceNameRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    ResourceId string `json:"resourceId"`

    /*  (Optional) */
    Region *string `json:"region"`

    /*  (Optional) */
    ServiceCode *string `json:"serviceCode"`
}

/*
 * param regionId:  (Required)
 * param resourceId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetResourceNameRequest(
    regionId string,
    resourceId string,
) *GetResourceNameRequest {

	return &GetResourceNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/resourceOrder/{resourceId}/resourceName",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ResourceId: resourceId,
	}
}

/*
 * param regionId:  (Required)
 * param resourceId:  (Required)
 * param region:  (Optional)
 * param serviceCode:  (Optional)
 */
func NewGetResourceNameRequestWithAllParams(
    regionId string,
    resourceId string,
    region *string,
    serviceCode *string,
) *GetResourceNameRequest {

    return &GetResourceNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceOrder/{resourceId}/resourceName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ResourceId: resourceId,
        Region: region,
        ServiceCode: serviceCode,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetResourceNameRequestWithoutParam() *GetResourceNameRequest {

    return &GetResourceNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourceOrder/{resourceId}/resourceName",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *GetResourceNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param resourceId: (Required) */
func (r *GetResourceNameRequest) SetResourceId(resourceId string) {
    r.ResourceId = resourceId
}

/* param region: (Optional) */
func (r *GetResourceNameRequest) SetRegion(region string) {
    r.Region = &region
}

/* param serviceCode: (Optional) */
func (r *GetResourceNameRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetResourceNameRequest) GetRegionId() string {
    return r.RegionId
}

type GetResourceNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetResourceNameResult `json:"result"`
}

type GetResourceNameResult struct {
    ResourceName string `json:"resourceName"`
}