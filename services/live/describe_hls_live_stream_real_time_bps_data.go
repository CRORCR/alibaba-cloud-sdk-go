package live

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

// DescribeHlsLiveStreamRealTimeBpsData invokes the live.DescribeHlsLiveStreamRealTimeBpsData API synchronously
// api document: https://help.aliyun.com/api/live/describehlslivestreamrealtimebpsdata.html
func (client *Client) DescribeHlsLiveStreamRealTimeBpsData(request *DescribeHlsLiveStreamRealTimeBpsDataRequest) (response *DescribeHlsLiveStreamRealTimeBpsDataResponse, err error) {
	response = CreateDescribeHlsLiveStreamRealTimeBpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHlsLiveStreamRealTimeBpsDataWithChan invokes the live.DescribeHlsLiveStreamRealTimeBpsData API asynchronously
// api document: https://help.aliyun.com/api/live/describehlslivestreamrealtimebpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHlsLiveStreamRealTimeBpsDataWithChan(request *DescribeHlsLiveStreamRealTimeBpsDataRequest) (<-chan *DescribeHlsLiveStreamRealTimeBpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeHlsLiveStreamRealTimeBpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHlsLiveStreamRealTimeBpsData(request)
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

// DescribeHlsLiveStreamRealTimeBpsDataWithCallback invokes the live.DescribeHlsLiveStreamRealTimeBpsData API asynchronously
// api document: https://help.aliyun.com/api/live/describehlslivestreamrealtimebpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHlsLiveStreamRealTimeBpsDataWithCallback(request *DescribeHlsLiveStreamRealTimeBpsDataRequest, callback func(response *DescribeHlsLiveStreamRealTimeBpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHlsLiveStreamRealTimeBpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeHlsLiveStreamRealTimeBpsData(request)
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

// DescribeHlsLiveStreamRealTimeBpsDataRequest is the request struct for api DescribeHlsLiveStreamRealTimeBpsData
type DescribeHlsLiveStreamRealTimeBpsDataRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Time       string           `position:"Query" name:"Time"`
}

// DescribeHlsLiveStreamRealTimeBpsDataResponse is the response struct for api DescribeHlsLiveStreamRealTimeBpsData
type DescribeHlsLiveStreamRealTimeBpsDataResponse struct {
	*responses.BaseResponse
	Time      string               `json:"Time" xml:"Time"`
	RequestId string               `json:"RequestId" xml:"RequestId"`
	UsageData []UsageDataPerDomain `json:"UsageData" xml:"UsageData"`
}

// CreateDescribeHlsLiveStreamRealTimeBpsDataRequest creates a request to invoke DescribeHlsLiveStreamRealTimeBpsData API
func CreateDescribeHlsLiveStreamRealTimeBpsDataRequest() (request *DescribeHlsLiveStreamRealTimeBpsDataRequest) {
	request = &DescribeHlsLiveStreamRealTimeBpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeHlsLiveStreamRealTimeBpsData", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeHlsLiveStreamRealTimeBpsDataResponse creates a response to parse from DescribeHlsLiveStreamRealTimeBpsData response
func CreateDescribeHlsLiveStreamRealTimeBpsDataResponse() (response *DescribeHlsLiveStreamRealTimeBpsDataResponse) {
	response = &DescribeHlsLiveStreamRealTimeBpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
