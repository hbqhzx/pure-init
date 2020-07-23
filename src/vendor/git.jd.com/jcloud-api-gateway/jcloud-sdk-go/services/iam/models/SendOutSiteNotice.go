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


type SendOutSiteNotice struct {

    /* 用户pin  */
    Pin string `json:"pin"`

    /* 邮件标题 (Optional) */
    EmailSubject *string `json:"emailSubject"`

    /* 邮件内容 (Optional) */
    EmailContent *string `json:"emailContent"`

    /* 消息类型  */
    NotifyBusinessTypeEnum string `json:"notifyBusinessTypeEnum"`

    /* 模版code  */
    TemplateId int `json:"templateId"`

    /* 模版参数  */
    TemplateParam []string `json:"templateParam"`

    /* 业务编码(和产品申请)  */
    SmsMessageSource string `json:"smsMessageSource"`
}