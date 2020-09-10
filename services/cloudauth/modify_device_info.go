package cloudauth

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

// ModifyDeviceInfo invokes the cloudauth.ModifyDeviceInfo API synchronously
// api document: https://help.aliyun.com/api/cloudauth/modifydeviceinfo.html
func (client *Client) ModifyDeviceInfo(request *ModifyDeviceInfoRequest) (response *ModifyDeviceInfoResponse, err error) {
	response = CreateModifyDeviceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDeviceInfoWithChan invokes the cloudauth.ModifyDeviceInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/modifydeviceinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDeviceInfoWithChan(request *ModifyDeviceInfoRequest) (<-chan *ModifyDeviceInfoResponse, <-chan error) {
	responseChan := make(chan *ModifyDeviceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDeviceInfo(request)
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

// ModifyDeviceInfoWithCallback invokes the cloudauth.ModifyDeviceInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/modifydeviceinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDeviceInfoWithCallback(request *ModifyDeviceInfoRequest, callback func(response *ModifyDeviceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDeviceInfoResponse
		var err error
		defer close(result)
		response, err = client.ModifyDeviceInfo(request)
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

// ModifyDeviceInfoRequest is the request struct for api ModifyDeviceInfo
type ModifyDeviceInfoRequest struct {
	*requests.RpcRequest
	UserDeviceId string `position:"Query" name:"UserDeviceId"`
	Duration     string `position:"Query" name:"Duration"`
	ExpiredDay   string `position:"Query" name:"ExpiredDay"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	Lang         string `position:"Query" name:"Lang"`
	DeviceId     string `position:"Query" name:"DeviceId"`
	BizType      string `position:"Query" name:"BizType"`
}

// ModifyDeviceInfoResponse is the response struct for api ModifyDeviceInfo
type ModifyDeviceInfoResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DeviceId     string `json:"DeviceId" xml:"DeviceId"`
	UserDeviceId string `json:"UserDeviceId" xml:"UserDeviceId"`
	BizType      string `json:"BizType" xml:"BizType"`
	BeginDay     string `json:"BeginDay" xml:"BeginDay"`
	ExpiredDay   string `json:"ExpiredDay" xml:"ExpiredDay"`
}

// CreateModifyDeviceInfoRequest creates a request to invoke ModifyDeviceInfo API
func CreateModifyDeviceInfoRequest() (request *ModifyDeviceInfoRequest) {
	request = &ModifyDeviceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2018-09-16", "ModifyDeviceInfo", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDeviceInfoResponse creates a response to parse from ModifyDeviceInfo response
func CreateModifyDeviceInfoResponse() (response *ModifyDeviceInfoResponse) {
	response = &ModifyDeviceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
