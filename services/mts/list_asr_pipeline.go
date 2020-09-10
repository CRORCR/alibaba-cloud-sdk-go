package mts

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

// ListAsrPipeline invokes the mts.ListAsrPipeline API synchronously
// api document: https://help.aliyun.com/api/mts/listasrpipeline.html
func (client *Client) ListAsrPipeline(request *ListAsrPipelineRequest) (response *ListAsrPipelineResponse, err error) {
	response = CreateListAsrPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// ListAsrPipelineWithChan invokes the mts.ListAsrPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/listasrpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAsrPipelineWithChan(request *ListAsrPipelineRequest) (<-chan *ListAsrPipelineResponse, <-chan error) {
	responseChan := make(chan *ListAsrPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAsrPipeline(request)
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

// ListAsrPipelineWithCallback invokes the mts.ListAsrPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/listasrpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListAsrPipelineWithCallback(request *ListAsrPipelineRequest, callback func(response *ListAsrPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAsrPipelineResponse
		var err error
		defer close(result)
		response, err = client.ListAsrPipeline(request)
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

// ListAsrPipelineRequest is the request struct for api ListAsrPipeline
type ListAsrPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	State                string           `position:"Query" name:"State"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ListAsrPipelineResponse is the response struct for api ListAsrPipeline
type ListAsrPipelineResponse struct {
	*responses.BaseResponse
	RequestId    string                        `json:"RequestId" xml:"RequestId"`
	TotalCount   int64                         `json:"TotalCount" xml:"TotalCount"`
	PageNumber   int64                         `json:"PageNumber" xml:"PageNumber"`
	PageSize     int64                         `json:"PageSize" xml:"PageSize"`
	PipelineList PipelineListInListAsrPipeline `json:"PipelineList" xml:"PipelineList"`
}

// CreateListAsrPipelineRequest creates a request to invoke ListAsrPipeline API
func CreateListAsrPipelineRequest() (request *ListAsrPipelineRequest) {
	request = &ListAsrPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ListAsrPipeline", "", "")
	return
}

// CreateListAsrPipelineResponse creates a response to parse from ListAsrPipeline response
func CreateListAsrPipelineResponse() (response *ListAsrPipelineResponse) {
	response = &ListAsrPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
