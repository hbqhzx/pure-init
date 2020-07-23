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

type UpdatePropertyNameRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    Name string `json:"name"`

    /*   */
    Id int `json:"id"`

    /*   */
    Description string `json:"description"`
}

/*
 * param regionId:  (Required)
 * param name:  (Required)
 * param id:  (Required)
 * param description:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdatePropertyNameRequest(
    regionId string,
    name string,
    id int,
    description string,
) *UpdatePropertyNameRequest {

	return &UpdatePropertyNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/category:updatePropertyName",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        Id: id,
        Description: description,
	}
}

/*
 * param regionId:  (Required)
 * param name:  (Required)
 * param id:  (Required)
 * param description:  (Required)
 */
func NewUpdatePropertyNameRequestWithAllParams(
    regionId string,
    name string,
    id int,
    description string,
) *UpdatePropertyNameRequest {

    return &UpdatePropertyNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/category:updatePropertyName",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        Id: id,
        Description: description,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdatePropertyNameRequestWithoutParam() *UpdatePropertyNameRequest {

    return &UpdatePropertyNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/category:updatePropertyName",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *UpdatePropertyNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: (Required) */
func (r *UpdatePropertyNameRequest) SetName(name string) {
    r.Name = name
}

/* param id: (Required) */
func (r *UpdatePropertyNameRequest) SetId(id int) {
    r.Id = id
}

/* param description: (Required) */
func (r *UpdatePropertyNameRequest) SetDescription(description string) {
    r.Description = description
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdatePropertyNameRequest) GetRegionId() string {
    return r.RegionId
}

type UpdatePropertyNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdatePropertyNameResult `json:"result"`
}

type UpdatePropertyNameResult struct {
    Count int `json:"count"`
}