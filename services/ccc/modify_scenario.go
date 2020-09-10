package ccc

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

// ModifyScenario invokes the ccc.ModifyScenario API synchronously
// api document: https://help.aliyun.com/api/ccc/modifyscenario.html
func (client *Client) ModifyScenario(request *ModifyScenarioRequest) (response *ModifyScenarioResponse, err error) {
	response = CreateModifyScenarioResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyScenarioWithChan invokes the ccc.ModifyScenario API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyscenario.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScenarioWithChan(request *ModifyScenarioRequest) (<-chan *ModifyScenarioResponse, <-chan error) {
	responseChan := make(chan *ModifyScenarioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScenario(request)
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

// ModifyScenarioWithCallback invokes the ccc.ModifyScenario API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyscenario.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyScenarioWithCallback(request *ModifyScenarioRequest, callback func(response *ModifyScenarioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScenarioResponse
		var err error
		defer close(result)
		response, err = client.ModifyScenario(request)
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

// ModifyScenarioRequest is the request struct for api ModifyScenario
type ModifyScenarioRequest struct {
	*requests.RpcRequest
	Variables   string `position:"Query" name:"Variables"`
	Description string `position:"Query" name:"Description"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	Name        string `position:"Query" name:"Name"`
	ScenarioId  string `position:"Query" name:"ScenarioId"`
}

// ModifyScenarioResponse is the response struct for api ModifyScenario
type ModifyScenarioResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Scenario       Scenario `json:"Scenario" xml:"Scenario"`
}

// CreateModifyScenarioRequest creates a request to invoke ModifyScenario API
func CreateModifyScenarioRequest() (request *ModifyScenarioRequest) {
	request = &ModifyScenarioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifyScenario", "", "")
	return
}

// CreateModifyScenarioResponse creates a response to parse from ModifyScenario response
func CreateModifyScenarioResponse() (response *ModifyScenarioResponse) {
	response = &ModifyScenarioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
