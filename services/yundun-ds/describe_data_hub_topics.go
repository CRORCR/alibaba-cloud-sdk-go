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

// DescribeDataHubTopics invokes the yundun_ds.DescribeDataHubTopics API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/describedatahubtopics.html
func (client *Client) DescribeDataHubTopics(request *DescribeDataHubTopicsRequest) (response *DescribeDataHubTopicsResponse, err error) {
	response = CreateDescribeDataHubTopicsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataHubTopicsWithChan invokes the yundun_ds.DescribeDataHubTopics API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describedatahubtopics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataHubTopicsWithChan(request *DescribeDataHubTopicsRequest) (<-chan *DescribeDataHubTopicsResponse, <-chan error) {
	responseChan := make(chan *DescribeDataHubTopicsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataHubTopics(request)
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

// DescribeDataHubTopicsWithCallback invokes the yundun_ds.DescribeDataHubTopics API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describedatahubtopics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDataHubTopicsWithCallback(request *DescribeDataHubTopicsRequest, callback func(response *DescribeDataHubTopicsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataHubTopicsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataHubTopics(request)
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

// DescribeDataHubTopicsRequest is the request struct for api DescribeDataHubTopics
type DescribeDataHubTopicsRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	DepartId    requests.Integer `position:"Query" name:"DepartId"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Lang        string           `position:"Query" name:"Lang"`
	ProjectId   requests.Integer `position:"Query" name:"ProjectId"`
	Key         string           `position:"Query" name:"Key"`
}

// DescribeDataHubTopicsResponse is the response struct for api DescribeDataHubTopics
type DescribeDataHubTopicsResponse struct {
	*responses.BaseResponse
	RequestId   string  `json:"RequestId" xml:"RequestId"`
	PageSize    int     `json:"PageSize" xml:"PageSize"`
	CurrentPage int     `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int     `json:"TotalCount" xml:"TotalCount"`
	Items       []Topic `json:"Items" xml:"Items"`
}

// CreateDescribeDataHubTopicsRequest creates a request to invoke DescribeDataHubTopics API
func CreateDescribeDataHubTopicsRequest() (request *DescribeDataHubTopicsRequest) {
	request = &DescribeDataHubTopicsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "DescribeDataHubTopics", "sddp", "openAPI")
	return
}

// CreateDescribeDataHubTopicsResponse creates a response to parse from DescribeDataHubTopics response
func CreateDescribeDataHubTopicsResponse() (response *DescribeDataHubTopicsResponse) {
	response = &DescribeDataHubTopicsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
