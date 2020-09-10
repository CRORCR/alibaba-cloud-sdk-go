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

// DescribeSagRouteProtocolOspf invokes the smartag.DescribeSagRouteProtocolOspf API synchronously
// api document: https://help.aliyun.com/api/smartag/describesagrouteprotocolospf.html
func (client *Client) DescribeSagRouteProtocolOspf(request *DescribeSagRouteProtocolOspfRequest) (response *DescribeSagRouteProtocolOspfResponse, err error) {
	response = CreateDescribeSagRouteProtocolOspfResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagRouteProtocolOspfWithChan invokes the smartag.DescribeSagRouteProtocolOspf API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesagrouteprotocolospf.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagRouteProtocolOspfWithChan(request *DescribeSagRouteProtocolOspfRequest) (<-chan *DescribeSagRouteProtocolOspfResponse, <-chan error) {
	responseChan := make(chan *DescribeSagRouteProtocolOspfResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagRouteProtocolOspf(request)
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

// DescribeSagRouteProtocolOspfWithCallback invokes the smartag.DescribeSagRouteProtocolOspf API asynchronously
// api document: https://help.aliyun.com/api/smartag/describesagrouteprotocolospf.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSagRouteProtocolOspfWithCallback(request *DescribeSagRouteProtocolOspfRequest, callback func(response *DescribeSagRouteProtocolOspfResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagRouteProtocolOspfResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagRouteProtocolOspf(request)
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

// DescribeSagRouteProtocolOspfRequest is the request struct for api DescribeSagRouteProtocolOspf
type DescribeSagRouteProtocolOspfRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
}

// DescribeSagRouteProtocolOspfResponse is the response struct for api DescribeSagRouteProtocolOspf
type DescribeSagRouteProtocolOspfResponse struct {
	*responses.BaseResponse
	RequestId          string      `json:"RequestId" xml:"RequestId"`
	RouterId           string      `json:"RouterId" xml:"RouterId"`
	AreaId             string      `json:"AreaId" xml:"AreaId"`
	AreaType           string      `json:"AreaType" xml:"AreaType"`
	DeadTime           int         `json:"DeadTime" xml:"DeadTime"`
	HelloTime          int         `json:"HelloTime" xml:"HelloTime"`
	AuthenticationType string      `json:"AuthenticationType" xml:"AuthenticationType"`
	Md5KeyId           int         `json:"Md5KeyId" xml:"Md5KeyId"`
	Md5Key             string      `json:"Md5Key" xml:"Md5Key"`
	TaskStates         []TaskState `json:"TaskStates" xml:"TaskStates"`
}

// CreateDescribeSagRouteProtocolOspfRequest creates a request to invoke DescribeSagRouteProtocolOspf API
func CreateDescribeSagRouteProtocolOspfRequest() (request *DescribeSagRouteProtocolOspfRequest) {
	request = &DescribeSagRouteProtocolOspfRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagRouteProtocolOspf", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagRouteProtocolOspfResponse creates a response to parse from DescribeSagRouteProtocolOspf response
func CreateDescribeSagRouteProtocolOspfResponse() (response *DescribeSagRouteProtocolOspfResponse) {
	response = &DescribeSagRouteProtocolOspfResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
