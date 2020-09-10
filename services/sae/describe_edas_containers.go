package sae

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

// DescribeEdasContainers invokes the sae.DescribeEdasContainers API synchronously
// api document: https://help.aliyun.com/api/sae/describeedascontainers.html
func (client *Client) DescribeEdasContainers(request *DescribeEdasContainersRequest) (response *DescribeEdasContainersResponse, err error) {
	response = CreateDescribeEdasContainersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEdasContainersWithChan invokes the sae.DescribeEdasContainers API asynchronously
// api document: https://help.aliyun.com/api/sae/describeedascontainers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEdasContainersWithChan(request *DescribeEdasContainersRequest) (<-chan *DescribeEdasContainersResponse, <-chan error) {
	responseChan := make(chan *DescribeEdasContainersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEdasContainers(request)
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

// DescribeEdasContainersWithCallback invokes the sae.DescribeEdasContainers API asynchronously
// api document: https://help.aliyun.com/api/sae/describeedascontainers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEdasContainersWithCallback(request *DescribeEdasContainersRequest, callback func(response *DescribeEdasContainersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEdasContainersResponse
		var err error
		defer close(result)
		response, err = client.DescribeEdasContainers(request)
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

// DescribeEdasContainersRequest is the request struct for api DescribeEdasContainers
type DescribeEdasContainersRequest struct {
	*requests.RoaRequest
}

// DescribeEdasContainersResponse is the response struct for api DescribeEdasContainers
type DescribeEdasContainersResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Code      string     `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	Success   bool       `json:"Success" xml:"Success"`
	ErrorCode string     `json:"ErrorCode" xml:"ErrorCode"`
	TraceId   string     `json:"TraceId" xml:"TraceId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeEdasContainersRequest creates a request to invoke DescribeEdasContainers API
func CreateDescribeEdasContainersRequest() (request *DescribeEdasContainersRequest) {
	request = &DescribeEdasContainersRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeEdasContainers", "/pop/v1/sam/resource/edasContainers", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeEdasContainersResponse creates a response to parse from DescribeEdasContainers response
func CreateDescribeEdasContainersResponse() (response *DescribeEdasContainersResponse) {
	response = &DescribeEdasContainersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
