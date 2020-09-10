package cloudwf

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

// GetDailyStatistic invokes the cloudwf.GetDailyStatistic API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getdailystatistic.html
func (client *Client) GetDailyStatistic(request *GetDailyStatisticRequest) (response *GetDailyStatisticResponse, err error) {
	response = CreateGetDailyStatisticResponse()
	err = client.DoAction(request, response)
	return
}

// GetDailyStatisticWithChan invokes the cloudwf.GetDailyStatistic API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getdailystatistic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDailyStatisticWithChan(request *GetDailyStatisticRequest) (<-chan *GetDailyStatisticResponse, <-chan error) {
	responseChan := make(chan *GetDailyStatisticResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDailyStatistic(request)
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

// GetDailyStatisticWithCallback invokes the cloudwf.GetDailyStatistic API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getdailystatistic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDailyStatisticWithCallback(request *GetDailyStatisticRequest, callback func(response *GetDailyStatisticResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDailyStatisticResponse
		var err error
		defer close(result)
		response, err = client.GetDailyStatistic(request)
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

// GetDailyStatisticRequest is the request struct for api GetDailyStatistic
type GetDailyStatisticRequest struct {
	*requests.RpcRequest
	ApgroupId requests.Integer `position:"Query" name:"ApgroupId"`
}

// GetDailyStatisticResponse is the response struct for api GetDailyStatistic
type GetDailyStatisticResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetDailyStatisticRequest creates a request to invoke GetDailyStatistic API
func CreateGetDailyStatisticRequest() (request *GetDailyStatisticRequest) {
	request = &GetDailyStatisticRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetDailyStatistic", "cloudwf", "openAPI")
	return
}

// CreateGetDailyStatisticResponse creates a response to parse from GetDailyStatistic response
func CreateGetDailyStatisticResponse() (response *GetDailyStatisticResponse) {
	response = &GetDailyStatisticResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
