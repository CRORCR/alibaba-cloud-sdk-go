package live

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

// DeleteCasterVideoResource invokes the live.DeleteCasterVideoResource API synchronously
// api document: https://help.aliyun.com/api/live/deletecastervideoresource.html
func (client *Client) DeleteCasterVideoResource(request *DeleteCasterVideoResourceRequest) (response *DeleteCasterVideoResourceResponse, err error) {
	response = CreateDeleteCasterVideoResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCasterVideoResourceWithChan invokes the live.DeleteCasterVideoResource API asynchronously
// api document: https://help.aliyun.com/api/live/deletecastervideoresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCasterVideoResourceWithChan(request *DeleteCasterVideoResourceRequest) (<-chan *DeleteCasterVideoResourceResponse, <-chan error) {
	responseChan := make(chan *DeleteCasterVideoResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCasterVideoResource(request)
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

// DeleteCasterVideoResourceWithCallback invokes the live.DeleteCasterVideoResource API asynchronously
// api document: https://help.aliyun.com/api/live/deletecastervideoresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCasterVideoResourceWithCallback(request *DeleteCasterVideoResourceRequest, callback func(response *DeleteCasterVideoResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCasterVideoResourceResponse
		var err error
		defer close(result)
		response, err = client.DeleteCasterVideoResource(request)
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

// DeleteCasterVideoResourceRequest is the request struct for api DeleteCasterVideoResource
type DeleteCasterVideoResourceRequest struct {
	*requests.RpcRequest
	ResourceId string           `position:"Query" name:"ResourceId"`
	CasterId   string           `position:"Query" name:"CasterId"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteCasterVideoResourceResponse is the response struct for api DeleteCasterVideoResource
type DeleteCasterVideoResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCasterVideoResourceRequest creates a request to invoke DeleteCasterVideoResource API
func CreateDeleteCasterVideoResourceRequest() (request *DeleteCasterVideoResourceRequest) {
	request = &DeleteCasterVideoResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteCasterVideoResource", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteCasterVideoResourceResponse creates a response to parse from DeleteCasterVideoResource response
func CreateDeleteCasterVideoResourceResponse() (response *DeleteCasterVideoResourceResponse) {
	response = &DeleteCasterVideoResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
