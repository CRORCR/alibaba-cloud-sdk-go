package ens

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

// ModifyImageAttribute invokes the ens.ModifyImageAttribute API synchronously
// api document: https://help.aliyun.com/api/ens/modifyimageattribute.html
func (client *Client) ModifyImageAttribute(request *ModifyImageAttributeRequest) (response *ModifyImageAttributeResponse, err error) {
	response = CreateModifyImageAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyImageAttributeWithChan invokes the ens.ModifyImageAttribute API asynchronously
// api document: https://help.aliyun.com/api/ens/modifyimageattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyImageAttributeWithChan(request *ModifyImageAttributeRequest) (<-chan *ModifyImageAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyImageAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyImageAttribute(request)
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

// ModifyImageAttributeWithCallback invokes the ens.ModifyImageAttribute API asynchronously
// api document: https://help.aliyun.com/api/ens/modifyimageattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyImageAttributeWithCallback(request *ModifyImageAttributeRequest, callback func(response *ModifyImageAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyImageAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyImageAttribute(request)
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

// ModifyImageAttributeRequest is the request struct for api ModifyImageAttribute
type ModifyImageAttributeRequest struct {
	*requests.RpcRequest
	ImageId   string `position:"Query" name:"ImageId"`
	ImageName string `position:"Query" name:"ImageName"`
	Product   string `position:"Query" name:"product"`
	Version   string `position:"Query" name:"Version"`
}

// ModifyImageAttributeResponse is the response struct for api ModifyImageAttribute
type ModifyImageAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
}

// CreateModifyImageAttributeRequest creates a request to invoke ModifyImageAttribute API
func CreateModifyImageAttributeRequest() (request *ModifyImageAttributeRequest) {
	request = &ModifyImageAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "ModifyImageAttribute", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyImageAttributeResponse creates a response to parse from ModifyImageAttribute response
func CreateModifyImageAttributeResponse() (response *ModifyImageAttributeResponse) {
	response = &ModifyImageAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
