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

type AddCNameRequest struct {

    core.JDCloudRequest

    /* 区域ID  */
    RegionId string `json:"regionId"`

    /* Bucket名称  */
    BucketName string `json:"bucketName"`

    /* 自定义域名  */
    Cname string `json:"cname"`

    /* http版本，0：http，1：https (Optional) */
    ProtoType *int `json:"protoType"`

    /* 域名  */
    EndPoint string `json:"endPoint"`

    /* 是否拦截内部域名添，任意值跳过拦截 (Optional) */
    Internal *string `json:"internal"`
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 * param cname: 自定义域名 (Required)
 * param endPoint: 域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddCNameRequest(
    regionId string,
    bucketName string,
    cname string,
    endPoint string,
) *AddCNameRequest {

	return &AddCNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/buckets/{bucketName}/cname/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BucketName: bucketName,
        Cname: cname,
        EndPoint: endPoint,
	}
}

/*
 * param regionId: 区域ID (Required)
 * param bucketName: Bucket名称 (Required)
 * param cname: 自定义域名 (Required)
 * param protoType: http版本，0：http，1：https (Optional)
 * param endPoint: 域名 (Required)
 * param internal: 是否拦截内部域名添，任意值跳过拦截 (Optional)
 */
func NewAddCNameRequestWithAllParams(
    regionId string,
    bucketName string,
    cname string,
    protoType *int,
    endPoint string,
    internal *string,
) *AddCNameRequest {

    return &AddCNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/cname/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BucketName: bucketName,
        Cname: cname,
        ProtoType: protoType,
        EndPoint: endPoint,
        Internal: internal,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddCNameRequestWithoutParam() *AddCNameRequest {

    return &AddCNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketName}/cname/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 区域ID(Required) */
func (r *AddCNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bucketName: Bucket名称(Required) */
func (r *AddCNameRequest) SetBucketName(bucketName string) {
    r.BucketName = bucketName
}

/* param cname: 自定义域名(Required) */
func (r *AddCNameRequest) SetCname(cname string) {
    r.Cname = cname
}

/* param protoType: http版本，0：http，1：https(Optional) */
func (r *AddCNameRequest) SetProtoType(protoType int) {
    r.ProtoType = &protoType
}

/* param endPoint: 域名(Required) */
func (r *AddCNameRequest) SetEndPoint(endPoint string) {
    r.EndPoint = endPoint
}

/* param internal: 是否拦截内部域名添，任意值跳过拦截(Optional) */
func (r *AddCNameRequest) SetInternal(internal string) {
    r.Internal = &internal
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddCNameRequest) GetRegionId() string {
    return r.RegionId
}

type AddCNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddCNameResult `json:"result"`
}

type AddCNameResult struct {
}