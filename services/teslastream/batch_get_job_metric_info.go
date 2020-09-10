package teslastream

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

// BatchGetJobMetricInfo invokes the teslastream.BatchGetJobMetricInfo API synchronously
// api document: https://help.aliyun.com/api/teslastream/batchgetjobmetricinfo.html
func (client *Client) BatchGetJobMetricInfo(request *BatchGetJobMetricInfoRequest) (response *BatchGetJobMetricInfoResponse, err error) {
	response = CreateBatchGetJobMetricInfoResponse()
	err = client.DoAction(request, response)
	return
}

// BatchGetJobMetricInfoWithChan invokes the teslastream.BatchGetJobMetricInfo API asynchronously
// api document: https://help.aliyun.com/api/teslastream/batchgetjobmetricinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetJobMetricInfoWithChan(request *BatchGetJobMetricInfoRequest) (<-chan *BatchGetJobMetricInfoResponse, <-chan error) {
	responseChan := make(chan *BatchGetJobMetricInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchGetJobMetricInfo(request)
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

// BatchGetJobMetricInfoWithCallback invokes the teslastream.BatchGetJobMetricInfo API asynchronously
// api document: https://help.aliyun.com/api/teslastream/batchgetjobmetricinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetJobMetricInfoWithCallback(request *BatchGetJobMetricInfoRequest, callback func(response *BatchGetJobMetricInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchGetJobMetricInfoResponse
		var err error
		defer close(result)
		response, err = client.BatchGetJobMetricInfo(request)
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

// BatchGetJobMetricInfoRequest is the request struct for api BatchGetJobMetricInfo
type BatchGetJobMetricInfoRequest struct {
	*requests.RpcRequest
	JobInfos string `position:"Query" name:"JobInfos"`
}

// BatchGetJobMetricInfoResponse is the response struct for api BatchGetJobMetricInfo
type BatchGetJobMetricInfoResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      []Job  `json:"Data" xml:"Data"`
}

// CreateBatchGetJobMetricInfoRequest creates a request to invoke BatchGetJobMetricInfo API
func CreateBatchGetJobMetricInfoRequest() (request *BatchGetJobMetricInfoRequest) {
	request = &BatchGetJobMetricInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaStream", "2018-01-15", "BatchGetJobMetricInfo", "teslastream", "openAPI")
	return
}

// CreateBatchGetJobMetricInfoResponse creates a response to parse from BatchGetJobMetricInfo response
func CreateBatchGetJobMetricInfoResponse() (response *BatchGetJobMetricInfoResponse) {
	response = &BatchGetJobMetricInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
