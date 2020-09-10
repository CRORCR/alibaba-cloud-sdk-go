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

// DescribeABTestExperiment invokes the opensearch.DescribeABTestExperiment API synchronously
// api document: https://help.aliyun.com/api/opensearch/describeabtestexperiment.html
func (client *Client) DescribeABTestExperiment(request *DescribeABTestExperimentRequest) (response *DescribeABTestExperimentResponse, err error) {
	response = CreateDescribeABTestExperimentResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeABTestExperimentWithChan invokes the opensearch.DescribeABTestExperiment API asynchronously
// api document: https://help.aliyun.com/api/opensearch/describeabtestexperiment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeABTestExperimentWithChan(request *DescribeABTestExperimentRequest) (<-chan *DescribeABTestExperimentResponse, <-chan error) {
	responseChan := make(chan *DescribeABTestExperimentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeABTestExperiment(request)
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

// DescribeABTestExperimentWithCallback invokes the opensearch.DescribeABTestExperiment API asynchronously
// api document: https://help.aliyun.com/api/opensearch/describeabtestexperiment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeABTestExperimentWithCallback(request *DescribeABTestExperimentRequest, callback func(response *DescribeABTestExperimentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeABTestExperimentResponse
		var err error
		defer close(result)
		response, err = client.DescribeABTestExperiment(request)
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

// DescribeABTestExperimentRequest is the request struct for api DescribeABTestExperiment
type DescribeABTestExperimentRequest struct {
	*requests.RoaRequest
	GroupId          requests.Integer `position:"Path" name:"groupId"`
	SceneId          requests.Integer `position:"Path" name:"sceneId"`
	ExperimentId     requests.Integer `position:"Path" name:"experimentId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// DescribeABTestExperimentResponse is the response struct for api DescribeABTestExperiment
type DescribeABTestExperimentResponse struct {
	*responses.BaseResponse
	RequestId string                           `json:"requestId" xml:"requestId"`
	Result    ResultInDescribeABTestExperiment `json:"result" xml:"result"`
}

// CreateDescribeABTestExperimentRequest creates a request to invoke DescribeABTestExperiment API
func CreateDescribeABTestExperimentRequest() (request *DescribeABTestExperimentRequest) {
	request = &DescribeABTestExperimentRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeABTestExperiment", "/v4/openapi/app-groups/[appGroupIdentity]/scenes/[sceneId]/groups/[groupId]/experiments/[experimentId]", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeABTestExperimentResponse creates a response to parse from DescribeABTestExperiment response
func CreateDescribeABTestExperimentResponse() (response *DescribeABTestExperimentResponse) {
	response = &DescribeABTestExperimentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
