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


type PasswordPolicyRule struct {

    /* 是否包含大写字母  */
    RequireUppercaseCharacters bool `json:"requireUppercaseCharacters"`

    /* 是否包含小写字母  */
    RequireLowercaseCharacters bool `json:"requireLowercaseCharacters"`

    /* 是否包含数字  */
    RequireNumbers bool `json:"requireNumbers"`

    /* 是否包含特殊字符  */
    RequireSpecialCharacters bool `json:"requireSpecialCharacters"`
}