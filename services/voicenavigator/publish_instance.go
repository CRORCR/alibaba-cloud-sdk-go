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

// PublishInstance invokes the voicenavigator.PublishInstance API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/publishinstance.html
func (client *Client) PublishInstance(request *PublishInstanceRequest) (response *PublishInstanceResponse, err error) {
	response = CreatePublishInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// PublishInstanceWithChan invokes the voicenavigator.PublishInstance API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/publishinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PublishInstanceWithChan(request *PublishInstanceRequest) (<-chan *PublishInstanceResponse, <-chan error) {
	responseChan := make(chan *PublishInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishInstance(request)
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

// PublishInstanceWithCallback invokes the voicenavigator.PublishInstance API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/publishinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PublishInstanceWithCallback(request *PublishInstanceRequest, callback func(response *PublishInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishInstanceResponse
		var err error
		defer close(result)
		response, err = client.PublishInstance(request)
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

// PublishInstanceRequest is the request struct for api PublishInstance
type PublishInstanceRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// PublishInstanceResponse is the response struct for api PublishInstance
type PublishInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
	Version   string `json:"Version" xml:"Version"`
}

// CreatePublishInstanceRequest creates a request to invoke PublishInstance API
func CreatePublishInstanceRequest() (request *PublishInstanceRequest) {
	request = &PublishInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "PublishInstance", "voicebot", "openAPI")
	return
}

// CreatePublishInstanceResponse creates a response to parse from PublishInstance response
func CreatePublishInstanceResponse() (response *PublishInstanceResponse) {
	response = &PublishInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
