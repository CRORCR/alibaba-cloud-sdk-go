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

// DescribeAntChainBlock invokes the baas.DescribeAntChainBlock API synchronously
// api document: https://help.aliyun.com/api/baas/describeantchainblock.html
func (client *Client) DescribeAntChainBlock(request *DescribeAntChainBlockRequest) (response *DescribeAntChainBlockResponse, err error) {
	response = CreateDescribeAntChainBlockResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAntChainBlockWithChan invokes the baas.DescribeAntChainBlock API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchainblock.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainBlockWithChan(request *DescribeAntChainBlockRequest) (<-chan *DescribeAntChainBlockResponse, <-chan error) {
	responseChan := make(chan *DescribeAntChainBlockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAntChainBlock(request)
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

// DescribeAntChainBlockWithCallback invokes the baas.DescribeAntChainBlock API asynchronously
// api document: https://help.aliyun.com/api/baas/describeantchainblock.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAntChainBlockWithCallback(request *DescribeAntChainBlockRequest, callback func(response *DescribeAntChainBlockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAntChainBlockResponse
		var err error
		defer close(result)
		response, err = client.DescribeAntChainBlock(request)
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

// DescribeAntChainBlockRequest is the request struct for api DescribeAntChainBlock
type DescribeAntChainBlockRequest struct {
	*requests.RpcRequest
	AntChainId string           `position:"Body" name:"AntChainId"`
	Height     requests.Integer `position:"Body" name:"Height"`
}

// DescribeAntChainBlockResponse is the response struct for api DescribeAntChainBlock
type DescribeAntChainBlockResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeAntChainBlockRequest creates a request to invoke DescribeAntChainBlock API
func CreateDescribeAntChainBlockRequest() (request *DescribeAntChainBlockRequest) {
	request = &DescribeAntChainBlockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeAntChainBlock", "baas", "openAPI")
	return
}

// CreateDescribeAntChainBlockResponse creates a response to parse from DescribeAntChainBlock response
func CreateDescribeAntChainBlockResponse() (response *DescribeAntChainBlockResponse) {
	response = &DescribeAntChainBlockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
