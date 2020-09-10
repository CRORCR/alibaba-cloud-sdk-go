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

// GetFaceEntity invokes the facebody.GetFaceEntity API synchronously
func (client *Client) GetFaceEntity(request *GetFaceEntityRequest) (response *GetFaceEntityResponse, err error) {
	response = CreateGetFaceEntityResponse()
	err = client.DoAction(request, response)
	return
}

// GetFaceEntityWithChan invokes the facebody.GetFaceEntity API asynchronously
func (client *Client) GetFaceEntityWithChan(request *GetFaceEntityRequest) (<-chan *GetFaceEntityResponse, <-chan error) {
	responseChan := make(chan *GetFaceEntityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFaceEntity(request)
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

// GetFaceEntityWithCallback invokes the facebody.GetFaceEntity API asynchronously
func (client *Client) GetFaceEntityWithCallback(request *GetFaceEntityRequest, callback func(response *GetFaceEntityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFaceEntityResponse
		var err error
		defer close(result)
		response, err = client.GetFaceEntity(request)
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

// GetFaceEntityRequest is the request struct for api GetFaceEntity
type GetFaceEntityRequest struct {
	*requests.RpcRequest
	EntityId string `position:"Body" name:"EntityId"`
	DbName   string `position:"Body" name:"DbName"`
}

// GetFaceEntityResponse is the response struct for api GetFaceEntity
type GetFaceEntityResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetFaceEntityRequest creates a request to invoke GetFaceEntity API
func CreateGetFaceEntityRequest() (request *GetFaceEntityRequest) {
	request = &GetFaceEntityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "GetFaceEntity", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetFaceEntityResponse creates a response to parse from GetFaceEntity response
func CreateGetFaceEntityResponse() (response *GetFaceEntityResponse) {
	response = &GetFaceEntityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
