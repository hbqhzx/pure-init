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


type Credentials struct {

    /* 临时accessKey (Optional) */
    AccessKey string `json:"accessKey"`

    /* 临时secretKey (Optional) */
    SecretKey string `json:"secretKey"`

    /* 临时安全令牌 (Optional) */
    SessionToken string `json:"sessionToken"`

    /* 有效期 (Optional) */
    Expiration string `json:"expiration"`
}
