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

// SaveSingleTaskForCreatingDnsHost invokes the domain_intl.SaveSingleTaskForCreatingDnsHost API synchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskforcreatingdnshost.html
func (client *Client) SaveSingleTaskForCreatingDnsHost(request *SaveSingleTaskForCreatingDnsHostRequest) (response *SaveSingleTaskForCreatingDnsHostResponse, err error) {
	response = CreateSaveSingleTaskForCreatingDnsHostResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForCreatingDnsHostWithChan invokes the domain_intl.SaveSingleTaskForCreatingDnsHost API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskforcreatingdnshost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForCreatingDnsHostWithChan(request *SaveSingleTaskForCreatingDnsHostRequest) (<-chan *SaveSingleTaskForCreatingDnsHostResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForCreatingDnsHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForCreatingDnsHost(request)
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

// SaveSingleTaskForCreatingDnsHostWithCallback invokes the domain_intl.SaveSingleTaskForCreatingDnsHost API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savesingletaskforcreatingdnshost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForCreatingDnsHostWithCallback(request *SaveSingleTaskForCreatingDnsHostRequest, callback func(response *SaveSingleTaskForCreatingDnsHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForCreatingDnsHostResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForCreatingDnsHost(request)
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

// SaveSingleTaskForCreatingDnsHostRequest is the request struct for api SaveSingleTaskForCreatingDnsHost
type SaveSingleTaskForCreatingDnsHostRequest struct {
	*requests.RpcRequest
	InstanceId   string    `position:"Query" name:"InstanceId"`
	Ip           *[]string `position:"Query" name:"Ip"  type:"Repeated"`
	DnsName      string    `position:"Query" name:"DnsName"`
	UserClientIp string    `position:"Query" name:"UserClientIp"`
	Lang         string    `position:"Query" name:"Lang"`
}

// SaveSingleTaskForCreatingDnsHostResponse is the response struct for api SaveSingleTaskForCreatingDnsHost
type SaveSingleTaskForCreatingDnsHostResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForCreatingDnsHostRequest creates a request to invoke SaveSingleTaskForCreatingDnsHost API
func CreateSaveSingleTaskForCreatingDnsHostRequest() (request *SaveSingleTaskForCreatingDnsHostRequest) {
	request = &SaveSingleTaskForCreatingDnsHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForCreatingDnsHost", "domain", "openAPI")
	return
}

// CreateSaveSingleTaskForCreatingDnsHostResponse creates a response to parse from SaveSingleTaskForCreatingDnsHost response
func CreateSaveSingleTaskForCreatingDnsHostResponse() (response *SaveSingleTaskForCreatingDnsHostResponse) {
	response = &SaveSingleTaskForCreatingDnsHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
