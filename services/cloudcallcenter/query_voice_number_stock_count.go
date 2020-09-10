package cloudcallcenter

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

// QueryVoiceNumberStockCount invokes the cloudcallcenter.QueryVoiceNumberStockCount API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryvoicenumberstockcount.html
func (client *Client) QueryVoiceNumberStockCount(request *QueryVoiceNumberStockCountRequest) (response *QueryVoiceNumberStockCountResponse, err error) {
	response = CreateQueryVoiceNumberStockCountResponse()
	err = client.DoAction(request, response)
	return
}

// QueryVoiceNumberStockCountWithChan invokes the cloudcallcenter.QueryVoiceNumberStockCount API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryvoicenumberstockcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryVoiceNumberStockCountWithChan(request *QueryVoiceNumberStockCountRequest) (<-chan *QueryVoiceNumberStockCountResponse, <-chan error) {
	responseChan := make(chan *QueryVoiceNumberStockCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryVoiceNumberStockCount(request)
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

// QueryVoiceNumberStockCountWithCallback invokes the cloudcallcenter.QueryVoiceNumberStockCount API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryvoicenumberstockcount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryVoiceNumberStockCountWithCallback(request *QueryVoiceNumberStockCountRequest, callback func(response *QueryVoiceNumberStockCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryVoiceNumberStockCountResponse
		var err error
		defer close(result)
		response, err = client.QueryVoiceNumberStockCount(request)
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

// QueryVoiceNumberStockCountRequest is the request struct for api QueryVoiceNumberStockCount
type QueryVoiceNumberStockCountRequest struct {
	*requests.RpcRequest
	SpecId   *[]string `position:"Query" name:"SpecId"  type:"Repeated"`
	Province string    `position:"Query" name:"Province"`
	City     string    `position:"Query" name:"City"`
}

// QueryVoiceNumberStockCountResponse is the response struct for api QueryVoiceNumberStockCount
type QueryVoiceNumberStockCountResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Count          int64  `json:"Count" xml:"Count"`
	CuteCount      int64  `json:"CuteCount" xml:"CuteCount"`
	CommonCount    int64  `json:"CommonCount" xml:"CommonCount"`
}

// CreateQueryVoiceNumberStockCountRequest creates a request to invoke QueryVoiceNumberStockCount API
func CreateQueryVoiceNumberStockCountRequest() (request *QueryVoiceNumberStockCountRequest) {
	request = &QueryVoiceNumberStockCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "QueryVoiceNumberStockCount", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryVoiceNumberStockCountResponse creates a response to parse from QueryVoiceNumberStockCount response
func CreateQueryVoiceNumberStockCountResponse() (response *QueryVoiceNumberStockCountResponse) {
	response = &QueryVoiceNumberStockCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
