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


type UserConfig struct {

    /* 自定义规则ID (Optional) */
    Id *string `json:"id"`

    /* 所属的Waf实例ID (Optional) */
    WafId *string `json:"wafId"`

    /* 防护优先级(数值越小优先级越高，整数取值0-10000)，同一waf实例下所有自定义规则优先级不能相同。  */
    Weight int `json:"weight"`

    /* 规则名称，名称不可为空，只支持中文、数字、大小写字母、英文下划线“_”及中划线“-”，且不能超过32字符。  */
    Name string `json:"name"`

    /* ID编号 (Optional) */
    UserConfigId *int `json:"userConfigId"`

    /* 动作, 0:放行;1:检测模式;2:拦截模式;  */
    Action int `json:"action"`

    /* 开关, 0:关闭;1:开启  */
    Status int `json:"status"`

    /* 匹配字段,可取值(USER_AGENT,REFERER,URI,REMOTE_ADDR,COOKIE,REQUEST_URI,REQUEST_BODY)  */
    Position string `json:"position"`

    /* 逻辑符,y:包含;n:不包含  */
    Logic string `json:"logic"`

    /* 匹配内容,不支持正则,不能包含(<>"')特殊字符,最大长度1024
当position为REMOTE_ADDR时,匹配内容应为一个有效的IPV4地址。如 192.168.1.1
当position为REQUEST_URI时,匹配内容应为一个完整的URL地址,且协议只能为http或https。如 https://www.jdcloud.com/
当position为URI时,匹配内容应为一个有效的URL中PATH的一部分。如 /a.html、/abc/ 、abc/
  */
    MatchStr string `json:"matchStr"`

    /* 创建时间 (Optional) */
    CreateTime *string `json:"createTime"`
}
