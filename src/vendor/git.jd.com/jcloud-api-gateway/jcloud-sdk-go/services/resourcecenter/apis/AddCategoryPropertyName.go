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

type AddCategoryPropertyNameRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    Name string `json:"name"`

    /*   */
    Type int `json:"type"`

    /*   */
    DataType int `json:"dataType"`

    /*   */
    CategoryId int `json:"categoryId"`

    /*   */
    Description string `json:"description"`

    /*  (Optional) */
    ParentId *int `json:"parentId"`
}

/*
 * param regionId:  (Required)
 * param name:  (Required)
 * param type_:  (Required)
 * param dataType:  (Required)
 * param categoryId:  (Required)
 * param description:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddCategoryPropertyNameRequest(
    regionId string,
    name string,
    type_ int,
    dataType int,
    categoryId int,
    description string,
) *AddCategoryPropertyNameRequest {

	return &AddCategoryPropertyNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/category/propertyName",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Name: name,
        Type: type_,
        DataType: dataType,
        CategoryId: categoryId,
        Description: description,
	}
}

/*
 * param regionId:  (Required)
 * param name:  (Required)
 * param type_:  (Required)
 * param dataType:  (Required)
 * param categoryId:  (Required)
 * param description:  (Required)
 * param parentId:  (Optional)
 */
func NewAddCategoryPropertyNameRequestWithAllParams(
    regionId string,
    name string,
    type_ int,
    dataType int,
    categoryId int,
    description string,
    parentId *int,
) *AddCategoryPropertyNameRequest {

    return &AddCategoryPropertyNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/category/propertyName",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Name: name,
        Type: type_,
        DataType: dataType,
        CategoryId: categoryId,
        Description: description,
        ParentId: parentId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddCategoryPropertyNameRequestWithoutParam() *AddCategoryPropertyNameRequest {

    return &AddCategoryPropertyNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/category/propertyName",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *AddCategoryPropertyNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param name: (Required) */
func (r *AddCategoryPropertyNameRequest) SetName(name string) {
    r.Name = name
}

/* param type_: (Required) */
func (r *AddCategoryPropertyNameRequest) SetType(type_ int) {
    r.Type = type_
}

/* param dataType: (Required) */
func (r *AddCategoryPropertyNameRequest) SetDataType(dataType int) {
    r.DataType = dataType
}

/* param categoryId: (Required) */
func (r *AddCategoryPropertyNameRequest) SetCategoryId(categoryId int) {
    r.CategoryId = categoryId
}

/* param description: (Required) */
func (r *AddCategoryPropertyNameRequest) SetDescription(description string) {
    r.Description = description
}

/* param parentId: (Optional) */
func (r *AddCategoryPropertyNameRequest) SetParentId(parentId int) {
    r.ParentId = &parentId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddCategoryPropertyNameRequest) GetRegionId() string {
    return r.RegionId
}

type AddCategoryPropertyNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddCategoryPropertyNameResult `json:"result"`
}

type AddCategoryPropertyNameResult struct {
    Count int `json:"count"`
}