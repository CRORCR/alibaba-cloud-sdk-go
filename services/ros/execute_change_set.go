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

// ExecuteChangeSet invokes the ros.ExecuteChangeSet API synchronously
// api document: https://help.aliyun.com/api/ros/executechangeset.html
func (client *Client) ExecuteChangeSet(request *ExecuteChangeSetRequest) (response *ExecuteChangeSetResponse, err error) {
	response = CreateExecuteChangeSetResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteChangeSetWithChan invokes the ros.ExecuteChangeSet API asynchronously
// api document: https://help.aliyun.com/api/ros/executechangeset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExecuteChangeSetWithChan(request *ExecuteChangeSetRequest) (<-chan *ExecuteChangeSetResponse, <-chan error) {
	responseChan := make(chan *ExecuteChangeSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteChangeSet(request)
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

// ExecuteChangeSetWithCallback invokes the ros.ExecuteChangeSet API asynchronously
// api document: https://help.aliyun.com/api/ros/executechangeset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExecuteChangeSetWithCallback(request *ExecuteChangeSetRequest, callback func(response *ExecuteChangeSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteChangeSetResponse
		var err error
		defer close(result)
		response, err = client.ExecuteChangeSet(request)
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

// ExecuteChangeSetRequest is the request struct for api ExecuteChangeSet
type ExecuteChangeSetRequest struct {
	*requests.RoaRequest
	ChangeSetName string `position:"Path" name:"ChangeSetName"`
	StackId       string `position:"Path" name:"StackId"`
	StackName     string `position:"Path" name:"StackName"`
}

// ExecuteChangeSetResponse is the response struct for api ExecuteChangeSet
type ExecuteChangeSetResponse struct {
	*responses.BaseResponse
	Dummy string `json:"Dummy" xml:"Dummy"`
}

// CreateExecuteChangeSetRequest creates a request to invoke ExecuteChangeSet API
func CreateExecuteChangeSetRequest() (request *ExecuteChangeSetRequest) {
	request = &ExecuteChangeSetRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "ExecuteChangeSet", "/stacks/[StackName]/[StackId]/changeSets/[ChangeSetName]/execute", "ROS", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateExecuteChangeSetResponse creates a response to parse from ExecuteChangeSet response
func CreateExecuteChangeSetResponse() (response *ExecuteChangeSetResponse) {
	response = &ExecuteChangeSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
