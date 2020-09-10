package kms

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

// DescribeService invokes the kms.DescribeService API synchronously
// api document: https://help.aliyun.com/api/kms/describeservice.html
func (client *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
	response = CreateDescribeServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeServiceWithChan invokes the kms.DescribeService API asynchronously
// api document: https://help.aliyun.com/api/kms/describeservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServiceWithChan(request *DescribeServiceRequest) (<-chan *DescribeServiceResponse, <-chan error) {
	responseChan := make(chan *DescribeServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeService(request)
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

// DescribeServiceWithCallback invokes the kms.DescribeService API asynchronously
// api document: https://help.aliyun.com/api/kms/describeservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServiceWithCallback(request *DescribeServiceRequest, callback func(response *DescribeServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeServiceResponse
		var err error
		defer close(result)
		response, err = client.DescribeService(request)
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

// DescribeServiceRequest is the request struct for api DescribeService
type DescribeServiceRequest struct {
	*requests.RpcRequest
}

// DescribeServiceResponse is the response struct for api DescribeService
type DescribeServiceResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	ProtectionLevels ProtectionLevels `json:"ProtectionLevels" xml:"ProtectionLevels"`
	KeySpecs         KeySpecs         `json:"KeySpecs" xml:"KeySpecs"`
}

// CreateDescribeServiceRequest creates a request to invoke DescribeService API
func CreateDescribeServiceRequest() (request *DescribeServiceRequest) {
	request = &DescribeServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "DescribeService", "kms-service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeServiceResponse creates a response to parse from DescribeService response
func CreateDescribeServiceResponse() (response *DescribeServiceResponse) {
	response = &DescribeServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
