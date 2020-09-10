package vod

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

// DescribeVodRefreshTasks invokes the vod.DescribeVodRefreshTasks API synchronously
func (client *Client) DescribeVodRefreshTasks(request *DescribeVodRefreshTasksRequest) (response *DescribeVodRefreshTasksResponse, err error) {
	response = CreateDescribeVodRefreshTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodRefreshTasksWithChan invokes the vod.DescribeVodRefreshTasks API asynchronously
func (client *Client) DescribeVodRefreshTasksWithChan(request *DescribeVodRefreshTasksRequest) (<-chan *DescribeVodRefreshTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeVodRefreshTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodRefreshTasks(request)
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

// DescribeVodRefreshTasksWithCallback invokes the vod.DescribeVodRefreshTasks API asynchronously
func (client *Client) DescribeVodRefreshTasksWithCallback(request *DescribeVodRefreshTasksRequest, callback func(response *DescribeVodRefreshTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodRefreshTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodRefreshTasks(request)
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

// DescribeVodRefreshTasksRequest is the request struct for api DescribeVodRefreshTasks
type DescribeVodRefreshTasksRequest struct {
	*requests.RpcRequest
	ObjectPath      string           `position:"Query" name:"ObjectPath"`
	StartTime       string           `position:"Query" name:"StartTime"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ObjectType      string           `position:"Query" name:"ObjectType"`
	TaskId          string           `position:"Query" name:"TaskId"`
	DomainName      string           `position:"Query" name:"DomainName"`
	EndTime         string           `position:"Query" name:"EndTime"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	Status          string           `position:"Query" name:"Status"`
}

// DescribeVodRefreshTasksResponse is the response struct for api DescribeVodRefreshTasks
type DescribeVodRefreshTasksResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	PageNumber int64  `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64  `json:"PageSize" xml:"PageSize"`
	TotalCount int64  `json:"TotalCount" xml:"TotalCount"`
	Tasks      Tasks  `json:"Tasks" xml:"Tasks"`
}

// CreateDescribeVodRefreshTasksRequest creates a request to invoke DescribeVodRefreshTasks API
func CreateDescribeVodRefreshTasksRequest() (request *DescribeVodRefreshTasksRequest) {
	request = &DescribeVodRefreshTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodRefreshTasks", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodRefreshTasksResponse creates a response to parse from DescribeVodRefreshTasks response
func CreateDescribeVodRefreshTasksResponse() (response *DescribeVodRefreshTasksResponse) {
	response = &DescribeVodRefreshTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
