package aegis

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

// DescribeAttackAnalysisData invokes the aegis.DescribeAttackAnalysisData API synchronously
// api document: https://help.aliyun.com/api/aegis/describeattackanalysisdata.html
func (client *Client) DescribeAttackAnalysisData(request *DescribeAttackAnalysisDataRequest) (response *DescribeAttackAnalysisDataResponse, err error) {
	response = CreateDescribeAttackAnalysisDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAttackAnalysisDataWithChan invokes the aegis.DescribeAttackAnalysisData API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeattackanalysisdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAttackAnalysisDataWithChan(request *DescribeAttackAnalysisDataRequest) (<-chan *DescribeAttackAnalysisDataResponse, <-chan error) {
	responseChan := make(chan *DescribeAttackAnalysisDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAttackAnalysisData(request)
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

// DescribeAttackAnalysisDataWithCallback invokes the aegis.DescribeAttackAnalysisData API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeattackanalysisdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeAttackAnalysisDataWithCallback(request *DescribeAttackAnalysisDataRequest, callback func(response *DescribeAttackAnalysisDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAttackAnalysisDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeAttackAnalysisData(request)
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

// DescribeAttackAnalysisDataRequest is the request struct for api DescribeAttackAnalysisData
type DescribeAttackAnalysisDataRequest struct {
	*requests.RpcRequest
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Data        string           `position:"Query" name:"Data"`
	Base64      string           `position:"Query" name:"Base64"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	EndTime     requests.Integer `position:"Query" name:"EndTime"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	StartTime   requests.Integer `position:"Query" name:"StartTime"`
	Lang        string           `position:"Query" name:"Lang"`
	Type        string           `position:"Query" name:"Type"`
}

// DescribeAttackAnalysisDataResponse is the response struct for api DescribeAttackAnalysisData
type DescribeAttackAnalysisDataResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Total     int    `json:"Total" xml:"Total"`
	Page      int    `json:"Page" xml:"Page"`
	PageSize  int    `json:"PageSize" xml:"PageSize"`
}

// CreateDescribeAttackAnalysisDataRequest creates a request to invoke DescribeAttackAnalysisData API
func CreateDescribeAttackAnalysisDataRequest() (request *DescribeAttackAnalysisDataRequest) {
	request = &DescribeAttackAnalysisDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeAttackAnalysisData", "vipaegis", "openAPI")
	return
}

// CreateDescribeAttackAnalysisDataResponse creates a response to parse from DescribeAttackAnalysisData response
func CreateDescribeAttackAnalysisDataResponse() (response *DescribeAttackAnalysisDataResponse) {
	response = &DescribeAttackAnalysisDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
