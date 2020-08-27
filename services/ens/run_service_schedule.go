package ens

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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// RunServiceSchedule invokes the ens.RunServiceSchedule API synchronously
// api document: https://help.aliyun.com/api/ens/runserviceschedule.html
func (client *Client) RunServiceSchedule(request *RunServiceScheduleRequest) (response *RunServiceScheduleResponse, err error) {
	response = CreateRunServiceScheduleResponse()
	err = client.DoAction(request, response)
	return
}

// RunServiceScheduleWithChan invokes the ens.RunServiceSchedule API asynchronously
// api document: https://help.aliyun.com/api/ens/runserviceschedule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RunServiceScheduleWithChan(request *RunServiceScheduleRequest) (<-chan *RunServiceScheduleResponse, <-chan error) {
	responseChan := make(chan *RunServiceScheduleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunServiceSchedule(request)
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

// RunServiceScheduleWithCallback invokes the ens.RunServiceSchedule API asynchronously
// api document: https://help.aliyun.com/api/ens/runserviceschedule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RunServiceScheduleWithCallback(request *RunServiceScheduleRequest, callback func(response *RunServiceScheduleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunServiceScheduleResponse
		var err error
		defer close(result)
		response, err = client.RunServiceSchedule(request)
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

// RunServiceScheduleRequest is the request struct for api RunServiceSchedule
type RunServiceScheduleRequest struct {
	*requests.RpcRequest
	Directorys       string           `position:"Query" name:"Directorys"`
	PreLockedTimeout requests.Integer `position:"Query" name:"PreLockedTimeout"`
	Uuid             string           `position:"Query" name:"Uuid"`
	ClientIp         string           `position:"Query" name:"ClientIp"`
	PodConfigName    string           `position:"Query" name:"PodConfigName"`
	ServiceAction    string           `position:"Query" name:"ServiceAction"`
	ServiceCommands  string           `position:"Query" name:"ServiceCommands"`
	ScheduleStrategy string           `position:"Query" name:"ScheduleStrategy"`
	AppId            string           `position:"Query" name:"AppId"`
}

// RunServiceScheduleResponse is the response struct for api RunServiceSchedule
type RunServiceScheduleResponse struct {
	*responses.BaseResponse
	RequestId       string         `json:"RequestId" xml:"RequestId"`
	RequestRepeated string         `json:"RequestRepeated" xml:"RequestRepeated"`
	TcpPorts        bool           `json:"TcpPorts" xml:"TcpPorts"`
	InstanceId      string         `json:"InstanceId" xml:"InstanceId"`
	InstancePort    int            `json:"InstancePort" xml:"InstancePort"`
	InstanceIp      string         `json:"InstanceIp" xml:"InstanceIp"`
	Index           int            `json:"Index" xml:"Index"`
	CommandResults  CommandResults `json:"CommandResults" xml:"CommandResults"`
}

// CreateRunServiceScheduleRequest creates a request to invoke RunServiceSchedule API
func CreateRunServiceScheduleRequest() (request *RunServiceScheduleRequest) {
	request = &RunServiceScheduleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "RunServiceSchedule", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRunServiceScheduleResponse creates a response to parse from RunServiceSchedule response
func CreateRunServiceScheduleResponse() (response *RunServiceScheduleResponse) {
	response = &RunServiceScheduleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
