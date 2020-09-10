package emr

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

// ListClusterServiceConfigHistory invokes the emr.ListClusterServiceConfigHistory API synchronously
// api document: https://help.aliyun.com/api/emr/listclusterserviceconfighistory.html
func (client *Client) ListClusterServiceConfigHistory(request *ListClusterServiceConfigHistoryRequest) (response *ListClusterServiceConfigHistoryResponse, err error) {
	response = CreateListClusterServiceConfigHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterServiceConfigHistoryWithChan invokes the emr.ListClusterServiceConfigHistory API asynchronously
// api document: https://help.aliyun.com/api/emr/listclusterserviceconfighistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClusterServiceConfigHistoryWithChan(request *ListClusterServiceConfigHistoryRequest) (<-chan *ListClusterServiceConfigHistoryResponse, <-chan error) {
	responseChan := make(chan *ListClusterServiceConfigHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterServiceConfigHistory(request)
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

// ListClusterServiceConfigHistoryWithCallback invokes the emr.ListClusterServiceConfigHistory API asynchronously
// api document: https://help.aliyun.com/api/emr/listclusterserviceconfighistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClusterServiceConfigHistoryWithCallback(request *ListClusterServiceConfigHistoryRequest, callback func(response *ListClusterServiceConfigHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterServiceConfigHistoryResponse
		var err error
		defer close(result)
		response, err = client.ListClusterServiceConfigHistory(request)
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

// ListClusterServiceConfigHistoryRequest is the request struct for api ListClusterServiceConfigHistory
type ListClusterServiceConfigHistoryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ConfigVersion   string           `position:"Query" name:"ConfigVersion"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ServiceName     string           `position:"Query" name:"ServiceName"`
}

// ListClusterServiceConfigHistoryResponse is the response struct for api ListClusterServiceConfigHistory
type ListClusterServiceConfigHistoryResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	TotalCount        int               `json:"TotalCount" xml:"TotalCount"`
	PageNumber        int               `json:"PageNumber" xml:"PageNumber"`
	PageSize          int               `json:"PageSize" xml:"PageSize"`
	ConfigHistoryList ConfigHistoryList `json:"ConfigHistoryList" xml:"ConfigHistoryList"`
}

// CreateListClusterServiceConfigHistoryRequest creates a request to invoke ListClusterServiceConfigHistory API
func CreateListClusterServiceConfigHistoryRequest() (request *ListClusterServiceConfigHistoryRequest) {
	request = &ListClusterServiceConfigHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceConfigHistory", "emr", "openAPI")
	return
}

// CreateListClusterServiceConfigHistoryResponse creates a response to parse from ListClusterServiceConfigHistory response
func CreateListClusterServiceConfigHistoryResponse() (response *ListClusterServiceConfigHistoryResponse) {
	response = &ListClusterServiceConfigHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
