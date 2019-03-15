package finmall

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
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetOverdueRecordList invokes the finmall.GetOverdueRecordList API synchronously
// api document: https://help.aliyun.com/api/finmall/getoverduerecordlist.html
func (client *Client) GetOverdueRecordList(request *GetOverdueRecordListRequest) (response *GetOverdueRecordListResponse, err error) {
	response = CreateGetOverdueRecordListResponse()
	err = client.DoAction(request, response)
	return
}

// GetOverdueRecordListWithChan invokes the finmall.GetOverdueRecordList API asynchronously
// api document: https://help.aliyun.com/api/finmall/getoverduerecordlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOverdueRecordListWithChan(request *GetOverdueRecordListRequest) (<-chan *GetOverdueRecordListResponse, <-chan error) {
	responseChan := make(chan *GetOverdueRecordListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOverdueRecordList(request)
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

// GetOverdueRecordListWithCallback invokes the finmall.GetOverdueRecordList API asynchronously
// api document: https://help.aliyun.com/api/finmall/getoverduerecordlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOverdueRecordListWithCallback(request *GetOverdueRecordListRequest, callback func(response *GetOverdueRecordListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOverdueRecordListResponse
		var err error
		defer close(result)
		response, err = client.GetOverdueRecordList(request)
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

// GetOverdueRecordListRequest is the request struct for api GetOverdueRecordList
type GetOverdueRecordListRequest struct {
	*requests.RpcRequest
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	QueryExpression string           `position:"Query" name:"QueryExpression"`
	UserId          string           `position:"Query" name:"UserId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
}

// GetOverdueRecordListResponse is the response struct for api GetOverdueRecordList
type GetOverdueRecordListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetOverdueRecordListRequest creates a request to invoke GetOverdueRecordList API
func CreateGetOverdueRecordListRequest() (request *GetOverdueRecordListRequest) {
	request = &GetOverdueRecordListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("finmall", "2018-07-23", "GetOverdueRecordList", "finmall", "openAPI")
	return
}

// CreateGetOverdueRecordListResponse creates a response to parse from GetOverdueRecordList response
func CreateGetOverdueRecordListResponse() (response *GetOverdueRecordListResponse) {
	response = &GetOverdueRecordListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}