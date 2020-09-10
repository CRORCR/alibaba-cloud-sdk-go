package iot

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

// StartRule invokes the iot.StartRule API synchronously
// api document: https://help.aliyun.com/api/iot/startrule.html
func (client *Client) StartRule(request *StartRuleRequest) (response *StartRuleResponse, err error) {
	response = CreateStartRuleResponse()
	err = client.DoAction(request, response)
	return
}

// StartRuleWithChan invokes the iot.StartRule API asynchronously
// api document: https://help.aliyun.com/api/iot/startrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartRuleWithChan(request *StartRuleRequest) (<-chan *StartRuleResponse, <-chan error) {
	responseChan := make(chan *StartRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartRule(request)
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

// StartRuleWithCallback invokes the iot.StartRule API asynchronously
// api document: https://help.aliyun.com/api/iot/startrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartRuleWithCallback(request *StartRuleRequest, callback func(response *StartRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartRuleResponse
		var err error
		defer close(result)
		response, err = client.StartRule(request)
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

// StartRuleRequest is the request struct for api StartRule
type StartRuleRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	RuleId        requests.Integer `position:"Query" name:"RuleId"`
}

// StartRuleResponse is the response struct for api StartRule
type StartRuleResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateStartRuleRequest creates a request to invoke StartRule API
func CreateStartRuleRequest() (request *StartRuleRequest) {
	request = &StartRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "StartRule", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartRuleResponse creates a response to parse from StartRule response
func CreateStartRuleResponse() (response *StartRuleResponse) {
	response = &StartRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
