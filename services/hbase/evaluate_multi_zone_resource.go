package hbase

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

// EvaluateMultiZoneResource invokes the hbase.EvaluateMultiZoneResource API synchronously
// api document: https://help.aliyun.com/api/hbase/evaluatemultizoneresource.html
func (client *Client) EvaluateMultiZoneResource(request *EvaluateMultiZoneResourceRequest) (response *EvaluateMultiZoneResourceResponse, err error) {
	response = CreateEvaluateMultiZoneResourceResponse()
	err = client.DoAction(request, response)
	return
}

// EvaluateMultiZoneResourceWithChan invokes the hbase.EvaluateMultiZoneResource API asynchronously
// api document: https://help.aliyun.com/api/hbase/evaluatemultizoneresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EvaluateMultiZoneResourceWithChan(request *EvaluateMultiZoneResourceRequest) (<-chan *EvaluateMultiZoneResourceResponse, <-chan error) {
	responseChan := make(chan *EvaluateMultiZoneResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EvaluateMultiZoneResource(request)
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

// EvaluateMultiZoneResourceWithCallback invokes the hbase.EvaluateMultiZoneResource API asynchronously
// api document: https://help.aliyun.com/api/hbase/evaluatemultizoneresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EvaluateMultiZoneResourceWithCallback(request *EvaluateMultiZoneResourceRequest, callback func(response *EvaluateMultiZoneResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EvaluateMultiZoneResourceResponse
		var err error
		defer close(result)
		response, err = client.EvaluateMultiZoneResource(request)
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

// EvaluateMultiZoneResourceRequest is the request struct for api EvaluateMultiZoneResource
type EvaluateMultiZoneResourceRequest struct {
	*requests.RpcRequest
	ArchVersion          string           `position:"Query" name:"ArchVersion"`
	ClusterName          string           `position:"Query" name:"ClusterName"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	LogDiskType          string           `position:"Query" name:"LogDiskType"`
	PrimaryVSwitchId     string           `position:"Query" name:"PrimaryVSwitchId"`
	LogInstanceType      string           `position:"Query" name:"LogInstanceType"`
	AutoRenewPeriod      requests.Integer `position:"Query" name:"AutoRenewPeriod"`
	Period               requests.Integer `position:"Query" name:"Period"`
	LogNodeCount         requests.Integer `position:"Query" name:"LogNodeCount"`
	SecurityIPList       string           `position:"Query" name:"SecurityIPList"`
	PeriodUnit           string           `position:"Query" name:"PeriodUnit"`
	CoreDiskType         string           `position:"Query" name:"CoreDiskType"`
	ArbiterZoneId        string           `position:"Query" name:"ArbiterZoneId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	MultiZoneCombination string           `position:"Query" name:"MultiZoneCombination"`
	PrimaryZoneId        string           `position:"Query" name:"PrimaryZoneId"`
	Engine               string           `position:"Query" name:"Engine"`
	StandbyVSwitchId     string           `position:"Query" name:"StandbyVSwitchId"`
	StandbyZoneId        string           `position:"Query" name:"StandbyZoneId"`
	MasterInstanceType   string           `position:"Query" name:"MasterInstanceType"`
	CoreNodeCount        requests.Integer `position:"Query" name:"CoreNodeCount"`
	LogDiskSize          requests.Integer `position:"Query" name:"LogDiskSize"`
	CoreInstanceType     string           `position:"Query" name:"CoreInstanceType"`
	CoreDiskSize         requests.Integer `position:"Query" name:"CoreDiskSize"`
	VpcId                string           `position:"Query" name:"VpcId"`
	PayType              string           `position:"Query" name:"PayType"`
	ArbiterVSwitchId     string           `position:"Query" name:"ArbiterVSwitchId"`
}

// EvaluateMultiZoneResourceResponse is the response struct for api EvaluateMultiZoneResource
type EvaluateMultiZoneResourceResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEvaluateMultiZoneResourceRequest creates a request to invoke EvaluateMultiZoneResource API
func CreateEvaluateMultiZoneResourceRequest() (request *EvaluateMultiZoneResourceRequest) {
	request = &EvaluateMultiZoneResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "EvaluateMultiZoneResource", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEvaluateMultiZoneResourceResponse creates a response to parse from EvaluateMultiZoneResource response
func CreateEvaluateMultiZoneResourceResponse() (response *EvaluateMultiZoneResourceResponse) {
	response = &EvaluateMultiZoneResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
