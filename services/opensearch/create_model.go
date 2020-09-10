package opensearch

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

// CreateModel invokes the opensearch.CreateModel API synchronously
// api document: https://help.aliyun.com/api/opensearch/createmodel.html
func (client *Client) CreateModel(request *CreateModelRequest) (response *CreateModelResponse, err error) {
	response = CreateCreateModelResponse()
	err = client.DoAction(request, response)
	return
}

// CreateModelWithChan invokes the opensearch.CreateModel API asynchronously
// api document: https://help.aliyun.com/api/opensearch/createmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateModelWithChan(request *CreateModelRequest) (<-chan *CreateModelResponse, <-chan error) {
	responseChan := make(chan *CreateModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateModel(request)
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

// CreateModelWithCallback invokes the opensearch.CreateModel API asynchronously
// api document: https://help.aliyun.com/api/opensearch/createmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateModelWithCallback(request *CreateModelRequest, callback func(response *CreateModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateModelResponse
		var err error
		defer close(result)
		response, err = client.CreateModel(request)
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

// CreateModelRequest is the request struct for api CreateModel
type CreateModelRequest struct {
	*requests.RoaRequest
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// CreateModelResponse is the response struct for api CreateModel
type CreateModelResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    map[string]interface{} `json:"result" xml:"result"`
}

// CreateCreateModelRequest creates a request to invoke CreateModel API
func CreateCreateModelRequest() (request *CreateModelRequest) {
	request = &CreateModelRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "CreateModel", "/v4/openapi/app-groups/[appGroupIdentity]/algorithm/models", "opensearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateModelResponse creates a response to parse from CreateModel response
func CreateCreateModelResponse() (response *CreateModelResponse) {
	response = &CreateModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
