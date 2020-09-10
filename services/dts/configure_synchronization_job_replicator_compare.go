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

package dts

import (
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/CRORCR/alibaba-cloud-sdk-go/sdk/responses"
)

// ConfigureSynchronizationJobReplicatorCompare invokes the dts.ConfigureSynchronizationJobReplicatorCompare API synchronously
// api document: https://help.aliyun.com/api/dts/configuresynchronizationjobreplicatorcompare.html
func (client *Client) ConfigureSynchronizationJobReplicatorCompare(request *ConfigureSynchronizationJobReplicatorCompareRequest) (response *ConfigureSynchronizationJobReplicatorCompareResponse, err error) {
	response = CreateConfigureSynchronizationJobReplicatorCompareResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureSynchronizationJobReplicatorCompareWithChan invokes the dts.ConfigureSynchronizationJobReplicatorCompare API asynchronously
// api document: https://help.aliyun.com/api/dts/configuresynchronizationjobreplicatorcompare.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigureSynchronizationJobReplicatorCompareWithChan(request *ConfigureSynchronizationJobReplicatorCompareRequest) (<-chan *ConfigureSynchronizationJobReplicatorCompareResponse, <-chan error) {
	responseChan := make(chan *ConfigureSynchronizationJobReplicatorCompareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureSynchronizationJobReplicatorCompare(request)
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

// ConfigureSynchronizationJobReplicatorCompareWithCallback invokes the dts.ConfigureSynchronizationJobReplicatorCompare API asynchronously
// api document: https://help.aliyun.com/api/dts/configuresynchronizationjobreplicatorcompare.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfigureSynchronizationJobReplicatorCompareWithCallback(request *ConfigureSynchronizationJobReplicatorCompareRequest, callback func(response *ConfigureSynchronizationJobReplicatorCompareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureSynchronizationJobReplicatorCompareResponse
		var err error
		defer close(result)
		response, err = client.ConfigureSynchronizationJobReplicatorCompare(request)
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

// ConfigureSynchronizationJobReplicatorCompareRequest is the request struct for api ConfigureSynchronizationJobReplicatorCompare
type ConfigureSynchronizationJobReplicatorCompareRequest struct {
	*requests.RpcRequest
	SynchronizationJobId                   string           `position:"Query" name:"SynchronizationJobId"`
	SynchronizationDirection               string           `position:"Query" name:"SynchronizationDirection"`
	SynchronizationReplicatorCompareEnable requests.Boolean `position:"Query" name:"SynchronizationReplicatorCompareEnable"`
	ClientToken                            string           `position:"Query" name:"ClientToken"`
	OwnerId                                string           `position:"Query" name:"OwnerId"`
}

// ConfigureSynchronizationJobReplicatorCompareResponse is the response struct for api ConfigureSynchronizationJobReplicatorCompare
type ConfigureSynchronizationJobReplicatorCompareResponse struct {
	*responses.BaseResponse
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigureSynchronizationJobReplicatorCompareRequest creates a request to invoke ConfigureSynchronizationJobReplicatorCompare API
func CreateConfigureSynchronizationJobReplicatorCompareRequest() (request *ConfigureSynchronizationJobReplicatorCompareRequest) {
	request = &ConfigureSynchronizationJobReplicatorCompareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "ConfigureSynchronizationJobReplicatorCompare", "dts", "openAPI")
	return
}

// CreateConfigureSynchronizationJobReplicatorCompareResponse creates a response to parse from ConfigureSynchronizationJobReplicatorCompare response
func CreateConfigureSynchronizationJobReplicatorCompareResponse() (response *ConfigureSynchronizationJobReplicatorCompareResponse) {
	response = &ConfigureSynchronizationJobReplicatorCompareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
