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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type SetBackupPolicyRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* RDS 实例ID，唯一标识一个RDS实例  */
    InstanceId string `json:"instanceId"`

    /* 自动备份开始时间窗口,例如：00:00-01:00，表示0点到1点开始进行数据库自动备份，备份完成时间则跟实例大小有关，不一定在这个时间范围中<br>SQL Server:范围00:00-23:59，时间范围差不得小于30分钟。<br>MySQL,只能是以下取值:<br>00:00-01:00<br>01:00-02:00<br>......<br>23:00-24:00 (Optional) */
    StartWindow *string `json:""`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewSetBackupPolicyRequest(
    regionId string,
    instanceId string,
) *SetBackupPolicyRequest {

	return &SetBackupPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:setBackupPolicy",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param instanceId: RDS 实例ID，唯一标识一个RDS实例 (Required)
 * param startWindow: 自动备份开始时间窗口,例如：00:00-01:00，表示0点到1点开始进行数据库自动备份，备份完成时间则跟实例大小有关，不一定在这个时间范围中<br>SQL Server:范围00:00-23:59，时间范围差不得小于30分钟。<br>MySQL,只能是以下取值:<br>00:00-01:00<br>01:00-02:00<br>......<br>23:00-24:00 (Optional)
 */
func NewSetBackupPolicyRequestWithAllParams(
    regionId string,
    instanceId string,
    startWindow *string,
) *SetBackupPolicyRequest {

    return &SetBackupPolicyRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:setBackupPolicy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        StartWindow: startWindow,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewSetBackupPolicyRequestWithoutParam() *SetBackupPolicyRequest {

    return &SetBackupPolicyRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:setBackupPolicy",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *SetBackupPolicyRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: RDS 实例ID，唯一标识一个RDS实例(Required) */
func (r *SetBackupPolicyRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param startWindow: 自动备份开始时间窗口,例如：00:00-01:00，表示0点到1点开始进行数据库自动备份，备份完成时间则跟实例大小有关，不一定在这个时间范围中<br>SQL Server:范围00:00-23:59，时间范围差不得小于30分钟。<br>MySQL,只能是以下取值:<br>00:00-01:00<br>01:00-02:00<br>......<br>23:00-24:00(Optional) */
func (r *SetBackupPolicyRequest) SetStartWindow(startWindow string) {
    r.StartWindow = &startWindow
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r SetBackupPolicyRequest) GetRegionId() string {
    return r.RegionId
}

type SetBackupPolicyResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result SetBackupPolicyResult `json:"result"`
}

type SetBackupPolicyResult struct {
}