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

// ModifyVnKnowledge invokes the cloudcallcenter.ModifyVnKnowledge API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnknowledge.html
func (client *Client) ModifyVnKnowledge(request *ModifyVnKnowledgeRequest) (response *ModifyVnKnowledgeResponse, err error) {
	response = CreateModifyVnKnowledgeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVnKnowledgeWithChan invokes the cloudcallcenter.ModifyVnKnowledge API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnknowledge.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnKnowledgeWithChan(request *ModifyVnKnowledgeRequest) (<-chan *ModifyVnKnowledgeResponse, <-chan error) {
	responseChan := make(chan *ModifyVnKnowledgeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVnKnowledge(request)
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

// ModifyVnKnowledgeWithCallback invokes the cloudcallcenter.ModifyVnKnowledge API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifyvnknowledge.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVnKnowledgeWithCallback(request *ModifyVnKnowledgeRequest, callback func(response *ModifyVnKnowledgeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVnKnowledgeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVnKnowledge(request)
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

// ModifyVnKnowledgeRequest is the request struct for api ModifyVnKnowledge
type ModifyVnKnowledgeRequest struct {
	*requests.RpcRequest
	UserUtterance      string           `position:"Query" name:"UserUtterance"`
	Interruptible      requests.Boolean `position:"Query" name:"Interruptible"`
	InstanceId         string           `position:"Query" name:"InstanceId"`
	NavigationScriptId string           `position:"Query" name:"NavigationScriptId"`
	Answer             string           `position:"Query" name:"Answer"`
	SimilarUtterances  *[]string        `position:"Query" name:"SimilarUtterances"  type:"Repeated"`
}

// ModifyVnKnowledgeResponse is the response struct for api ModifyVnKnowledge
type ModifyVnKnowledgeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVnKnowledgeRequest creates a request to invoke ModifyVnKnowledge API
func CreateModifyVnKnowledgeRequest() (request *ModifyVnKnowledgeRequest) {
	request = &ModifyVnKnowledgeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyVnKnowledge", "", "")
	request.Method = requests.GET
	return
}

// CreateModifyVnKnowledgeResponse creates a response to parse from ModifyVnKnowledge response
func CreateModifyVnKnowledgeResponse() (response *ModifyVnKnowledgeResponse) {
	response = &ModifyVnKnowledgeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
