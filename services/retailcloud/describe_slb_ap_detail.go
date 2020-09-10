package retailcloud

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

// DescribeSlbAPDetail invokes the retailcloud.DescribeSlbAPDetail API synchronously
// api document: https://help.aliyun.com/api/retailcloud/describeslbapdetail.html
func (client *Client) DescribeSlbAPDetail(request *DescribeSlbAPDetailRequest) (response *DescribeSlbAPDetailResponse, err error) {
	response = CreateDescribeSlbAPDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSlbAPDetailWithChan invokes the retailcloud.DescribeSlbAPDetail API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describeslbapdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlbAPDetailWithChan(request *DescribeSlbAPDetailRequest) (<-chan *DescribeSlbAPDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeSlbAPDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSlbAPDetail(request)
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

// DescribeSlbAPDetailWithCallback invokes the retailcloud.DescribeSlbAPDetail API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describeslbapdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSlbAPDetailWithCallback(request *DescribeSlbAPDetailRequest, callback func(response *DescribeSlbAPDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSlbAPDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeSlbAPDetail(request)
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

// DescribeSlbAPDetailRequest is the request struct for api DescribeSlbAPDetail
type DescribeSlbAPDetailRequest struct {
	*requests.RpcRequest
	SlbAPId requests.Integer `position:"Query" name:"SlbAPId"`
}

// DescribeSlbAPDetailResponse is the response struct for api DescribeSlbAPDetail
type DescribeSlbAPDetailResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeSlbAPDetailRequest creates a request to invoke DescribeSlbAPDetail API
func CreateDescribeSlbAPDetailRequest() (request *DescribeSlbAPDetailRequest) {
	request = &DescribeSlbAPDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DescribeSlbAPDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeSlbAPDetailResponse creates a response to parse from DescribeSlbAPDetail response
func CreateDescribeSlbAPDetailResponse() (response *DescribeSlbAPDetailResponse) {
	response = &DescribeSlbAPDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
