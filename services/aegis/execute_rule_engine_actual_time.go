package aegis

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

// ExecuteRuleEngineActualTime invokes the aegis.ExecuteRuleEngineActualTime API synchronously
// api document: https://help.aliyun.com/api/aegis/executeruleengineactualtime.html
func (client *Client) ExecuteRuleEngineActualTime(request *ExecuteRuleEngineActualTimeRequest) (response *ExecuteRuleEngineActualTimeResponse, err error) {
	response = CreateExecuteRuleEngineActualTimeResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteRuleEngineActualTimeWithChan invokes the aegis.ExecuteRuleEngineActualTime API asynchronously
// api document: https://help.aliyun.com/api/aegis/executeruleengineactualtime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExecuteRuleEngineActualTimeWithChan(request *ExecuteRuleEngineActualTimeRequest) (<-chan *ExecuteRuleEngineActualTimeResponse, <-chan error) {
	responseChan := make(chan *ExecuteRuleEngineActualTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteRuleEngineActualTime(request)
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

// ExecuteRuleEngineActualTimeWithCallback invokes the aegis.ExecuteRuleEngineActualTime API asynchronously
// api document: https://help.aliyun.com/api/aegis/executeruleengineactualtime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExecuteRuleEngineActualTimeWithCallback(request *ExecuteRuleEngineActualTimeRequest, callback func(response *ExecuteRuleEngineActualTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteRuleEngineActualTimeResponse
		var err error
		defer close(result)
		response, err = client.ExecuteRuleEngineActualTime(request)
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

// ExecuteRuleEngineActualTimeRequest is the request struct for api ExecuteRuleEngineActualTime
type ExecuteRuleEngineActualTimeRequest struct {
	*requests.RpcRequest
	SourceIp string           `position:"Query" name:"SourceIp"`
	RuleId   requests.Integer `position:"Query" name:"RuleId"`
	Message  string           `position:"Query" name:"Message"`
}

// ExecuteRuleEngineActualTimeResponse is the response struct for api ExecuteRuleEngineActualTime
type ExecuteRuleEngineActualTimeResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ExecResult string `json:"ExecResult" xml:"ExecResult"`
}

// CreateExecuteRuleEngineActualTimeRequest creates a request to invoke ExecuteRuleEngineActualTime API
func CreateExecuteRuleEngineActualTimeRequest() (request *ExecuteRuleEngineActualTimeRequest) {
	request = &ExecuteRuleEngineActualTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ExecuteRuleEngineActualTime", "vipaegis", "openAPI")
	return
}

// CreateExecuteRuleEngineActualTimeResponse creates a response to parse from ExecuteRuleEngineActualTime response
func CreateExecuteRuleEngineActualTimeResponse() (response *ExecuteRuleEngineActualTimeResponse) {
	response = &ExecuteRuleEngineActualTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
