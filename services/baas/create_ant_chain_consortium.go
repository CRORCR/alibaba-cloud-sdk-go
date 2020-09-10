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

// CreateAntChainConsortium invokes the baas.CreateAntChainConsortium API synchronously
// api document: https://help.aliyun.com/api/baas/createantchainconsortium.html
func (client *Client) CreateAntChainConsortium(request *CreateAntChainConsortiumRequest) (response *CreateAntChainConsortiumResponse, err error) {
	response = CreateCreateAntChainConsortiumResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAntChainConsortiumWithChan invokes the baas.CreateAntChainConsortium API asynchronously
// api document: https://help.aliyun.com/api/baas/createantchainconsortium.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAntChainConsortiumWithChan(request *CreateAntChainConsortiumRequest) (<-chan *CreateAntChainConsortiumResponse, <-chan error) {
	responseChan := make(chan *CreateAntChainConsortiumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAntChainConsortium(request)
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

// CreateAntChainConsortiumWithCallback invokes the baas.CreateAntChainConsortium API asynchronously
// api document: https://help.aliyun.com/api/baas/createantchainconsortium.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAntChainConsortiumWithCallback(request *CreateAntChainConsortiumRequest, callback func(response *CreateAntChainConsortiumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAntChainConsortiumResponse
		var err error
		defer close(result)
		response, err = client.CreateAntChainConsortium(request)
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

// CreateAntChainConsortiumRequest is the request struct for api CreateAntChainConsortium
type CreateAntChainConsortiumRequest struct {
	*requests.RpcRequest
	ConsortiumName        string `position:"Body" name:"ConsortiumName"`
	ConsortiumDescription string `position:"Body" name:"ConsortiumDescription"`
}

// CreateAntChainConsortiumResponse is the response struct for api CreateAntChainConsortium
type CreateAntChainConsortiumResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateCreateAntChainConsortiumRequest creates a request to invoke CreateAntChainConsortium API
func CreateCreateAntChainConsortiumRequest() (request *CreateAntChainConsortiumRequest) {
	request = &CreateAntChainConsortiumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "CreateAntChainConsortium", "baas", "openAPI")
	return
}

// CreateCreateAntChainConsortiumResponse creates a response to parse from CreateAntChainConsortium response
func CreateCreateAntChainConsortiumResponse() (response *CreateAntChainConsortiumResponse) {
	response = &CreateAntChainConsortiumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
