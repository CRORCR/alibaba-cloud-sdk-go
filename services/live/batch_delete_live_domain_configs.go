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

// BatchDeleteLiveDomainConfigs invokes the live.BatchDeleteLiveDomainConfigs API synchronously
// api document: https://help.aliyun.com/api/live/batchdeletelivedomainconfigs.html
func (client *Client) BatchDeleteLiveDomainConfigs(request *BatchDeleteLiveDomainConfigsRequest) (response *BatchDeleteLiveDomainConfigsResponse, err error) {
	response = CreateBatchDeleteLiveDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchDeleteLiveDomainConfigsWithChan invokes the live.BatchDeleteLiveDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/live/batchdeletelivedomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteLiveDomainConfigsWithChan(request *BatchDeleteLiveDomainConfigsRequest) (<-chan *BatchDeleteLiveDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *BatchDeleteLiveDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchDeleteLiveDomainConfigs(request)
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

// BatchDeleteLiveDomainConfigsWithCallback invokes the live.BatchDeleteLiveDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/live/batchdeletelivedomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteLiveDomainConfigsWithCallback(request *BatchDeleteLiveDomainConfigsRequest, callback func(response *BatchDeleteLiveDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchDeleteLiveDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.BatchDeleteLiveDomainConfigs(request)
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

// BatchDeleteLiveDomainConfigsRequest is the request struct for api BatchDeleteLiveDomainConfigs
type BatchDeleteLiveDomainConfigsRequest struct {
	*requests.RpcRequest
	FunctionNames string           `position:"Query" name:"FunctionNames"`
	DomainNames   string           `position:"Query" name:"DomainNames"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// BatchDeleteLiveDomainConfigsResponse is the response struct for api BatchDeleteLiveDomainConfigs
type BatchDeleteLiveDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchDeleteLiveDomainConfigsRequest creates a request to invoke BatchDeleteLiveDomainConfigs API
func CreateBatchDeleteLiveDomainConfigsRequest() (request *BatchDeleteLiveDomainConfigsRequest) {
	request = &BatchDeleteLiveDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "BatchDeleteLiveDomainConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchDeleteLiveDomainConfigsResponse creates a response to parse from BatchDeleteLiveDomainConfigs response
func CreateBatchDeleteLiveDomainConfigsResponse() (response *BatchDeleteLiveDomainConfigsResponse) {
	response = &BatchDeleteLiveDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
