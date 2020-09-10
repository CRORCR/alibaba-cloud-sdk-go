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

// RegisterMediaDetailScenario invokes the mts.RegisterMediaDetailScenario API synchronously
// api document: https://help.aliyun.com/api/mts/registermediadetailscenario.html
func (client *Client) RegisterMediaDetailScenario(request *RegisterMediaDetailScenarioRequest) (response *RegisterMediaDetailScenarioResponse, err error) {
	response = CreateRegisterMediaDetailScenarioResponse()
	err = client.DoAction(request, response)
	return
}

// RegisterMediaDetailScenarioWithChan invokes the mts.RegisterMediaDetailScenario API asynchronously
// api document: https://help.aliyun.com/api/mts/registermediadetailscenario.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegisterMediaDetailScenarioWithChan(request *RegisterMediaDetailScenarioRequest) (<-chan *RegisterMediaDetailScenarioResponse, <-chan error) {
	responseChan := make(chan *RegisterMediaDetailScenarioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterMediaDetailScenario(request)
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

// RegisterMediaDetailScenarioWithCallback invokes the mts.RegisterMediaDetailScenario API asynchronously
// api document: https://help.aliyun.com/api/mts/registermediadetailscenario.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RegisterMediaDetailScenarioWithCallback(request *RegisterMediaDetailScenarioRequest, callback func(response *RegisterMediaDetailScenarioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterMediaDetailScenarioResponse
		var err error
		defer close(result)
		response, err = client.RegisterMediaDetailScenario(request)
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

// RegisterMediaDetailScenarioRequest is the request struct for api RegisterMediaDetailScenario
type RegisterMediaDetailScenarioRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	JobId                string           `position:"Query" name:"JobId"`
	Scenario             string           `position:"Query" name:"Scenario"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RegisterMediaDetailScenarioResponse is the response struct for api RegisterMediaDetailScenario
type RegisterMediaDetailScenarioResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ScenarioId string `json:"ScenarioId" xml:"ScenarioId"`
}

// CreateRegisterMediaDetailScenarioRequest creates a request to invoke RegisterMediaDetailScenario API
func CreateRegisterMediaDetailScenarioRequest() (request *RegisterMediaDetailScenarioRequest) {
	request = &RegisterMediaDetailScenarioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "RegisterMediaDetailScenario", "", "")
	return
}

// CreateRegisterMediaDetailScenarioResponse creates a response to parse from RegisterMediaDetailScenario response
func CreateRegisterMediaDetailScenarioResponse() (response *RegisterMediaDetailScenarioResponse) {
	response = &RegisterMediaDetailScenarioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
