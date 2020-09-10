package cloudcallcenter

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

// DownloadAllTypeRecording invokes the cloudcallcenter.DownloadAllTypeRecording API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadalltyperecording.html
func (client *Client) DownloadAllTypeRecording(request *DownloadAllTypeRecordingRequest) (response *DownloadAllTypeRecordingResponse, err error) {
	response = CreateDownloadAllTypeRecordingResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadAllTypeRecordingWithChan invokes the cloudcallcenter.DownloadAllTypeRecording API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadalltyperecording.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadAllTypeRecordingWithChan(request *DownloadAllTypeRecordingRequest) (<-chan *DownloadAllTypeRecordingResponse, <-chan error) {
	responseChan := make(chan *DownloadAllTypeRecordingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadAllTypeRecording(request)
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

// DownloadAllTypeRecordingWithCallback invokes the cloudcallcenter.DownloadAllTypeRecording API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadalltyperecording.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadAllTypeRecordingWithCallback(request *DownloadAllTypeRecordingRequest, callback func(response *DownloadAllTypeRecordingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadAllTypeRecordingResponse
		var err error
		defer close(result)
		response, err = client.DownloadAllTypeRecording(request)
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

// DownloadAllTypeRecordingRequest is the request struct for api DownloadAllTypeRecording
type DownloadAllTypeRecordingRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	ContactId  string `position:"Query" name:"ContactId"`
	Channel    string `position:"Query" name:"Channel"`
}

// DownloadAllTypeRecordingResponse is the response struct for api DownloadAllTypeRecording
type DownloadAllTypeRecordingResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	Success                bool                   `json:"Success" xml:"Success"`
	Code                   string                 `json:"Code" xml:"Code"`
	Message                string                 `json:"Message" xml:"Message"`
	HttpStatusCode         int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	MediaDownloadParamList MediaDownloadParamList `json:"MediaDownloadParamList" xml:"MediaDownloadParamList"`
}

// CreateDownloadAllTypeRecordingRequest creates a request to invoke DownloadAllTypeRecording API
func CreateDownloadAllTypeRecordingRequest() (request *DownloadAllTypeRecordingRequest) {
	request = &DownloadAllTypeRecordingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DownloadAllTypeRecording", "", "")
	request.Method = requests.POST
	return
}

// CreateDownloadAllTypeRecordingResponse creates a response to parse from DownloadAllTypeRecording response
func CreateDownloadAllTypeRecordingResponse() (response *DownloadAllTypeRecordingResponse) {
	response = &DownloadAllTypeRecordingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
