package ehpc

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

// CreateGWSInstance invokes the ehpc.CreateGWSInstance API synchronously
// api document: https://help.aliyun.com/api/ehpc/creategwsinstance.html
func (client *Client) CreateGWSInstance(request *CreateGWSInstanceRequest) (response *CreateGWSInstanceResponse, err error) {
	response = CreateCreateGWSInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGWSInstanceWithChan invokes the ehpc.CreateGWSInstance API asynchronously
// api document: https://help.aliyun.com/api/ehpc/creategwsinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGWSInstanceWithChan(request *CreateGWSInstanceRequest) (<-chan *CreateGWSInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateGWSInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGWSInstance(request)
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

// CreateGWSInstanceWithCallback invokes the ehpc.CreateGWSInstance API asynchronously
// api document: https://help.aliyun.com/api/ehpc/creategwsinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateGWSInstanceWithCallback(request *CreateGWSInstanceRequest, callback func(response *CreateGWSInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGWSInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateGWSInstance(request)
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

// CreateGWSInstanceRequest is the request struct for api CreateGWSInstance
type CreateGWSInstanceRequest struct {
	*requests.RpcRequest
	ImageId                 string           `position:"Query" name:"ImageId"`
	AllocatePublicAddress   requests.Boolean `position:"Query" name:"AllocatePublicAddress"`
	AppList                 string           `position:"Query" name:"AppList"`
	InternetMaxBandwidthOut requests.Integer `position:"Query" name:"InternetMaxBandwidthOut"`
	SystemDiskCategory      string           `position:"Query" name:"SystemDiskCategory"`
	SystemDiskSize          requests.Integer `position:"Query" name:"SystemDiskSize"`
	InstanceType            string           `position:"Query" name:"InstanceType"`
	InstanceChargeType      string           `position:"Query" name:"InstanceChargeType"`
	Period                  string           `position:"Query" name:"Period"`
	ClusterId               string           `position:"Query" name:"ClusterId"`
	WorkMode                string           `position:"Query" name:"WorkMode"`
	VSwitchId               string           `position:"Query" name:"VSwitchId"`
	PeriodUnit              string           `position:"Query" name:"PeriodUnit"`
	AutoRenew               requests.Boolean `position:"Query" name:"AutoRenew"`
	InternetChargeType      string           `position:"Query" name:"InternetChargeType"`
	Name                    string           `position:"Query" name:"Name"`
	InternetMaxBandwidthIn  requests.Integer `position:"Query" name:"InternetMaxBandwidthIn"`
}

// CreateGWSInstanceResponse is the response struct for api CreateGWSInstance
type CreateGWSInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
}

// CreateCreateGWSInstanceRequest creates a request to invoke CreateGWSInstance API
func CreateCreateGWSInstanceRequest() (request *CreateGWSInstanceRequest) {
	request = &CreateGWSInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "CreateGWSInstance", "", "")
	request.Method = requests.GET
	return
}

// CreateCreateGWSInstanceResponse creates a response to parse from CreateGWSInstance response
func CreateCreateGWSInstanceResponse() (response *CreateGWSInstanceResponse) {
	response = &CreateGWSInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
