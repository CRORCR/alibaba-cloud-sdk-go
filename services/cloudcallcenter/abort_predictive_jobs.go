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

// AbortPredictiveJobs invokes the cloudcallcenter.AbortPredictiveJobs API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/abortpredictivejobs.html
func (client *Client) AbortPredictiveJobs(request *AbortPredictiveJobsRequest) (response *AbortPredictiveJobsResponse, err error) {
	response = CreateAbortPredictiveJobsResponse()
	err = client.DoAction(request, response)
	return
}

// AbortPredictiveJobsWithChan invokes the cloudcallcenter.AbortPredictiveJobs API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/abortpredictivejobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbortPredictiveJobsWithChan(request *AbortPredictiveJobsRequest) (<-chan *AbortPredictiveJobsResponse, <-chan error) {
	responseChan := make(chan *AbortPredictiveJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AbortPredictiveJobs(request)
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

// AbortPredictiveJobsWithCallback invokes the cloudcallcenter.AbortPredictiveJobs API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/abortpredictivejobs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbortPredictiveJobsWithCallback(request *AbortPredictiveJobsRequest, callback func(response *AbortPredictiveJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AbortPredictiveJobsResponse
		var err error
		defer close(result)
		response, err = client.AbortPredictiveJobs(request)
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

// AbortPredictiveJobsRequest is the request struct for api AbortPredictiveJobs
type AbortPredictiveJobsRequest struct {
	*requests.RpcRequest
	All          requests.Boolean `position:"Query" name:"All"`
	JobId        *[]string        `position:"Query" name:"JobId"  type:"Repeated"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	JobGroupId   string           `position:"Query" name:"JobGroupId"`
}

// AbortPredictiveJobsResponse is the response struct for api AbortPredictiveJobs
type AbortPredictiveJobsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateAbortPredictiveJobsRequest creates a request to invoke AbortPredictiveJobs API
func CreateAbortPredictiveJobsRequest() (request *AbortPredictiveJobsRequest) {
	request = &AbortPredictiveJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "AbortPredictiveJobs", "", "")
	request.Method = requests.POST
	return
}

// CreateAbortPredictiveJobsResponse creates a response to parse from AbortPredictiveJobs response
func CreateAbortPredictiveJobsResponse() (response *AbortPredictiveJobsResponse) {
	response = &AbortPredictiveJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
