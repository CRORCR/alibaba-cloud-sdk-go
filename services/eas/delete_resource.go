package eas

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

// DeleteResource invokes the eas.DeleteResource API synchronously
// api document: https://help.aliyun.com/api/eas/deleteresource.html
func (client *Client) DeleteResource(request *DeleteResourceRequest) (response *DeleteResourceResponse, err error) {
	response = CreateDeleteResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteResourceWithChan invokes the eas.DeleteResource API asynchronously
// api document: https://help.aliyun.com/api/eas/deleteresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteResourceWithChan(request *DeleteResourceRequest) (<-chan *DeleteResourceResponse, <-chan error) {
	responseChan := make(chan *DeleteResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteResource(request)
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

// DeleteResourceWithCallback invokes the eas.DeleteResource API asynchronously
// api document: https://help.aliyun.com/api/eas/deleteresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteResourceWithCallback(request *DeleteResourceRequest, callback func(response *DeleteResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteResourceResponse
		var err error
		defer close(result)
		response, err = client.DeleteResource(request)
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

// DeleteResourceRequest is the request struct for api DeleteResource
type DeleteResourceRequest struct {
	*requests.RoaRequest
	ClusterId    string `position:"Path" name:"cluster_id"`
	ResourceName string `position:"Path" name:"resource_name"`
}

// DeleteResourceResponse is the response struct for api DeleteResource
type DeleteResourceResponse struct {
	*responses.BaseResponse
}

// CreateDeleteResourceRequest creates a request to invoke DeleteResource API
func CreateDeleteResourceRequest() (request *DeleteResourceRequest) {
	request = &DeleteResourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2018-05-22", "DeleteResource", "/api/resources/[cluster_id]/[resource_name]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteResourceResponse creates a response to parse from DeleteResource response
func CreateDeleteResourceResponse() (response *DeleteResourceResponse) {
	response = &DeleteResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
