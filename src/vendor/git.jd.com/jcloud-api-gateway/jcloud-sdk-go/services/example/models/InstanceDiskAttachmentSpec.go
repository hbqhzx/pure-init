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

import disk "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/disk/models"

type InstanceDiskAttachmentSpec struct {

    /* 磁盘分类，local或者cloud；系统盘仅支持local；数据盘仅支持cloud (Optional) */
    DiskCategory *string `json:"diskCategory"`

    /* 自动删除，删除主机时自动删除此磁盘，默认为True，本地盘不能更改此值 (Optional) */
    AutoDelete *bool `json:"autoDelete"`

    /* 云硬盘规格 (Optional) */
    CloudDiskSpec *disk.DiskSpec `json:"cloudDiskSpec"`
}