package qualitycheck

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

// ListTaskAssignRules invokes the qualitycheck.ListTaskAssignRules API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/listtaskassignrules.html
func (client *Client) ListTaskAssignRules(request *ListTaskAssignRulesRequest) (response *ListTaskAssignRulesResponse, err error) {
	response = CreateListTaskAssignRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListTaskAssignRulesWithChan invokes the qualitycheck.ListTaskAssignRules API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/listtaskassignrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTaskAssignRulesWithChan(request *ListTaskAssignRulesRequest) (<-chan *ListTaskAssignRulesResponse, <-chan error) {
	responseChan := make(chan *ListTaskAssignRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTaskAssignRules(request)
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

// ListTaskAssignRulesWithCallback invokes the qualitycheck.ListTaskAssignRules API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/listtaskassignrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListTaskAssignRulesWithCallback(request *ListTaskAssignRulesRequest, callback func(response *ListTaskAssignRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTaskAssignRulesResponse
		var err error
		defer close(result)
		response, err = client.ListTaskAssignRules(request)
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

// ListTaskAssignRulesRequest is the request struct for api ListTaskAssignRules
type ListTaskAssignRulesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// ListTaskAssignRulesResponse is the response struct for api ListTaskAssignRules
type ListTaskAssignRulesResponse struct {
	*responses.BaseResponse
	RequestId  string                    `json:"RequestId" xml:"RequestId"`
	Success    bool                      `json:"Success" xml:"Success"`
	Code       string                    `json:"Code" xml:"Code"`
	Message    string                    `json:"Message" xml:"Message"`
	PageNumber int                       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                       `json:"PageSize" xml:"PageSize"`
	Count      int                       `json:"Count" xml:"Count"`
	Data       DataInListTaskAssignRules `json:"Data" xml:"Data"`
}

// CreateListTaskAssignRulesRequest creates a request to invoke ListTaskAssignRules API
func CreateListTaskAssignRulesRequest() (request *ListTaskAssignRulesRequest) {
	request = &ListTaskAssignRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "ListTaskAssignRules", "", "")
	return
}

// CreateListTaskAssignRulesResponse creates a response to parse from ListTaskAssignRules response
func CreateListTaskAssignRulesResponse() (response *ListTaskAssignRulesResponse) {
	response = &ListTaskAssignRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
