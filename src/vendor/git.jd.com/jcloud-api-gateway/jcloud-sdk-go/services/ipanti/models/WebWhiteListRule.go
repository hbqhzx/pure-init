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


type WebWhiteListRule struct {

    /* 白名单规则 Id (Optional) */
    Id string `json:"id"`

    /* 白名单规则名称 (Optional) */
    Name string `json:"name"`

    /* 匹配 key (Optional) */
    Key string `json:"key"`

    /* 匹配 value (Optional) */
    Value string `json:"value"`

    /* 匹配模式:<br>- 0: uri<br>- 1: ip<br>- 2: cookie<br>- 3: geo<br>- 4: header (Optional) */
    Mode int `json:"mode"`

    /* 匹配规则:<br>- 0: 完全匹配<br>- 1: 前缀匹配<br>- 2: 包含<br>- 3: 正则匹配<br>- 4: 大于 (Optional) */
    Pattern int `json:"pattern"`

    /* 命中后处理动作<br>- :0: 放行<br>- 1: CC 防护 (Optional) */
    Action int `json:"action"`

    /* 规则状态<br>- :0: 关闭<br>- 1: 开启 (Optional) */
    Status int `json:"status"`
}
