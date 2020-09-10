package outboundbot

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

// CreateIntent invokes the outboundbot.CreateIntent API synchronously
// api document: https://help.aliyun.com/api/outboundbot/createintent.html
func (client *Client) CreateIntent(request *CreateIntentRequest) (response *CreateIntentResponse, err error) {
	response = CreateCreateIntentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateIntentWithChan invokes the outboundbot.CreateIntent API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/createintent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateIntentWithChan(request *CreateIntentRequest) (<-chan *CreateIntentResponse, <-chan error) {
	responseChan := make(chan *CreateIntentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateIntent(request)
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

// CreateIntentWithCallback invokes the outboundbot.CreateIntent API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/createintent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateIntentWithCallback(request *CreateIntentRequest, callback func(response *CreateIntentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateIntentResponse
		var err error
		defer close(result)
		response, err = client.CreateIntent(request)
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

// CreateIntentRequest is the request struct for api CreateIntent
type CreateIntentRequest struct {
	*requests.RpcRequest
	Utterances        string `position:"Query" name:"Utterances"`
	Keywords          string `position:"Query" name:"Keywords"`
	IntentDescription string `position:"Query" name:"IntentDescription"`
	ScriptId          string `position:"Query" name:"ScriptId"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	IntentName        string `position:"Query" name:"IntentName"`
}

// CreateIntentResponse is the response struct for api CreateIntent
type CreateIntentResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	IntentId       string `json:"IntentId" xml:"IntentId"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateCreateIntentRequest creates a request to invoke CreateIntent API
func CreateCreateIntentRequest() (request *CreateIntentRequest) {
	request = &CreateIntentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "CreateIntent", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateIntentResponse creates a response to parse from CreateIntent response
func CreateCreateIntentResponse() (response *CreateIntentResponse) {
	response = &CreateIntentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
