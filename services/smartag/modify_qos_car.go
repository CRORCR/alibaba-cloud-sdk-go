package smartag

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

// ModifyQosCar invokes the smartag.ModifyQosCar API synchronously
// api document: https://help.aliyun.com/api/smartag/modifyqoscar.html
func (client *Client) ModifyQosCar(request *ModifyQosCarRequest) (response *ModifyQosCarResponse, err error) {
	response = CreateModifyQosCarResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyQosCarWithChan invokes the smartag.ModifyQosCar API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifyqoscar.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyQosCarWithChan(request *ModifyQosCarRequest) (<-chan *ModifyQosCarResponse, <-chan error) {
	responseChan := make(chan *ModifyQosCarResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyQosCar(request)
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

// ModifyQosCarWithCallback invokes the smartag.ModifyQosCar API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifyqoscar.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyQosCarWithCallback(request *ModifyQosCarRequest, callback func(response *ModifyQosCarResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyQosCarResponse
		var err error
		defer close(result)
		response, err = client.ModifyQosCar(request)
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

// ModifyQosCarRequest is the request struct for api ModifyQosCar
type ModifyQosCarRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MinBandwidthAbs      requests.Integer `position:"Query" name:"MinBandwidthAbs"`
	Description          string           `position:"Query" name:"Description"`
	PercentSourceType    string           `position:"Query" name:"PercentSourceType"`
	QosId                string           `position:"Query" name:"QosId"`
	MaxBandwidthAbs      requests.Integer `position:"Query" name:"MaxBandwidthAbs"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	MaxBandwidthPercent  requests.Integer `position:"Query" name:"MaxBandwidthPercent"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	QosCarId             string           `position:"Query" name:"QosCarId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	MinBandwidthPercent  requests.Integer `position:"Query" name:"MinBandwidthPercent"`
	LimitType            string           `position:"Query" name:"LimitType"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyQosCarResponse is the response struct for api ModifyQosCar
type ModifyQosCarResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyQosCarRequest creates a request to invoke ModifyQosCar API
func CreateModifyQosCarRequest() (request *ModifyQosCarRequest) {
	request = &ModifyQosCarRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifyQosCar", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyQosCarResponse creates a response to parse from ModifyQosCar response
func CreateModifyQosCarResponse() (response *ModifyQosCarResponse) {
	response = &ModifyQosCarResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
