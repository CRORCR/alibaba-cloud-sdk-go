package emr

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

// ListFlowNodeInstanceContainerStatus invokes the emr.ListFlowNodeInstanceContainerStatus API synchronously
// api document: https://help.aliyun.com/api/emr/listflownodeinstancecontainerstatus.html
func (client *Client) ListFlowNodeInstanceContainerStatus(request *ListFlowNodeInstanceContainerStatusRequest) (response *ListFlowNodeInstanceContainerStatusResponse, err error) {
	response = CreateListFlowNodeInstanceContainerStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlowNodeInstanceContainerStatusWithChan invokes the emr.ListFlowNodeInstanceContainerStatus API asynchronously
// api document: https://help.aliyun.com/api/emr/listflownodeinstancecontainerstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFlowNodeInstanceContainerStatusWithChan(request *ListFlowNodeInstanceContainerStatusRequest) (<-chan *ListFlowNodeInstanceContainerStatusResponse, <-chan error) {
	responseChan := make(chan *ListFlowNodeInstanceContainerStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlowNodeInstanceContainerStatus(request)
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

// ListFlowNodeInstanceContainerStatusWithCallback invokes the emr.ListFlowNodeInstanceContainerStatus API asynchronously
// api document: https://help.aliyun.com/api/emr/listflownodeinstancecontainerstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFlowNodeInstanceContainerStatusWithCallback(request *ListFlowNodeInstanceContainerStatusRequest, callback func(response *ListFlowNodeInstanceContainerStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlowNodeInstanceContainerStatusResponse
		var err error
		defer close(result)
		response, err = client.ListFlowNodeInstanceContainerStatus(request)
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

// ListFlowNodeInstanceContainerStatusRequest is the request struct for api ListFlowNodeInstanceContainerStatus
type ListFlowNodeInstanceContainerStatusRequest struct {
	*requests.RpcRequest
	NodeInstanceId string           `position:"Query" name:"NodeInstanceId"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	ProjectId      string           `position:"Query" name:"ProjectId"`
}

// ListFlowNodeInstanceContainerStatusResponse is the response struct for api ListFlowNodeInstanceContainerStatus
type ListFlowNodeInstanceContainerStatusResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	PageNumber          int                 `json:"PageNumber" xml:"PageNumber"`
	PageSize            int                 `json:"PageSize" xml:"PageSize"`
	Total               int                 `json:"Total" xml:"Total"`
	ContainerStatusList ContainerStatusList `json:"ContainerStatusList" xml:"ContainerStatusList"`
}

// CreateListFlowNodeInstanceContainerStatusRequest creates a request to invoke ListFlowNodeInstanceContainerStatus API
func CreateListFlowNodeInstanceContainerStatusRequest() (request *ListFlowNodeInstanceContainerStatusRequest) {
	request = &ListFlowNodeInstanceContainerStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListFlowNodeInstanceContainerStatus", "emr", "openAPI")
	return
}

// CreateListFlowNodeInstanceContainerStatusResponse creates a response to parse from ListFlowNodeInstanceContainerStatus response
func CreateListFlowNodeInstanceContainerStatusResponse() (response *ListFlowNodeInstanceContainerStatusResponse) {
	response = &ListFlowNodeInstanceContainerStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
