package trademark

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

// DeleteMaterial invokes the trademark.DeleteMaterial API synchronously
// api document: https://help.aliyun.com/api/trademark/deletematerial.html
func (client *Client) DeleteMaterial(request *DeleteMaterialRequest) (response *DeleteMaterialResponse, err error) {
	response = CreateDeleteMaterialResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMaterialWithChan invokes the trademark.DeleteMaterial API asynchronously
// api document: https://help.aliyun.com/api/trademark/deletematerial.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMaterialWithChan(request *DeleteMaterialRequest) (<-chan *DeleteMaterialResponse, <-chan error) {
	responseChan := make(chan *DeleteMaterialResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMaterial(request)
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

// DeleteMaterialWithCallback invokes the trademark.DeleteMaterial API asynchronously
// api document: https://help.aliyun.com/api/trademark/deletematerial.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMaterialWithCallback(request *DeleteMaterialRequest, callback func(response *DeleteMaterialResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMaterialResponse
		var err error
		defer close(result)
		response, err = client.DeleteMaterial(request)
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

// DeleteMaterialRequest is the request struct for api DeleteMaterial
type DeleteMaterialRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// DeleteMaterialResponse is the response struct for api DeleteMaterial
type DeleteMaterialResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateDeleteMaterialRequest creates a request to invoke DeleteMaterial API
func CreateDeleteMaterialRequest() (request *DeleteMaterialRequest) {
	request = &DeleteMaterialRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "DeleteMaterial", "trademark", "openAPI")
	return
}

// CreateDeleteMaterialResponse creates a response to parse from DeleteMaterial response
func CreateDeleteMaterialResponse() (response *DeleteMaterialResponse) {
	response = &DeleteMaterialResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
