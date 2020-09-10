package mse

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

// CreateZnode invokes the mse.CreateZnode API synchronously
// api document: https://help.aliyun.com/api/mse/createznode.html
func (client *Client) CreateZnode(request *CreateZnodeRequest) (response *CreateZnodeResponse, err error) {
	response = CreateCreateZnodeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateZnodeWithChan invokes the mse.CreateZnode API asynchronously
// api document: https://help.aliyun.com/api/mse/createznode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateZnodeWithChan(request *CreateZnodeRequest) (<-chan *CreateZnodeResponse, <-chan error) {
	responseChan := make(chan *CreateZnodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateZnode(request)
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

// CreateZnodeWithCallback invokes the mse.CreateZnode API asynchronously
// api document: https://help.aliyun.com/api/mse/createznode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateZnodeWithCallback(request *CreateZnodeRequest, callback func(response *CreateZnodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateZnodeResponse
		var err error
		defer close(result)
		response, err = client.CreateZnode(request)
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

// CreateZnodeRequest is the request struct for api CreateZnode
type CreateZnodeRequest struct {
	*requests.RpcRequest
	Data      string `position:"Query" name:"Data"`
	ClusterId string `position:"Query" name:"ClusterId"`
	Path      string `position:"Query" name:"Path"`
}

// CreateZnodeResponse is the response struct for api CreateZnode
type CreateZnodeResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	HttpCode  string `json:"HttpCode" xml:"HttpCode"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateZnodeRequest creates a request to invoke CreateZnode API
func CreateCreateZnodeRequest() (request *CreateZnodeRequest) {
	request = &CreateZnodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "CreateZnode", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateZnodeResponse creates a response to parse from CreateZnode response
func CreateCreateZnodeResponse() (response *CreateZnodeResponse) {
	response = &CreateZnodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
