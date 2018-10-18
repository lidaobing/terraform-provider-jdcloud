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

type CheckRegistryNameRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 待验证的注册表名。  */
    RegistryName string `json:"registryName"`
}

/*
 * param regionId: Region ID (Required)
 * param registryName: 待验证的注册表名。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCheckRegistryNameRequest(
    regionId string,
    registryName string,
) *CheckRegistryNameRequest {

	return &CheckRegistryNameRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/registries:checkRegistryName",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        RegistryName: registryName,
	}
}

/*
 * param regionId: Region ID (Required)
 * param registryName: 待验证的注册表名。 (Required)
 */
func NewCheckRegistryNameRequestWithAllParams(
    regionId string,
    registryName string,
) *CheckRegistryNameRequest {

    return &CheckRegistryNameRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries:checkRegistryName",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        RegistryName: registryName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCheckRegistryNameRequestWithoutParam() *CheckRegistryNameRequest {

    return &CheckRegistryNameRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/registries:checkRegistryName",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CheckRegistryNameRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param registryName: 待验证的注册表名。(Required) */
func (r *CheckRegistryNameRequest) SetRegistryName(registryName string) {
    r.RegistryName = registryName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CheckRegistryNameRequest) GetRegionId() string {
    return r.RegionId
}

type CheckRegistryNameResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CheckRegistryNameResult `json:"result"`
}

type CheckRegistryNameResult struct {
    Code int `json:"code"`
    Reason string `json:"reason"`
}