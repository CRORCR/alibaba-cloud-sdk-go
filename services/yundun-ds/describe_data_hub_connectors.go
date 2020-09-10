package yundun_ds

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

// DescribeDataHubConnectors invokes the yundun_ds.DescribeDataHubConnectors API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/describedatahubconnectors.html
func (client *Client) DescribeDataHubConnectors(request *DescribeDataHubConnectorsRequest) (response *DescribeDataHubConnectorsResponse, err error) {
	response = CreateDescribeDataHubConnectorsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataHubConnectorsWithChan invokes the yundun_ds.DescribeDataHubConnectors API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describedatahubconnectors.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataHubConnectorsWithChan(request *DescribeDataHubConnectorsRequest) (<-chan *DescribeDataHubConnectorsResponse, <-chan error) {
	responseChan := make(chan *DescribeDataHubConnectorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataHubConnectors(request)
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

// DescribeDataHubConnectorsWithCallback invokes the yundun_ds.DescribeDataHubConnectors API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describedatahubconnectors.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataHubConnectorsWithCallback(request *DescribeDataHubConnectorsRequest, callback func(response *DescribeDataHubConnectorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataHubConnectorsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataHubConnectors(request)
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

// DescribeDataHubConnectorsRequest is the request struct for api DescribeDataHubConnectors
type DescribeDataHubConnectorsRequest struct {
	*requests.RpcRequest
	TopicId     requests.Integer `position:"Query" name:"TopicId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	DepartId    requests.Integer `position:"Query" name:"DepartId"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Lang        string           `position:"Query" name:"Lang"`
	ProjectId   requests.Integer `position:"Query" name:"ProjectId"`
	Key         string           `position:"Query" name:"Key"`
}

// DescribeDataHubConnectorsResponse is the response struct for api DescribeDataHubConnectors
type DescribeDataHubConnectorsResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	CurrentPage int         `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	Items       []Connector `json:"Items" xml:"Items"`
}

// CreateDescribeDataHubConnectorsRequest creates a request to invoke DescribeDataHubConnectors API
func CreateDescribeDataHubConnectorsRequest() (request *DescribeDataHubConnectorsRequest) {
	request = &DescribeDataHubConnectorsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "DescribeDataHubConnectors", "sddp", "openAPI")
	return
}

// CreateDescribeDataHubConnectorsResponse creates a response to parse from DescribeDataHubConnectors response
func CreateDescribeDataHubConnectorsResponse() (response *DescribeDataHubConnectorsResponse) {
	response = &DescribeDataHubConnectorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
