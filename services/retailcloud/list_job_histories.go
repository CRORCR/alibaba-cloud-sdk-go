package retailcloud

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

// ListJobHistories invokes the retailcloud.ListJobHistories API synchronously
// api document: https://help.aliyun.com/api/retailcloud/listjobhistories.html
func (client *Client) ListJobHistories(request *ListJobHistoriesRequest) (response *ListJobHistoriesResponse, err error) {
	response = CreateListJobHistoriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobHistoriesWithChan invokes the retailcloud.ListJobHistories API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listjobhistories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobHistoriesWithChan(request *ListJobHistoriesRequest) (<-chan *ListJobHistoriesResponse, <-chan error) {
	responseChan := make(chan *ListJobHistoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobHistories(request)
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

// ListJobHistoriesWithCallback invokes the retailcloud.ListJobHistories API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listjobhistories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobHistoriesWithCallback(request *ListJobHistoriesRequest, callback func(response *ListJobHistoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobHistoriesResponse
		var err error
		defer close(result)
		response, err = client.ListJobHistories(request)
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

// ListJobHistoriesRequest is the request struct for api ListJobHistories
type ListJobHistoriesRequest struct {
	*requests.RpcRequest
	AppId      requests.Integer `position:"Query" name:"AppId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	EnvId      requests.Integer `position:"Query" name:"EnvId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Status     string           `position:"Query" name:"Status"`
}

// ListJobHistoriesResponse is the response struct for api ListJobHistories
type ListJobHistoriesResponse struct {
	*responses.BaseResponse
	Code       int         `json:"Code" xml:"Code"`
	PageNumber int         `json:"PageNumber" xml:"PageNumber"`
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	PageSize   int         `json:"PageSize" xml:"PageSize"`
	TotalCount int64       `json:"TotalCount" xml:"TotalCount"`
	ErrorMsg   string      `json:"ErrorMsg" xml:"ErrorMsg"`
	Data       []JobDetail `json:"Data" xml:"Data"`
}

// CreateListJobHistoriesRequest creates a request to invoke ListJobHistories API
func CreateListJobHistoriesRequest() (request *ListJobHistoriesRequest) {
	request = &ListJobHistoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListJobHistories", "", "")
	request.Method = requests.GET
	return
}

// CreateListJobHistoriesResponse creates a response to parse from ListJobHistories response
func CreateListJobHistoriesResponse() (response *ListJobHistoriesResponse) {
	response = &ListJobHistoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
