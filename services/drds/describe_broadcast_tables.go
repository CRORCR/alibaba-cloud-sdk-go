package drds

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

// DescribeBroadcastTables invokes the drds.DescribeBroadcastTables API synchronously
// api document: https://help.aliyun.com/api/drds/describebroadcasttables.html
func (client *Client) DescribeBroadcastTables(request *DescribeBroadcastTablesRequest) (response *DescribeBroadcastTablesResponse, err error) {
	response = CreateDescribeBroadcastTablesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBroadcastTablesWithChan invokes the drds.DescribeBroadcastTables API asynchronously
// api document: https://help.aliyun.com/api/drds/describebroadcasttables.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBroadcastTablesWithChan(request *DescribeBroadcastTablesRequest) (<-chan *DescribeBroadcastTablesResponse, <-chan error) {
	responseChan := make(chan *DescribeBroadcastTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBroadcastTables(request)
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

// DescribeBroadcastTablesWithCallback invokes the drds.DescribeBroadcastTables API asynchronously
// api document: https://help.aliyun.com/api/drds/describebroadcasttables.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBroadcastTablesWithCallback(request *DescribeBroadcastTablesRequest, callback func(response *DescribeBroadcastTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBroadcastTablesResponse
		var err error
		defer close(result)
		response, err = client.DescribeBroadcastTables(request)
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

// DescribeBroadcastTablesRequest is the request struct for api DescribeBroadcastTables
type DescribeBroadcastTablesRequest struct {
	*requests.RpcRequest
	Query          string           `position:"Query" name:"Query"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	DrdsInstanceId string           `position:"Query" name:"DrdsInstanceId"`
	DbName         string           `position:"Query" name:"DbName"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeBroadcastTablesResponse is the response struct for api DescribeBroadcastTables
type DescribeBroadcastTablesResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Success    bool       `json:"Success" xml:"Success"`
	IsShard    bool       `json:"IsShard" xml:"IsShard"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	Total      int        `json:"Total" xml:"Total"`
	List       []ListItem `json:"List" xml:"List"`
}

// CreateDescribeBroadcastTablesRequest creates a request to invoke DescribeBroadcastTables API
func CreateDescribeBroadcastTablesRequest() (request *DescribeBroadcastTablesRequest) {
	request = &DescribeBroadcastTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeBroadcastTables", "Drds", "openAPI")
	return
}

// CreateDescribeBroadcastTablesResponse creates a response to parse from DescribeBroadcastTables response
func CreateDescribeBroadcastTablesResponse() (response *DescribeBroadcastTablesResponse) {
	response = &DescribeBroadcastTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
