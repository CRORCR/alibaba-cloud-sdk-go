package green

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

// DescribeOssCallbackSetting invokes the green.DescribeOssCallbackSetting API synchronously
// api document: https://help.aliyun.com/api/green/describeosscallbacksetting.html
func (client *Client) DescribeOssCallbackSetting(request *DescribeOssCallbackSettingRequest) (response *DescribeOssCallbackSettingResponse, err error) {
	response = CreateDescribeOssCallbackSettingResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOssCallbackSettingWithChan invokes the green.DescribeOssCallbackSetting API asynchronously
// api document: https://help.aliyun.com/api/green/describeosscallbacksetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOssCallbackSettingWithChan(request *DescribeOssCallbackSettingRequest) (<-chan *DescribeOssCallbackSettingResponse, <-chan error) {
	responseChan := make(chan *DescribeOssCallbackSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOssCallbackSetting(request)
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

// DescribeOssCallbackSettingWithCallback invokes the green.DescribeOssCallbackSetting API asynchronously
// api document: https://help.aliyun.com/api/green/describeosscallbacksetting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOssCallbackSettingWithCallback(request *DescribeOssCallbackSettingRequest, callback func(response *DescribeOssCallbackSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOssCallbackSettingResponse
		var err error
		defer close(result)
		response, err = client.DescribeOssCallbackSetting(request)
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

// DescribeOssCallbackSettingRequest is the request struct for api DescribeOssCallbackSetting
type DescribeOssCallbackSettingRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// DescribeOssCallbackSettingResponse is the response struct for api DescribeOssCallbackSetting
type DescribeOssCallbackSettingResponse struct {
	*responses.BaseResponse
	RequestId               string   `json:"RequestId" xml:"RequestId"`
	CallbackUrl             string   `json:"CallbackUrl" xml:"CallbackUrl"`
	ScanCallback            bool     `json:"ScanCallback" xml:"ScanCallback"`
	AuditCallback           bool     `json:"AuditCallback" xml:"AuditCallback"`
	CallbackSeed            string   `json:"CallbackSeed" xml:"CallbackSeed"`
	ScanCallbackSuggestions []string `json:"ScanCallbackSuggestions" xml:"ScanCallbackSuggestions"`
	ServiceModules          []string `json:"ServiceModules" xml:"ServiceModules"`
}

// CreateDescribeOssCallbackSettingRequest creates a request to invoke DescribeOssCallbackSetting API
func CreateDescribeOssCallbackSettingRequest() (request *DescribeOssCallbackSettingRequest) {
	request = &DescribeOssCallbackSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeOssCallbackSetting", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeOssCallbackSettingResponse creates a response to parse from DescribeOssCallbackSetting response
func CreateDescribeOssCallbackSettingResponse() (response *DescribeOssCallbackSettingResponse) {
	response = &DescribeOssCallbackSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
