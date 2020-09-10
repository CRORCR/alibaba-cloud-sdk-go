package edas

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

// ChangeDeployGroup invokes the edas.ChangeDeployGroup API synchronously
// api document: https://help.aliyun.com/api/edas/changedeploygroup.html
func (client *Client) ChangeDeployGroup(request *ChangeDeployGroupRequest) (response *ChangeDeployGroupResponse, err error) {
	response = CreateChangeDeployGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeDeployGroupWithChan invokes the edas.ChangeDeployGroup API asynchronously
// api document: https://help.aliyun.com/api/edas/changedeploygroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ChangeDeployGroupWithChan(request *ChangeDeployGroupRequest) (<-chan *ChangeDeployGroupResponse, <-chan error) {
	responseChan := make(chan *ChangeDeployGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeDeployGroup(request)
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

// ChangeDeployGroupWithCallback invokes the edas.ChangeDeployGroup API asynchronously
// api document: https://help.aliyun.com/api/edas/changedeploygroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ChangeDeployGroupWithCallback(request *ChangeDeployGroupRequest, callback func(response *ChangeDeployGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeDeployGroupResponse
		var err error
		defer close(result)
		response, err = client.ChangeDeployGroup(request)
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

// ChangeDeployGroupRequest is the request struct for api ChangeDeployGroup
type ChangeDeployGroupRequest struct {
	*requests.RoaRequest
	ForceStatus requests.Boolean `position:"Query" name:"ForceStatus"`
	AppId       string           `position:"Query" name:"AppId"`
	EccInfo     string           `position:"Query" name:"EccInfo"`
	GroupName   string           `position:"Query" name:"GroupName"`
}

// ChangeDeployGroupResponse is the response struct for api ChangeDeployGroup
type ChangeDeployGroupResponse struct {
	*responses.BaseResponse
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateChangeDeployGroupRequest creates a request to invoke ChangeDeployGroup API
func CreateChangeDeployGroupRequest() (request *ChangeDeployGroupRequest) {
	request = &ChangeDeployGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ChangeDeployGroup", "/pop/v5/changeorder/co_change_group", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateChangeDeployGroupResponse creates a response to parse from ChangeDeployGroup response
func CreateChangeDeployGroupResponse() (response *ChangeDeployGroupResponse) {
	response = &ChangeDeployGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
