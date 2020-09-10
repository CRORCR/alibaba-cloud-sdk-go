package teambition_aliyun

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

// UpdateProject invokes the teambition_aliyun.UpdateProject API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/updateproject.html
func (client *Client) UpdateProject(request *UpdateProjectRequest) (response *UpdateProjectResponse, err error) {
	response = CreateUpdateProjectResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateProjectWithChan invokes the teambition_aliyun.UpdateProject API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/updateproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateProjectWithChan(request *UpdateProjectRequest) (<-chan *UpdateProjectResponse, <-chan error) {
	responseChan := make(chan *UpdateProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateProject(request)
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

// UpdateProjectWithCallback invokes the teambition_aliyun.UpdateProject API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/updateproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateProjectWithCallback(request *UpdateProjectRequest, callback func(response *UpdateProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateProjectResponse
		var err error
		defer close(result)
		response, err = client.UpdateProject(request)
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

// UpdateProjectRequest is the request struct for api UpdateProject
type UpdateProjectRequest struct {
	*requests.RpcRequest
	Name        string `position:"Body" name:"Name"`
	Description string `position:"Body" name:"Description"`
	ProjectId   string `position:"Body" name:"ProjectId"`
	OrgId       string `position:"Body" name:"OrgId"`
}

// UpdateProjectResponse is the response struct for api UpdateProject
type UpdateProjectResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       string `json:"Object" xml:"Object"`
}

// CreateUpdateProjectRequest creates a request to invoke UpdateProject API
func CreateUpdateProjectRequest() (request *UpdateProjectRequest) {
	request = &UpdateProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "UpdateProject", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateProjectResponse creates a response to parse from UpdateProject response
func CreateUpdateProjectResponse() (response *UpdateProjectResponse) {
	response = &UpdateProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
