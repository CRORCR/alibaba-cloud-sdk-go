package baas

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

// ApplyAntChainCertificate invokes the baas.ApplyAntChainCertificate API synchronously
// api document: https://help.aliyun.com/api/baas/applyantchaincertificate.html
func (client *Client) ApplyAntChainCertificate(request *ApplyAntChainCertificateRequest) (response *ApplyAntChainCertificateResponse, err error) {
	response = CreateApplyAntChainCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// ApplyAntChainCertificateWithChan invokes the baas.ApplyAntChainCertificate API asynchronously
// api document: https://help.aliyun.com/api/baas/applyantchaincertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyAntChainCertificateWithChan(request *ApplyAntChainCertificateRequest) (<-chan *ApplyAntChainCertificateResponse, <-chan error) {
	responseChan := make(chan *ApplyAntChainCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyAntChainCertificate(request)
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

// ApplyAntChainCertificateWithCallback invokes the baas.ApplyAntChainCertificate API asynchronously
// api document: https://help.aliyun.com/api/baas/applyantchaincertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ApplyAntChainCertificateWithCallback(request *ApplyAntChainCertificateRequest, callback func(response *ApplyAntChainCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyAntChainCertificateResponse
		var err error
		defer close(result)
		response, err = client.ApplyAntChainCertificate(request)
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

// ApplyAntChainCertificateRequest is the request struct for api ApplyAntChainCertificate
type ApplyAntChainCertificateRequest struct {
	*requests.RpcRequest
	UploadReq  string `position:"Body" name:"UploadReq"`
	AntChainId string `position:"Body" name:"AntChainId"`
}

// ApplyAntChainCertificateResponse is the response struct for api ApplyAntChainCertificate
type ApplyAntChainCertificateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateApplyAntChainCertificateRequest creates a request to invoke ApplyAntChainCertificate API
func CreateApplyAntChainCertificateRequest() (request *ApplyAntChainCertificateRequest) {
	request = &ApplyAntChainCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "ApplyAntChainCertificate", "baas", "openAPI")
	return
}

// CreateApplyAntChainCertificateResponse creates a response to parse from ApplyAntChainCertificate response
func CreateApplyAntChainCertificateResponse() (response *ApplyAntChainCertificateResponse) {
	response = &ApplyAntChainCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
