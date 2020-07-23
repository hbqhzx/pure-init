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

type AddServiceLinePropertyValueRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    Value string `json:"value"`

    /*   */
    Description string `json:"description"`

    /*   */
    ValueType int `json:"valueType"`

    /*   */
    NameId int `json:"nameId"`

    /*  (Optional) */
    ParentId *int `json:"parentId"`
}

/*
 * param regionId:  (Required)
 * param value:  (Required)
 * param description:  (Required)
 * param valueType:  (Required)
 * param nameId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddServiceLinePropertyValueRequest(
    regionId string,
    value string,
    description string,
    valueType int,
    nameId int,
) *AddServiceLinePropertyValueRequest {

	return &AddServiceLinePropertyValueRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/serviceline/propertyValue",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Value: value,
        Description: description,
        ValueType: valueType,
        NameId: nameId,
	}
}

/*
 * param regionId:  (Required)
 * param value:  (Required)
 * param description:  (Required)
 * param valueType:  (Required)
 * param nameId:  (Required)
 * param parentId:  (Optional)
 */
func NewAddServiceLinePropertyValueRequestWithAllParams(
    regionId string,
    value string,
    description string,
    valueType int,
    nameId int,
    parentId *int,
) *AddServiceLinePropertyValueRequest {

    return &AddServiceLinePropertyValueRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serviceline/propertyValue",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Value: value,
        Description: description,
        ValueType: valueType,
        NameId: nameId,
        ParentId: parentId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddServiceLinePropertyValueRequestWithoutParam() *AddServiceLinePropertyValueRequest {

    return &AddServiceLinePropertyValueRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/serviceline/propertyValue",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *AddServiceLinePropertyValueRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param value: (Required) */
func (r *AddServiceLinePropertyValueRequest) SetValue(value string) {
    r.Value = value
}

/* param description: (Required) */
func (r *AddServiceLinePropertyValueRequest) SetDescription(description string) {
    r.Description = description
}

/* param valueType: (Required) */
func (r *AddServiceLinePropertyValueRequest) SetValueType(valueType int) {
    r.ValueType = valueType
}

/* param nameId: (Required) */
func (r *AddServiceLinePropertyValueRequest) SetNameId(nameId int) {
    r.NameId = nameId
}

/* param parentId: (Optional) */
func (r *AddServiceLinePropertyValueRequest) SetParentId(parentId int) {
    r.ParentId = &parentId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddServiceLinePropertyValueRequest) GetRegionId() string {
    return r.RegionId
}

type AddServiceLinePropertyValueResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddServiceLinePropertyValueResult `json:"result"`
}

type AddServiceLinePropertyValueResult struct {
    Count int `json:"count"`
}