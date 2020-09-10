package smartag

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

// DescribeSagCurrentDns invokes the smartag.DescribeSagCurrentDns API synchronously
// api document: https://help.aliyun.com/api/smartag/describesagcurrentdns.html
func (client *Client) DescribeSagCurrentDns(request *DescribeSagCurrentDnsRequest) (response *DescribeSagCurrentDnsResponse, err error) {
	response = CreateDescribeSagCurrentDnsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagCurrentDnsWithChan invokes the smartag.DescribeSagCurrentDns API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesagcurrentdns.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagCurrentDnsWithChan(request *DescribeSagCurrentDnsRequest) (<-chan *DescribeSagCurrentDnsResponse, <-chan error) {
	responseChan := make(chan *DescribeSagCurrentDnsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagCurrentDns(request)
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

// DescribeSagCurrentDnsWithCallback invokes the smartag.DescribeSagCurrentDns API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesagcurrentdns.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagCurrentDnsWithCallback(request *DescribeSagCurrentDnsRequest, callback func(response *DescribeSagCurrentDnsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagCurrentDnsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagCurrentDns(request)
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

// DescribeSagCurrentDnsRequest is the request struct for api DescribeSagCurrentDns
type DescribeSagCurrentDnsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
}

// DescribeSagCurrentDnsResponse is the response struct for api DescribeSagCurrentDns
type DescribeSagCurrentDnsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MasterDns string `json:"MasterDns" xml:"MasterDns"`
	SlaveDns  string `json:"SlaveDns" xml:"SlaveDns"`
}

// CreateDescribeSagCurrentDnsRequest creates a request to invoke DescribeSagCurrentDns API
func CreateDescribeSagCurrentDnsRequest() (request *DescribeSagCurrentDnsRequest) {
	request = &DescribeSagCurrentDnsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagCurrentDns", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagCurrentDnsResponse creates a response to parse from DescribeSagCurrentDns response
func CreateDescribeSagCurrentDnsResponse() (response *DescribeSagCurrentDnsResponse) {
	response = &DescribeSagCurrentDnsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
