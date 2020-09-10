package webplus

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

// DescribeInstanceHealth invokes the webplus.DescribeInstanceHealth API synchronously
// api document: https://help.aliyun.com/api/webplus/describeinstancehealth.html
func (client *Client) DescribeInstanceHealth(request *DescribeInstanceHealthRequest) (response *DescribeInstanceHealthResponse, err error) {
	response = CreateDescribeInstanceHealthResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceHealthWithChan invokes the webplus.DescribeInstanceHealth API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeinstancehealth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceHealthWithChan(request *DescribeInstanceHealthRequest) (<-chan *DescribeInstanceHealthResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceHealthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceHealth(request)
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

// DescribeInstanceHealthWithCallback invokes the webplus.DescribeInstanceHealth API asynchronously
// api document: https://help.aliyun.com/api/webplus/describeinstancehealth.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceHealthWithCallback(request *DescribeInstanceHealthRequest, callback func(response *DescribeInstanceHealthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceHealthResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceHealth(request)
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

// DescribeInstanceHealthRequest is the request struct for api DescribeInstanceHealth
type DescribeInstanceHealthRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeInstanceHealthResponse is the response struct for api DescribeInstanceHealth
type DescribeInstanceHealthResponse struct {
	*responses.BaseResponse
	RequestId      string                                 `json:"RequestId" xml:"RequestId"`
	Code           string                                 `json:"Code" xml:"Code"`
	Message        string                                 `json:"Message" xml:"Message"`
	InstanceHealth InstanceHealthInDescribeInstanceHealth `json:"InstanceHealth" xml:"InstanceHealth"`
}

// CreateDescribeInstanceHealthRequest creates a request to invoke DescribeInstanceHealth API
func CreateDescribeInstanceHealthRequest() (request *DescribeInstanceHealthRequest) {
	request = &DescribeInstanceHealthRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("WebPlus", "2019-03-20", "DescribeInstanceHealth", "/pop/v1/wam/instance/health", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeInstanceHealthResponse creates a response to parse from DescribeInstanceHealth response
func CreateDescribeInstanceHealthResponse() (response *DescribeInstanceHealthResponse) {
	response = &DescribeInstanceHealthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
