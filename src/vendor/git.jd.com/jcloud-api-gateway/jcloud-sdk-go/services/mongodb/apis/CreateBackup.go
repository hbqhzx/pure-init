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

type CreateBackupRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 备份名称 (Optional) */
    BackupName *string `json:"backupName"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateBackupRequest(
    regionId string,
    instanceId string,
) *CreateBackupRequest {

	return &CreateBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backups",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: 实例ID (Required)
 * param backupName: 备份名称 (Optional)
 */
func NewCreateBackupRequestWithAllParams(
    regionId string,
    instanceId string,
    backupName *string,
) *CreateBackupRequest {

    return &CreateBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backups",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        BackupName: backupName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateBackupRequestWithoutParam() *CreateBackupRequest {

    return &CreateBackupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/backups",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateBackupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 实例ID(Required) */
func (r *CreateBackupRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param backupName: 备份名称(Optional) */
func (r *CreateBackupRequest) SetBackupName(backupName string) {
    r.BackupName = &backupName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateBackupRequest) GetRegionId() string {
    return r.RegionId
}

type CreateBackupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateBackupResult `json:"result"`
}

type CreateBackupResult struct {
    BackupId string `json:"backupId"`
}