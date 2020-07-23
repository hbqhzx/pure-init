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


type Listener struct {

    /* Listener的Id (Optional) */
    ListenerId string `json:"listenerId"`

    /* Listener的名字 (Optional) */
    ListenerName string `json:"listenerName"`

    /* Listener状态, 取值为On或者为Off (Optional) */
    Status string `json:"status"`

    /* Listener所属loadBalancer的Id (Optional) */
    LoadBalancerId string `json:"loadBalancerId"`

    /* Listener所属负载均衡类型，取值为：alb、nlb (Optional) */
    LoadBalancerType string `json:"loadBalancerType"`

    /* 监听协议, 取值为Tcp, Tls, Http, Https (Optional) */
    Protocol string `json:"protocol"`

    /* 监听端口，取值范围为[1, 65535] (Optional) */
    Port int `json:"port"`

    /* 默认后端服务的转发策略,取值为Forward或Redirect, 现只支持Forward (Optional) */
    Action string `json:"action"`

    /* 默认的后端服务Id (Optional) */
    BackendId string `json:"backendId"`

    /* 【仅alb支持】转发规则组Id (Optional) */
    UrlMapId string `json:"urlMapId"`

    /* 连接保持时间, >=1，范围为[1,86400]【Tcp和Tls协议】默认为：1800s【Http和Https协议】默认为：60s (Optional) */
    ConnectionIdleTimeSeconds int `json:"connectionIdleTimeSeconds"`

    /* 【alb Https和Tls协议】ssl server证书列表，现只支持一个证书 (Optional) */
    CertificateSpecs []CertificateSpec `json:"certificateSpecs"`

    /* Listener的描述信息 (Optional) */
    Description string `json:"description"`

    /* Listener的创建时间 (Optional) */
    CreatedTime string `json:"createdTime"`
}
