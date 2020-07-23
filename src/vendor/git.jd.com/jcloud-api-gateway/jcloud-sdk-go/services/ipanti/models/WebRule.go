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


type WebRule struct {

    /* 规则 Id (Optional) */
    Id int64 `json:"id"`

    /* 实例 Id (Optional) */
    InstanceId int64 `json:"instanceId"`

    /* 子域名 (Optional) */
    Domain string `json:"domain"`

    /* 规则的 cname (Optional) */
    Cname string `json:"cname"`

    /*  (Optional) */
    Protocol WebRuleProtocol `json:"protocol"`

    /* 是否为自定义端口号, 0: 为默认, 1: 为自定义 (Optional) */
    CustomPortStatus int `json:"customPortStatus"`

    /* HTTP 协议的端口号, 如 80,81 (Optional) */
    Port []int `json:"port"`

    /* HTTPS 协议的端口号, 如 443,8443 (Optional) */
    HttpsPort []int `json:"httpsPort"`

    /* 是否开启 http 回源, 0: 为不开启, 1: 为开启, 当勾选 HTTPS 时可以配置该属性 (Optional) */
    HttpOrigin int `json:"httpOrigin"`

    /* 0: 防御状态, 1: 回源状态 (Optional) */
    Status int `json:"status"`

    /* 回源类型: A 或者 CNAME (Optional) */
    OriginType string `json:"originType"`

    /*  (Optional) */
    OriginAddr []OriginAddrItem `json:"originAddr"`

    /* 回源域名, originType 为 CNAME 时返回该字段 (Optional) */
    OriginDomain string `json:"originDomain"`

    /*  (Optional) */
    OnlineAddr []string `json:"onlineAddr"`

    /* 证书状态, 0: 异常, 1: 正常, 2: 证书未上传 (Optional) */
    HttpCertStatus int `json:"httpCertStatus"`

    /* 证书 Id (Optional) */
    CertId int64 `json:"certId"`

    /* 证书名称 (Optional) */
    CertName string `json:"certName"`

    /* 证书内容 (Optional) */
    HttpsCertContent string `json:"httpsCertContent"`

    /* 证书私钥 (Optional) */
    HttpsRsaKey string `json:"httpsRsaKey"`

    /* 是否开启https强制跳转, 当 protocol 为 HTTP_HTTPS 时可以配置该属性
  - 0 不强跳
  - 1 开启强跳
 (Optional) */
    ForceJump int `json:"forceJump"`

    /* 转发规则,  wrr: 带权重的轮询, rr: 不带权重的轮询 (Optional) */
    Algorithm string `json:"algorithm"`

    /* CC 状态, 0: CC 关闭, 1: CC 开启 (Optional) */
    CcStatus int `json:"ccStatus"`

    /* webSocketStatus, 0: 关闭, 1: 开启 (Optional) */
    WebSocketStatus int `json:"webSocketStatus"`
}
