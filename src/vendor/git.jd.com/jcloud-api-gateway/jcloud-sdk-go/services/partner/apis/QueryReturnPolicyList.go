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
    partner "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/partner/models"
)

type QueryReturnPolicyListRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* ID (Optional) */
    Id *int `json:"id"`

    /* 部门ID (Optional) */
    DeptId *int `json:"deptId"`

    /* 渠道商类型 (Optional) */
    DistributorType *int `json:"distributorType"`

    /* 返还类型 (Optional) */
    ReturnType *int `json:"returnType"`

    /* 项目编码 (Optional) */
    ItemId *int `json:"itemId"`

    /* 项目名称 (Optional) */
    ItemName *string `json:"itemName"`

    /* 周期类型 (Optional) */
    CircleType *int `json:"circleType"`

    /* 周期名称 (Optional) */
    CircleName *string `json:"circleName"`

    /* 指定周期标识 (Optional) */
    CircleFlag *int `json:"circleFlag"`

    /* 周期值 (Optional) */
    CircleValue *int `json:"circleValue"`

    /* 条件 (Optional) */
    Condition *string `json:"condition"`

    /* 说明 (Optional) */
    ConditionRemark *string `json:"conditionRemark"`

    /* 返还比例 (Optional) */
    ReturnRatio *int `json:"returnRatio"`

    /* 状态 (Optional) */
    Status *int `json:"status"`

    /* 创建时间 (Optional) */
    CreateTime *string `json:"createTime"`

    /* 创建人 (Optional) */
    CreateUser *string `json:"createUser"`

    /* 修改时间 (Optional) */
    UpdateTime *string `json:"updateTime"`

    /* 修改人 (Optional) */
    UpdateUser *string `json:"updateUser"`

    /* 是否删除0未删除,1已删除 (Optional) */
    Yn *int `json:"yn"`

    /* 项目名称 (Optional) */
    ItemNameLike *string `json:"itemNameLike"`

    /* 当前页序号 (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 每页结果数量 (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    Offset *int `json:"offset"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryReturnPolicyListRequest(
    regionId string,
) *QueryReturnPolicyListRequest {

	return &QueryReturnPolicyListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/returnPolicy/queryReturnPolicyList",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param id: ID (Optional)
 * param deptId: 部门ID (Optional)
 * param distributorType: 渠道商类型 (Optional)
 * param returnType: 返还类型 (Optional)
 * param itemId: 项目编码 (Optional)
 * param itemName: 项目名称 (Optional)
 * param circleType: 周期类型 (Optional)
 * param circleName: 周期名称 (Optional)
 * param circleFlag: 指定周期标识 (Optional)
 * param circleValue: 周期值 (Optional)
 * param condition: 条件 (Optional)
 * param conditionRemark: 说明 (Optional)
 * param returnRatio: 返还比例 (Optional)
 * param status: 状态 (Optional)
 * param createTime: 创建时间 (Optional)
 * param createUser: 创建人 (Optional)
 * param updateTime: 修改时间 (Optional)
 * param updateUser: 修改人 (Optional)
 * param yn: 是否删除0未删除,1已删除 (Optional)
 * param itemNameLike: 项目名称 (Optional)
 * param pageIndex: 当前页序号 (Optional)
 * param pageSize: 每页结果数量 (Optional)
 * param offset:  (Optional)
 */
