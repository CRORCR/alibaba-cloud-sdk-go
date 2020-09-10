package trademark

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

// RefundProduce invokes the trademark.RefundProduce API synchronously
// api document: https://help.aliyun.com/api/trademark/refundproduce.html
func (client *Client) RefundProduce(request *RefundProduceRequest) (response *RefundProduceResponse, err error) {
	response = CreateRefundProduceResponse()
	err = client.DoAction(request, response)
	return
}

// RefundProduceWithChan invokes the trademark.RefundProduce API asynchronously
// api document: https://help.aliyun.com/api/trademark/refundproduce.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefundProduceWithChan(request *RefundProduceRequest) (<-chan *RefundProduceResponse, <-chan error) {
	responseChan := make(chan *RefundProduceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefundProduce(request)
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

// RefundProduceWithCallback invokes the trademark.RefundProduce API asynchronously
// api document: https://help.aliyun.com/api/trademark/refundproduce.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefundProduceWithCallback(request *RefundProduceRequest, callback func(response *RefundProduceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefundProduceResponse
		var err error
		defer close(result)
		response, err = client.RefundProduce(request)
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

// RefundProduceRequest is the request struct for api RefundProduce
type RefundProduceRequest struct {
	*requests.RpcRequest
	BizId string `position:"Query" name:"BizId"`
}

// RefundProduceResponse is the response struct for api RefundProduce
type RefundProduceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateRefundProduceRequest creates a request to invoke RefundProduce API
func CreateRefundProduceRequest() (request *RefundProduceRequest) {
	request = &RefundProduceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "RefundProduce", "trademark", "openAPI")
	return
}

// CreateRefundProduceResponse creates a response to parse from RefundProduce response
func CreateRefundProduceResponse() (response *RefundProduceResponse) {
	response = &RefundProduceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
