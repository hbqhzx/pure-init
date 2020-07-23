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
    ucapi "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/ucapi/models"
)

type UpdateContactPersonRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 联系人信息 (Optional) */
    Person *ucapi.ContactPersonParam `json:"person"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateContactPersonRequest(
    regionId string,
) *UpdateContactPersonRequest {

	return &UpdateContactPersonRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/contact/person",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param person: 联系人信息 (Optional)
 */
func NewUpdateContactPersonRequestWithAllParams(
    regionId string,
    person *ucapi.ContactPersonParam,
) *UpdateContactPersonRequest {

    return &UpdateContactPersonRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/contact/person",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Person: person,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateContactPersonRequestWithoutParam() *UpdateContactPersonRequest {

    return &UpdateContactPersonRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/contact/person",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *UpdateContactPersonRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param person: 联系人信息(Optional) */
func (r *UpdateContactPersonRequest) SetPerson(person *ucapi.ContactPersonParam) {
    r.Person = person
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateContactPersonRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateContactPersonResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateContactPersonResult `json:"result"`
}

type UpdateContactPersonResult struct {
}