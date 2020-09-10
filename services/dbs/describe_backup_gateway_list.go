package dbs

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

// DescribeBackupGatewayList invokes the dbs.DescribeBackupGatewayList API synchronously
func (client *Client) DescribeBackupGatewayList(request *DescribeBackupGatewayListRequest) (response *DescribeBackupGatewayListResponse, err error) {
	response = CreateDescribeBackupGatewayListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupGatewayListWithChan invokes the dbs.DescribeBackupGatewayList API asynchronously
func (client *Client) DescribeBackupGatewayListWithChan(request *DescribeBackupGatewayListRequest) (<-chan *DescribeBackupGatewayListResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupGatewayListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupGatewayList(request)
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

// DescribeBackupGatewayListWithCallback invokes the dbs.DescribeBackupGatewayList API asynchronously
func (client *Client) DescribeBackupGatewayListWithCallback(request *DescribeBackupGatewayListRequest, callback func(response *DescribeBackupGatewayListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupGatewayListResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupGatewayList(request)
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

// DescribeBackupGatewayListRequest is the request struct for api DescribeBackupGatewayList
type DescribeBackupGatewayListRequest struct {
	*requests.RpcRequest
	Identifier  string           `position:"Query" name:"Identifier"`
	ClientToken string           `position:"Query" name:"ClientToken"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	OwnerId     string           `position:"Query" name:"OwnerId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Region      string           `position:"Query" name:"Region"`
}

// DescribeBackupGatewayListResponse is the response struct for api DescribeBackupGatewayList
type DescribeBackupGatewayListResponse struct {
	*responses.BaseResponse
	Success        bool                             `json:"Success" xml:"Success"`
	ErrCode        string                           `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string                           `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int                              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                           `json:"RequestId" xml:"RequestId"`
	TotalPages     int                              `json:"TotalPages" xml:"TotalPages"`
	PageSize       int                              `json:"PageSize" xml:"PageSize"`
	PageNum        int                              `json:"PageNum" xml:"PageNum"`
	TotalElements  int                              `json:"TotalElements" xml:"TotalElements"`
	Items          ItemsInDescribeBackupGatewayList `json:"Items" xml:"Items"`
}

// CreateDescribeBackupGatewayListRequest creates a request to invoke DescribeBackupGatewayList API
func CreateDescribeBackupGatewayListRequest() (request *DescribeBackupGatewayListRequest) {
	request = &DescribeBackupGatewayListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DescribeBackupGatewayList", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupGatewayListResponse creates a response to parse from DescribeBackupGatewayList response
func CreateDescribeBackupGatewayListResponse() (response *DescribeBackupGatewayListResponse) {
	response = &DescribeBackupGatewayListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
