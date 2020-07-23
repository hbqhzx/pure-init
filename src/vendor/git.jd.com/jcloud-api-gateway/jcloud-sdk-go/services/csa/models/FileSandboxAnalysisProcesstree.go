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


type FileSandboxAnalysisProcesstree struct {

    /* 子信息 (Optional) */
    Children []FileSandboxAnalysisProcesstree `json:"children"`

    /* 命令行 (Optional) */
    CommandLine string `json:"commandLine"`

    /* 时间戳 (Optional) */
    FirstSeen float32 `json:"firstSeen"`

    /* 进程ID (Optional) */
    Pid int `json:"pid"`

    /* 父进程ID (Optional) */
    Ppid int `json:"ppid"`

    /* 进程名 (Optional) */
    ProcessName string `json:"processName"`

    /* 是否跟踪 (Optional) */
    Track bool `json:"track"`
}