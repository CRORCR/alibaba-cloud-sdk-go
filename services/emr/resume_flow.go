package emr

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

// ResumeFlow invokes the emr.ResumeFlow API synchronously
// api document: https://help.aliyun.com/api/emr/resumeflow.html
func (client *Client) ResumeFlow(request *ResumeFlowRequest) (response *ResumeFlowResponse, err error) {
	response = CreateResumeFlowResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeFlowWithChan invokes the emr.ResumeFlow API asynchronously
// api document: https://help.aliyun.com/api/emr/resumeflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeFlowWithChan(request *ResumeFlowRequest) (<-chan *ResumeFlowResponse, <-chan error) {
	responseChan := make(chan *ResumeFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeFlow(request)
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

// ResumeFlowWithCallback invokes the emr.ResumeFlow API asynchronously
// api document: https://help.aliyun.com/api/emr/resumeflow.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeFlowWithCallback(request *ResumeFlowRequest, callback func(response *ResumeFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeFlowResponse
		var err error
		defer close(result)
		response, err = client.ResumeFlow(request)
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

// ResumeFlowRequest is the request struct for api ResumeFlow
type ResumeFlowRequest struct {
	*requests.RpcRequest
	FlowInstanceId string `position:"Query" name:"FlowInstanceId"`
	ProjectId      string `position:"Query" name:"ProjectId"`
}

// ResumeFlowResponse is the response struct for api ResumeFlow
type ResumeFlowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateResumeFlowRequest creates a request to invoke ResumeFlow API
func CreateResumeFlowRequest() (request *ResumeFlowRequest) {
	request = &ResumeFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ResumeFlow", "emr", "openAPI")
	return
}

// CreateResumeFlowResponse creates a response to parse from ResumeFlow response
func CreateResumeFlowResponse() (response *ResumeFlowResponse) {
	response = &ResumeFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
