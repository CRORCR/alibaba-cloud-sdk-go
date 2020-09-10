package voicenavigator

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

// QueryPerformanceIndicators invokes the voicenavigator.QueryPerformanceIndicators API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/queryperformanceindicators.html
func (client *Client) QueryPerformanceIndicators(request *QueryPerformanceIndicatorsRequest) (response *QueryPerformanceIndicatorsResponse, err error) {
	response = CreateQueryPerformanceIndicatorsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPerformanceIndicatorsWithChan invokes the voicenavigator.QueryPerformanceIndicators API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/queryperformanceindicators.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPerformanceIndicatorsWithChan(request *QueryPerformanceIndicatorsRequest) (<-chan *QueryPerformanceIndicatorsResponse, <-chan error) {
	responseChan := make(chan *QueryPerformanceIndicatorsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPerformanceIndicators(request)
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

// QueryPerformanceIndicatorsWithCallback invokes the voicenavigator.QueryPerformanceIndicators API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/queryperformanceindicators.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPerformanceIndicatorsWithCallback(request *QueryPerformanceIndicatorsRequest, callback func(response *QueryPerformanceIndicatorsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPerformanceIndicatorsResponse
		var err error
		defer close(result)
		response, err = client.QueryPerformanceIndicators(request)
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

// QueryPerformanceIndicatorsRequest is the request struct for api QueryPerformanceIndicators
type QueryPerformanceIndicatorsRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	DateUnit   string `position:"Query" name:"DateUnit"`
}

// QueryPerformanceIndicatorsResponse is the response struct for api QueryPerformanceIndicators
type QueryPerformanceIndicatorsResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	ResolutionRate   ResolutionRate   `json:"ResolutionRate" xml:"ResolutionRate"`
	ValidAnswerRate  ValidAnswerRate  `json:"ValidAnswerRate" xml:"ValidAnswerRate"`
	DialoguePassRate DialoguePassRate `json:"DialoguePassRate" xml:"DialoguePassRate"`
	KnowledgeHitRate KnowledgeHitRate `json:"KnowledgeHitRate" xml:"KnowledgeHitRate"`
}

// CreateQueryPerformanceIndicatorsRequest creates a request to invoke QueryPerformanceIndicators API
func CreateQueryPerformanceIndicatorsRequest() (request *QueryPerformanceIndicatorsRequest) {
	request = &QueryPerformanceIndicatorsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "QueryPerformanceIndicators", "voicebot", "openAPI")
	return
}

// CreateQueryPerformanceIndicatorsResponse creates a response to parse from QueryPerformanceIndicators response
func CreateQueryPerformanceIndicatorsResponse() (response *QueryPerformanceIndicatorsResponse) {
	response = &QueryPerformanceIndicatorsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
