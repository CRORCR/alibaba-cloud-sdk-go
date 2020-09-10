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

// DescribeLiveStreamSnapshotInfo invokes the cdn.DescribeLiveStreamSnapshotInfo API synchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamsnapshotinfo.html
func (client *Client) DescribeLiveStreamSnapshotInfo(request *DescribeLiveStreamSnapshotInfoRequest) (response *DescribeLiveStreamSnapshotInfoResponse, err error) {
	response = CreateDescribeLiveStreamSnapshotInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamSnapshotInfoWithChan invokes the cdn.DescribeLiveStreamSnapshotInfo API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamsnapshotinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamSnapshotInfoWithChan(request *DescribeLiveStreamSnapshotInfoRequest) (<-chan *DescribeLiveStreamSnapshotInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamSnapshotInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamSnapshotInfo(request)
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

// DescribeLiveStreamSnapshotInfoWithCallback invokes the cdn.DescribeLiveStreamSnapshotInfo API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamsnapshotinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamSnapshotInfoWithCallback(request *DescribeLiveStreamSnapshotInfoRequest, callback func(response *DescribeLiveStreamSnapshotInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamSnapshotInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamSnapshotInfo(request)
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

// DescribeLiveStreamSnapshotInfoRequest is the request struct for api DescribeLiveStreamSnapshotInfo
type DescribeLiveStreamSnapshotInfoRequest struct {
	*requests.RpcRequest
	StartTime     string           `position:"Query" name:"StartTime"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	Limit         requests.Integer `position:"Query" name:"Limit"`
	StreamName    string           `position:"Query" name:"StreamName"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamSnapshotInfoResponse is the response struct for api DescribeLiveStreamSnapshotInfo
type DescribeLiveStreamSnapshotInfoResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	NextStartTime              string                     `json:"NextStartTime" xml:"NextStartTime"`
	LiveStreamSnapshotInfoList LiveStreamSnapshotInfoList `json:"LiveStreamSnapshotInfoList" xml:"LiveStreamSnapshotInfoList"`
}

// CreateDescribeLiveStreamSnapshotInfoRequest creates a request to invoke DescribeLiveStreamSnapshotInfo API
func CreateDescribeLiveStreamSnapshotInfoRequest() (request *DescribeLiveStreamSnapshotInfoRequest) {
	request = &DescribeLiveStreamSnapshotInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamSnapshotInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveStreamSnapshotInfoResponse creates a response to parse from DescribeLiveStreamSnapshotInfo response
func CreateDescribeLiveStreamSnapshotInfoResponse() (response *DescribeLiveStreamSnapshotInfoResponse) {
	response = &DescribeLiveStreamSnapshotInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
