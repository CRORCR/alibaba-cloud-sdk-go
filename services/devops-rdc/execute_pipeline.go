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

// ExecutePipeline invokes the devops_rdc.ExecutePipeline API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/executepipeline.html
func (client *Client) ExecutePipeline(request *ExecutePipelineRequest) (response *ExecutePipelineResponse, err error) {
	response = CreateExecutePipelineResponse()
	err = client.DoAction(request, response)
	return
}

// ExecutePipelineWithChan invokes the devops_rdc.ExecutePipeline API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/executepipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExecutePipelineWithChan(request *ExecutePipelineRequest) (<-chan *ExecutePipelineResponse, <-chan error) {
	responseChan := make(chan *ExecutePipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecutePipeline(request)
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

// ExecutePipelineWithCallback invokes the devops_rdc.ExecutePipeline API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/executepipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExecutePipelineWithCallback(request *ExecutePipelineRequest, callback func(response *ExecutePipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecutePipelineResponse
		var err error
		defer close(result)
		response, err = client.ExecutePipeline(request)
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

// ExecutePipelineRequest is the request struct for api ExecutePipeline
type ExecutePipelineRequest struct {
	*requests.RpcRequest
	Parameters string           `position:"Body" name:"Parameters"`
	UserPk     string           `position:"Body" name:"UserPk"`
	OrgId      string           `position:"Body" name:"OrgId"`
	PipelineId requests.Integer `position:"Body" name:"PipelineId"`
}

// ExecutePipelineResponse is the response struct for api ExecutePipeline
type ExecutePipelineResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       int64  `json:"Object" xml:"Object"`
}

// CreateExecutePipelineRequest creates a request to invoke ExecutePipeline API
func CreateExecutePipelineRequest() (request *ExecutePipelineRequest) {
	request = &ExecutePipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "ExecutePipeline", "", "")
	request.Method = requests.POST
	return
}

// CreateExecutePipelineResponse creates a response to parse from ExecutePipeline response
func CreateExecutePipelineResponse() (response *ExecutePipelineResponse) {
	response = &ExecutePipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
