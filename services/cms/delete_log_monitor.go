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

// DeleteLogMonitor invokes the cms.DeleteLogMonitor API synchronously
// api document: https://help.aliyun.com/api/cms/deletelogmonitor.html
func (client *Client) DeleteLogMonitor(request *DeleteLogMonitorRequest) (response *DeleteLogMonitorResponse, err error) {
	response = CreateDeleteLogMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLogMonitorWithChan invokes the cms.DeleteLogMonitor API asynchronously
// api document: https://help.aliyun.com/api/cms/deletelogmonitor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLogMonitorWithChan(request *DeleteLogMonitorRequest) (<-chan *DeleteLogMonitorResponse, <-chan error) {
	responseChan := make(chan *DeleteLogMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLogMonitor(request)
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

// DeleteLogMonitorWithCallback invokes the cms.DeleteLogMonitor API asynchronously
// api document: https://help.aliyun.com/api/cms/deletelogmonitor.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLogMonitorWithCallback(request *DeleteLogMonitorRequest, callback func(response *DeleteLogMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLogMonitorResponse
		var err error
		defer close(result)
		response, err = client.DeleteLogMonitor(request)
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

// DeleteLogMonitorRequest is the request struct for api DeleteLogMonitor
type DeleteLogMonitorRequest struct {
	*requests.RpcRequest
	LogId requests.Integer `position:"Query" name:"LogId"`
}

// DeleteLogMonitorResponse is the response struct for api DeleteLogMonitor
type DeleteLogMonitorResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLogMonitorRequest creates a request to invoke DeleteLogMonitor API
func CreateDeleteLogMonitorRequest() (request *DeleteLogMonitorRequest) {
	request = &DeleteLogMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteLogMonitor", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLogMonitorResponse creates a response to parse from DeleteLogMonitor response
func CreateDeleteLogMonitorResponse() (response *DeleteLogMonitorResponse) {
	response = &DeleteLogMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
