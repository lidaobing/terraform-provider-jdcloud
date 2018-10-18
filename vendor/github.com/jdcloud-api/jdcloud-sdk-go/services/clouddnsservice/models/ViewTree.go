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


type ViewTree struct {

    /* 此解析线路是否禁用 (Optional) */
    Disabled bool `json:"disabled"`

    /* 解析线路的名称 (Optional) */
    Label string `json:"label"`

    /* 此数据是否是叶子节点 (Optional) */
    Leaf bool `json:"leaf"`

    /* 解析线路ID (Optional) */
    Value int `json:"value"`

    /*  (Optional) */
    Children []ViewTree `json:"children"`
}
