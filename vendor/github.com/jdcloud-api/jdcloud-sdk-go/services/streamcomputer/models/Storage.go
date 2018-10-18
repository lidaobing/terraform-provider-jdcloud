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


type Storage struct {

    /*  (Optional) */
    Id *int `json:"id"`

    /*  (Optional) */
    Name *string `json:"name"`

    /*  (Optional) */
    Uid *string `json:"uid"`

    /*  (Optional) */
    Type *string `json:"type"`

    /* 这个参数有input和ouput两个可选值，取决于创建输入还是输出 (Optional) */
    StorageType *string `json:"storageType"`

    /*  (Optional) */
    UserName *string `json:"userName"`

    /*  (Optional) */
    CreateTime *string `json:"createTime"`

    /*  (Optional) */
    UpdateTime *string `json:"updateTime"`

    /*  (Optional) */
    NamespaceId *string `json:"namespaceId"`

    /*  (Optional) */
    Deleted *string `json:"deleted"`

    /* Storage的具体参数。<br>1、创建源类型为流式数据输入时，则需要传输source，topicName，duration，format，delimiter，schema 。<br> 2、创建输出如果输出位置为流数据总线，需要传topicName，format，delimiter，ossFlag，bucketName，directory，objectName。<br>3、创建输出如果输出位置为数据计算服务，则需要传输database，table，ossFlag，bucketName，directory，objectName。 (Optional) */
    StorageParameterList []StorageParameter `json:"storageParameterList"`
}
