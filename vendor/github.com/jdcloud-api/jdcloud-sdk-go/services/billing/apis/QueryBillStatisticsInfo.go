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

type QueryBillStatisticsInfoRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* appCode  */
    AppCode string `json:"appCode"`

    /* serviceCode  */
    ServiceCode string `json:"serviceCode"`

    /* 支付状态  */
    PayState int `json:"payState"`

    /* 时间类型  */
    TimeType int `json:"timeType"`

    /* 开始时间  */
    StartTime string `json:"startTime"`

    /* 结束时间  */
    EndTime string `json:"endTime"`

    /* 查询类型  */
    QueryType int `json:"queryType"`

    /* 支付类型  */
    PayType int `json:"payType"`

    /* 计费类型  */
    BillingType int `json:"billingType"`
}

/*
 * param regionId:  (Required)
 * param appCode: appCode (Required)
 * param serviceCode: serviceCode (Required)
 * param payState: 支付状态 (Required)
 * param timeType: 时间类型 (Required)
 * param startTime: 开始时间 (Required)
 * param endTime: 结束时间 (Required)
 * param queryType: 查询类型 (Required)
 * param payType: 支付类型 (Required)
 * param billingType: 计费类型 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryBillStatisticsInfoRequest(
    regionId string,
    appCode string,
    serviceCode string,
    payState int,
    timeType int,
    startTime string,
    endTime string,
    queryType int,
    payType int,
    billingType int,
) *QueryBillStatisticsInfoRequest {

	return &QueryBillStatisticsInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/billStatisticsInfo",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        AppCode: appCode,
        ServiceCode: serviceCode,
        PayState: payState,
        TimeType: timeType,
        StartTime: startTime,
        EndTime: endTime,
        QueryType: queryType,
        PayType: payType,
        BillingType: billingType,
	}
}

/*
 * param regionId:  (Required)
 * param appCode: appCode (Required)
 * param serviceCode: serviceCode (Required)
 * param payState: 支付状态 (Required)
 * param timeType: 时间类型 (Required)
 * param startTime: 开始时间 (Required)
 * param endTime: 结束时间 (Required)
 * param queryType: 查询类型 (Required)
 * param payType: 支付类型 (Required)
 * param billingType: 计费类型 (Required)
 */
func NewQueryBillStatisticsInfoRequestWithAllParams(
    regionId string,
    appCode string,
    serviceCode string,
    payState int,
    timeType int,
    startTime string,
    endTime string,
    queryType int,
    payType int,
    billingType int,
) *QueryBillStatisticsInfoRequest {

    return &QueryBillStatisticsInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/billStatisticsInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        AppCode: appCode,
        ServiceCode: serviceCode,
        PayState: payState,
        TimeType: timeType,
        StartTime: startTime,
        EndTime: endTime,
        QueryType: queryType,
        PayType: payType,
        BillingType: billingType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryBillStatisticsInfoRequestWithoutParam() *QueryBillStatisticsInfoRequest {

    return &QueryBillStatisticsInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/billStatisticsInfo",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *QueryBillStatisticsInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appCode: appCode(Required) */
func (r *QueryBillStatisticsInfoRequest) SetAppCode(appCode string) {
    r.AppCode = appCode
}

/* param serviceCode: serviceCode(Required) */
func (r *QueryBillStatisticsInfoRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = serviceCode
}

/* param payState: 支付状态(Required) */
func (r *QueryBillStatisticsInfoRequest) SetPayState(payState int) {
    r.PayState = payState
}

/* param timeType: 时间类型(Required) */
func (r *QueryBillStatisticsInfoRequest) SetTimeType(timeType int) {
    r.TimeType = timeType
}

/* param startTime: 开始时间(Required) */
func (r *QueryBillStatisticsInfoRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 结束时间(Required) */
func (r *QueryBillStatisticsInfoRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param queryType: 查询类型(Required) */
func (r *QueryBillStatisticsInfoRequest) SetQueryType(queryType int) {
    r.QueryType = queryType
}

/* param payType: 支付类型(Required) */
func (r *QueryBillStatisticsInfoRequest) SetPayType(payType int) {
    r.PayType = payType
}

/* param billingType: 计费类型(Required) */
func (r *QueryBillStatisticsInfoRequest) SetBillingType(billingType int) {
    r.BillingType = billingType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryBillStatisticsInfoRequest) GetRegionId() string {
    return r.RegionId
}

type QueryBillStatisticsInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryBillStatisticsInfoResult `json:"result"`
}

type QueryBillStatisticsInfoResult struct {
    TotalFee int `json:"totalFee"`
    CashPayFee int `json:"cashPayFee"`
    CashCouponPayFee int `json:"cashCouponPayFee"`
    BalancePayFee int `json:"balancePayFee"`
    ArrearFee int `json:"arrearFee"`
    BillFee int `json:"billFee"`
    DiscountFee int `json:"discountFee"`
}