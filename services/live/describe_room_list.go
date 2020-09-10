package live

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

// DescribeRoomList invokes the live.DescribeRoomList API synchronously
// api document: https://help.aliyun.com/api/live/describeroomlist.html
func (client *Client) DescribeRoomList(request *DescribeRoomListRequest) (response *DescribeRoomListResponse, err error) {
	response = CreateDescribeRoomListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRoomListWithChan invokes the live.DescribeRoomList API asynchronously
// api document: https://help.aliyun.com/api/live/describeroomlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRoomListWithChan(request *DescribeRoomListRequest) (<-chan *DescribeRoomListResponse, <-chan error) {
	responseChan := make(chan *DescribeRoomListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRoomList(request)
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

// DescribeRoomListWithCallback invokes the live.DescribeRoomList API asynchronously
// api document: https://help.aliyun.com/api/live/describeroomlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRoomListWithCallback(request *DescribeRoomListRequest, callback func(response *DescribeRoomListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRoomListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRoomList(request)
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

// DescribeRoomListRequest is the request struct for api DescribeRoomList
type DescribeRoomListRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	AnchorId   string           `position:"Query" name:"AnchorId"`
	PageNum    requests.Integer `position:"Query" name:"PageNum"`
	RoomStatus requests.Integer `position:"Query" name:"RoomStatus"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Order      string           `position:"Query" name:"Order"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	RoomId     string           `position:"Query" name:"RoomId"`
	AppId      string           `position:"Query" name:"AppId"`
}

// DescribeRoomListResponse is the response struct for api DescribeRoomList
type DescribeRoomListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TotalNum  int    `json:"TotalNum" xml:"TotalNum"`
	TotalPage int    `json:"TotalPage" xml:"TotalPage"`
	RoomList  []Room `json:"RoomList" xml:"RoomList"`
}

// CreateDescribeRoomListRequest creates a request to invoke DescribeRoomList API
func CreateDescribeRoomListRequest() (request *DescribeRoomListRequest) {
	request = &DescribeRoomListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeRoomList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRoomListResponse creates a response to parse from DescribeRoomList response
func CreateDescribeRoomListResponse() (response *DescribeRoomListResponse) {
	response = &DescribeRoomListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
