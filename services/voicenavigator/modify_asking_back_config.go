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

// ModifyAskingBackConfig invokes the voicenavigator.ModifyAskingBackConfig API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifyaskingbackconfig.html
func (client *Client) ModifyAskingBackConfig(request *ModifyAskingBackConfigRequest) (response *ModifyAskingBackConfigResponse, err error) {
	response = CreateModifyAskingBackConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAskingBackConfigWithChan invokes the voicenavigator.ModifyAskingBackConfig API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifyaskingbackconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAskingBackConfigWithChan(request *ModifyAskingBackConfigRequest) (<-chan *ModifyAskingBackConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyAskingBackConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAskingBackConfig(request)
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

// ModifyAskingBackConfigWithCallback invokes the voicenavigator.ModifyAskingBackConfig API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifyaskingbackconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAskingBackConfigWithCallback(request *ModifyAskingBackConfigRequest, callback func(response *ModifyAskingBackConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAskingBackConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyAskingBackConfig(request)
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

// ModifyAskingBackConfigRequest is the request struct for api ModifyAskingBackConfig
type ModifyAskingBackConfigRequest struct {
	*requests.RpcRequest
	NegativeFeedbackPrompt       string           `position:"Query" name:"NegativeFeedbackPrompt"`
	NegativeFeedbackAction       string           `position:"Query" name:"NegativeFeedbackAction"`
	Enabled                      requests.Boolean `position:"Query" name:"Enabled"`
	EnableNegativeFeedback       requests.Boolean `position:"Query" name:"EnableNegativeFeedback"`
	InstanceId                   string           `position:"Query" name:"InstanceId"`
	Prompt                       string           `position:"Query" name:"Prompt"`
	NegativeFeedbackUtterances   *[]string        `position:"Query" name:"NegativeFeedbackUtterances"  type:"Repeated"`
	NegativeFeedbackActionParams string           `position:"Query" name:"NegativeFeedbackActionParams"`
}

// ModifyAskingBackConfigResponse is the response struct for api ModifyAskingBackConfig
type ModifyAskingBackConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAskingBackConfigRequest creates a request to invoke ModifyAskingBackConfig API
func CreateModifyAskingBackConfigRequest() (request *ModifyAskingBackConfigRequest) {
	request = &ModifyAskingBackConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "ModifyAskingBackConfig", "voicebot", "openAPI")
	return
}

// CreateModifyAskingBackConfigResponse creates a response to parse from ModifyAskingBackConfig response
func CreateModifyAskingBackConfigResponse() (response *ModifyAskingBackConfigResponse) {
	response = &ModifyAskingBackConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
