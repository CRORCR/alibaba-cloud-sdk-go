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

// CreatePipeline invokes the devops_rdc.CreatePipeline API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/createpipeline.html
func (client *Client) CreatePipeline(request *CreatePipelineRequest) (response *CreatePipelineResponse, err error) {
	response = CreateCreatePipelineResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePipelineWithChan invokes the devops_rdc.CreatePipeline API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/createpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePipelineWithChan(request *CreatePipelineRequest) (<-chan *CreatePipelineResponse, <-chan error) {
	responseChan := make(chan *CreatePipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePipeline(request)
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

// CreatePipelineWithCallback invokes the devops_rdc.CreatePipeline API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/createpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePipelineWithCallback(request *CreatePipelineRequest, callback func(response *CreatePipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePipelineResponse
		var err error
		defer close(result)
		response, err = client.CreatePipeline(request)
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

// CreatePipelineRequest is the request struct for api CreatePipeline
type CreatePipelineRequest struct {
	*requests.RpcRequest
	Pipeline string `position:"Body" name:"Pipeline"`
	UserPk   string `position:"Body" name:"UserPk"`
	OrgId    string `position:"Body" name:"OrgId"`
}

// CreatePipelineResponse is the response struct for api CreatePipeline
type CreatePipelineResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       int64  `json:"Object" xml:"Object"`
}

// CreateCreatePipelineRequest creates a request to invoke CreatePipeline API
func CreateCreatePipelineRequest() (request *CreatePipelineRequest) {
	request = &CreatePipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "CreatePipeline", "", "")
	request.Method = requests.POST
	return
}

// CreateCreatePipelineResponse creates a response to parse from CreatePipeline response
func CreateCreatePipelineResponse() (response *CreatePipelineResponse) {
	response = &CreatePipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
