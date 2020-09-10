package aliyuncvc

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

// JoinLive invokes the aliyuncvc.JoinLive API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/joinlive.html
func (client *Client) JoinLive(request *JoinLiveRequest) (response *JoinLiveResponse, err error) {
	response = CreateJoinLiveResponse()
	err = client.DoAction(request, response)
	return
}

// JoinLiveWithChan invokes the aliyuncvc.JoinLive API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/joinlive.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinLiveWithChan(request *JoinLiveRequest) (<-chan *JoinLiveResponse, <-chan error) {
	responseChan := make(chan *JoinLiveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.JoinLive(request)
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

// JoinLiveWithCallback invokes the aliyuncvc.JoinLive API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/joinlive.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinLiveWithCallback(request *JoinLiveRequest, callback func(response *JoinLiveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *JoinLiveResponse
		var err error
		defer close(result)
		response, err = client.JoinLive(request)
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

// JoinLiveRequest is the request struct for api JoinLive
type JoinLiveRequest struct {
	*requests.RpcRequest
	LiveUUID string `position:"Body" name:"LiveUUID"`
	UserId   string `position:"Body" name:"UserId"`
	Password string `position:"Body" name:"Password"`
}

// JoinLiveResponse is the response struct for api JoinLive
type JoinLiveResponse struct {
	*responses.BaseResponse
	ErrorCode   int         `json:"ErrorCode" xml:"ErrorCode"`
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Message     string      `json:"Message" xml:"Message"`
	MeetingInfo MeetingInfo `json:"MeetingInfo" xml:"MeetingInfo"`
}

// CreateJoinLiveRequest creates a request to invoke JoinLive API
func CreateJoinLiveRequest() (request *JoinLiveRequest) {
	request = &JoinLiveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "JoinLive", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateJoinLiveResponse creates a response to parse from JoinLive response
func CreateJoinLiveResponse() (response *JoinLiveResponse) {
	response = &JoinLiveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
