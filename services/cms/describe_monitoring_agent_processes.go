package cms

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

// DescribeMonitoringAgentProcesses invokes the cms.DescribeMonitoringAgentProcesses API synchronously
// api document: https://help.aliyun.com/api/cms/describemonitoringagentprocesses.html
func (client *Client) DescribeMonitoringAgentProcesses(request *DescribeMonitoringAgentProcessesRequest) (response *DescribeMonitoringAgentProcessesResponse, err error) {
	response = CreateDescribeMonitoringAgentProcessesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMonitoringAgentProcessesWithChan invokes the cms.DescribeMonitoringAgentProcesses API asynchronously
// api document: https://help.aliyun.com/api/cms/describemonitoringagentprocesses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonitoringAgentProcessesWithChan(request *DescribeMonitoringAgentProcessesRequest) (<-chan *DescribeMonitoringAgentProcessesResponse, <-chan error) {
	responseChan := make(chan *DescribeMonitoringAgentProcessesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMonitoringAgentProcesses(request)
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

// DescribeMonitoringAgentProcessesWithCallback invokes the cms.DescribeMonitoringAgentProcesses API asynchronously
// api document: https://help.aliyun.com/api/cms/describemonitoringagentprocesses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMonitoringAgentProcessesWithCallback(request *DescribeMonitoringAgentProcessesRequest, callback func(response *DescribeMonitoringAgentProcessesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMonitoringAgentProcessesResponse
		var err error
		defer close(result)
		response, err = client.DescribeMonitoringAgentProcesses(request)
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

// DescribeMonitoringAgentProcessesRequest is the request struct for api DescribeMonitoringAgentProcesses
type DescribeMonitoringAgentProcessesRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeMonitoringAgentProcessesResponse is the response struct for api DescribeMonitoringAgentProcesses
type DescribeMonitoringAgentProcessesResponse struct {
	*responses.BaseResponse
	Code          string        `json:"Code" xml:"Code"`
	Message       string        `json:"Message" xml:"Message"`
	Success       bool          `json:"Success" xml:"Success"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	NodeProcesses NodeProcesses `json:"NodeProcesses" xml:"NodeProcesses"`
}

// CreateDescribeMonitoringAgentProcessesRequest creates a request to invoke DescribeMonitoringAgentProcesses API
func CreateDescribeMonitoringAgentProcessesRequest() (request *DescribeMonitoringAgentProcessesRequest) {
	request = &DescribeMonitoringAgentProcessesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeMonitoringAgentProcesses", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMonitoringAgentProcessesResponse creates a response to parse from DescribeMonitoringAgentProcesses response
func CreateDescribeMonitoringAgentProcessesResponse() (response *DescribeMonitoringAgentProcessesResponse) {
	response = &DescribeMonitoringAgentProcessesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
