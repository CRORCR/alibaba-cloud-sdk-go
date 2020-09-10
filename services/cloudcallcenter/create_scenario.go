package cloudcallcenter

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

// CreateScenario invokes the cloudcallcenter.CreateScenario API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createscenario.html
func (client *Client) CreateScenario(request *CreateScenarioRequest) (response *CreateScenarioResponse, err error) {
	response = CreateCreateScenarioResponse()
	err = client.DoAction(request, response)
	return
}

// CreateScenarioWithChan invokes the cloudcallcenter.CreateScenario API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createscenario.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateScenarioWithChan(request *CreateScenarioRequest) (<-chan *CreateScenarioResponse, <-chan error) {
	responseChan := make(chan *CreateScenarioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateScenario(request)
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

// CreateScenarioWithCallback invokes the cloudcallcenter.CreateScenario API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createscenario.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateScenarioWithCallback(request *CreateScenarioRequest, callback func(response *CreateScenarioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateScenarioResponse
		var err error
		defer close(result)
		response, err = client.CreateScenario(request)
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

// CreateScenarioRequest is the request struct for api CreateScenario
type CreateScenarioRequest struct {
	*requests.RpcRequest
	SurveysJson   *[]string `position:"Query" name:"SurveysJson"  type:"Repeated"`
	Description   string    `position:"Query" name:"Description"`
	Type          string    `position:"Query" name:"Type"`
	InstanceId    string    `position:"Query" name:"InstanceId"`
	BeebotVersion string    `position:"Query" name:"BeebotVersion"`
	StrategyJson  string    `position:"Query" name:"StrategyJson"`
	Name          string    `position:"Query" name:"Name"`
}

// CreateScenarioResponse is the response struct for api CreateScenario
type CreateScenarioResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Scenario       Scenario `json:"Scenario" xml:"Scenario"`
}

// CreateCreateScenarioRequest creates a request to invoke CreateScenario API
func CreateCreateScenarioRequest() (request *CreateScenarioRequest) {
	request = &CreateScenarioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateScenario", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateScenarioResponse creates a response to parse from CreateScenario response
func CreateCreateScenarioResponse() (response *CreateScenarioResponse) {
	response = &CreateScenarioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
