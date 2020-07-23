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


type DomainConfig struct {

    /* 域名 (Optional) */
    Domain string `json:"domain"`

    /* cname域名 (Optional) */
    Cname string `json:"cname"`

    /* 绑定的证书名称 (Optional) */
    CertName string `json:"certName"`

    /* 网站lb配置 (Optional) */
    LbConf LbConf `json:"lbConf"`

    /* 网站dns配置 (Optional) */
    DnsStatus DnsStatus `json:"dnsStatus"`

    /* 网站waf防护配置 (Optional) */
    WafConf WafConf `json:"wafConf"`

    /* 网站cc防护配置 (Optional) */
    CcConf CcConf `json:"ccConf"`

    /* 网站恶意ip防护配置 (Optional) */
    IpbanConf IpbanConf `json:"ipbanConf"`

    /* 网站防爬虫防护配置 (Optional) */
    AntispiderConf AntispiderConf `json:"antispiderConf"`
}
