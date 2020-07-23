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

type ShareSnapshotRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 快照ID  */
    SnapshotId string `json:"snapshotId"`

    /* 被共享的用户的pin  */
    DestPin string `json:"destPin"`
}

/*
 * param regionId: 地域ID (Required)
 * param snapshotId: 快照ID (Required)
 * param destPin: 被共享的用户的pin (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewShareSnapshotRequest(
    regionId string,
    snapshotId string,
    destPin string,
) *ShareSnapshotRequest {

	return &ShareSnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/snapshots/{snapshotId}:share",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SnapshotId: snapshotId,
        DestPin: destPin,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param snapshotId: 快照ID (Required)
 * param destPin: 被共享的用户的pin (Required)
 */
func NewShareSnapshotRequestWithAllParams(
    regionId string,
    snapshotId string,
    destPin string,
) *ShareSnapshotRequest {

    return &ShareSnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshots/{snapshotId}:share",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SnapshotId: snapshotId,
        DestPin: destPin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewShareSnapshotRequestWithoutParam() *ShareSnapshotRequest {

    return &ShareSnapshotRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshots/{snapshotId}:share",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ShareSnapshotRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param snapshotId: 快照ID(Required) */
func (r *ShareSnapshotRequest) SetSnapshotId(snapshotId string) {
    r.SnapshotId = snapshotId
}

/* param destPin: 被共享的用户的pin(Required) */
func (r *ShareSnapshotRequest) SetDestPin(destPin string) {
    r.DestPin = destPin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ShareSnapshotRequest) GetRegionId() string {
    return r.RegionId
}

type ShareSnapshotResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ShareSnapshotResult `json:"result"`
}

type ShareSnapshotResult struct {
}