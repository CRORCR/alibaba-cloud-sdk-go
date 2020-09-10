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

// BeginVnDialogue invokes the cloudcallcenter.BeginVnDialogue API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/beginvndialogue.html
func (client *Client) BeginVnDialogue(request *BeginVnDialogueRequest) (response *BeginVnDialogueResponse, err error) {
	response = CreateBeginVnDialogueResponse()
	err = client.DoAction(request, response)
	return
}

// BeginVnDialogueWithChan invokes the cloudcallcenter.BeginVnDialogue API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/beginvndialogue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BeginVnDialogueWithChan(request *BeginVnDialogueRequest) (<-chan *BeginVnDialogueResponse, <-chan error) {
	responseChan := make(chan *BeginVnDialogueResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BeginVnDialogue(request)
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

// BeginVnDialogueWithCallback invokes the cloudcallcenter.BeginVnDialogue API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/beginvndialogue.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BeginVnDialogueWithCallback(request *BeginVnDialogueRequest, callback func(response *BeginVnDialogueResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BeginVnDialogueResponse
		var err error
		defer close(result)
		response, err = client.BeginVnDialogue(request)
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

// BeginVnDialogueRequest is the request struct for api BeginVnDialogue
type BeginVnDialogueRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	InitialContext string `position:"Query" name:"InitialContext"`
	CallingNumber  string `position:"Query" name:"CallingNumber"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	CalledNumber   string `position:"Query" name:"CalledNumber"`
}

// BeginVnDialogueResponse is the response struct for api BeginVnDialogue
type BeginVnDialogueResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	Action        string `json:"Action" xml:"Action"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
}

// CreateBeginVnDialogueRequest creates a request to invoke BeginVnDialogue API
func CreateBeginVnDialogueRequest() (request *BeginVnDialogueRequest) {
	request = &BeginVnDialogueRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "BeginVnDialogue", "", "")
	request.Method = requests.GET
	return
}

// CreateBeginVnDialogueResponse creates a response to parse from BeginVnDialogue response
func CreateBeginVnDialogueResponse() (response *BeginVnDialogueResponse) {
	response = &BeginVnDialogueResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
