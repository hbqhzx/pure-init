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


type TagsResourcesInfo struct {

    /*  (Optional) */
    TagKey string `json:"tagKey"`

    /*  (Optional) */
    TagValue string `json:"tagValue"`

    /*  (Optional) */
    VmResourceCount int64 `json:"vmResourceCount"`

    /*  (Optional) */
    DiskResourceCount int64 `json:"diskResourceCount"`

    /*  (Optional) */
    SqlServerResourceCount int64 `json:"sqlServerResourceCount"`

    /*  (Optional) */
    MongodbResourceCount int64 `json:"mongodbResourceCount"`

    /*  (Optional) */
    IpResourceCount int64 `json:"ipResourceCount"`

    /*  (Optional) */
    EsResourceCount int64 `json:"esResourceCount"`

    /*  (Optional) */
    JcqResourceCount int64 `json:"jcqResourceCount"`

    /*  (Optional) */
    DrdsResourceCount int64 `json:"drdsResourceCount"`

    /*  (Optional) */
    MemcachedResourceCount int64 `json:"memcachedResourceCount"`

    /*  (Optional) */
    RedisResourceCount int64 `json:"redisResourceCount"`

    /*  (Optional) */
    DatabaseResourceCount int64 `json:"databaseResourceCount"`

    /*  (Optional) */
    DbRoResourceCount int64 `json:"dbRoResourceCount"`

    /*  (Optional) */
    PerconaResourceCount int64 `json:"perconaResourceCount"`

    /*  (Optional) */
    MariadbResourceCount int64 `json:"mariadbResourceCount"`

    /*  (Optional) */
    MariadbRoResourceCount int64 `json:"mariadbRoResourceCount"`

    /*  (Optional) */
    PgResourceCount int64 `json:"pgResourceCount"`
}
