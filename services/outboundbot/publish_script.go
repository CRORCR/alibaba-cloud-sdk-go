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

// PublishScript invokes the outboundbot.PublishScript API synchronously
// api document: https://help.aliyun.com/api/outboundbot/publishscript.html
func (client *Client) PublishScript(request *PublishScriptRequest) (response *PublishScriptResponse, err error) {
	response = CreatePublishScriptResponse()
	err = client.DoAction(request, response)
	return
}

// PublishScriptWithChan invokes the outboundbot.PublishScript API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/publishscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PublishScriptWithChan(request *PublishScriptRequest) (<-chan *PublishScriptResponse, <-chan error) {
	responseChan := make(chan *PublishScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishScript(request)
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

// PublishScriptWithCallback invokes the outboundbot.PublishScript API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/publishscript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PublishScriptWithCallback(request *PublishScriptRequest, callback func(response *PublishScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishScriptResponse
		var err error
		defer close(result)
		response, err = client.PublishScript(request)
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

// PublishScriptRequest is the request struct for api PublishScript
type PublishScriptRequest struct {
	*requests.RpcRequest
	Description     string           `position:"Query" name:"Description"`
	ScriptId        string           `position:"Query" name:"ScriptId"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	InstanceOwnerId requests.Integer `position:"Query" name:"InstanceOwnerId"`
}

// PublishScriptResponse is the response struct for api PublishScript
type PublishScriptResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreatePublishScriptRequest creates a request to invoke PublishScript API
func CreatePublishScriptRequest() (request *PublishScriptRequest) {
	request = &PublishScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "PublishScript", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePublishScriptResponse creates a response to parse from PublishScript response
func CreatePublishScriptResponse() (response *PublishScriptResponse) {
	response = &PublishScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
