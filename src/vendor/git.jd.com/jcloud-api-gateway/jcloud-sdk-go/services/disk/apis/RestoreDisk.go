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

type RestoreDiskRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云硬盘ID  */
    DiskId string `json:"diskId"`

    /* 用于恢复云盘的快照ID  */
    SnapshotId string `json:"snapshotId"`
}

/*
 * param regionId: 地域ID (Required)
 * param diskId: 云硬盘ID (Required)
 * param snapshotId: 用于恢复云盘的快照ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRestoreDiskRequest(
    regionId string,
    diskId string,
    snapshotId string,
) *RestoreDiskRequest {

	return &RestoreDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/disks/{diskId}:restore",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DiskId: diskId,
        SnapshotId: snapshotId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param diskId: 云硬盘ID (Required)
 * param snapshotId: 用于恢复云盘的快照ID (Required)
 */
func NewRestoreDiskRequestWithAllParams(
    regionId string,
    diskId string,
    snapshotId string,
) *RestoreDiskRequest {

    return &RestoreDiskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks/{diskId}:restore",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DiskId: diskId,
        SnapshotId: snapshotId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRestoreDiskRequestWithoutParam() *RestoreDiskRequest {

    return &RestoreDiskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/disks/{diskId}:restore",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *RestoreDiskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param diskId: 云硬盘ID(Required) */
func (r *RestoreDiskRequest) SetDiskId(diskId string) {
    r.DiskId = diskId
}

/* param snapshotId: 用于恢复云盘的快照ID(Required) */
func (r *RestoreDiskRequest) SetSnapshotId(snapshotId string) {
    r.SnapshotId = snapshotId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RestoreDiskRequest) GetRegionId() string {
    return r.RegionId
}

type RestoreDiskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RestoreDiskResult `json:"result"`
}

type RestoreDiskResult struct {
}