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
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// PushNoticeToiOS invokes the push.PushNoticeToiOS API synchronously
// api document: https://help.aliyun.com/api/push/pushnoticetoios.html
func (client *Client) PushNoticeToiOS(request *PushNoticeToiOSRequest) (response *PushNoticeToiOSResponse, err error) {
	response = CreatePushNoticeToiOSResponse()
	err = client.DoAction(request, response)
	return
}

// PushNoticeToiOSWithChan invokes the push.PushNoticeToiOS API asynchronously
// api document: https://help.aliyun.com/api/push/pushnoticetoios.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PushNoticeToiOSWithChan(request *PushNoticeToiOSRequest) (<-chan *PushNoticeToiOSResponse, <-chan error) {
	responseChan := make(chan *PushNoticeToiOSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushNoticeToiOS(request)
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

// PushNoticeToiOSWithCallback invokes the push.PushNoticeToiOS API asynchronously
// api document: https://help.aliyun.com/api/push/pushnoticetoios.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PushNoticeToiOSWithCallback(request *PushNoticeToiOSRequest, callback func(response *PushNoticeToiOSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushNoticeToiOSResponse
		var err error
		defer close(result)
		response, err = client.PushNoticeToiOS(request)
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

// PushNoticeToiOSRequest is the request struct for api PushNoticeToiOS
type PushNoticeToiOSRequest struct {
	*requests.RpcRequest
	ExtParameters         string           `position:"Query" name:"ExtParameters"`
	ApnsEnv               string           `position:"Query" name:"ApnsEnv"`
	Title                 string           `position:"Query" name:"Title"`
	Body                  string           `position:"Query" name:"Body"`
	JobKey                string           `position:"Query" name:"JobKey"`
	Target                string           `position:"Query" name:"Target"`
	AppKey                requests.Integer `position:"Query" name:"AppKey"`
	TargetValue           string           `position:"Query" name:"TargetValue"`
	Badge                 requests.Integer `position:"Query" name:"Badge"`
	IOSBadge              requests.Integer `position:"Query" name:"IOSBadge"`
	IOSBadgeAutoIncrement bool             `position:"Query" name:"IOSBadgeAutoIncrement"`
}

// PushNoticeToiOSResponse is the response struct for api PushNoticeToiOS
type PushNoticeToiOSResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MessageId string `json:"MessageId" xml:"MessageId"`
}

// CreatePushNoticeToiOSRequest creates a request to invoke PushNoticeToiOS API
func CreatePushNoticeToiOSRequest() (request *PushNoticeToiOSRequest) {
	request = &PushNoticeToiOSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "PushNoticeToiOS", "", "")
	request.Method = requests.POST
	return
}

// CreatePushNoticeToiOSResponse creates a response to parse from PushNoticeToiOS response
func CreatePushNoticeToiOSResponse() (response *PushNoticeToiOSResponse) {
	response = &PushNoticeToiOSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
