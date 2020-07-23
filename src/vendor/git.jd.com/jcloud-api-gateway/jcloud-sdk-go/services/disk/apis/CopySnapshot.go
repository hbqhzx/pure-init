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

type CopySnapshotRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 快照ID  */
    SnapshotId string `json:"snapshotId"`

    /* 目标Region  */
    DestRegion string `json:"destRegion"`

    /* snapshot在新region的名字  */
    DestSnapshotName string `json:"destSnapshotName"`

    /* snapshot在新region的描述 (Optional) */
    DestSnapshotDescription *string `json:"destSnapshotDescription"`
}

/*
 * param regionId: 地域ID (Required)
 * param snapshotId: 快照ID (Required)
 * param destRegion: 目标Region (Required)
 * param destSnapshotName: snapshot在新region的名字 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCopySnapshotRequest(
    regionId string,
    snapshotId string,
    destRegion string,
    destSnapshotName string,
) *CopySnapshotRequest {

	return &CopySnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/snapshots/{snapshotId}:copySnapshot",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SnapshotId: snapshotId,
        DestRegion: destRegion,
        DestSnapshotName: destSnapshotName,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param snapshotId: 快照ID (Required)
 * param destRegion: 目标Region (Required)
 * param destSnapshotName: snapshot在新region的名字 (Required)
 * param destSnapshotDescription: snapshot在新region的描述 (Optional)
 */
func NewCopySnapshotRequestWithAllParams(
    regionId string,
    snapshotId string,
    destRegion string,
    destSnapshotName string,
    destSnapshotDescription *string,
) *CopySnapshotRequest {

    return &CopySnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshots/{snapshotId}:copySnapshot",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SnapshotId: snapshotId,
        DestRegion: destRegion,
        DestSnapshotName: destSnapshotName,
        DestSnapshotDescription: destSnapshotDescription,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCopySnapshotRequestWithoutParam() *CopySnapshotRequest {

    return &CopySnapshotRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshots/{snapshotId}:copySnapshot",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *CopySnapshotRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param snapshotId: 快照ID(Required) */
func (r *CopySnapshotRequest) SetSnapshotId(snapshotId string) {
    r.SnapshotId = snapshotId
}

/* param destRegion: 目标Region(Required) */
func (r *CopySnapshotRequest) SetDestRegion(destRegion string) {
    r.DestRegion = destRegion
}

/* param destSnapshotName: snapshot在新region的名字(Required) */
func (r *CopySnapshotRequest) SetDestSnapshotName(destSnapshotName string) {
    r.DestSnapshotName = destSnapshotName
}

/* param destSnapshotDescription: snapshot在新region的描述(Optional) */
func (r *CopySnapshotRequest) SetDestSnapshotDescription(destSnapshotDescription string) {
    r.DestSnapshotDescription = &destSnapshotDescription
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CopySnapshotRequest) GetRegionId() string {
    return r.RegionId
}

type CopySnapshotResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CopySnapshotResult `json:"result"`
}

type CopySnapshotResult struct {
    DestSnapshotId string `json:"destSnapshotId"`
}