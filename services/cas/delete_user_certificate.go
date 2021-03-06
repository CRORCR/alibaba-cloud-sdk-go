package cas

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

// DeleteUserCertificate invokes the cas.DeleteUserCertificate API synchronously
// api document: https://help.aliyun.com/api/cas/deleteusercertificate.html
func (client *Client) DeleteUserCertificate(request *DeleteUserCertificateRequest) (response *DeleteUserCertificateResponse, err error) {
	response = CreateDeleteUserCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUserCertificateWithChan invokes the cas.DeleteUserCertificate API asynchronously
// api document: https://help.aliyun.com/api/cas/deleteusercertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUserCertificateWithChan(request *DeleteUserCertificateRequest) (<-chan *DeleteUserCertificateResponse, <-chan error) {
	responseChan := make(chan *DeleteUserCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUserCertificate(request)
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

// DeleteUserCertificateWithCallback invokes the cas.DeleteUserCertificate API asynchronously
// api document: https://help.aliyun.com/api/cas/deleteusercertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteUserCertificateWithCallback(request *DeleteUserCertificateRequest, callback func(response *DeleteUserCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUserCertificateResponse
		var err error
		defer close(result)
		response, err = client.DeleteUserCertificate(request)
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

// DeleteUserCertificateRequest is the request struct for api DeleteUserCertificate
type DeleteUserCertificateRequest struct {
	*requests.RpcRequest
	CertId   requests.Integer `position:"Query" name:"CertId"`
	SourceIp string           `position:"Query" name:"SourceIp"`
	Lang     string           `position:"Query" name:"Lang"`
}

// DeleteUserCertificateResponse is the response struct for api DeleteUserCertificate
type DeleteUserCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteUserCertificateRequest creates a request to invoke DeleteUserCertificate API
func CreateDeleteUserCertificateRequest() (request *DeleteUserCertificateRequest) {
	request = &DeleteUserCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2018-07-13", "DeleteUserCertificate", "cas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteUserCertificateResponse creates a response to parse from DeleteUserCertificate response
func CreateDeleteUserCertificateResponse() (response *DeleteUserCertificateResponse) {
	response = &DeleteUserCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
