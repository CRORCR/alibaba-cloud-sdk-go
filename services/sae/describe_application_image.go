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

// DescribeApplicationImage invokes the sae.DescribeApplicationImage API synchronously
// api document: https://help.aliyun.com/api/sae/describeapplicationimage.html
func (client *Client) DescribeApplicationImage(request *DescribeApplicationImageRequest) (response *DescribeApplicationImageResponse, err error) {
	response = CreateDescribeApplicationImageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApplicationImageWithChan invokes the sae.DescribeApplicationImage API asynchronously
// api document: https://help.aliyun.com/api/sae/describeapplicationimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApplicationImageWithChan(request *DescribeApplicationImageRequest) (<-chan *DescribeApplicationImageResponse, <-chan error) {
	responseChan := make(chan *DescribeApplicationImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApplicationImage(request)
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

// DescribeApplicationImageWithCallback invokes the sae.DescribeApplicationImage API asynchronously
// api document: https://help.aliyun.com/api/sae/describeapplicationimage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApplicationImageWithCallback(request *DescribeApplicationImageRequest, callback func(response *DescribeApplicationImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApplicationImageResponse
		var err error
		defer close(result)
		response, err = client.DescribeApplicationImage(request)
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

// DescribeApplicationImageRequest is the request struct for api DescribeApplicationImage
type DescribeApplicationImageRequest struct {
	*requests.RoaRequest
	AppId    string `position:"Query" name:"AppId"`
	ImageUrl string `position:"Query" name:"ImageUrl"`
}

// DescribeApplicationImageResponse is the response struct for api DescribeApplicationImage
type DescribeApplicationImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeApplicationImageRequest creates a request to invoke DescribeApplicationImage API
func CreateDescribeApplicationImageRequest() (request *DescribeApplicationImageRequest) {
	request = &DescribeApplicationImageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DescribeApplicationImage", "/pop/v1/sam/container/describeApplicationImage", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeApplicationImageResponse creates a response to parse from DescribeApplicationImage response
func CreateDescribeApplicationImageResponse() (response *DescribeApplicationImageResponse) {
	response = &DescribeApplicationImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
