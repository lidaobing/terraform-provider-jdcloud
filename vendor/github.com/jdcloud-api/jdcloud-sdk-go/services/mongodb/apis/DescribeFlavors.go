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
    mongodb "github.com/jdcloud-api/jdcloud-sdk-go/services/mongodb/models"
)

type DescribeFlavorsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeFlavorsRequest(
    regionId string,
) *DescribeFlavorsRequest {

	return &DescribeFlavorsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/flavors",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 */
func NewDescribeFlavorsRequestWithAllParams(
    regionId string,
) *DescribeFlavorsRequest {

    return &DescribeFlavorsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/flavors",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeFlavorsRequestWithoutParam() *DescribeFlavorsRequest {

    return &DescribeFlavorsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/flavors",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeFlavorsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeFlavorsRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeFlavorsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeFlavorsResult `json:"result"`
}

type DescribeFlavorsResult struct {
    Flavors []mongodb.Flavor `json:"flavors"`
}