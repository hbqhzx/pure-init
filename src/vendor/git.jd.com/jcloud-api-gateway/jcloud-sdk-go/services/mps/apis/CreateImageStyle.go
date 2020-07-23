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

type CreateImageStyleRequest struct {

    core.JDCloudRequest

    /* 区域ID  */
    RegionId string `json:"regionId"`

    /* Bucket名称  */
    BucketName string `json:"bucketName"`

    /* 图片样式id(readOnly) (Optional) */
    Id *int64 `json:"id"`

    /* 用户id(readOnly) (Optional) */
    UserId *string `json:"userId"`

    /* 图片样式名称 (Optional) */
    StyleName *string `json:"styleName"`

    /* 图片样式参数 (Optional) */
    Params *string `json:"params"`

    /* 图片样式参数别名 (Optional) */
    ParamAlias *string `json:"paramAlias"`

    /* 所属区域(readOnly) (Optional) */
    RegionId *string `json:"regionId"`

    /* 所属Bucket(readOnly) (Optional) */
    BucketName *string `json:"bucketName"`

    /* 图片样式状态(readOnly) (Optional) */
    Status *int `json:"status"`

    /* 修改时间(readOnly) (Optional) */
    ModifyTime *string `json:"modifyTime"`

    /* 创建时间(readOnly) (Optional) */
    CreatedTime *string `json:"createdTime"`
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateImageStyleRequest(
    regionId string,
    bucketName string,
) *CreateImageStyleRequest {

	return &CreateImageStyleRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/buckets/{bucketName}/imageStyles",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BucketName: bucketName,
	}
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 * param id: 图片样式id(readOnly) (Optional)
 * param userId: 用户id(readOnly) (Optional)
 * param styleName: 图片样式名称 (Optional)
 * param params: 图片样式参数 (Optional)
 * param paramAlias: 图片样式参数别名 (Optional)
 * param regionId: 所属区域(readOnly) (Optional)
 * param bucketName: 所属Bucket(readOnly) (Optional)
 * param status: 图片样式状态(readOnly) (Optional)
 * param modifyTime: 修改时间(readOnly) (Optional)
 * param createdTime: 创建时间(readOnly) (Optional)
 */
func NewCreateImageStyleRequestWithAllParams(
    regionId string,
    bucketName string,
    id *int64,
    userId *string,
    styleName *string,
    params *string,
    paramAlias *string,
    regionId *string,
    bucketName *string,
    status *int,
    modifyTime *string,
    createdTime *string,
) *CreateImageStyleRequest {

    return &CreateImageStyleRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/imageStyles",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BucketName: bucketName,
        Id: id,
        UserId: userId,
        StyleName: styleName,
        Params: params,
        ParamAlias: paramAlias,
        RegionId: regionId,
        BucketName: bucketName,
        Status: status,
        ModifyTime: modifyTime,
        CreatedTime: createdTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateImageStyleRequestWithoutParam() *CreateImageStyleRequest {

    return &CreateImageStyleRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/imageStyles",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域ID(Required) */
func (r *CreateImageStyleRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bucketName: Bucket名称(Required) */
func (r *CreateImageStyleRequest) SetBucketName(bucketName string) {
    r.BucketName = bucketName
}

/* param id: 图片样式id(readOnly)(Optional) */
func (r *CreateImageStyleRequest) SetId(id int64) {
    r.Id = &id
}

/* param userId: 用户id(readOnly)(Optional) */
func (r *CreateImageStyleRequest) SetUserId(userId string) {
    r.UserId = &userId
}

/* param styleName: 图片样式名称(Optional) */
func (r *CreateImageStyleRequest) SetStyleName(styleName string) {
    r.StyleName = &styleName
}

/* param params: 图片样式参数(Optional) */
func (r *CreateImageStyleRequest) SetParams(params string) {
    r.Params = &params
}

/* param paramAlias: 图片样式参数别名(Optional) */
func (r *CreateImageStyleRequest) SetParamAlias(paramAlias string) {
    r.ParamAlias = &paramAlias
}

/* param regionId: 所属区域(readOnly)(Optional) */
func (r *CreateImageStyleRequest) SetRegionId(regionId string) {
    r.RegionId = &regionId
}

/* param bucketName: 所属Bucket(readOnly)(Optional) */
func (r *CreateImageStyleRequest) SetBucketName(bucketName string) {
    r.BucketName = &bucketName
}

/* param status: 图片样式状态(readOnly)(Optional) */
func (r *CreateImageStyleRequest) SetStatus(status int) {
    r.Status = &status
}

/* param modifyTime: 修改时间(readOnly)(Optional) */
func (r *CreateImageStyleRequest) SetModifyTime(modifyTime string) {
    r.ModifyTime = &modifyTime
}

/* param createdTime: 创建时间(readOnly)(Optional) */
func (r *CreateImageStyleRequest) SetCreatedTime(createdTime string) {
    r.CreatedTime = &createdTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateImageStyleRequest) GetRegionId() string {
    return r.RegionId
}

type CreateImageStyleResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateImageStyleResult `json:"result"`
}

type CreateImageStyleResult struct {
    Id int64 `json:"id"`
}