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

// DeleteABTestExperiment invokes the opensearch.DeleteABTestExperiment API synchronously
// api document: https://help.aliyun.com/api/opensearch/deleteabtestexperiment.html
func (client *Client) DeleteABTestExperiment(request *DeleteABTestExperimentRequest) (response *DeleteABTestExperimentResponse, err error) {
	response = CreateDeleteABTestExperimentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteABTestExperimentWithChan invokes the opensearch.DeleteABTestExperiment API asynchronously
// api document: https://help.aliyun.com/api/opensearch/deleteabtestexperiment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteABTestExperimentWithChan(request *DeleteABTestExperimentRequest) (<-chan *DeleteABTestExperimentResponse, <-chan error) {
	responseChan := make(chan *DeleteABTestExperimentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteABTestExperiment(request)
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

// DeleteABTestExperimentWithCallback invokes the opensearch.DeleteABTestExperiment API asynchronously
// api document: https://help.aliyun.com/api/opensearch/deleteabtestexperiment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteABTestExperimentWithCallback(request *DeleteABTestExperimentRequest, callback func(response *DeleteABTestExperimentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteABTestExperimentResponse
		var err error
		defer close(result)
		response, err = client.DeleteABTestExperiment(request)
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

// DeleteABTestExperimentRequest is the request struct for api DeleteABTestExperiment
type DeleteABTestExperimentRequest struct {
	*requests.RoaRequest
	GroupId          requests.Integer `position:"Path" name:"groupId"`
	SceneId          requests.Integer `position:"Path" name:"sceneId"`
	ExperimentId     requests.Integer `position:"Path" name:"experimentId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// DeleteABTestExperimentResponse is the response struct for api DeleteABTestExperiment
type DeleteABTestExperimentResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    map[string]interface{} `json:"result" xml:"result"`
}

// CreateDeleteABTestExperimentRequest creates a request to invoke DeleteABTestExperiment API
func CreateDeleteABTestExperimentRequest() (request *DeleteABTestExperimentRequest) {
	request = &DeleteABTestExperimentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DeleteABTestExperiment", "/v4/openapi/app-groups/[appGroupIdentity]/scenes/[sceneId]/groups/[groupId]/experiments/[experimentId]", "opensearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteABTestExperimentResponse creates a response to parse from DeleteABTestExperiment response
func CreateDeleteABTestExperimentResponse() (response *DeleteABTestExperimentResponse) {
	response = &DeleteABTestExperimentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
