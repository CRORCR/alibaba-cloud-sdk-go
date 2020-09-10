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

// DescribeAntChainTransaction invokes the baas.DescribeAntChainTransaction API synchronously
// api document: https://help.aliyun.com/api/baas/describeantchaintransaction.html
func (client *Client) DescribeAntChainTransaction(request *DescribeAntChainTransactionRequest) (response *DescribeAntChainTransactionResponse, err error) {
	response = CreateDescribeAntChainTransactionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAntChainTransactionWithChan invokes the baas.DescribeAntChainTransaction API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchaintransaction.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainTransactionWithChan(request *DescribeAntChainTransactionRequest) (<-chan *DescribeAntChainTransactionResponse, <-chan error) {
	responseChan := make(chan *DescribeAntChainTransactionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAntChainTransaction(request)
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

// DescribeAntChainTransactionWithCallback invokes the baas.DescribeAntChainTransaction API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchaintransaction.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainTransactionWithCallback(request *DescribeAntChainTransactionRequest, callback func(response *DescribeAntChainTransactionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAntChainTransactionResponse
		var err error
		defer close(result)
		response, err = client.DescribeAntChainTransaction(request)
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

// DescribeAntChainTransactionRequest is the request struct for api DescribeAntChainTransaction
type DescribeAntChainTransactionRequest struct {
	*requests.RpcRequest
	AntChainId string `position:"Body" name:"AntChainId"`
	Hash       string `position:"Body" name:"Hash"`
}

// DescribeAntChainTransactionResponse is the response struct for api DescribeAntChainTransaction
type DescribeAntChainTransactionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeAntChainTransactionRequest creates a request to invoke DescribeAntChainTransaction API
func CreateDescribeAntChainTransactionRequest() (request *DescribeAntChainTransactionRequest) {
	request = &DescribeAntChainTransactionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeAntChainTransaction", "baas", "openAPI")
	return
}

// CreateDescribeAntChainTransactionResponse creates a response to parse from DescribeAntChainTransaction response
func CreateDescribeAntChainTransactionResponse() (response *DescribeAntChainTransactionResponse) {
	response = &DescribeAntChainTransactionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
