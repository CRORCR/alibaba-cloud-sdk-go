package cr_ee

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

// ListRepoTrigger invokes the cr.ListRepoTrigger API synchronously
// api document: https://help.aliyun.com/api/cr/listrepotrigger.html
func (client *Client) ListRepoTrigger(request *ListRepoTriggerRequest) (response *ListRepoTriggerResponse, err error) {
	response = CreateListRepoTriggerResponse()
	err = client.DoAction(request, response)
	return
}

// ListRepoTriggerWithChan invokes the cr.ListRepoTrigger API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepotrigger.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoTriggerWithChan(request *ListRepoTriggerRequest) (<-chan *ListRepoTriggerResponse, <-chan error) {
	responseChan := make(chan *ListRepoTriggerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRepoTrigger(request)
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

// ListRepoTriggerWithCallback invokes the cr.ListRepoTrigger API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepotrigger.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoTriggerWithCallback(request *ListRepoTriggerRequest, callback func(response *ListRepoTriggerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRepoTriggerResponse
		var err error
		defer close(result)
		response, err = client.ListRepoTrigger(request)
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

// ListRepoTriggerRequest is the request struct for api ListRepoTrigger
type ListRepoTriggerRequest struct {
	*requests.RpcRequest
	RepoId     string `position:"Query" name:"RepoId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListRepoTriggerResponse is the response struct for api ListRepoTrigger
type ListRepoTriggerResponse struct {
	*responses.BaseResponse
	ListRepoTriggerIsSuccess bool           `json:"IsSuccess" xml:"IsSuccess"`
	Code                     string         `json:"Code" xml:"Code"`
	RequestId                string         `json:"RequestId" xml:"RequestId"`
	Triggers                 []TriggersItem `json:"Triggers" xml:"Triggers"`
}

// CreateListRepoTriggerRequest creates a request to invoke ListRepoTrigger API
func CreateListRepoTriggerRequest() (request *ListRepoTriggerRequest) {
	request = &ListRepoTriggerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "ListRepoTrigger", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListRepoTriggerResponse creates a response to parse from ListRepoTrigger response
func CreateListRepoTriggerResponse() (response *ListRepoTriggerResponse) {
	response = &ListRepoTriggerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
