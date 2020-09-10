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

// DeleteScript invokes the outboundbot.DeleteScript API synchronously
// api document: https://help.aliyun.com/api/outboundbot/deletescript.html
func (client *Client) DeleteScript(request *DeleteScriptRequest) (response *DeleteScriptResponse, err error) {
	response = CreateDeleteScriptResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteScriptWithChan invokes the outboundbot.DeleteScript API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/deletescript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteScriptWithChan(request *DeleteScriptRequest) (<-chan *DeleteScriptResponse, <-chan error) {
	responseChan := make(chan *DeleteScriptResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteScript(request)
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

// DeleteScriptWithCallback invokes the outboundbot.DeleteScript API asynchronously
// api document: https://help.aliyun.com/api/outboundbot/deletescript.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteScriptWithCallback(request *DeleteScriptRequest, callback func(response *DeleteScriptResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteScriptResponse
		var err error
		defer close(result)
		response, err = client.DeleteScript(request)
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

// DeleteScriptRequest is the request struct for api DeleteScript
type DeleteScriptRequest struct {
	*requests.RpcRequest
	ScriptId   string `position:"Query" name:"ScriptId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DeleteScriptResponse is the response struct for api DeleteScript
type DeleteScriptResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeleteScriptRequest creates a request to invoke DeleteScript API
func CreateDeleteScriptRequest() (request *DeleteScriptRequest) {
	request = &DeleteScriptRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DeleteScript", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteScriptResponse creates a response to parse from DeleteScript response
func CreateDeleteScriptResponse() (response *DeleteScriptResponse) {
	response = &DeleteScriptResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
