package netana

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeNetworkAnalyticsDataTransfer invokes the netana.DescribeNetworkAnalyticsDataTransfer API synchronously
// api document: https://help.aliyun.com/api/netana/describenetworkanalyticsdatatransfer.html
func (client *Client) DescribeNetworkAnalyticsDataTransfer(request *DescribeNetworkAnalyticsDataTransferRequest) (response *DescribeNetworkAnalyticsDataTransferResponse, err error) {
	response = CreateDescribeNetworkAnalyticsDataTransferResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeNetworkAnalyticsDataTransferWithChan invokes the netana.DescribeNetworkAnalyticsDataTransfer API asynchronously
// api document: https://help.aliyun.com/api/netana/describenetworkanalyticsdatatransfer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNetworkAnalyticsDataTransferWithChan(request *DescribeNetworkAnalyticsDataTransferRequest) (<-chan *DescribeNetworkAnalyticsDataTransferResponse, <-chan error) {
	responseChan := make(chan *DescribeNetworkAnalyticsDataTransferResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNetworkAnalyticsDataTransfer(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeNetworkAnalyticsDataTransferWithCallback invokes the netana.DescribeNetworkAnalyticsDataTransfer API asynchronously
// api document: https://help.aliyun.com/api/netana/describenetworkanalyticsdatatransfer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeNetworkAnalyticsDataTransferWithCallback(request *DescribeNetworkAnalyticsDataTransferRequest, callback func(response *DescribeNetworkAnalyticsDataTransferResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNetworkAnalyticsDataTransferResponse
		var err error
		defer close(result)
		response, err = client.DescribeNetworkAnalyticsDataTransfer(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeNetworkAnalyticsDataTransferRequest is the request struct for api DescribeNetworkAnalyticsDataTransfer
type DescribeNetworkAnalyticsDataTransferRequest struct {
	*requests.RpcRequest
	Country              string           `position:"Query" name:"Country"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Product              string           `position:"Query" name:"Product"`
	Period               string           `position:"Query" name:"Period"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Ip                   string           `position:"Query" name:"Ip"`
	EndTime              string           `position:"Query" name:"EndTime"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           string           `position:"Query" name:"PageNumber"`
	Carrier              string           `position:"Query" name:"Carrier"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	Province             string           `position:"Query" name:"Province"`
	InternetChargeType   string           `position:"Query" name:"InternetChargeType"`
	Grade                string           `position:"Query" name:"Grade"`
	PageSize             string           `position:"Query" name:"PageSize"`
	Direction            string           `position:"Query" name:"Direction"`
}

// DescribeNetworkAnalyticsDataTransferResponse is the response struct for api DescribeNetworkAnalyticsDataTransfer
type DescribeNetworkAnalyticsDataTransferResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	TotalCount        int               `json:"TotalCount" xml:"TotalCount"`
	PageNumber        int               `json:"PageNumber" xml:"PageNumber"`
	PageSize          int               `json:"PageSize" xml:"PageSize"`
	DataTransferInfos DataTransferInfos `json:"DataTransferInfos" xml:"DataTransferInfos"`
}

// CreateDescribeNetworkAnalyticsDataTransferRequest creates a request to invoke DescribeNetworkAnalyticsDataTransfer API
func CreateDescribeNetworkAnalyticsDataTransferRequest() (request *DescribeNetworkAnalyticsDataTransferRequest) {
	request = &DescribeNetworkAnalyticsDataTransferRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Netana", "2018-10-18", "DescribeNetworkAnalyticsDataTransfer", "Netana", "openAPI")
	return
}

// CreateDescribeNetworkAnalyticsDataTransferResponse creates a response to parse from DescribeNetworkAnalyticsDataTransfer response
func CreateDescribeNetworkAnalyticsDataTransferResponse() (response *DescribeNetworkAnalyticsDataTransferResponse) {
	response = &DescribeNetworkAnalyticsDataTransferResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
