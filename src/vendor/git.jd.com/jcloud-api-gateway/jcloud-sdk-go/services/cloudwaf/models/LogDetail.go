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


type LogDetail struct {

    /* 时间，例如 2018-08-08 18:14:17 (Optional) */
    Time string `json:"time"`

    /* 客户端地址 (Optional) */
    Remote_addr string `json:"remote_addr"`

    /* 请求的域名 (Optional) */
    Domain string `json:"domain"`

    /* Referer (Optional) */
    Http_referer string `json:"http_referer"`

    /* 客户端的UserAgent (Optional) */
    Http_user_agent string `json:"http_user_agent"`

    /* 状态码 (Optional) */
    Status string `json:"status"`

    /* 请求行 (Optional) */
    Req_line string `json:"req_line"`

    /* 请求头 (Optional) */
    Req_headers string `json:"req_headers"`

    /* 请求body (Optional) */
    Req_body string `json:"req_body"`

    /* 对应请求中的日志类型 (Optional) */
    Log_type string `json:"log_type"`

    /* 请求的Uri (Optional) */
    Document_uri string `json:"document_uri"`

    /* 执行动作，例如 'WAF攻击 拦截' (Optional) */
    Action string `json:"action"`

    /* 回源ip (Optional) */
    Upstream_ip string `json:"upstream_ip"`

    /* 回源时间，单位毫秒 (Optional) */
    Upstream_time string `json:"upstream_time"`

    /* 请求内容 (Optional) */
    Req_content string `json:"req_content"`

    /* x-forwarded-for (Optional) */
    X_forwarded_for string `json:"x_forwarded_for"`

    /* 请求cookie (Optional) */
    Cookie string `json:"cookie"`
}