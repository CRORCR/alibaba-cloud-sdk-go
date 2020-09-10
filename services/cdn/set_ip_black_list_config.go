package cdn

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

// SetIpBlackListConfig invokes the cdn.SetIpBlackListConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/setipblacklistconfig.html
func (client *Client) SetIpBlackListConfig(request *SetIpBlackListConfigRequest) (response *SetIpBlackListConfigResponse, err error) {
	response = CreateSetIpBlackListConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetIpBlackListConfigWithChan invokes the cdn.SetIpBlackListConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setipblacklistconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetIpBlackListConfigWithChan(request *SetIpBlackListConfigRequest) (<-chan *SetIpBlackListConfigResponse, <-chan error) {
	responseChan := make(chan *SetIpBlackListConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetIpBlackListConfig(request)
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

// SetIpBlackListConfigWithCallback invokes the cdn.SetIpBlackListConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/setipblacklistconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetIpBlackListConfigWithCallback(request *SetIpBlackListConfigRequest, callback func(response *SetIpBlackListConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetIpBlackListConfigResponse
		var err error
		defer close(result)
		response, err = client.SetIpBlackListConfig(request)
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

// SetIpBlackListConfigRequest is the request struct for api SetIpBlackListConfig
type SetIpBlackListConfigRequest struct {
	*requests.RpcRequest
	BlockIps   string           `position:"Query" name:"BlockIps"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	ConfigId   requests.Integer `position:"Query" name:"ConfigId"`
}

// SetIpBlackListConfigResponse is the response struct for api SetIpBlackListConfig
type SetIpBlackListConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetIpBlackListConfigRequest creates a request to invoke SetIpBlackListConfig API
func CreateSetIpBlackListConfigRequest() (request *SetIpBlackListConfigRequest) {
	request = &SetIpBlackListConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "SetIpBlackListConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateSetIpBlackListConfigResponse creates a response to parse from SetIpBlackListConfig response
func CreateSetIpBlackListConfigResponse() (response *SetIpBlackListConfigResponse) {
	response = &SetIpBlackListConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
