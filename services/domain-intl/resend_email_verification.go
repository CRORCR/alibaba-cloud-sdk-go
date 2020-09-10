package domain_intl

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

// ResendEmailVerification invokes the domain_intl.ResendEmailVerification API synchronously
// api document: https://help.aliyun.com/api/domain-intl/resendemailverification.html
func (client *Client) ResendEmailVerification(request *ResendEmailVerificationRequest) (response *ResendEmailVerificationResponse, err error) {
	response = CreateResendEmailVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// ResendEmailVerificationWithChan invokes the domain_intl.ResendEmailVerification API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/resendemailverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResendEmailVerificationWithChan(request *ResendEmailVerificationRequest) (<-chan *ResendEmailVerificationResponse, <-chan error) {
	responseChan := make(chan *ResendEmailVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResendEmailVerification(request)
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

// ResendEmailVerificationWithCallback invokes the domain_intl.ResendEmailVerification API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/resendemailverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResendEmailVerificationWithCallback(request *ResendEmailVerificationRequest, callback func(response *ResendEmailVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResendEmailVerificationResponse
		var err error
		defer close(result)
		response, err = client.ResendEmailVerification(request)
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

// ResendEmailVerificationRequest is the request struct for api ResendEmailVerification
type ResendEmailVerificationRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Email        string `position:"Query" name:"Email"`
}

// ResendEmailVerificationResponse is the response struct for api ResendEmailVerification
type ResendEmailVerificationResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	SuccessList []SendResult `json:"SuccessList" xml:"SuccessList"`
	FailList    []SendResult `json:"FailList" xml:"FailList"`
}

// CreateResendEmailVerificationRequest creates a request to invoke ResendEmailVerification API
func CreateResendEmailVerificationRequest() (request *ResendEmailVerificationRequest) {
	request = &ResendEmailVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "ResendEmailVerification", "domain", "openAPI")
	return
}

// CreateResendEmailVerificationResponse creates a response to parse from ResendEmailVerification response
func CreateResendEmailVerificationResponse() (response *ResendEmailVerificationResponse) {
	response = &ResendEmailVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
