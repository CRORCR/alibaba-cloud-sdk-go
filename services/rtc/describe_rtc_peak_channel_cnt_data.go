package rtc

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

// DescribeRtcPeakChannelCntData invokes the rtc.DescribeRtcPeakChannelCntData API synchronously
// api document: https://help.aliyun.com/api/rtc/describertcpeakchannelcntdata.html
func (client *Client) DescribeRtcPeakChannelCntData(request *DescribeRtcPeakChannelCntDataRequest) (response *DescribeRtcPeakChannelCntDataResponse, err error) {
	response = CreateDescribeRtcPeakChannelCntDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRtcPeakChannelCntDataWithChan invokes the rtc.DescribeRtcPeakChannelCntData API asynchronously
// api document: https://help.aliyun.com/api/rtc/describertcpeakchannelcntdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRtcPeakChannelCntDataWithChan(request *DescribeRtcPeakChannelCntDataRequest) (<-chan *DescribeRtcPeakChannelCntDataResponse, <-chan error) {
	responseChan := make(chan *DescribeRtcPeakChannelCntDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRtcPeakChannelCntData(request)
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

// DescribeRtcPeakChannelCntDataWithCallback invokes the rtc.DescribeRtcPeakChannelCntData API asynchronously
// api document: https://help.aliyun.com/api/rtc/describertcpeakchannelcntdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRtcPeakChannelCntDataWithCallback(request *DescribeRtcPeakChannelCntDataRequest, callback func(response *DescribeRtcPeakChannelCntDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRtcPeakChannelCntDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeRtcPeakChannelCntData(request)
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

// DescribeRtcPeakChannelCntDataRequest is the request struct for api DescribeRtcPeakChannelCntData
type DescribeRtcPeakChannelCntDataRequest struct {
	*requests.RpcRequest
	StartTime   string           `position:"Query" name:"StartTime"`
	ServiceArea string           `position:"Query" name:"ServiceArea"`
	EndTime     string           `position:"Query" name:"EndTime"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	AppId       string           `position:"Query" name:"AppId"`
	Interval    string           `position:"Query" name:"Interval"`
}

// DescribeRtcPeakChannelCntDataResponse is the response struct for api DescribeRtcPeakChannelCntData
type DescribeRtcPeakChannelCntDataResponse struct {
	*responses.BaseResponse
	RequestId                     string                        `json:"RequestId" xml:"RequestId"`
	PeakChannelCntDataPerInterval PeakChannelCntDataPerInterval `json:"PeakChannelCntDataPerInterval" xml:"PeakChannelCntDataPerInterval"`
}

// CreateDescribeRtcPeakChannelCntDataRequest creates a request to invoke DescribeRtcPeakChannelCntData API
func CreateDescribeRtcPeakChannelCntDataRequest() (request *DescribeRtcPeakChannelCntDataRequest) {
	request = &DescribeRtcPeakChannelCntDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeRtcPeakChannelCntData", "rtc", "openAPI")
	return
}

// CreateDescribeRtcPeakChannelCntDataResponse creates a response to parse from DescribeRtcPeakChannelCntData response
func CreateDescribeRtcPeakChannelCntDataResponse() (response *DescribeRtcPeakChannelCntDataResponse) {
	response = &DescribeRtcPeakChannelCntDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
