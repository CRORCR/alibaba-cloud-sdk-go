package cloudcallcenter

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

// ListPrivacyNumberCallDetails invokes the cloudcallcenter.ListPrivacyNumberCallDetails API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listprivacynumbercalldetails.html
func (client *Client) ListPrivacyNumberCallDetails(request *ListPrivacyNumberCallDetailsRequest) (response *ListPrivacyNumberCallDetailsResponse, err error) {
	response = CreateListPrivacyNumberCallDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ListPrivacyNumberCallDetailsWithChan invokes the cloudcallcenter.ListPrivacyNumberCallDetails API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listprivacynumbercalldetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPrivacyNumberCallDetailsWithChan(request *ListPrivacyNumberCallDetailsRequest) (<-chan *ListPrivacyNumberCallDetailsResponse, <-chan error) {
	responseChan := make(chan *ListPrivacyNumberCallDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPrivacyNumberCallDetails(request)
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

// ListPrivacyNumberCallDetailsWithCallback invokes the cloudcallcenter.ListPrivacyNumberCallDetails API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listprivacynumbercalldetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPrivacyNumberCallDetailsWithCallback(request *ListPrivacyNumberCallDetailsRequest, callback func(response *ListPrivacyNumberCallDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPrivacyNumberCallDetailsResponse
		var err error
		defer close(result)
		response, err = client.ListPrivacyNumberCallDetails(request)
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

// ListPrivacyNumberCallDetailsRequest is the request struct for api ListPrivacyNumberCallDetails
type ListPrivacyNumberCallDetailsRequest struct {
	*requests.RpcRequest
	AgentId    string           `position:"Query" name:"AgentId"`
	ContactId  string           `position:"Query" name:"ContactId"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	AgentName  string           `position:"Query" name:"AgentName"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListPrivacyNumberCallDetailsResponse is the response struct for api ListPrivacyNumberCallDetails
type ListPrivacyNumberCallDetailsResponse struct {
	*responses.BaseResponse
	RequestId                string                   `json:"RequestId" xml:"RequestId"`
	Success                  bool                     `json:"Success" xml:"Success"`
	Code                     string                   `json:"Code" xml:"Code"`
	Message                  string                   `json:"Message" xml:"Message"`
	HttpStatusCode           int                      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PrivacyNumberCallDetails PrivacyNumberCallDetails `json:"PrivacyNumberCallDetails" xml:"PrivacyNumberCallDetails"`
}

// CreateListPrivacyNumberCallDetailsRequest creates a request to invoke ListPrivacyNumberCallDetails API
func CreateListPrivacyNumberCallDetailsRequest() (request *ListPrivacyNumberCallDetailsRequest) {
	request = &ListPrivacyNumberCallDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListPrivacyNumberCallDetails", "", "")
	request.Method = requests.POST
	return
}

// CreateListPrivacyNumberCallDetailsResponse creates a response to parse from ListPrivacyNumberCallDetails response
func CreateListPrivacyNumberCallDetailsResponse() (response *ListPrivacyNumberCallDetailsResponse) {
	response = &ListPrivacyNumberCallDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
