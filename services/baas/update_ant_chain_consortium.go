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

// UpdateAntChainConsortium invokes the baas.UpdateAntChainConsortium API synchronously
// api document: https://help.aliyun.com/api/baas/updateantchainconsortium.html
func (client *Client) UpdateAntChainConsortium(request *UpdateAntChainConsortiumRequest) (response *UpdateAntChainConsortiumResponse, err error) {
	response = CreateUpdateAntChainConsortiumResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAntChainConsortiumWithChan invokes the baas.UpdateAntChainConsortium API asynchronously
// api document: https://help.aliyun.com/api/baas/updateantchainconsortium.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAntChainConsortiumWithChan(request *UpdateAntChainConsortiumRequest) (<-chan *UpdateAntChainConsortiumResponse, <-chan error) {
	responseChan := make(chan *UpdateAntChainConsortiumResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAntChainConsortium(request)
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

// UpdateAntChainConsortiumWithCallback invokes the baas.UpdateAntChainConsortium API asynchronously
// api document: https://help.aliyun.com/api/baas/updateantchainconsortium.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAntChainConsortiumWithCallback(request *UpdateAntChainConsortiumRequest, callback func(response *UpdateAntChainConsortiumResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAntChainConsortiumResponse
		var err error
		defer close(result)
		response, err = client.UpdateAntChainConsortium(request)
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

// UpdateAntChainConsortiumRequest is the request struct for api UpdateAntChainConsortium
type UpdateAntChainConsortiumRequest struct {
	*requests.RpcRequest
	ConsortiumName        string `position:"Body" name:"ConsortiumName"`
	ConsortiumDescription string `position:"Body" name:"ConsortiumDescription"`
	ConsortiumId          string `position:"Body" name:"ConsortiumId"`
}

// UpdateAntChainConsortiumResponse is the response struct for api UpdateAntChainConsortium
type UpdateAntChainConsortiumResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateUpdateAntChainConsortiumRequest creates a request to invoke UpdateAntChainConsortium API
func CreateUpdateAntChainConsortiumRequest() (request *UpdateAntChainConsortiumRequest) {
	request = &UpdateAntChainConsortiumRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "UpdateAntChainConsortium", "baas", "openAPI")
	return
}

// CreateUpdateAntChainConsortiumResponse creates a response to parse from UpdateAntChainConsortium response
func CreateUpdateAntChainConsortiumResponse() (response *UpdateAntChainConsortiumResponse) {
	response = &UpdateAntChainConsortiumResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
