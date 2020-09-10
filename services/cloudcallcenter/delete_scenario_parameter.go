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

// DeleteScenarioParameter invokes the cloudcallcenter.DeleteScenarioParameter API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletescenarioparameter.html
func (client *Client) DeleteScenarioParameter(request *DeleteScenarioParameterRequest) (response *DeleteScenarioParameterResponse, err error) {
	response = CreateDeleteScenarioParameterResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteScenarioParameterWithChan invokes the cloudcallcenter.DeleteScenarioParameter API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletescenarioparameter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteScenarioParameterWithChan(request *DeleteScenarioParameterRequest) (<-chan *DeleteScenarioParameterResponse, <-chan error) {
	responseChan := make(chan *DeleteScenarioParameterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteScenarioParameter(request)
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

// DeleteScenarioParameterWithCallback invokes the cloudcallcenter.DeleteScenarioParameter API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletescenarioparameter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteScenarioParameterWithCallback(request *DeleteScenarioParameterRequest, callback func(response *DeleteScenarioParameterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteScenarioParameterResponse
		var err error
		defer close(result)
		response, err = client.DeleteScenarioParameter(request)
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

// DeleteScenarioParameterRequest is the request struct for api DeleteScenarioParameter
type DeleteScenarioParameterRequest struct {
	*requests.RpcRequest
	ScenarioParameterId string `position:"Query" name:"ScenarioParameterId"`
	InstanceId          string `position:"Query" name:"InstanceId"`
	ScenarioId          string `position:"Query" name:"ScenarioId"`
}

// DeleteScenarioParameterResponse is the response struct for api DeleteScenarioParameter
type DeleteScenarioParameterResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDeleteScenarioParameterRequest creates a request to invoke DeleteScenarioParameter API
func CreateDeleteScenarioParameterRequest() (request *DeleteScenarioParameterRequest) {
	request = &DeleteScenarioParameterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DeleteScenarioParameter", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteScenarioParameterResponse creates a response to parse from DeleteScenarioParameter response
func CreateDeleteScenarioParameterResponse() (response *DeleteScenarioParameterResponse) {
	response = &DeleteScenarioParameterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
