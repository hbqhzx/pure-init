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


type WarnCount struct {

    /* 总事件数 (Optional) */
    TotalCount int `json:"totalCount"`

    /* 弱口令事件数 (Optional) */
    WeakCount int `json:"weakCount"`

    /* 暴力破解事件数 (Optional) */
    CrackCount int `json:"crackCount"`

    /* 远程登录事件数 (Optional) */
    RemoteCount int `json:"remoteCount"`

    /* webShell事件数 (Optional) */
    WebShellCount int `json:"webShellCount"`
}