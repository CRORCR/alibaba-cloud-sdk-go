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

// DeleteLiveStreamsNotifyUrlConfig invokes the live.DeleteLiveStreamsNotifyUrlConfig API synchronously
// api document: https://help.aliyun.com/api/live/deletelivestreamsnotifyurlconfig.html
func (client *Client) DeleteLiveStreamsNotifyUrlConfig(request *DeleteLiveStreamsNotifyUrlConfigRequest) (response *DeleteLiveStreamsNotifyUrlConfigResponse, err error) {
	response = CreateDeleteLiveStreamsNotifyUrlConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveStreamsNotifyUrlConfigWithChan invokes the live.DeleteLiveStreamsNotifyUrlConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deletelivestreamsnotifyurlconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveStreamsNotifyUrlConfigWithChan(request *DeleteLiveStreamsNotifyUrlConfigRequest) (<-chan *DeleteLiveStreamsNotifyUrlConfigResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveStreamsNotifyUrlConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveStreamsNotifyUrlConfig(request)
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

// DeleteLiveStreamsNotifyUrlConfigWithCallback invokes the live.DeleteLiveStreamsNotifyUrlConfig API asynchronously
// api document: https://help.aliyun.com/api/live/deletelivestreamsnotifyurlconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLiveStreamsNotifyUrlConfigWithCallback(request *DeleteLiveStreamsNotifyUrlConfigRequest, callback func(response *DeleteLiveStreamsNotifyUrlConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveStreamsNotifyUrlConfigResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveStreamsNotifyUrlConfig(request)
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

// DeleteLiveStreamsNotifyUrlConfigRequest is the request struct for api DeleteLiveStreamsNotifyUrlConfig
type DeleteLiveStreamsNotifyUrlConfigRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteLiveStreamsNotifyUrlConfigResponse is the response struct for api DeleteLiveStreamsNotifyUrlConfig
type DeleteLiveStreamsNotifyUrlConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLiveStreamsNotifyUrlConfigRequest creates a request to invoke DeleteLiveStreamsNotifyUrlConfig API
func CreateDeleteLiveStreamsNotifyUrlConfigRequest() (request *DeleteLiveStreamsNotifyUrlConfigRequest) {
	request = &DeleteLiveStreamsNotifyUrlConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveStreamsNotifyUrlConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteLiveStreamsNotifyUrlConfigResponse creates a response to parse from DeleteLiveStreamsNotifyUrlConfig response
func CreateDeleteLiveStreamsNotifyUrlConfigResponse() (response *DeleteLiveStreamsNotifyUrlConfigResponse) {
	response = &DeleteLiveStreamsNotifyUrlConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
