package dysmsapi

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

// AddSmsSign invokes the dysmsapi.AddSmsSign API synchronously
// api document: https://help.aliyun.com/api/dysmsapi/addsmssign.html
func (client *Client) AddSmsSign(request *AddSmsSignRequest) (response *AddSmsSignResponse, err error) {
	response = CreateAddSmsSignResponse()
	err = client.DoAction(request, response)
	return
}

// AddSmsSignWithChan invokes the dysmsapi.AddSmsSign API asynchronously
// api document: https://help.aliyun.com/api/dysmsapi/addsmssign.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSmsSignWithChan(request *AddSmsSignRequest) (<-chan *AddSmsSignResponse, <-chan error) {
	responseChan := make(chan *AddSmsSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddSmsSign(request)
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

// AddSmsSignWithCallback invokes the dysmsapi.AddSmsSign API asynchronously
// api document: https://help.aliyun.com/api/dysmsapi/addsmssign.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddSmsSignWithCallback(request *AddSmsSignRequest, callback func(response *AddSmsSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddSmsSignResponse
		var err error
		defer close(result)
		response, err = client.AddSmsSign(request)
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

// AddSmsSignRequest is the request struct for api AddSmsSign
type AddSmsSignRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer          `position:"Query" name:"ResourceOwnerId"`
	Remark               string                    `position:"Query" name:"Remark"`
	SignName             string                    `position:"Query" name:"SignName"`
	SignFileList         *[]AddSmsSignSignFileList `position:"Query" name:"SignFileList"  type:"Repeated"`
	ResourceOwnerAccount string                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer          `position:"Query" name:"OwnerId"`
	SignSource           requests.Integer          `position:"Query" name:"SignSource"`
}

// AddSmsSignSignFileList is a repeated param struct in AddSmsSignRequest
type AddSmsSignSignFileList struct {
	FileContents string `name:"FileContents"`
	FileSuffix   string `name:"FileSuffix"`
}

// AddSmsSignResponse is the response struct for api AddSmsSign
type AddSmsSignResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	SignName  string `json:"SignName" xml:"SignName"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateAddSmsSignRequest creates a request to invoke AddSmsSign API
func CreateAddSmsSignRequest() (request *AddSmsSignRequest) {
	request = &AddSmsSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "AddSmsSign", "dysms", "openAPI")
	return
}

// CreateAddSmsSignResponse creates a response to parse from AddSmsSign response
func CreateAddSmsSignResponse() (response *AddSmsSignResponse) {
	response = &AddSmsSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
