package cdn

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

// DescribeCdnDomainByCertificate invokes the cdn.DescribeCdnDomainByCertificate API synchronously
// api document: https://help.aliyun.com/api/cdn/describecdndomainbycertificate.html
func (client *Client) DescribeCdnDomainByCertificate(request *DescribeCdnDomainByCertificateRequest) (response *DescribeCdnDomainByCertificateResponse, err error) {
	response = CreateDescribeCdnDomainByCertificateResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnDomainByCertificateWithChan invokes the cdn.DescribeCdnDomainByCertificate API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdndomainbycertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnDomainByCertificateWithChan(request *DescribeCdnDomainByCertificateRequest) (<-chan *DescribeCdnDomainByCertificateResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnDomainByCertificateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnDomainByCertificate(request)
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

// DescribeCdnDomainByCertificateWithCallback invokes the cdn.DescribeCdnDomainByCertificate API asynchronously
// api document: https://help.aliyun.com/api/cdn/describecdndomainbycertificate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCdnDomainByCertificateWithCallback(request *DescribeCdnDomainByCertificateRequest, callback func(response *DescribeCdnDomainByCertificateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnDomainByCertificateResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnDomainByCertificate(request)
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

// DescribeCdnDomainByCertificateRequest is the request struct for api DescribeCdnDomainByCertificate
type DescribeCdnDomainByCertificateRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	SSLPub  string           `position:"Query" name:"SSLPub"`
}

// DescribeCdnDomainByCertificateResponse is the response struct for api DescribeCdnDomainByCertificate
type DescribeCdnDomainByCertificateResponse struct {
	*responses.BaseResponse
	RequestId string                                    `json:"RequestId" xml:"RequestId"`
	CertInfos CertInfosInDescribeCdnDomainByCertificate `json:"CertInfos" xml:"CertInfos"`
}

// CreateDescribeCdnDomainByCertificateRequest creates a request to invoke DescribeCdnDomainByCertificate API
func CreateDescribeCdnDomainByCertificateRequest() (request *DescribeCdnDomainByCertificateRequest) {
	request = &DescribeCdnDomainByCertificateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnDomainByCertificate", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCdnDomainByCertificateResponse creates a response to parse from DescribeCdnDomainByCertificate response
func CreateDescribeCdnDomainByCertificateResponse() (response *DescribeCdnDomainByCertificateResponse) {
	response = &DescribeCdnDomainByCertificateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
