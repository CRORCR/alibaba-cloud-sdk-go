package facebody

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

// EnhanceFace invokes the facebody.EnhanceFace API synchronously
func (client *Client) EnhanceFace(request *EnhanceFaceRequest) (response *EnhanceFaceResponse, err error) {
	response = CreateEnhanceFaceResponse()
	err = client.DoAction(request, response)
	return
}

// EnhanceFaceWithChan invokes the facebody.EnhanceFace API asynchronously
func (client *Client) EnhanceFaceWithChan(request *EnhanceFaceRequest) (<-chan *EnhanceFaceResponse, <-chan error) {
	responseChan := make(chan *EnhanceFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnhanceFace(request)
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

// EnhanceFaceWithCallback invokes the facebody.EnhanceFace API asynchronously
func (client *Client) EnhanceFaceWithCallback(request *EnhanceFaceRequest, callback func(response *EnhanceFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnhanceFaceResponse
		var err error
		defer close(result)
		response, err = client.EnhanceFace(request)
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

// EnhanceFaceRequest is the request struct for api EnhanceFace
type EnhanceFaceRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// EnhanceFaceResponse is the response struct for api EnhanceFace
type EnhanceFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateEnhanceFaceRequest creates a request to invoke EnhanceFace API
func CreateEnhanceFaceRequest() (request *EnhanceFaceRequest) {
	request = &EnhanceFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "EnhanceFace", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnhanceFaceResponse creates a response to parse from EnhanceFace response
func CreateEnhanceFaceResponse() (response *EnhanceFaceResponse) {
	response = &EnhanceFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
