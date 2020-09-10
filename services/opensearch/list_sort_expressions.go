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

// ListSortExpressions invokes the opensearch.ListSortExpressions API synchronously
// api document: https://help.aliyun.com/api/opensearch/listsortexpressions.html
func (client *Client) ListSortExpressions(request *ListSortExpressionsRequest) (response *ListSortExpressionsResponse, err error) {
	response = CreateListSortExpressionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSortExpressionsWithChan invokes the opensearch.ListSortExpressions API asynchronously
// api document: https://help.aliyun.com/api/opensearch/listsortexpressions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSortExpressionsWithChan(request *ListSortExpressionsRequest) (<-chan *ListSortExpressionsResponse, <-chan error) {
	responseChan := make(chan *ListSortExpressionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSortExpressions(request)
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

// ListSortExpressionsWithCallback invokes the opensearch.ListSortExpressions API asynchronously
// api document: https://help.aliyun.com/api/opensearch/listsortexpressions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSortExpressionsWithCallback(request *ListSortExpressionsRequest, callback func(response *ListSortExpressionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSortExpressionsResponse
		var err error
		defer close(result)
		response, err = client.ListSortExpressions(request)
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

// ListSortExpressionsRequest is the request struct for api ListSortExpressions
type ListSortExpressionsRequest struct {
	*requests.RoaRequest
	AppId            requests.Integer `position:"Path" name:"appId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// ListSortExpressionsResponse is the response struct for api ListSortExpressions
type ListSortExpressionsResponse struct {
	*responses.BaseResponse
	RequestId string          `json:"requestId" xml:"requestId"`
	Result    []FirstRankItem `json:"result" xml:"result"`
}

// CreateListSortExpressionsRequest creates a request to invoke ListSortExpressions API
func CreateListSortExpressionsRequest() (request *ListSortExpressionsRequest) {
	request = &ListSortExpressionsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListSortExpressions", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/sort-expressions", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListSortExpressionsResponse creates a response to parse from ListSortExpressions response
func CreateListSortExpressionsResponse() (response *ListSortExpressionsResponse) {
	response = &ListSortExpressionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
