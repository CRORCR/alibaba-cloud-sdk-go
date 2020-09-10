package domain

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

// QueryDomainTransferStatus invokes the domain.QueryDomainTransferStatus API synchronously
// api document: https://help.aliyun.com/api/domain/querydomaintransferstatus.html
func (client *Client) QueryDomainTransferStatus(request *QueryDomainTransferStatusRequest) (response *QueryDomainTransferStatusResponse, err error) {
	response = CreateQueryDomainTransferStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainTransferStatusWithChan invokes the domain.QueryDomainTransferStatus API asynchronously
// api document: https://help.aliyun.com/api/domain/querydomaintransferstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainTransferStatusWithChan(request *QueryDomainTransferStatusRequest) (<-chan *QueryDomainTransferStatusResponse, <-chan error) {
	responseChan := make(chan *QueryDomainTransferStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainTransferStatus(request)
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

// QueryDomainTransferStatusWithCallback invokes the domain.QueryDomainTransferStatus API asynchronously
// api document: https://help.aliyun.com/api/domain/querydomaintransferstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDomainTransferStatusWithCallback(request *QueryDomainTransferStatusRequest, callback func(response *QueryDomainTransferStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainTransferStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainTransferStatus(request)
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

// QueryDomainTransferStatusRequest is the request struct for api QueryDomainTransferStatus
type QueryDomainTransferStatusRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Body" name:"DomainName"`
}

// QueryDomainTransferStatusResponse is the response struct for api QueryDomainTransferStatus
type QueryDomainTransferStatusResponse struct {
	*responses.BaseResponse
	RequestId            string                     `json:"RequestId" xml:"RequestId"`
	DomainTransferStatus []DomainTransferStatusItem `json:"DomainTransferStatus" xml:"DomainTransferStatus"`
}

// CreateQueryDomainTransferStatusRequest creates a request to invoke QueryDomainTransferStatus API
func CreateQueryDomainTransferStatusRequest() (request *QueryDomainTransferStatusRequest) {
	request = &QueryDomainTransferStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "QueryDomainTransferStatus", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDomainTransferStatusResponse creates a response to parse from QueryDomainTransferStatus response
func CreateQueryDomainTransferStatusResponse() (response *QueryDomainTransferStatusResponse) {
	response = &QueryDomainTransferStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
