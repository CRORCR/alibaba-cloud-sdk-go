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

// GetSkillGroupIdByTaskId invokes the cloudcallcenter.GetSkillGroupIdByTaskId API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getskillgroupidbytaskid.html
func (client *Client) GetSkillGroupIdByTaskId(request *GetSkillGroupIdByTaskIdRequest) (response *GetSkillGroupIdByTaskIdResponse, err error) {
	response = CreateGetSkillGroupIdByTaskIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetSkillGroupIdByTaskIdWithChan invokes the cloudcallcenter.GetSkillGroupIdByTaskId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getskillgroupidbytaskid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSkillGroupIdByTaskIdWithChan(request *GetSkillGroupIdByTaskIdRequest) (<-chan *GetSkillGroupIdByTaskIdResponse, <-chan error) {
	responseChan := make(chan *GetSkillGroupIdByTaskIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSkillGroupIdByTaskId(request)
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

// GetSkillGroupIdByTaskIdWithCallback invokes the cloudcallcenter.GetSkillGroupIdByTaskId API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getskillgroupidbytaskid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetSkillGroupIdByTaskIdWithCallback(request *GetSkillGroupIdByTaskIdRequest, callback func(response *GetSkillGroupIdByTaskIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSkillGroupIdByTaskIdResponse
		var err error
		defer close(result)
		response, err = client.GetSkillGroupIdByTaskId(request)
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

// GetSkillGroupIdByTaskIdRequest is the request struct for api GetSkillGroupIdByTaskId
type GetSkillGroupIdByTaskIdRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	TaskId     string `position:"Query" name:"TaskId"`
}

// GetSkillGroupIdByTaskIdResponse is the response struct for api GetSkillGroupIdByTaskId
type GetSkillGroupIdByTaskIdResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	SkillGroupId   string `json:"SkillGroupId" xml:"SkillGroupId"`
}

// CreateGetSkillGroupIdByTaskIdRequest creates a request to invoke GetSkillGroupIdByTaskId API
func CreateGetSkillGroupIdByTaskIdRequest() (request *GetSkillGroupIdByTaskIdRequest) {
	request = &GetSkillGroupIdByTaskIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetSkillGroupIdByTaskId", "", "")
	request.Method = requests.POST
	return
}

// CreateGetSkillGroupIdByTaskIdResponse creates a response to parse from GetSkillGroupIdByTaskId response
func CreateGetSkillGroupIdByTaskIdResponse() (response *GetSkillGroupIdByTaskIdResponse) {
	response = &GetSkillGroupIdByTaskIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
