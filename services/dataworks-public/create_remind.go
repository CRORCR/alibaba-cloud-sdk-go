package dataworks_public

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

// CreateRemind invokes the dataworks_public.CreateRemind API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/createremind.html
func (client *Client) CreateRemind(request *CreateRemindRequest) (response *CreateRemindResponse, err error) {
	response = CreateCreateRemindResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRemindWithChan invokes the dataworks_public.CreateRemind API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/createremind.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRemindWithChan(request *CreateRemindRequest) (<-chan *CreateRemindResponse, <-chan error) {
	responseChan := make(chan *CreateRemindResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRemind(request)
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

// CreateRemindWithCallback invokes the dataworks_public.CreateRemind API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/createremind.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRemindWithCallback(request *CreateRemindRequest, callback func(response *CreateRemindResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRemindResponse
		var err error
		defer close(result)
		response, err = client.CreateRemind(request)
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

// CreateRemindRequest is the request struct for api CreateRemind
type CreateRemindRequest struct {
	*requests.RpcRequest
	DndEnd        string           `position:"Body" name:"DndEnd"`
	AlertUnit     string           `position:"Body" name:"AlertUnit"`
	RemindUnit    string           `position:"Body" name:"RemindUnit"`
	AlertInterval requests.Integer `position:"Body" name:"AlertInterval"`
	AlertMethods  string           `position:"Body" name:"AlertMethods"`
	RobotUrls     string           `position:"Body" name:"RobotUrls"`
	MaxAlertTimes requests.Integer `position:"Body" name:"MaxAlertTimes"`
	BizProcessIds string           `position:"Body" name:"BizProcessIds"`
	RemindType    string           `position:"Body" name:"RemindType"`
	AlertTargets  string           `position:"Body" name:"AlertTargets"`
	BaselineIds   string           `position:"Body" name:"BaselineIds"`
	Detail        string           `position:"Body" name:"Detail"`
	RemindName    string           `position:"Body" name:"RemindName"`
	ProjectId     requests.Integer `position:"Body" name:"ProjectId"`
	NodeIds       string           `position:"Body" name:"NodeIds"`
}

// CreateRemindResponse is the response struct for api CreateRemind
type CreateRemindResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateCreateRemindRequest creates a request to invoke CreateRemind API
func CreateCreateRemindRequest() (request *CreateRemindRequest) {
	request = &CreateRemindRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateRemind", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRemindResponse creates a response to parse from CreateRemind response
func CreateCreateRemindResponse() (response *CreateRemindResponse) {
	response = &CreateRemindResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
