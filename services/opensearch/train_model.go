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

// TrainModel invokes the opensearch.TrainModel API synchronously
// api document: https://help.aliyun.com/api/opensearch/trainmodel.html
func (client *Client) TrainModel(request *TrainModelRequest) (response *TrainModelResponse, err error) {
	response = CreateTrainModelResponse()
	err = client.DoAction(request, response)
	return
}

// TrainModelWithChan invokes the opensearch.TrainModel API asynchronously
// api document: https://help.aliyun.com/api/opensearch/trainmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TrainModelWithChan(request *TrainModelRequest) (<-chan *TrainModelResponse, <-chan error) {
	responseChan := make(chan *TrainModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TrainModel(request)
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

// TrainModelWithCallback invokes the opensearch.TrainModel API asynchronously
// api document: https://help.aliyun.com/api/opensearch/trainmodel.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TrainModelWithCallback(request *TrainModelRequest, callback func(response *TrainModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TrainModelResponse
		var err error
		defer close(result)
		response, err = client.TrainModel(request)
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

// TrainModelRequest is the request struct for api TrainModel
type TrainModelRequest struct {
	*requests.RoaRequest
	ModelName        string `position:"Path" name:"modelName"`
	AppGroupIdentity string `position:"Path" name:"appGroupIdentity"`
}

// TrainModelResponse is the response struct for api TrainModel
type TrainModelResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"requestId" xml:"requestId"`
	Result    map[string]interface{} `json:"result" xml:"result"`
}

// CreateTrainModelRequest creates a request to invoke TrainModel API
func CreateTrainModelRequest() (request *TrainModelRequest) {
	request = &TrainModelRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "TrainModel", "/v4/openapi/app-groups/[appGroupIdentity]/algorithm/models/[modelName]/actions/train", "opensearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTrainModelResponse creates a response to parse from TrainModel response
func CreateTrainModelResponse() (response *TrainModelResponse) {
	response = &TrainModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
