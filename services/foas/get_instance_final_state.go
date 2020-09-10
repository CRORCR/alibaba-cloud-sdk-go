package foas

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

// GetInstanceFinalState invokes the foas.GetInstanceFinalState API synchronously
// api document: https://help.aliyun.com/api/foas/getinstancefinalstate.html
func (client *Client) GetInstanceFinalState(request *GetInstanceFinalStateRequest) (response *GetInstanceFinalStateResponse, err error) {
	response = CreateGetInstanceFinalStateResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceFinalStateWithChan invokes the foas.GetInstanceFinalState API asynchronously
// api document: https://help.aliyun.com/api/foas/getinstancefinalstate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceFinalStateWithChan(request *GetInstanceFinalStateRequest) (<-chan *GetInstanceFinalStateResponse, <-chan error) {
	responseChan := make(chan *GetInstanceFinalStateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceFinalState(request)
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

// GetInstanceFinalStateWithCallback invokes the foas.GetInstanceFinalState API asynchronously
// api document: https://help.aliyun.com/api/foas/getinstancefinalstate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstanceFinalStateWithCallback(request *GetInstanceFinalStateRequest, callback func(response *GetInstanceFinalStateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceFinalStateResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceFinalState(request)
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

// GetInstanceFinalStateRequest is the request struct for api GetInstanceFinalState
type GetInstanceFinalStateRequest struct {
	*requests.RoaRequest
	ProjectName string           `position:"Path" name:"projectName"`
	InstanceId  requests.Integer `position:"Path" name:"instanceId"`
	JobName     string           `position:"Path" name:"jobName"`
}

// GetInstanceFinalStateResponse is the response struct for api GetInstanceFinalState
type GetInstanceFinalStateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Finalstate string `json:"Finalstate" xml:"Finalstate"`
}

// CreateGetInstanceFinalStateRequest creates a request to invoke GetInstanceFinalState API
func CreateGetInstanceFinalStateRequest() (request *GetInstanceFinalStateRequest) {
	request = &GetInstanceFinalStateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "GetInstanceFinalState", "/api/v2/projects/[projectName]/jobs/[jobName]/instances/[instanceId]/finalstate", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetInstanceFinalStateResponse creates a response to parse from GetInstanceFinalState response
func CreateGetInstanceFinalStateResponse() (response *GetInstanceFinalStateResponse) {
	response = &GetInstanceFinalStateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
