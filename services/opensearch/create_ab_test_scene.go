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

// CreateABTestScene invokes the opensearch.CreateABTestScene API synchronously
// api document: https://help.aliyun.com/api/opensearch/createabtestscene.html
func (client *Client) CreateABTestScene(request *CreateABTestSceneRequest) (response *CreateABTestSceneResponse, err error) {
	response = CreateCreateABTestSceneResponse()
	err = client.DoAction(request, response)
	return
}

// CreateABTestSceneWithChan invokes the opensearch.CreateABTestScene API asynchronously
// api document: https://help.aliyun.com/api/opensearch/createabtestscene.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateABTestSceneWithChan(request *CreateABTestSceneRequest) (<-chan *CreateABTestSceneResponse, <-chan error) {
	responseChan := make(chan *CreateABTestSceneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateABTestScene(request)
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

// CreateABTestSceneWithCallback invokes the opensearch.CreateABTestScene API asynchronously
// api document: https://help.aliyun.com/api/opensearch/createabtestscene.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateABTestSceneWithCallback(request *CreateABTestSceneRequest, callback func(response *CreateABTestSceneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateABTestSceneResponse
		var err error
		defer close(result)
		response, err = client.CreateABTestScene(request)
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

// CreateABTestSceneRequest is the request struct for api CreateABTestScene
type CreateABTestSceneRequest struct {
	*requests.RoaRequest
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// CreateABTestSceneResponse is the response struct for api CreateABTestScene
type CreateABTestSceneResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateCreateABTestSceneRequest creates a request to invoke CreateABTestScene API
func CreateCreateABTestSceneRequest() (request *CreateABTestSceneRequest) {
	request = &CreateABTestSceneRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "CreateABTestScene", "/v4/openapi/app-groups/[appGroupIdentity]/scenes", "opensearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateABTestSceneResponse creates a response to parse from CreateABTestScene response
func CreateCreateABTestSceneResponse() (response *CreateABTestSceneResponse) {
	response = &CreateABTestSceneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
