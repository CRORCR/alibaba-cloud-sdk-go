package push

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

// invoke QueryPushList api with *QueryPushListRequest synchronously
// api document: https://help.aliyun.com/api/push/querypushlist.html
func (client *Client) QueryPushList(request *QueryPushListRequest) (response *QueryPushListResponse, err error) {
	response = CreateQueryPushListResponse()
	err = client.DoAction(request, response)
	return
}

// invoke QueryPushList api with *QueryPushListRequest asynchronously
// api document: https://help.aliyun.com/api/push/querypushlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPushListWithChan(request *QueryPushListRequest) (<-chan *QueryPushListResponse, <-chan error) {
	responseChan := make(chan *QueryPushListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPushList(request)
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

// invoke QueryPushList api with *QueryPushListRequest asynchronously
// api document: https://help.aliyun.com/api/push/querypushlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPushListWithCallback(request *QueryPushListRequest, callback func(response *QueryPushListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPushListResponse
		var err error
		defer close(result)
		response, err = client.QueryPushList(request)
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

type QueryPushListRequest struct {
	*requests.RpcRequest
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
	PushType  string           `position:"Query" name:"PushType"`
	StartTime string           `position:"Query" name:"StartTime"`
	EndTime   string           `position:"Query" name:"EndTime"`
	Page      requests.Integer `position:"Query" name:"Page"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
}

type QueryPushListResponse struct {
	*responses.BaseResponse
	RequestId        string                          `json:"RequestId" xml:"RequestId"`
	HasNext          bool                            `json:"HasNext" xml:"HasNext"`
	Page             int                             `json:"Page" xml:"Page"`
	PageSize         int                             `json:"PageSize" xml:"PageSize"`
	PushMessageInfos PushMessageInfosInQueryPushList `json:"PushMessageInfos" xml:"PushMessageInfos"`
}

// create a request to invoke QueryPushList API
func CreateQueryPushListRequest() (request *QueryPushListRequest) {
	request = &QueryPushListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "QueryPushList", "", "")
	return
}

// create a response to parse from QueryPushList response
func CreateQueryPushListResponse() (response *QueryPushListResponse) {
	response = &QueryPushListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}