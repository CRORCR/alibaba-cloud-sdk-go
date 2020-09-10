package cloudcallcenter

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

// CreatePredictiveJobGroup invokes the cloudcallcenter.CreatePredictiveJobGroup API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createpredictivejobgroup.html
func (client *Client) CreatePredictiveJobGroup(request *CreatePredictiveJobGroupRequest) (response *CreatePredictiveJobGroupResponse, err error) {
	response = CreateCreatePredictiveJobGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePredictiveJobGroupWithChan invokes the cloudcallcenter.CreatePredictiveJobGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createpredictivejobgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePredictiveJobGroupWithChan(request *CreatePredictiveJobGroupRequest) (<-chan *CreatePredictiveJobGroupResponse, <-chan error) {
	responseChan := make(chan *CreatePredictiveJobGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePredictiveJobGroup(request)
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

// CreatePredictiveJobGroupWithCallback invokes the cloudcallcenter.CreatePredictiveJobGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createpredictivejobgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePredictiveJobGroupWithCallback(request *CreatePredictiveJobGroupRequest, callback func(response *CreatePredictiveJobGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePredictiveJobGroupResponse
		var err error
		defer close(result)
		response, err = client.CreatePredictiveJobGroup(request)
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

// CreatePredictiveJobGroupRequest is the request struct for api CreatePredictiveJobGroup
type CreatePredictiveJobGroupRequest struct {
	*requests.RpcRequest
	Description    string           `position:"Query" name:"Description"`
	TimingSchedule requests.Boolean `position:"Query" name:"TimingSchedule"`
	JobsJson       *[]string        `position:"Body" name:"JobsJson"  type:"Repeated"`
	JobFilePath    string           `position:"Query" name:"JobFilePath"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	IsDraft        requests.Boolean `position:"Query" name:"IsDraft"`
	SkillGroupId   string           `position:"Query" name:"SkillGroupId"`
	StrategyJson   string           `position:"Query" name:"StrategyJson"`
	Name           string           `position:"Query" name:"Name"`
}

// CreatePredictiveJobGroupResponse is the response struct for api CreatePredictiveJobGroup
type CreatePredictiveJobGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	JobGroupId     string `json:"JobGroupId" xml:"JobGroupId"`
}

// CreateCreatePredictiveJobGroupRequest creates a request to invoke CreatePredictiveJobGroup API
func CreateCreatePredictiveJobGroupRequest() (request *CreatePredictiveJobGroupRequest) {
	request = &CreatePredictiveJobGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreatePredictiveJobGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateCreatePredictiveJobGroupResponse creates a response to parse from CreatePredictiveJobGroup response
func CreateCreatePredictiveJobGroupResponse() (response *CreatePredictiveJobGroupResponse) {
	response = &CreatePredictiveJobGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
