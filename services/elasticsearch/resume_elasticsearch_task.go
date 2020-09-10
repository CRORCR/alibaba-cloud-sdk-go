package elasticsearch

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

// ResumeElasticsearchTask invokes the elasticsearch.ResumeElasticsearchTask API synchronously
// api document: https://help.aliyun.com/api/elasticsearch/resumeelasticsearchtask.html
func (client *Client) ResumeElasticsearchTask(request *ResumeElasticsearchTaskRequest) (response *ResumeElasticsearchTaskResponse, err error) {
	response = CreateResumeElasticsearchTaskResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeElasticsearchTaskWithChan invokes the elasticsearch.ResumeElasticsearchTask API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/resumeelasticsearchtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeElasticsearchTaskWithChan(request *ResumeElasticsearchTaskRequest) (<-chan *ResumeElasticsearchTaskResponse, <-chan error) {
	responseChan := make(chan *ResumeElasticsearchTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeElasticsearchTask(request)
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

// ResumeElasticsearchTaskWithCallback invokes the elasticsearch.ResumeElasticsearchTask API asynchronously
// api document: https://help.aliyun.com/api/elasticsearch/resumeelasticsearchtask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeElasticsearchTaskWithCallback(request *ResumeElasticsearchTaskRequest, callback func(response *ResumeElasticsearchTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeElasticsearchTaskResponse
		var err error
		defer close(result)
		response, err = client.ResumeElasticsearchTask(request)
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

// ResumeElasticsearchTaskRequest is the request struct for api ResumeElasticsearchTask
type ResumeElasticsearchTaskRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// ResumeElasticsearchTaskResponse is the response struct for api ResumeElasticsearchTask
type ResumeElasticsearchTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateResumeElasticsearchTaskRequest creates a request to invoke ResumeElasticsearchTask API
func CreateResumeElasticsearchTaskRequest() (request *ResumeElasticsearchTaskRequest) {
	request = &ResumeElasticsearchTaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ResumeElasticsearchTask", "/openapi/instances/[InstanceId]/actions/resume", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResumeElasticsearchTaskResponse creates a response to parse from ResumeElasticsearchTask response
func CreateResumeElasticsearchTaskResponse() (response *ResumeElasticsearchTaskResponse) {
	response = &ResumeElasticsearchTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
