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

// PushInterventionDictionaryEntries invokes the opensearch.PushInterventionDictionaryEntries API synchronously
// api document: https://help.aliyun.com/api/opensearch/pushinterventiondictionaryentries.html
func (client *Client) PushInterventionDictionaryEntries(request *PushInterventionDictionaryEntriesRequest) (response *PushInterventionDictionaryEntriesResponse, err error) {
	response = CreatePushInterventionDictionaryEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// PushInterventionDictionaryEntriesWithChan invokes the opensearch.PushInterventionDictionaryEntries API asynchronously
// api document: https://help.aliyun.com/api/opensearch/pushinterventiondictionaryentries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PushInterventionDictionaryEntriesWithChan(request *PushInterventionDictionaryEntriesRequest) (<-chan *PushInterventionDictionaryEntriesResponse, <-chan error) {
	responseChan := make(chan *PushInterventionDictionaryEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushInterventionDictionaryEntries(request)
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

// PushInterventionDictionaryEntriesWithCallback invokes the opensearch.PushInterventionDictionaryEntries API asynchronously
// api document: https://help.aliyun.com/api/opensearch/pushinterventiondictionaryentries.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PushInterventionDictionaryEntriesWithCallback(request *PushInterventionDictionaryEntriesRequest, callback func(response *PushInterventionDictionaryEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushInterventionDictionaryEntriesResponse
		var err error
		defer close(result)
		response, err = client.PushInterventionDictionaryEntries(request)
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

// PushInterventionDictionaryEntriesRequest is the request struct for api PushInterventionDictionaryEntries
type PushInterventionDictionaryEntriesRequest struct {
	*requests.RoaRequest
	Name string `position:"Path" name:"name"`
}

// PushInterventionDictionaryEntriesResponse is the response struct for api PushInterventionDictionaryEntries
type PushInterventionDictionaryEntriesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"requestId" xml:"requestId"`
	Result    []string `json:"result" xml:"result"`
}

// CreatePushInterventionDictionaryEntriesRequest creates a request to invoke PushInterventionDictionaryEntries API
func CreatePushInterventionDictionaryEntriesRequest() (request *PushInterventionDictionaryEntriesRequest) {
	request = &PushInterventionDictionaryEntriesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "PushInterventionDictionaryEntries", "/v4/openapi/intervention-dictionaries/[name]/entries/actions/bulk", "opensearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePushInterventionDictionaryEntriesResponse creates a response to parse from PushInterventionDictionaryEntries response
func CreatePushInterventionDictionaryEntriesResponse() (response *PushInterventionDictionaryEntriesResponse) {
	response = &PushInterventionDictionaryEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
