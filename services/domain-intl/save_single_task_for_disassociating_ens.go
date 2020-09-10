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

// SaveSingleTaskForDisassociatingEns invokes the domain_intl.SaveSingleTaskForDisassociatingEns API synchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskfordisassociatingens.html
func (client *Client) SaveSingleTaskForDisassociatingEns(request *SaveSingleTaskForDisassociatingEnsRequest) (response *SaveSingleTaskForDisassociatingEnsResponse, err error) {
	response = CreateSaveSingleTaskForDisassociatingEnsResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForDisassociatingEnsWithChan invokes the domain_intl.SaveSingleTaskForDisassociatingEns API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskfordisassociatingens.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForDisassociatingEnsWithChan(request *SaveSingleTaskForDisassociatingEnsRequest) (<-chan *SaveSingleTaskForDisassociatingEnsResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForDisassociatingEnsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForDisassociatingEns(request)
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

// SaveSingleTaskForDisassociatingEnsWithCallback invokes the domain_intl.SaveSingleTaskForDisassociatingEns API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskfordisassociatingens.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForDisassociatingEnsWithCallback(request *SaveSingleTaskForDisassociatingEnsRequest, callback func(response *SaveSingleTaskForDisassociatingEnsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForDisassociatingEnsResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForDisassociatingEns(request)
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

// SaveSingleTaskForDisassociatingEnsRequest is the request struct for api SaveSingleTaskForDisassociatingEns
type SaveSingleTaskForDisassociatingEnsRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// SaveSingleTaskForDisassociatingEnsResponse is the response struct for api SaveSingleTaskForDisassociatingEns
type SaveSingleTaskForDisassociatingEnsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForDisassociatingEnsRequest creates a request to invoke SaveSingleTaskForDisassociatingEns API
func CreateSaveSingleTaskForDisassociatingEnsRequest() (request *SaveSingleTaskForDisassociatingEnsRequest) {
	request = &SaveSingleTaskForDisassociatingEnsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForDisassociatingEns", "domain", "openAPI")
	return
}

// CreateSaveSingleTaskForDisassociatingEnsResponse creates a response to parse from SaveSingleTaskForDisassociatingEns response
func CreateSaveSingleTaskForDisassociatingEnsResponse() (response *SaveSingleTaskForDisassociatingEnsResponse) {
	response = &SaveSingleTaskForDisassociatingEnsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
