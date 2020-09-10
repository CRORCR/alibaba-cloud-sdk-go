package ddospro

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

// DescribeCcMode invokes the ddospro.DescribeCcMode API synchronously
// api document: https://help.aliyun.com/api/ddospro/describeccmode.html
func (client *Client) DescribeCcMode(request *DescribeCcModeRequest) (response *DescribeCcModeResponse, err error) {
	response = CreateDescribeCcModeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCcModeWithChan invokes the ddospro.DescribeCcMode API asynchronously
// api document: https://help.aliyun.com/api/ddospro/describeccmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCcModeWithChan(request *DescribeCcModeRequest) (<-chan *DescribeCcModeResponse, <-chan error) {
	responseChan := make(chan *DescribeCcModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCcMode(request)
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

// DescribeCcModeWithCallback invokes the ddospro.DescribeCcMode API asynchronously
// api document: https://help.aliyun.com/api/ddospro/describeccmode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCcModeWithCallback(request *DescribeCcModeRequest, callback func(response *DescribeCcModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCcModeResponse
		var err error
		defer close(result)
		response, err = client.DescribeCcMode(request)
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

// DescribeCcModeRequest is the request struct for api DescribeCcMode
type DescribeCcModeRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Vips            *[]string        `position:"Query" name:"Vips"  type:"Repeated"`
}

// DescribeCcModeResponse is the response struct for api DescribeCcMode
type DescribeCcModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      []Data `json:"Data" xml:"Data"`
}

// CreateDescribeCcModeRequest creates a request to invoke DescribeCcMode API
func CreateDescribeCcModeRequest() (request *DescribeCcModeRequest) {
	request = &DescribeCcModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "DescribeCcMode", "", "")
	return
}

// CreateDescribeCcModeResponse creates a response to parse from DescribeCcMode response
func CreateDescribeCcModeResponse() (response *DescribeCcModeResponse) {
	response = &DescribeCcModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
