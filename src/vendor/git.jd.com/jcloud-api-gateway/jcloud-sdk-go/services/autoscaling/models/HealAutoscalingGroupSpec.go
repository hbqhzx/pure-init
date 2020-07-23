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


type HealAutoscalingGroupSpec struct {

    /*  (Optional) */
    AsConfigId string `json:"asConfigId"`

    /*  (Optional) */
    Billing string `json:"billing"`

    /*  (Optional) */
    Lbs []Lb `json:"lbs"`

    /*  (Optional) */
    SecurityGroups string `json:"securityGroups"`

    /*  (Optional) */
    SubnetIds string `json:"subnetIds"`

    /*  (Optional) */
    VpcId string `json:"vpcId"`
}
