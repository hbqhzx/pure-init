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


type Identity struct {

    /* type (Optional) */
    Type string `json:"type"`

    /* principal (Optional) */
    Principal string `json:"principal"`

    /* erpPrincipal (Optional) */
    ErpPrincipal string `json:"erpPrincipal"`

    /* account (Optional) */
    Account string `json:"account"`

    /* previousPrincipal (Optional) */
    PreviousPrincipal string `json:"previousPrincipal"`

    /* invokedBy (Optional) */
    InvokedBy string `json:"invokedBy"`

    /* mfa (Optional) */
    Mfa string `json:"mfa"`
}