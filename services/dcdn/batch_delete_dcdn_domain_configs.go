package dcdn

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

// BatchDeleteDcdnDomainConfigs invokes the dcdn.BatchDeleteDcdnDomainConfigs API synchronously
// api document: https://help.aliyun.com/api/dcdn/batchdeletedcdndomainconfigs.html
func (client *Client) BatchDeleteDcdnDomainConfigs(request *BatchDeleteDcdnDomainConfigsRequest) (response *BatchDeleteDcdnDomainConfigsResponse, err error) {
	response = CreateBatchDeleteDcdnDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchDeleteDcdnDomainConfigsWithChan invokes the dcdn.BatchDeleteDcdnDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/dcdn/batchdeletedcdndomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteDcdnDomainConfigsWithChan(request *BatchDeleteDcdnDomainConfigsRequest) (<-chan *BatchDeleteDcdnDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *BatchDeleteDcdnDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchDeleteDcdnDomainConfigs(request)
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

// BatchDeleteDcdnDomainConfigsWithCallback invokes the dcdn.BatchDeleteDcdnDomainConfigs API asynchronously
// api document: https://help.aliyun.com/api/dcdn/batchdeletedcdndomainconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchDeleteDcdnDomainConfigsWithCallback(request *BatchDeleteDcdnDomainConfigsRequest, callback func(response *BatchDeleteDcdnDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchDeleteDcdnDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.BatchDeleteDcdnDomainConfigs(request)
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

// BatchDeleteDcdnDomainConfigsRequest is the request struct for api BatchDeleteDcdnDomainConfigs
type BatchDeleteDcdnDomainConfigsRequest struct {
	*requests.RpcRequest
	FunctionNames string           `position:"Query" name:"FunctionNames"`
	DomainNames   string           `position:"Query" name:"DomainNames"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// BatchDeleteDcdnDomainConfigsResponse is the response struct for api BatchDeleteDcdnDomainConfigs
type BatchDeleteDcdnDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBatchDeleteDcdnDomainConfigsRequest creates a request to invoke BatchDeleteDcdnDomainConfigs API
func CreateBatchDeleteDcdnDomainConfigsRequest() (request *BatchDeleteDcdnDomainConfigsRequest) {
	request = &BatchDeleteDcdnDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "BatchDeleteDcdnDomainConfigs", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchDeleteDcdnDomainConfigsResponse creates a response to parse from BatchDeleteDcdnDomainConfigs response
func CreateBatchDeleteDcdnDomainConfigsResponse() (response *BatchDeleteDcdnDomainConfigsResponse) {
	response = &BatchDeleteDcdnDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
