package devops_rdc

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

// ListDevopsProjectTaskFlowStatus invokes the devops_rdc.ListDevopsProjectTaskFlowStatus API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/listdevopsprojecttaskflowstatus.html
func (client *Client) ListDevopsProjectTaskFlowStatus(request *ListDevopsProjectTaskFlowStatusRequest) (response *ListDevopsProjectTaskFlowStatusResponse, err error) {
	response = CreateListDevopsProjectTaskFlowStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ListDevopsProjectTaskFlowStatusWithChan invokes the devops_rdc.ListDevopsProjectTaskFlowStatus API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/listdevopsprojecttaskflowstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDevopsProjectTaskFlowStatusWithChan(request *ListDevopsProjectTaskFlowStatusRequest) (<-chan *ListDevopsProjectTaskFlowStatusResponse, <-chan error) {
	responseChan := make(chan *ListDevopsProjectTaskFlowStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDevopsProjectTaskFlowStatus(request)
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

// ListDevopsProjectTaskFlowStatusWithCallback invokes the devops_rdc.ListDevopsProjectTaskFlowStatus API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/listdevopsprojecttaskflowstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDevopsProjectTaskFlowStatusWithCallback(request *ListDevopsProjectTaskFlowStatusRequest, callback func(response *ListDevopsProjectTaskFlowStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDevopsProjectTaskFlowStatusResponse
		var err error
		defer close(result)
		response, err = client.ListDevopsProjectTaskFlowStatus(request)
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

// ListDevopsProjectTaskFlowStatusRequest is the request struct for api ListDevopsProjectTaskFlowStatus
type ListDevopsProjectTaskFlowStatusRequest struct {
	*requests.RpcRequest
	TaskFlowId string `position:"Body" name:"TaskFlowId"`
	OrgId      string `position:"Body" name:"OrgId"`
}

// ListDevopsProjectTaskFlowStatusResponse is the response struct for api ListDevopsProjectTaskFlowStatus
type ListDevopsProjectTaskFlowStatusResponse struct {
	*responses.BaseResponse
	Successful bool             `json:"Successful" xml:"Successful"`
	ErrorCode  string           `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string           `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	Object     []TaskflowStatus `json:"Object" xml:"Object"`
}

// CreateListDevopsProjectTaskFlowStatusRequest creates a request to invoke ListDevopsProjectTaskFlowStatus API
func CreateListDevopsProjectTaskFlowStatusRequest() (request *ListDevopsProjectTaskFlowStatusRequest) {
	request = &ListDevopsProjectTaskFlowStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "ListDevopsProjectTaskFlowStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateListDevopsProjectTaskFlowStatusResponse creates a response to parse from ListDevopsProjectTaskFlowStatus response
func CreateListDevopsProjectTaskFlowStatusResponse() (response *ListDevopsProjectTaskFlowStatusResponse) {
	response = &ListDevopsProjectTaskFlowStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