func NewQueryReturnPolicyListRequestWithAllParams(
    regionId string,
    id *int,
    deptId *int,
    distributorType *int,
    returnType *int,
    itemId *int,
    itemName *string,
    circleType *int,
    circleName *string,
    circleFlag *int,
    circleValue *int,
    condition *string,
    conditionRemark *string,
    returnRatio *int,
    status *int,
    createTime *string,
    createUser *string,
    updateTime *string,
    updateUser *string,
    yn *int,
    itemNameLike *string,
    pageIndex *int,
    pageSize *int,
    offset *int,
) *QueryReturnPolicyListRequest {

    return &QueryReturnPolicyListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/returnPolicy/queryReturnPolicyList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        DeptId: deptId,
        DistributorType: distributorType,
        ReturnType: returnType,
        ItemId: itemId,
        ItemName: itemName,
        CircleType: circleType,
        CircleName: circleName,
        CircleFlag: circleFlag,
        CircleValue: circleValue,
        Condition: condition,
        ConditionRemark: conditionRemark,
        ReturnRatio: returnRatio,
        Status: status,
        CreateTime: createTime,
        CreateUser: createUser,
        UpdateTime: updateTime,
        UpdateUser: updateUser,
        Yn: yn,
        ItemNameLike: itemNameLike,
        PageIndex: pageIndex,
        PageSize: pageSize,
        Offset: offset,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryReturnPolicyListRequestWithoutParam() *QueryReturnPolicyListRequest {

    return &QueryReturnPolicyListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/returnPolicy/queryReturnPolicyList",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *QueryReturnPolicyListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: ID(Optional) */
func (r *QueryReturnPolicyListRequest) SetId(id int) {
    r.Id = &id
}

/* param deptId: 部门ID(Optional) */
func (r *QueryReturnPolicyListRequest) SetDeptId(deptId int) {
    r.DeptId = &deptId
}

/* param distributorType: 渠道商类型(Optional) */
func (r *QueryReturnPolicyListRequest) SetDistributorType(distributorType int) {
    r.DistributorType = &distributorType
}

/* param returnType: 返还类型(Optional) */
func (r *QueryReturnPolicyListRequest) SetReturnType(returnType int) {
    r.ReturnType = &returnType
}

/* param itemId: 项目编码(Optional) */
func (r *QueryReturnPolicyListRequest) SetItemId(itemId int) {
    r.ItemId = &itemId
}

/* param itemName: 项目名称(Optional) */
func (r *QueryReturnPolicyListRequest) SetItemName(itemName string) {
    r.ItemName = &itemName
}

/* param circleType: 周期类型(Optional) */
func (r *QueryReturnPolicyListRequest) SetCircleType(circleType int) {
    r.CircleType = &circleType
}

/* param circleName: 周期名称(Optional) */
func (r *QueryReturnPolicyListRequest) SetCircleName(circleName string) {
    r.CircleName = &circleName
}

/* param circleFlag: 指定周期标识(Optional) */
func (r *QueryReturnPolicyListRequest) SetCircleFlag(circleFlag int) {
    r.CircleFlag = &circleFlag
}

/* param circleValue: 周期值(Optional) */
func (r *QueryReturnPolicyListRequest) SetCircleValue(circleValue int) {
    r.CircleValue = &circleValue
}

/* param condition: 条件(Optional) */
func (r *QueryReturnPolicyListRequest) SetCondition(condition string) {
    r.Condition = &condition
}

/* param conditionRemark: 说明(Optional) */
func (r *QueryReturnPolicyListRequest) SetConditionRemark(conditionRemark string) {
    r.ConditionRemark = &conditionRemark
}

/* param returnRatio: 返还比例(Optional) */
func (r *QueryReturnPolicyListRequest) SetReturnRatio(returnRatio int) {
    r.ReturnRatio = &returnRatio
}

/* param status: 状态(Optional) */
func (r *QueryReturnPolicyListRequest) SetStatus(status int) {
    r.Status = &status
}

/* param createTime: 创建时间(Optional) */
func (r *QueryReturnPolicyListRequest) SetCreateTime(createTime string) {
    r.CreateTime = &createTime
}

/* param createUser: 创建人(Optional) */
func (r *QueryReturnPolicyListRequest) SetCreateUser(createUser string) {
    r.CreateUser = &createUser
}

/* param updateTime: 修改时间(Optional) */
func (r *QueryReturnPolicyListRequest) SetUpdateTime(updateTime string) {
    r.UpdateTime = &updateTime
}

/* param updateUser: 修改人(Optional) */
func (r *QueryReturnPolicyListRequest) SetUpdateUser(updateUser string) {
    r.UpdateUser = &updateUser
}

/* param yn: 是否删除0未删除,1已删除(Optional) */
func (r *QueryReturnPolicyListRequest) SetYn(yn int) {
    r.Yn = &yn
}

/* param itemNameLike: 项目名称(Optional) */
func (r *QueryReturnPolicyListRequest) SetItemNameLike(itemNameLike string) {
    r.ItemNameLike = &itemNameLike
}

/* param pageIndex: 当前页序号(Optional) */
func (r *QueryReturnPolicyListRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 每页结果数量(Optional) */
func (r *QueryReturnPolicyListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param offset: (Optional) */
func (r *QueryReturnPolicyListRequest) SetOffset(offset int) {
    r.Offset = &offset
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryReturnPolicyListRequest) GetRegionId() string {
    return r.RegionId
}

type QueryReturnPolicyListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryReturnPolicyListResult `json:"result"`
}

type QueryReturnPolicyListResult struct {
    Pagination partner.Pagination `json:"pagination"`
    Result []partner.ReturnPolicyResult `json:"result"`
}