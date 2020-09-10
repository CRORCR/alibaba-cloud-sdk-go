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

// DescribeLiveAudioAuditNotifyConfig invokes the live.DescribeLiveAudioAuditNotifyConfig API synchronously
// api document: https://help.aliyun.com/api/live/describeliveaudioauditnotifyconfig.html
func (client *Client) DescribeLiveAudioAuditNotifyConfig(request *DescribeLiveAudioAuditNotifyConfigRequest) (response *DescribeLiveAudioAuditNotifyConfigResponse, err error) {
	response = CreateDescribeLiveAudioAuditNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveAudioAuditNotifyConfigWithChan invokes the live.DescribeLiveAudioAuditNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/describeliveaudioauditnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveAudioAuditNotifyConfigWithChan(request *DescribeLiveAudioAuditNotifyConfigRequest) (<-chan *DescribeLiveAudioAuditNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveAudioAuditNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveAudioAuditNotifyConfig(request)
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

// DescribeLiveAudioAuditNotifyConfigWithCallback invokes the live.DescribeLiveAudioAuditNotifyConfig API asynchronously
// api document: https://help.aliyun.com/api/live/describeliveaudioauditnotifyconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveAudioAuditNotifyConfigWithCallback(request *DescribeLiveAudioAuditNotifyConfigRequest, callback func(response *DescribeLiveAudioAuditNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveAudioAuditNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveAudioAuditNotifyConfig(request)
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

// DescribeLiveAudioAuditNotifyConfigRequest is the request struct for api DescribeLiveAudioAuditNotifyConfig
type DescribeLiveAudioAuditNotifyConfigRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveAudioAuditNotifyConfigResponse is the response struct for api DescribeLiveAudioAuditNotifyConfig
type DescribeLiveAudioAuditNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId                      string                         `json:"RequestId" xml:"RequestId"`
	LiveAudioAuditNotifyConfigList LiveAudioAuditNotifyConfigList `json:"LiveAudioAuditNotifyConfigList" xml:"LiveAudioAuditNotifyConfigList"`
}

// CreateDescribeLiveAudioAuditNotifyConfigRequest creates a request to invoke DescribeLiveAudioAuditNotifyConfig API
func CreateDescribeLiveAudioAuditNotifyConfigRequest() (request *DescribeLiveAudioAuditNotifyConfigRequest) {
	request = &DescribeLiveAudioAuditNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveAudioAuditNotifyConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveAudioAuditNotifyConfigResponse creates a response to parse from DescribeLiveAudioAuditNotifyConfig response
func CreateDescribeLiveAudioAuditNotifyConfigResponse() (response *DescribeLiveAudioAuditNotifyConfigResponse) {
	response = &DescribeLiveAudioAuditNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
