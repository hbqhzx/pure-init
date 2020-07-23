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


type FileSandboxAnalysisBehavior struct {

    /* 键数据 (Optional) */
    ReadKeys []string `json:"readKeys"`

    /* 键数据 (Optional) */
    WriteKeys []string `json:"writeKeys"`

    /* 键数据 (Optional) */
    DeleteKeys []string `json:"deleteKeys"`

    /* 文件组 (Optional) */
    ReadFiles []string `json:"readFiles"`

    /* 文件组 (Optional) */
    WriteFiles []string `json:"writeFiles"`

    /* 文件组 (Optional) */
    DeleteFiles []string `json:"deleteFiles"`

    /* 文件组 (Optional) */
    Files []string `json:"files"`

    /* 系统相关 (Optional) */
    ExecutedCommands []string `json:"executedCommands"`

    /* 系统相关 (Optional) */
    ResolvedApis []string `json:"resolvedApis"`

    /* 系统相关 (Optional) */
    Mutexes []string `json:"mutexes"`

    /* 系统相关 (Optional) */
    CreatedServices []string `json:"createdServices"`

    /* 系统相关 (Optional) */
    StartServices []string `json:"startServices"`

    /* 时间相关 (Optional) */
    Time []string `json:"time"`

    /* 网络相关 (Optional) */
    Network []string `json:"network"`

    /* wmi信息 (Optional) */
    Wmi []string `json:"wmi"`
}