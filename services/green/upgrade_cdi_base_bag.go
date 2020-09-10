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

// UpgradeCdiBaseBag invokes the green.UpgradeCdiBaseBag API synchronously
// api document: https://help.aliyun.com/api/green/upgradecdibasebag.html
func (client *Client) UpgradeCdiBaseBag(request *UpgradeCdiBaseBagRequest) (response *UpgradeCdiBaseBagResponse, err error) {
	response = CreateUpgradeCdiBaseBagResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeCdiBaseBagWithChan invokes the green.UpgradeCdiBaseBag API asynchronously
// api document: https://help.aliyun.com/api/green/upgradecdibasebag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeCdiBaseBagWithChan(request *UpgradeCdiBaseBagRequest) (<-chan *UpgradeCdiBaseBagResponse, <-chan error) {
	responseChan := make(chan *UpgradeCdiBaseBagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeCdiBaseBag(request)
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

// UpgradeCdiBaseBagWithCallback invokes the green.UpgradeCdiBaseBag API asynchronously
// api document: https://help.aliyun.com/api/green/upgradecdibasebag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpgradeCdiBaseBagWithCallback(request *UpgradeCdiBaseBagRequest, callback func(response *UpgradeCdiBaseBagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeCdiBaseBagResponse
		var err error
		defer close(result)
		response, err = client.UpgradeCdiBaseBag(request)
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

// UpgradeCdiBaseBagRequest is the request struct for api UpgradeCdiBaseBag
type UpgradeCdiBaseBagRequest struct {
	*requests.RpcRequest
	InstanceId    string           `position:"Query" name:"InstanceId"`
	ClientToken   string           `position:"Query" name:"ClientToken"`
	CommodityCode string           `position:"Query" name:"CommodityCode"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	FlowOutSpec   requests.Integer `position:"Query" name:"FlowOutSpec"`
	OrderType     string           `position:"Query" name:"OrderType"`
}

// UpgradeCdiBaseBagResponse is the response struct for api UpgradeCdiBaseBag
type UpgradeCdiBaseBagResponse struct {
	*responses.BaseResponse
	Code       string `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	OrderId    string `json:"OrderId" xml:"OrderId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateUpgradeCdiBaseBagRequest creates a request to invoke UpgradeCdiBaseBag API
func CreateUpgradeCdiBaseBagRequest() (request *UpgradeCdiBaseBagRequest) {
	request = &UpgradeCdiBaseBagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "UpgradeCdiBaseBag", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradeCdiBaseBagResponse creates a response to parse from UpgradeCdiBaseBag response
func CreateUpgradeCdiBaseBagResponse() (response *UpgradeCdiBaseBagResponse) {
	response = &UpgradeCdiBaseBagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
