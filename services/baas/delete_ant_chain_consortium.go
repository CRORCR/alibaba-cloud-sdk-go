package baas

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

// DeleteAntChainConsortium invokes the baas.DeleteAntChainConsortium API synchronously
// api document: https://help.aliyun.com/api/baas/deleteantchainconsortium.html
func (client *Client) DeleteAntChainConsortium(request *DeleteAntChainConsortiumRequest) (response *DeleteAntChainConsortiumResponse, err error) {
	response = CreateDeleteAntChainConsortiumResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAntChainConsortiumWithChan invokes the baas.DeleteAntChainConsortium API asynchronously
// api document: https://help.aliyun.com/api/baas/deleteantchainconsortium.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAntChainConsortiumWithChan(request *DeleteAntChainConsortiumRequest) (<-chan *DeleteAntChainConsortiumResponse, <-chan error) {
	responseChan := make(chan *DeleteAntChainConsortiumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAntChainConsortium(request)
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

// DeleteAntChainConsortiumWithCallback invokes the baas.DeleteAntChainConsortium API asynchronously
// api document: https://help.aliyun.com/api/baas/deleteantchainconsortium.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAntChainConsortiumWithCallback(request *DeleteAntChainConsortiumRequest, callback func(response *DeleteAntChainConsortiumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAntChainConsortiumResponse
		var err error
		defer close(result)
		response, err = client.DeleteAntChainConsortium(request)
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

// DeleteAntChainConsortiumRequest is the request struct for api DeleteAntChainConsortium
type DeleteAntChainConsortiumRequest struct {
	*requests.RpcRequest
	ConsortiumId string `position:"Body" name:"ConsortiumId"`
}

// DeleteAntChainConsortiumResponse is the response struct for api DeleteAntChainConsortium
type DeleteAntChainConsortiumResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateDeleteAntChainConsortiumRequest creates a request to invoke DeleteAntChainConsortium API
func CreateDeleteAntChainConsortiumRequest() (request *DeleteAntChainConsortiumRequest) {
	request = &DeleteAntChainConsortiumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DeleteAntChainConsortium", "baas", "openAPI")
	return
}

// CreateDeleteAntChainConsortiumResponse creates a response to parse from DeleteAntChainConsortium response
func CreateDeleteAntChainConsortiumResponse() (response *DeleteAntChainConsortiumResponse) {
	response = &DeleteAntChainConsortiumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
