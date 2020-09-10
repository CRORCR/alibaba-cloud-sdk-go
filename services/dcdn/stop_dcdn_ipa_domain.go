package dcdn

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

// StopDcdnIpaDomain invokes the dcdn.StopDcdnIpaDomain API synchronously
// api document: https://help.aliyun.com/api/dcdn/stopdcdnipadomain.html
func (client *Client) StopDcdnIpaDomain(request *StopDcdnIpaDomainRequest) (response *StopDcdnIpaDomainResponse, err error) {
	response = CreateStopDcdnIpaDomainResponse()
	err = client.DoAction(request, response)
	return
}

// StopDcdnIpaDomainWithChan invokes the dcdn.StopDcdnIpaDomain API asynchronously
// api document: https://help.aliyun.com/api/dcdn/stopdcdnipadomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopDcdnIpaDomainWithChan(request *StopDcdnIpaDomainRequest) (<-chan *StopDcdnIpaDomainResponse, <-chan error) {
	responseChan := make(chan *StopDcdnIpaDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopDcdnIpaDomain(request)
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

// StopDcdnIpaDomainWithCallback invokes the dcdn.StopDcdnIpaDomain API asynchronously
// api document: https://help.aliyun.com/api/dcdn/stopdcdnipadomain.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StopDcdnIpaDomainWithCallback(request *StopDcdnIpaDomainRequest, callback func(response *StopDcdnIpaDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopDcdnIpaDomainResponse
		var err error
		defer close(result)
		response, err = client.StopDcdnIpaDomain(request)
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

// StopDcdnIpaDomainRequest is the request struct for api StopDcdnIpaDomain
type StopDcdnIpaDomainRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// StopDcdnIpaDomainResponse is the response struct for api StopDcdnIpaDomain
type StopDcdnIpaDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopDcdnIpaDomainRequest creates a request to invoke StopDcdnIpaDomain API
func CreateStopDcdnIpaDomainRequest() (request *StopDcdnIpaDomainRequest) {
	request = &StopDcdnIpaDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "StopDcdnIpaDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateStopDcdnIpaDomainResponse creates a response to parse from StopDcdnIpaDomain response
func CreateStopDcdnIpaDomainResponse() (response *StopDcdnIpaDomainResponse) {
	response = &StopDcdnIpaDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
