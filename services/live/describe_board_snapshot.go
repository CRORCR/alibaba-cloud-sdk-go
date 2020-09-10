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

// DescribeBoardSnapshot invokes the live.DescribeBoardSnapshot API synchronously
// api document: https://help.aliyun.com/api/live/describeboardsnapshot.html
func (client *Client) DescribeBoardSnapshot(request *DescribeBoardSnapshotRequest) (response *DescribeBoardSnapshotResponse, err error) {
	response = CreateDescribeBoardSnapshotResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBoardSnapshotWithChan invokes the live.DescribeBoardSnapshot API asynchronously
// api document: https://help.aliyun.com/api/live/describeboardsnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBoardSnapshotWithChan(request *DescribeBoardSnapshotRequest) (<-chan *DescribeBoardSnapshotResponse, <-chan error) {
	responseChan := make(chan *DescribeBoardSnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBoardSnapshot(request)
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

// DescribeBoardSnapshotWithCallback invokes the live.DescribeBoardSnapshot API asynchronously
// api document: https://help.aliyun.com/api/live/describeboardsnapshot.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBoardSnapshotWithCallback(request *DescribeBoardSnapshotRequest, callback func(response *DescribeBoardSnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBoardSnapshotResponse
		var err error
		defer close(result)
		response, err = client.DescribeBoardSnapshot(request)
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

// DescribeBoardSnapshotRequest is the request struct for api DescribeBoardSnapshot
type DescribeBoardSnapshotRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
	BoardId string           `position:"Query" name:"BoardId"`
}

// DescribeBoardSnapshotResponse is the response struct for api DescribeBoardSnapshot
type DescribeBoardSnapshotResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Snapshot  Snapshot `json:"Snapshot" xml:"Snapshot"`
}

// CreateDescribeBoardSnapshotRequest creates a request to invoke DescribeBoardSnapshot API
func CreateDescribeBoardSnapshotRequest() (request *DescribeBoardSnapshotRequest) {
	request = &DescribeBoardSnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeBoardSnapshot", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeBoardSnapshotResponse creates a response to parse from DescribeBoardSnapshot response
func CreateDescribeBoardSnapshotResponse() (response *DescribeBoardSnapshotResponse) {
	response = &DescribeBoardSnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
