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

// RemoveDataCollection invokes the opensearch.RemoveDataCollection API synchronously
// api document: https://help.aliyun.com/api/opensearch/removedatacollection.html
func (client *Client) RemoveDataCollection(request *RemoveDataCollectionRequest) (response *RemoveDataCollectionResponse, err error) {
	response = CreateRemoveDataCollectionResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveDataCollectionWithChan invokes the opensearch.RemoveDataCollection API asynchronously
// api document: https://help.aliyun.com/api/opensearch/removedatacollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveDataCollectionWithChan(request *RemoveDataCollectionRequest) (<-chan *RemoveDataCollectionResponse, <-chan error) {
	responseChan := make(chan *RemoveDataCollectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveDataCollection(request)
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

// RemoveDataCollectionWithCallback invokes the opensearch.RemoveDataCollection API asynchronously
// api document: https://help.aliyun.com/api/opensearch/removedatacollection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveDataCollectionWithCallback(request *RemoveDataCollectionRequest, callback func(response *RemoveDataCollectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveDataCollectionResponse
		var err error
		defer close(result)
		response, err = client.RemoveDataCollection(request)
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

// RemoveDataCollectionRequest is the request struct for api RemoveDataCollection
type RemoveDataCollectionRequest struct {
	*requests.RoaRequest
	DataCollectionIdentity string `position:"Path" name:"dataCollectionIdentity"`
	AppGroupIdentity       string `position:"Path" name:"appGroupIdentity"`
}

// RemoveDataCollectionResponse is the response struct for api RemoveDataCollection
type RemoveDataCollectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    string `json:"result" xml:"result"`
}

// CreateRemoveDataCollectionRequest creates a request to invoke RemoveDataCollection API
func CreateRemoveDataCollectionRequest() (request *RemoveDataCollectionRequest) {
	request = &RemoveDataCollectionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "RemoveDataCollection", "/v4/openapi/app-groups/[appGroupIdentity]/data-collections/[dataCollectionIdentity]", "opensearch", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateRemoveDataCollectionResponse creates a response to parse from RemoveDataCollection response
func CreateRemoveDataCollectionResponse() (response *RemoveDataCollectionResponse) {
	response = &RemoveDataCollectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
