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


type ClusterListViewModel struct {

    /* 用户名 (Optional) */
    UserName *string `json:"userName"`

    /* 主节点信息 (Optional) */
    MasterNodeInfo *string `json:"masterNodeInfo"`

    /* 主节点磁盘类型 (Optional) */
    MasterNodeDiskType *string `json:"masterNodeDiskType"`

    /* 主节点磁盘容量 (Optional) */
    MasterNodeDiskVolume *int `json:"masterNodeDiskVolume"`

    /* 主节点数量 (Optional) */
    MasterNodeNumber *int `json:"masterNodeNumber"`

    /* 主节点实例类型 (Optional) */
    MasterInstanceType *string `json:"masterInstanceType"`

    /* 从节点信息 (Optional) */
    SlaveNodeInfo *string `json:"slaveNodeInfo"`

    /* 从节点磁盘类型 (Optional) */
    SlaveNodeDiskType *string `json:"slaveNodeDiskType"`

    /* 从节点磁盘容量 (Optional) */
    SlaveNodeDiskVolume *int `json:"slaveNodeDiskVolume"`

    /* 从节点数量 (Optional) */
    SlaveNodeNumber *int `json:"slaveNodeNumber"`

    /* 核心实例类型 (Optional) */
    CoreInstanceType *string `json:"coreInstanceType"`

    /* 网络带宽 (Optional) */
    BandwidthOut *int `json:"bandwidthOut"`

    /* 数据中心位置 (Optional) */
    DataCenter *string `json:"dataCenter"`
}
