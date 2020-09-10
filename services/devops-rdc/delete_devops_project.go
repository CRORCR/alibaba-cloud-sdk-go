package devops_rdc

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

// DeleteDevopsProject invokes the devops_rdc.DeleteDevopsProject API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/deletedevopsproject.html
func (client *Client) DeleteDevopsProject(request *DeleteDevopsProjectRequest) (response *DeleteDevopsProjectResponse, err error) {
	response = CreateDeleteDevopsProjectResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDevopsProjectWithChan invokes the devops_rdc.DeleteDevopsProject API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/deletedevopsproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDevopsProjectWithChan(request *DeleteDevopsProjectRequest) (<-chan *DeleteDevopsProjectResponse, <-chan error) {
	responseChan := make(chan *DeleteDevopsProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDevopsProject(request)
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

// DeleteDevopsProjectWithCallback invokes the devops_rdc.DeleteDevopsProject API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/deletedevopsproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDevopsProjectWithCallback(request *DeleteDevopsProjectRequest, callback func(response *DeleteDevopsProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDevopsProjectResponse
		var err error
		defer close(result)
		response, err = client.DeleteDevopsProject(request)
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

// DeleteDevopsProjectRequest is the request struct for api DeleteDevopsProject
type DeleteDevopsProjectRequest struct {
	*requests.RpcRequest
	ProjectId string `position:"Body" name:"ProjectId"`
	OrgId     string `position:"Body" name:"OrgId"`
}

// DeleteDevopsProjectResponse is the response struct for api DeleteDevopsProject
type DeleteDevopsProjectResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Object       string `json:"Object" xml:"Object"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateDeleteDevopsProjectRequest creates a request to invoke DeleteDevopsProject API
func CreateDeleteDevopsProjectRequest() (request *DeleteDevopsProjectRequest) {
	request = &DeleteDevopsProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "DeleteDevopsProject", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteDevopsProjectResponse creates a response to parse from DeleteDevopsProject response
func CreateDeleteDevopsProjectResponse() (response *DeleteDevopsProjectResponse) {
	response = &DeleteDevopsProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
