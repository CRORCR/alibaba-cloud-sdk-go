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

// CreateSurvey invokes the ccc.CreateSurvey API synchronously
// api document: https://help.aliyun.com/api/ccc/createsurvey.html
func (client *Client) CreateSurvey(request *CreateSurveyRequest) (response *CreateSurveyResponse, err error) {
	response = CreateCreateSurveyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSurveyWithChan invokes the ccc.CreateSurvey API asynchronously
// api document: https://help.aliyun.com/api/ccc/createsurvey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSurveyWithChan(request *CreateSurveyRequest) (<-chan *CreateSurveyResponse, <-chan error) {
	responseChan := make(chan *CreateSurveyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSurvey(request)
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

// CreateSurveyWithCallback invokes the ccc.CreateSurvey API asynchronously
// api document: https://help.aliyun.com/api/ccc/createsurvey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSurveyWithCallback(request *CreateSurveyRequest, callback func(response *CreateSurveyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSurveyResponse
		var err error
		defer close(result)
		response, err = client.CreateSurvey(request)
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

// CreateSurveyRequest is the request struct for api CreateSurvey
type CreateSurveyRequest struct {
	*requests.RpcRequest
	Role                    string           `position:"Query" name:"Role"`
	Description             string           `position:"Query" name:"Description"`
	SpeechOptimizationParam string           `position:"Query" name:"SpeechOptimizationParam"`
	InstanceId              string           `position:"Query" name:"InstanceId"`
	Round                   requests.Integer `position:"Query" name:"Round"`
	FlowJson                string           `position:"Query" name:"FlowJson"`
	Name                    string           `position:"Query" name:"Name"`
	GlobalQuestions         string           `position:"Query" name:"GlobalQuestions"`
	Corpora                 string           `position:"Query" name:"Corpora"`
	ScenarioId              string           `position:"Query" name:"ScenarioId"`
}

// CreateSurveyResponse is the response struct for api CreateSurvey
type CreateSurveyResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Survey         Survey `json:"Survey" xml:"Survey"`
}

// CreateCreateSurveyRequest creates a request to invoke CreateSurvey API
func CreateCreateSurveyRequest() (request *CreateSurveyRequest) {
	request = &CreateSurveyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "CreateSurvey", "", "")
	return
}

// CreateCreateSurveyResponse creates a response to parse from CreateSurvey response
func CreateCreateSurveyResponse() (response *CreateSurveyResponse) {
	response = &CreateSurveyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
