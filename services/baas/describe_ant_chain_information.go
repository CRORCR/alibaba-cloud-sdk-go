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

// DescribeAntChainInformation invokes the baas.DescribeAntChainInformation API synchronously
// api document: https://help.aliyun.com/api/baas/describeantchaininformation.html
func (client *Client) DescribeAntChainInformation(request *DescribeAntChainInformationRequest) (response *DescribeAntChainInformationResponse, err error) {
	response = CreateDescribeAntChainInformationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAntChainInformationWithChan invokes the baas.DescribeAntChainInformation API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchaininformation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainInformationWithChan(request *DescribeAntChainInformationRequest) (<-chan *DescribeAntChainInformationResponse, <-chan error) {
	responseChan := make(chan *DescribeAntChainInformationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAntChainInformation(request)
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

// DescribeAntChainInformationWithCallback invokes the baas.DescribeAntChainInformation API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchaininformation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainInformationWithCallback(request *DescribeAntChainInformationRequest, callback func(response *DescribeAntChainInformationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAntChainInformationResponse
		var err error
		defer close(result)
		response, err = client.DescribeAntChainInformation(request)
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

// DescribeAntChainInformationRequest is the request struct for api DescribeAntChainInformation
type DescribeAntChainInformationRequest struct {
	*requests.RpcRequest
	AntChainId string `position:"Body" name:"AntChainId"`
}

// DescribeAntChainInformationResponse is the response struct for api DescribeAntChainInformation
type DescribeAntChainInformationResponse struct {
	*responses.BaseResponse
	RequestId string                              `json:"RequestId" xml:"RequestId"`
	Result    ResultInDescribeAntChainInformation `json:"Result" xml:"Result"`
}

// CreateDescribeAntChainInformationRequest creates a request to invoke DescribeAntChainInformation API
func CreateDescribeAntChainInformationRequest() (request *DescribeAntChainInformationRequest) {
	request = &DescribeAntChainInformationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeAntChainInformation", "baas", "openAPI")
	return
}

// CreateDescribeAntChainInformationResponse creates a response to parse from DescribeAntChainInformation response
func CreateDescribeAntChainInformationResponse() (response *DescribeAntChainInformationResponse) {
	response = &DescribeAntChainInformationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
