package scdn

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

// StartScdnDomain invokes the scdn.StartScdnDomain API synchronously
// api document: https://help.aliyun.com/api/scdn/startscdndomain.html
func (client *Client) StartScdnDomain(request *StartScdnDomainRequest) (response *StartScdnDomainResponse, err error) {
	response = CreateStartScdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// StartScdnDomainWithChan invokes the scdn.StartScdnDomain API asynchronously
// api document: https://help.aliyun.com/api/scdn/startscdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartScdnDomainWithChan(request *StartScdnDomainRequest) (<-chan *StartScdnDomainResponse, <-chan error) {
	responseChan := make(chan *StartScdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartScdnDomain(request)
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

// StartScdnDomainWithCallback invokes the scdn.StartScdnDomain API asynchronously
// api document: https://help.aliyun.com/api/scdn/startscdndomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartScdnDomainWithCallback(request *StartScdnDomainRequest, callback func(response *StartScdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartScdnDomainResponse
		var err error
		defer close(result)
		response, err = client.StartScdnDomain(request)
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

// StartScdnDomainRequest is the request struct for api StartScdnDomain
type StartScdnDomainRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// StartScdnDomainResponse is the response struct for api StartScdnDomain
type StartScdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartScdnDomainRequest creates a request to invoke StartScdnDomain API
func CreateStartScdnDomainRequest() (request *StartScdnDomainRequest) {
	request = &StartScdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "StartScdnDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateStartScdnDomainResponse creates a response to parse from StartScdnDomain response
func CreateStartScdnDomainResponse() (response *StartScdnDomainResponse) {
	response = &StartScdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
