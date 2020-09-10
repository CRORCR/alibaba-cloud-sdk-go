package dm

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

// DeleteSign invokes the dm.DeleteSign API synchronously
// api document: https://help.aliyun.com/api/dm/deletesign.html
func (client *Client) DeleteSign(request *DeleteSignRequest) (response *DeleteSignResponse, err error) {
	response = CreateDeleteSignResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSignWithChan invokes the dm.DeleteSign API asynchronously
// api document: https://help.aliyun.com/api/dm/deletesign.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSignWithChan(request *DeleteSignRequest) (<-chan *DeleteSignResponse, <-chan error) {
	responseChan := make(chan *DeleteSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSign(request)
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

// DeleteSignWithCallback invokes the dm.DeleteSign API asynchronously
// api document: https://help.aliyun.com/api/dm/deletesign.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSignWithCallback(request *DeleteSignRequest, callback func(response *DeleteSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSignResponse
		var err error
		defer close(result)
		response, err = client.DeleteSign(request)
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

// DeleteSignRequest is the request struct for api DeleteSign
type DeleteSignRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SignId               requests.Integer `position:"Query" name:"SignId"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
}

// DeleteSignResponse is the response struct for api DeleteSign
type DeleteSignResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSignRequest creates a request to invoke DeleteSign API
func CreateDeleteSignRequest() (request *DeleteSignRequest) {
	request = &DeleteSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "DeleteSign", "", "")
	return
}

// CreateDeleteSignResponse creates a response to parse from DeleteSign response
func CreateDeleteSignResponse() (response *DeleteSignResponse) {
	response = &DeleteSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
