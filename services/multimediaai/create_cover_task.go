package multimediaai

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

// CreateCoverTask invokes the multimediaai.CreateCoverTask API synchronously
// api document: https://help.aliyun.com/api/multimediaai/createcovertask.html
func (client *Client) CreateCoverTask(request *CreateCoverTaskRequest) (response *CreateCoverTaskResponse, err error) {
	response = CreateCreateCoverTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCoverTaskWithChan invokes the multimediaai.CreateCoverTask API asynchronously
// api document: https://help.aliyun.com/api/multimediaai/createcovertask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCoverTaskWithChan(request *CreateCoverTaskRequest) (<-chan *CreateCoverTaskResponse, <-chan error) {
	responseChan := make(chan *CreateCoverTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCoverTask(request)
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

// CreateCoverTaskWithCallback invokes the multimediaai.CreateCoverTask API asynchronously
// api document: https://help.aliyun.com/api/multimediaai/createcovertask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCoverTaskWithCallback(request *CreateCoverTaskRequest, callback func(response *CreateCoverTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCoverTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateCoverTask(request)
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

// CreateCoverTaskRequest is the request struct for api CreateCoverTask
type CreateCoverTaskRequest struct {
	*requests.RpcRequest
	TemplateId    requests.Integer `position:"Query" name:"TemplateId"`
	VideoUrl      string           `position:"Query" name:"VideoUrl"`
	Scales        string           `position:"Body" name:"Scales"`
	VideoName     string           `position:"Query" name:"VideoName"`
	CallbackUrl   string           `position:"Query" name:"CallbackUrl"`
	ApplicationId string           `position:"Query" name:"ApplicationId"`
}

// CreateCoverTaskResponse is the response struct for api CreateCoverTask
type CreateCoverTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    int64  `json:"TaskId" xml:"TaskId"`
}

// CreateCreateCoverTaskRequest creates a request to invoke CreateCoverTask API
func CreateCreateCoverTaskRequest() (request *CreateCoverTaskRequest) {
	request = &CreateCoverTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("multimediaai", "2019-08-10", "CreateCoverTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCoverTaskResponse creates a response to parse from CreateCoverTask response
func CreateCreateCoverTaskResponse() (response *CreateCoverTaskResponse) {
	response = &CreateCoverTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
