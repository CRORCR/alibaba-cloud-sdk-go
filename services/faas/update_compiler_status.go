package faas

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

// UpdateCompilerStatus invokes the faas.UpdateCompilerStatus API synchronously
// api document: https://help.aliyun.com/api/faas/updatecompilerstatus.html
func (client *Client) UpdateCompilerStatus(request *UpdateCompilerStatusRequest) (response *UpdateCompilerStatusResponse, err error) {
	response = CreateUpdateCompilerStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCompilerStatusWithChan invokes the faas.UpdateCompilerStatus API asynchronously
// api document: https://help.aliyun.com/api/faas/updatecompilerstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateCompilerStatusWithChan(request *UpdateCompilerStatusRequest) (<-chan *UpdateCompilerStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateCompilerStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCompilerStatus(request)
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

// UpdateCompilerStatusWithCallback invokes the faas.UpdateCompilerStatus API asynchronously
// api document: https://help.aliyun.com/api/faas/updatecompilerstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateCompilerStatusWithCallback(request *UpdateCompilerStatusRequest, callback func(response *UpdateCompilerStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCompilerStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateCompilerStatus(request)
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

// UpdateCompilerStatusRequest is the request struct for api UpdateCompilerStatus
type UpdateCompilerStatusRequest struct {
	*requests.RpcRequest
	ErrorMessage string           `position:"Query" name:"ErrorMessage"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	TaskId       requests.Integer `position:"Query" name:"TaskId"`
	ErrorCode    string           `position:"Query" name:"ErrorCode"`
	Status       string           `position:"Query" name:"Status"`
}

// UpdateCompilerStatusResponse is the response struct for api UpdateCompilerStatus
type UpdateCompilerStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateUpdateCompilerStatusRequest creates a request to invoke UpdateCompilerStatus API
func CreateUpdateCompilerStatusRequest() (request *UpdateCompilerStatusRequest) {
	request = &UpdateCompilerStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2020-02-17", "UpdateCompilerStatus", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateCompilerStatusResponse creates a response to parse from UpdateCompilerStatus response
func CreateUpdateCompilerStatusResponse() (response *UpdateCompilerStatusResponse) {
	response = &UpdateCompilerStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
