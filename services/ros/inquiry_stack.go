package ros

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

// InquiryStack invokes the ros.InquiryStack API synchronously
// api document: https://help.aliyun.com/api/ros/inquirystack.html
func (client *Client) InquiryStack(request *InquiryStackRequest) (response *InquiryStackResponse, err error) {
	response = CreateInquiryStackResponse()
	err = client.DoAction(request, response)
	return
}

// InquiryStackWithChan invokes the ros.InquiryStack API asynchronously
// api document: https://help.aliyun.com/api/ros/inquirystack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InquiryStackWithChan(request *InquiryStackRequest) (<-chan *InquiryStackResponse, <-chan error) {
	responseChan := make(chan *InquiryStackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InquiryStack(request)
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

// InquiryStackWithCallback invokes the ros.InquiryStack API asynchronously
// api document: https://help.aliyun.com/api/ros/inquirystack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InquiryStackWithCallback(request *InquiryStackRequest, callback func(response *InquiryStackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InquiryStackResponse
		var err error
		defer close(result)
		response, err = client.InquiryStack(request)
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

// InquiryStackRequest is the request struct for api InquiryStack
type InquiryStackRequest struct {
	*requests.RoaRequest
}

// InquiryStackResponse is the response struct for api InquiryStack
type InquiryStackResponse struct {
	*responses.BaseResponse
}

// CreateInquiryStackRequest creates a request to invoke InquiryStack API
func CreateInquiryStackRequest() (request *InquiryStackRequest) {
	request = &InquiryStackRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "InquiryStack", "/stacks/inquiry", "ROS", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInquiryStackResponse creates a response to parse from InquiryStack response
func CreateInquiryStackResponse() (response *InquiryStackResponse) {
	response = &InquiryStackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
