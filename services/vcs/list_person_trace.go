package vcs

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

// ListPersonTrace invokes the vcs.ListPersonTrace API synchronously
// api document: https://help.aliyun.com/api/vcs/listpersontrace.html
func (client *Client) ListPersonTrace(request *ListPersonTraceRequest) (response *ListPersonTraceResponse, err error) {
	response = CreateListPersonTraceResponse()
	err = client.DoAction(request, response)
	return
}

// ListPersonTraceWithChan invokes the vcs.ListPersonTrace API asynchronously
// api document: https://help.aliyun.com/api/vcs/listpersontrace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPersonTraceWithChan(request *ListPersonTraceRequest) (<-chan *ListPersonTraceResponse, <-chan error) {
	responseChan := make(chan *ListPersonTraceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPersonTrace(request)
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

// ListPersonTraceWithCallback invokes the vcs.ListPersonTrace API asynchronously
// api document: https://help.aliyun.com/api/vcs/listpersontrace.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPersonTraceWithCallback(request *ListPersonTraceRequest, callback func(response *ListPersonTraceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPersonTraceResponse
		var err error
		defer close(result)
		response, err = client.ListPersonTrace(request)
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

// ListPersonTraceRequest is the request struct for api ListPersonTrace
type ListPersonTraceRequest struct {
	*requests.RpcRequest
	CorpId       string `position:"Body" name:"CorpId"`
	GroupId      string `position:"Body" name:"GroupId"`
	EndTime      string `position:"Body" name:"EndTime"`
	StartTime    string `position:"Body" name:"StartTime"`
	PageNumber   string `position:"Body" name:"PageNumber"`
	PageSize     string `position:"Body" name:"PageSize"`
	DataSourceId string `position:"Body" name:"DataSourceId"`
	PersonId     string `position:"Body" name:"PersonId"`
}

// ListPersonTraceResponse is the response struct for api ListPersonTrace
type ListPersonTraceResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    string `json:"Success" xml:"Success"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	Data       []Day  `json:"Data" xml:"Data"`
}

// CreateListPersonTraceRequest creates a request to invoke ListPersonTrace API
func CreateListPersonTraceRequest() (request *ListPersonTraceRequest) {
	request = &ListPersonTraceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListPersonTrace", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPersonTraceResponse creates a response to parse from ListPersonTrace response
func CreateListPersonTraceResponse() (response *ListPersonTraceResponse) {
	response = &ListPersonTraceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
