package aegis

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

// ModifyEmgVulSubmit invokes the aegis.ModifyEmgVulSubmit API synchronously
// api document: https://help.aliyun.com/api/aegis/modifyemgvulsubmit.html
func (client *Client) ModifyEmgVulSubmit(request *ModifyEmgVulSubmitRequest) (response *ModifyEmgVulSubmitResponse, err error) {
	response = CreateModifyEmgVulSubmitResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyEmgVulSubmitWithChan invokes the aegis.ModifyEmgVulSubmit API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyemgvulsubmit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyEmgVulSubmitWithChan(request *ModifyEmgVulSubmitRequest) (<-chan *ModifyEmgVulSubmitResponse, <-chan error) {
	responseChan := make(chan *ModifyEmgVulSubmitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyEmgVulSubmit(request)
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

// ModifyEmgVulSubmitWithCallback invokes the aegis.ModifyEmgVulSubmit API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyemgvulsubmit.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyEmgVulSubmitWithCallback(request *ModifyEmgVulSubmitRequest, callback func(response *ModifyEmgVulSubmitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyEmgVulSubmitResponse
		var err error
		defer close(result)
		response, err = client.ModifyEmgVulSubmit(request)
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

// ModifyEmgVulSubmitRequest is the request struct for api ModifyEmgVulSubmit
type ModifyEmgVulSubmitRequest struct {
	*requests.RpcRequest
	SourceIp      string `position:"Query" name:"SourceIp"`
	Name          string `position:"Query" name:"Name"`
	UserAgreement string `position:"Query" name:"UserAgreement"`
	Lang          string `position:"Query" name:"Lang"`
}

// ModifyEmgVulSubmitResponse is the response struct for api ModifyEmgVulSubmit
type ModifyEmgVulSubmitResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyEmgVulSubmitRequest creates a request to invoke ModifyEmgVulSubmit API
func CreateModifyEmgVulSubmitRequest() (request *ModifyEmgVulSubmitRequest) {
	request = &ModifyEmgVulSubmitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyEmgVulSubmit", "vipaegis", "openAPI")
	return
}

// CreateModifyEmgVulSubmitResponse creates a response to parse from ModifyEmgVulSubmit response
func CreateModifyEmgVulSubmitResponse() (response *ModifyEmgVulSubmitResponse) {
	response = &ModifyEmgVulSubmitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
