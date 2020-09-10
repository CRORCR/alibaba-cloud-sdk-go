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

// DescribeHotDbList invokes the drds.DescribeHotDbList API synchronously
// api document: https://help.aliyun.com/api/drds/describehotdblist.html
func (client *Client) DescribeHotDbList(request *DescribeHotDbListRequest) (response *DescribeHotDbListResponse, err error) {
	response = CreateDescribeHotDbListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHotDbListWithChan invokes the drds.DescribeHotDbList API asynchronously
// api document: https://help.aliyun.com/api/drds/describehotdblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHotDbListWithChan(request *DescribeHotDbListRequest) (<-chan *DescribeHotDbListResponse, <-chan error) {
	responseChan := make(chan *DescribeHotDbListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHotDbList(request)
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

// DescribeHotDbListWithCallback invokes the drds.DescribeHotDbList API asynchronously
// api document: https://help.aliyun.com/api/drds/describehotdblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHotDbListWithCallback(request *DescribeHotDbListRequest, callback func(response *DescribeHotDbListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHotDbListResponse
		var err error
		defer close(result)
		response, err = client.DescribeHotDbList(request)
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

// DescribeHotDbListRequest is the request struct for api DescribeHotDbList
type DescribeHotDbListRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// DescribeHotDbListResponse is the response struct for api DescribeHotDbList
type DescribeHotDbListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Msg       string `json:"Msg" xml:"Msg"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeHotDbListRequest creates a request to invoke DescribeHotDbList API
func CreateDescribeHotDbListRequest() (request *DescribeHotDbListRequest) {
	request = &DescribeHotDbListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeHotDbList", "Drds", "openAPI")
	return
}

// CreateDescribeHotDbListResponse creates a response to parse from DescribeHotDbList response
func CreateDescribeHotDbListResponse() (response *DescribeHotDbListResponse) {
	response = &DescribeHotDbListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
