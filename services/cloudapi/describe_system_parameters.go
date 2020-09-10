package cloudapi

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

// DescribeSystemParameters invokes the cloudapi.DescribeSystemParameters API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describesystemparameters.html
func (client *Client) DescribeSystemParameters(request *DescribeSystemParametersRequest) (response *DescribeSystemParametersResponse, err error) {
	response = CreateDescribeSystemParametersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSystemParametersWithChan invokes the cloudapi.DescribeSystemParameters API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describesystemparameters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSystemParametersWithChan(request *DescribeSystemParametersRequest) (<-chan *DescribeSystemParametersResponse, <-chan error) {
	responseChan := make(chan *DescribeSystemParametersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSystemParameters(request)
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

// DescribeSystemParametersWithCallback invokes the cloudapi.DescribeSystemParameters API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describesystemparameters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSystemParametersWithCallback(request *DescribeSystemParametersRequest, callback func(response *DescribeSystemParametersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSystemParametersResponse
		var err error
		defer close(result)
		response, err = client.DescribeSystemParameters(request)
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

// DescribeSystemParametersRequest is the request struct for api DescribeSystemParameters
type DescribeSystemParametersRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DescribeSystemParametersResponse is the response struct for api DescribeSystemParameters
type DescribeSystemParametersResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	SystemParams SystemParams `json:"SystemParams" xml:"SystemParams"`
}

// CreateDescribeSystemParametersRequest creates a request to invoke DescribeSystemParameters API
func CreateDescribeSystemParametersRequest() (request *DescribeSystemParametersRequest) {
	request = &DescribeSystemParametersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeSystemParameters", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSystemParametersResponse creates a response to parse from DescribeSystemParameters response
func CreateDescribeSystemParametersResponse() (response *DescribeSystemParametersResponse) {
	response = &DescribeSystemParametersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
