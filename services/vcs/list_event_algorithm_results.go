package vcs

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

// ListEventAlgorithmResults invokes the vcs.ListEventAlgorithmResults API synchronously
// api document: https://help.aliyun.com/api/vcs/listeventalgorithmresults.html
func (client *Client) ListEventAlgorithmResults(request *ListEventAlgorithmResultsRequest) (response *ListEventAlgorithmResultsResponse, err error) {
	response = CreateListEventAlgorithmResultsResponse()
	err = client.DoAction(request, response)
	return
}

// ListEventAlgorithmResultsWithChan invokes the vcs.ListEventAlgorithmResults API asynchronously
// api document: https://help.aliyun.com/api/vcs/listeventalgorithmresults.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEventAlgorithmResultsWithChan(request *ListEventAlgorithmResultsRequest) (<-chan *ListEventAlgorithmResultsResponse, <-chan error) {
	responseChan := make(chan *ListEventAlgorithmResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEventAlgorithmResults(request)
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

// ListEventAlgorithmResultsWithCallback invokes the vcs.ListEventAlgorithmResults API asynchronously
// api document: https://help.aliyun.com/api/vcs/listeventalgorithmresults.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEventAlgorithmResultsWithCallback(request *ListEventAlgorithmResultsRequest, callback func(response *ListEventAlgorithmResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEventAlgorithmResultsResponse
		var err error
		defer close(result)
		response, err = client.ListEventAlgorithmResults(request)
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

// ListEventAlgorithmResultsRequest is the request struct for api ListEventAlgorithmResults
type ListEventAlgorithmResultsRequest struct {
	*requests.RpcRequest
	CorpId       string `position:"Body" name:"CorpId"`
	ExtendValue  string `position:"Body" name:"ExtendValue"`
	EndTime      string `position:"Body" name:"EndTime"`
	StartTime    string `position:"Body" name:"StartTime"`
	PageNumber   string `position:"Body" name:"PageNumber"`
	DataSourceId string `position:"Body" name:"DataSourceId"`
	PageSize     string `position:"Body" name:"PageSize"`
	EventType    string `position:"Body" name:"EventType"`
}

// ListEventAlgorithmResultsResponse is the response struct for api ListEventAlgorithmResults
type ListEventAlgorithmResultsResponse struct {
	*responses.BaseResponse
	Code        string                          `json:"Code" xml:"Code"`
	Message     string                          `json:"Message" xml:"Message"`
	RequestId   string                          `json:"RequestId" xml:"RequestId"`
	ExtendValue string                          `json:"ExtendValue" xml:"ExtendValue"`
	Data        DataInListEventAlgorithmResults `json:"Data" xml:"Data"`
}

// CreateListEventAlgorithmResultsRequest creates a request to invoke ListEventAlgorithmResults API
func CreateListEventAlgorithmResultsRequest() (request *ListEventAlgorithmResultsRequest) {
	request = &ListEventAlgorithmResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListEventAlgorithmResults", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListEventAlgorithmResultsResponse creates a response to parse from ListEventAlgorithmResults response
func CreateListEventAlgorithmResultsResponse() (response *ListEventAlgorithmResultsResponse) {
	response = &ListEventAlgorithmResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
