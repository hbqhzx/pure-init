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

type PutBucketRequest struct {

    core.JDCloudRequest

    /* Region ID，例如：cn-north-1  */
    RegionId string `json:"regionId"`

    /* bucket名字，例如：test-bucket  */
    Bucketname string `json:"bucketname"`
}

/*
 * param regionId: Region ID，例如：cn-north-1 (Required)
 * param bucketname: bucket名字，例如：test-bucket (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewPutBucketRequest(
    regionId string,
    bucketname string,
) *PutBucketRequest {

	return &PutBucketRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/buckets/{bucketname}",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Bucketname: bucketname,
	}
}

/*
 * param regionId: Region ID，例如：cn-north-1 (Required)
 * param bucketname: bucket名字，例如：test-bucket (Required)
 */
func NewPutBucketRequestWithAllParams(
    regionId string,
    bucketname string,
) *PutBucketRequest {

    return &PutBucketRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketname}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Bucketname: bucketname,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewPutBucketRequestWithoutParam() *PutBucketRequest {

    return &PutBucketRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/buckets/{bucketname}",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID，例如：cn-north-1(Required) */
func (r *PutBucketRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bucketname: bucket名字，例如：test-bucket(Required) */
func (r *PutBucketRequest) SetBucketname(bucketname string) {
    r.Bucketname = bucketname
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r PutBucketRequest) GetRegionId() string {
    return r.RegionId
}

type PutBucketResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result PutBucketResult `json:"result"`
}

type PutBucketResult struct {
}