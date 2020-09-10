package hbase

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

// DescribeRestoreIncrDetail invokes the hbase.DescribeRestoreIncrDetail API synchronously
// api document: https://help.aliyun.com/api/hbase/describerestoreincrdetail.html
func (client *Client) DescribeRestoreIncrDetail(request *DescribeRestoreIncrDetailRequest) (response *DescribeRestoreIncrDetailResponse, err error) {
	response = CreateDescribeRestoreIncrDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRestoreIncrDetailWithChan invokes the hbase.DescribeRestoreIncrDetail API asynchronously
// api document: https://help.aliyun.com/api/hbase/describerestoreincrdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRestoreIncrDetailWithChan(request *DescribeRestoreIncrDetailRequest) (<-chan *DescribeRestoreIncrDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeRestoreIncrDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRestoreIncrDetail(request)
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

// DescribeRestoreIncrDetailWithCallback invokes the hbase.DescribeRestoreIncrDetail API asynchronously
// api document: https://help.aliyun.com/api/hbase/describerestoreincrdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRestoreIncrDetailWithCallback(request *DescribeRestoreIncrDetailRequest, callback func(response *DescribeRestoreIncrDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRestoreIncrDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeRestoreIncrDetail(request)
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

// DescribeRestoreIncrDetailRequest is the request struct for api DescribeRestoreIncrDetail
type DescribeRestoreIncrDetailRequest struct {
	*requests.RpcRequest
	ClusterId       string `position:"Query" name:"ClusterId"`
	RestoreRecordId string `position:"Query" name:"RestoreRecordId"`
}

// DescribeRestoreIncrDetailResponse is the response struct for api DescribeRestoreIncrDetail
type DescribeRestoreIncrDetailResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	RestoreIncrDetail RestoreIncrDetail `json:"RestoreIncrDetail" xml:"RestoreIncrDetail"`
}

// CreateDescribeRestoreIncrDetailRequest creates a request to invoke DescribeRestoreIncrDetail API
func CreateDescribeRestoreIncrDetailRequest() (request *DescribeRestoreIncrDetailRequest) {
	request = &DescribeRestoreIncrDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "DescribeRestoreIncrDetail", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeRestoreIncrDetailResponse creates a response to parse from DescribeRestoreIncrDetail response
func CreateDescribeRestoreIncrDetailResponse() (response *DescribeRestoreIncrDetailResponse) {
	response = &DescribeRestoreIncrDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
