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

// DescribeTemplate invokes the ros.DescribeTemplate API synchronously
// api document: https://help.aliyun.com/api/ros/describetemplate.html
func (client *Client) DescribeTemplate(request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
	response = CreateDescribeTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTemplateWithChan invokes the ros.DescribeTemplate API asynchronously
// api document: https://help.aliyun.com/api/ros/describetemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTemplateWithChan(request *DescribeTemplateRequest) (<-chan *DescribeTemplateResponse, <-chan error) {
	responseChan := make(chan *DescribeTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTemplate(request)
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

// DescribeTemplateWithCallback invokes the ros.DescribeTemplate API asynchronously
// api document: https://help.aliyun.com/api/ros/describetemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTemplateWithCallback(request *DescribeTemplateRequest, callback func(response *DescribeTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTemplateResponse
		var err error
		defer close(result)
		response, err = client.DescribeTemplate(request)
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

// DescribeTemplateRequest is the request struct for api DescribeTemplate
type DescribeTemplateRequest struct {
	*requests.RoaRequest
	StackId   string `position:"Path" name:"StackId"`
	StackName string `position:"Path" name:"StackName"`
}

// DescribeTemplateResponse is the response struct for api DescribeTemplate
type DescribeTemplateResponse struct {
	*responses.BaseResponse
}

// CreateDescribeTemplateRequest creates a request to invoke DescribeTemplate API
func CreateDescribeTemplateRequest() (request *DescribeTemplateRequest) {
	request = &DescribeTemplateRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "DescribeTemplate", "/stacks/[StackName]/[StackId]/template", "ROS", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeTemplateResponse creates a response to parse from DescribeTemplate response
func CreateDescribeTemplateResponse() (response *DescribeTemplateResponse) {
	response = &DescribeTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
