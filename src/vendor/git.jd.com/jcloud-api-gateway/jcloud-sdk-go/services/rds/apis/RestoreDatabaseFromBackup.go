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

type RestoreDatabaseFromBackupRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 库名称  */
    DbName string `json:"dbName"`

    /* 备份ID，可从备份查询接口describeBackups获取  */
    BackupId string `json:"backupId"`

    /* 指定该备份中用于恢复数据库的文件名称。通常情况下文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份  */
    BackupFileName string `json:"backupFileName"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param dbName: 库名称 (Required)
 * param backupId: 备份ID，可从备份查询接口describeBackups获取 (Required)
 * param backupFileName: 指定该备份中用于恢复数据库的文件名称。通常情况下文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewRestoreDatabaseFromBackupRequest(
    regionId string,
    instanceId string,
    dbName string,
    backupId string,
    backupFileName string,
) *RestoreDatabaseFromBackupRequest {

	return &RestoreDatabaseFromBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/databases/{dbName}:restoreDatabaseFromBackup",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        DbName: dbName,
        BackupId: backupId,
        BackupFileName: backupFileName,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param dbName: 库名称 (Required)
 * param backupId: 备份ID，可从备份查询接口describeBackups获取 (Required)
 * param backupFileName: 指定该备份中用于恢复数据库的文件名称。通常情况下文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份 (Required)
 */
func NewRestoreDatabaseFromBackupRequestWithAllParams(
    regionId string,
    instanceId string,
    dbName string,
    backupId string,
    backupFileName string,
) *RestoreDatabaseFromBackupRequest {

    return &RestoreDatabaseFromBackupRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/databases/{dbName}:restoreDatabaseFromBackup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        DbName: dbName,
        BackupId: backupId,
        BackupFileName: backupFileName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewRestoreDatabaseFromBackupRequestWithoutParam() *RestoreDatabaseFromBackupRequest {

    return &RestoreDatabaseFromBackupRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/databases/{dbName}:restoreDatabaseFromBackup",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *RestoreDatabaseFromBackupRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *RestoreDatabaseFromBackupRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param dbName: 库名称(Required) */
func (r *RestoreDatabaseFromBackupRequest) SetDbName(dbName string) {
    r.DbName = dbName
}

/* param backupId: 备份ID，可从备份查询接口describeBackups获取(Required) */
func (r *RestoreDatabaseFromBackupRequest) SetBackupId(backupId string) {
    r.BackupId = backupId
}

/* param backupFileName: 指定该备份中用于恢复数据库的文件名称。通常情况下文件名（不包括后缀）即为备份的数据库名。例如文件名为my_test_db.bak，表示该文件是my_test_db数据库的备份(Required) */
func (r *RestoreDatabaseFromBackupRequest) SetBackupFileName(backupFileName string) {
    r.BackupFileName = backupFileName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r RestoreDatabaseFromBackupRequest) GetRegionId() string {
    return r.RegionId
}

type RestoreDatabaseFromBackupResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result RestoreDatabaseFromBackupResult `json:"result"`
}

type RestoreDatabaseFromBackupResult struct {
}