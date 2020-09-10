package outboundbot

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

// ListJobGroups invokes the outboundbot.ListJobGroups API synchronously
// api document: https://help.aliyun.com/api/outboundbot/listjobgroups.html
func (client *Client) ListJobGroups(request *ListJobGroupsRequest) (response *ListJobGroupsResponse, err error) {
	response = CreateListJobGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListJobGroupsWithChan invokes the outboundbot.ListJobGroups API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/listjobgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobGroupsWithChan(request *ListJobGroupsRequest) (<-chan *ListJobGroupsResponse, <-chan error) {
	responseChan := make(chan *ListJobGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListJobGroups(request)
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

// ListJobGroupsWithCallback invokes the outboundbot.ListJobGroups API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/listjobgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListJobGroupsWithCallback(request *ListJobGroupsRequest, callback func(response *ListJobGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListJobGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListJobGroups(request)
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

// ListJobGroupsRequest is the request struct for api ListJobGroups
type ListJobGroupsRequest struct {
	*requests.RpcRequest
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListJobGroupsResponse is the response struct for api ListJobGroups
type ListJobGroupsResponse struct {
	*responses.BaseResponse
	Code           string    `json:"Code" xml:"Code"`
	HttpStatusCode int       `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string    `json:"Message" xml:"Message"`
	RequestId      string    `json:"RequestId" xml:"RequestId"`
	Success        bool      `json:"Success" xml:"Success"`
	JobGroups      JobGroups `json:"JobGroups" xml:"JobGroups"`
}

// CreateListJobGroupsRequest creates a request to invoke ListJobGroups API
func CreateListJobGroupsRequest() (request *ListJobGroupsRequest) {
	request = &ListJobGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ListJobGroups", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListJobGroupsResponse creates a response to parse from ListJobGroups response
func CreateListJobGroupsResponse() (response *ListJobGroupsResponse) {
	response = &ListJobGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
