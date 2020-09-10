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

// DescribeRecycleBinTables invokes the drds.DescribeRecycleBinTables API synchronously
// api document: https://help.aliyun.com/api/drds/describerecyclebintables.html
func (client *Client) DescribeRecycleBinTables(request *DescribeRecycleBinTablesRequest) (response *DescribeRecycleBinTablesResponse, err error) {
	response = CreateDescribeRecycleBinTablesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecycleBinTablesWithChan invokes the drds.DescribeRecycleBinTables API asynchronously
// api document: https://help.aliyun.com/api/drds/describerecyclebintables.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecycleBinTablesWithChan(request *DescribeRecycleBinTablesRequest) (<-chan *DescribeRecycleBinTablesResponse, <-chan error) {
	responseChan := make(chan *DescribeRecycleBinTablesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecycleBinTables(request)
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

// DescribeRecycleBinTablesWithCallback invokes the drds.DescribeRecycleBinTables API asynchronously
// api document: https://help.aliyun.com/api/drds/describerecyclebintables.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecycleBinTablesWithCallback(request *DescribeRecycleBinTablesRequest, callback func(response *DescribeRecycleBinTablesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecycleBinTablesResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecycleBinTables(request)
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

// DescribeRecycleBinTablesRequest is the request struct for api DescribeRecycleBinTables
type DescribeRecycleBinTablesRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// DescribeRecycleBinTablesResponse is the response struct for api DescribeRecycleBinTables
type DescribeRecycleBinTablesResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Success   bool       `json:"Success" xml:"Success"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateDescribeRecycleBinTablesRequest creates a request to invoke DescribeRecycleBinTables API
func CreateDescribeRecycleBinTablesRequest() (request *DescribeRecycleBinTablesRequest) {
	request = &DescribeRecycleBinTablesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeRecycleBinTables", "Drds", "openAPI")
	return
}

// CreateDescribeRecycleBinTablesResponse creates a response to parse from DescribeRecycleBinTables response
func CreateDescribeRecycleBinTablesResponse() (response *DescribeRecycleBinTablesResponse) {
	response = &DescribeRecycleBinTablesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
