package mts

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

// AddTerrorismPipeline invokes the mts.AddTerrorismPipeline API synchronously
// api document: https://help.aliyun.com/api/mts/addterrorismpipeline.html
func (client *Client) AddTerrorismPipeline(request *AddTerrorismPipelineRequest) (response *AddTerrorismPipelineResponse, err error) {
	response = CreateAddTerrorismPipelineResponse()
	err = client.DoAction(request, response)
	return
}

// AddTerrorismPipelineWithChan invokes the mts.AddTerrorismPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/addterrorismpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTerrorismPipelineWithChan(request *AddTerrorismPipelineRequest) (<-chan *AddTerrorismPipelineResponse, <-chan error) {
	responseChan := make(chan *AddTerrorismPipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTerrorismPipeline(request)
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

// AddTerrorismPipelineWithCallback invokes the mts.AddTerrorismPipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/addterrorismpipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTerrorismPipelineWithCallback(request *AddTerrorismPipelineRequest, callback func(response *AddTerrorismPipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTerrorismPipelineResponse
		var err error
		defer close(result)
		response, err = client.AddTerrorismPipeline(request)
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

// AddTerrorismPipelineRequest is the request struct for api AddTerrorismPipeline
type AddTerrorismPipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	Name                 string           `position:"Query" name:"Name"`
}

// AddTerrorismPipelineResponse is the response struct for api AddTerrorismPipeline
type AddTerrorismPipelineResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Pipeline  Pipeline `json:"Pipeline" xml:"Pipeline"`
}

// CreateAddTerrorismPipelineRequest creates a request to invoke AddTerrorismPipeline API
func CreateAddTerrorismPipelineRequest() (request *AddTerrorismPipelineRequest) {
	request = &AddTerrorismPipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddTerrorismPipeline", "", "")
	return
}

// CreateAddTerrorismPipelineResponse creates a response to parse from AddTerrorismPipeline response
func CreateAddTerrorismPipelineResponse() (response *AddTerrorismPipelineResponse) {
	response = &AddTerrorismPipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
