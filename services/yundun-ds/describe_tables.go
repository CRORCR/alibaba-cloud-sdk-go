package yundun_ds

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

// DescribeTables invokes the yundun_ds.DescribeTables API synchronously
// api document: https://help.aliyun.com/api/yundun-ds/describetables.html
func (client *Client) DescribeTables(request *DescribeTablesRequest) (response *DescribeTablesResponse, err error) {
	response = CreateDescribeTablesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTablesWithChan invokes the yundun_ds.DescribeTables API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describetables.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTablesWithChan(request *DescribeTablesRequest) (<-chan *DescribeTablesResponse, <-chan error) {
	responseChan := make(chan *DescribeTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTables(request)
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

// DescribeTablesWithCallback invokes the yundun_ds.DescribeTables API asynchronously
// api document: https://help.aliyun.com/api/yundun-ds/describetables.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTablesWithCallback(request *DescribeTablesRequest, callback func(response *DescribeTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTablesResponse
		var err error
		defer close(result)
		response, err = client.DescribeTables(request)
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

// DescribeTablesRequest is the request struct for api DescribeTables
type DescribeTablesRequest struct {
	*requests.RpcRequest
	ProductId   requests.Integer `position:"Query" name:"ProductId"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	PackageId   requests.Integer `position:"Query" name:"PackageId"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	QueryName   string           `position:"Query" name:"QueryName"`
	RiskLevelId requests.Integer `position:"Query" name:"RiskLevelId"`
	InstanceId  requests.Integer `position:"Query" name:"InstanceId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Name        string           `position:"Query" name:"Name"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Lang        string           `position:"Query" name:"Lang"`
	RuleId      requests.Integer `position:"Query" name:"RuleId"`
	QueryType   requests.Integer `position:"Query" name:"QueryType"`
}

// DescribeTablesResponse is the response struct for api DescribeTables
type DescribeTablesResponse struct {
	*responses.BaseResponse
	RequestId   string                  `json:"RequestId" xml:"RequestId"`
	PageSize    int                     `json:"PageSize" xml:"PageSize"`
	CurrentPage int                     `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int                     `json:"TotalCount" xml:"TotalCount"`
	Items       []TableInDescribeTables `json:"Items" xml:"Items"`
}

// CreateDescribeTablesRequest creates a request to invoke DescribeTables API
func CreateDescribeTablesRequest() (request *DescribeTablesRequest) {
	request = &DescribeTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-ds", "2019-01-03", "DescribeTables", "sddp", "openAPI")
	return
}

// CreateDescribeTablesResponse creates a response to parse from DescribeTables response
func CreateDescribeTablesResponse() (response *DescribeTablesResponse) {
	response = &DescribeTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
