package vcs

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

// AddDevice invokes the vcs.AddDevice API synchronously
// api document: https://help.aliyun.com/api/vcs/adddevice.html
func (client *Client) AddDevice(request *AddDeviceRequest) (response *AddDeviceResponse, err error) {
	response = CreateAddDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// AddDeviceWithChan invokes the vcs.AddDevice API asynchronously
// api document: https://help.aliyun.com/api/vcs/adddevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddDeviceWithChan(request *AddDeviceRequest) (<-chan *AddDeviceResponse, <-chan error) {
	responseChan := make(chan *AddDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDevice(request)
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

// AddDeviceWithCallback invokes the vcs.AddDevice API asynchronously
// api document: https://help.aliyun.com/api/vcs/adddevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddDeviceWithCallback(request *AddDeviceRequest, callback func(response *AddDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDeviceResponse
		var err error
		defer close(result)
		response, err = client.AddDevice(request)
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

// AddDeviceRequest is the request struct for api AddDevice
type AddDeviceRequest struct {
	*requests.RpcRequest
	DeviceSite       string `position:"Body" name:"DeviceSite"`
	CorpId           string `position:"Body" name:"CorpId"`
	GbId             string `position:"Body" name:"GbId"`
	BitRate          string `position:"Body" name:"BitRate"`
	DeviceDirection  string `position:"Body" name:"DeviceDirection"`
	DeviceAddress    string `position:"Body" name:"DeviceAddress"`
	DeviceType       string `position:"Body" name:"DeviceType"`
	DeviceResolution string `position:"Body" name:"DeviceResolution"`
	Vendor           string `position:"Body" name:"Vendor"`
	DeviceName       string `position:"Body" name:"DeviceName"`
}

// AddDeviceResponse is the response struct for api AddDevice
type AddDeviceResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateAddDeviceRequest creates a request to invoke AddDevice API
func CreateAddDeviceRequest() (request *AddDeviceRequest) {
	request = &AddDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "AddDevice", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddDeviceResponse creates a response to parse from AddDevice response
func CreateAddDeviceResponse() (response *AddDeviceResponse) {
	response = &AddDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
