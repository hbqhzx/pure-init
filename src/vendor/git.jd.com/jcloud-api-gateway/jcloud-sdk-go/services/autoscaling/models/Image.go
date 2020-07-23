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

package models


type Image struct {

    /* 镜像架构。取值：i386,x86_64  (Optional) */
    Architecture string `json:"architecture"`

    /* 创建时间  (Optional) */
    CreateTime string `json:"createTime"`

    /* 镜像数据盘映射信息  (Optional) */
    DataDisks []InstanceDiskAttachment `json:"dataDisks"`

    /* 镜像描述  (Optional) */
    Desc string `json:"desc"`

    /* 镜像ID  (Optional) */
    ImageId string `json:"imageId"`

    /* 镜像来源。取值：jcloud：官方镜像；marketplace：镜像市场镜像；self：用户自己的镜像；shared：其他用户分享的镜像  (Optional) */
    ImageSource string `json:"imageSource"`

    /* 镜像名称  (Optional) */
    Name string `json:"name"`

    /* 镜像的操作系统类型。取值：windows,linux  (Optional) */
    OsType string `json:"osType"`

    /* 镜像的操作系统版本。  (Optional) */
    OsVersion string `json:"osVersion"`

    /* 镜像的操作系统发行版。取值：Ubuntu,CentOS,Windows Server  (Optional) */
    Platform string `json:"platform"`

    /* 镜像复制和转换时的进度，仅显示数值，单位为百分比  (Optional) */
    Progress string `json:"progress"`

    /* 镜像支持的系统盘类型。取值：localDisk：本地盘系统盘；cloudDisk：云盘系统盘。  (Optional) */
    RootDeviceType string `json:"rootDeviceType"`

    /* 镜像文件实际大小  (Optional) */
    SizeMB int64 `json:"sizeMB"`

    /* 创建云盘系统盘所使用的云硬盘快照ID。系统盘类型为本地盘的镜像，此参数为空。  (Optional) */
    SnapshotId string `json:"snapshotId"`

    /* <a href="http://docs.jdcloud.com/virtual-machines/api/image_status">参考镜像状态</a>  (Optional) */
    Status string `json:"status"`

    /*  (Optional) */
    SystemDisk InstanceDiskAttachment `json:"systemDisk"`

    /* 镜像系统盘大小  (Optional) */
    SystemDiskSizeGB int64 `json:"systemDiskSizeGB"`
}
