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

// BatchUpdateTaobaoUserMap invokes the cloudcallcenter.BatchUpdateTaobaoUserMap API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/batchupdatetaobaousermap.html
func (client *Client) BatchUpdateTaobaoUserMap(request *BatchUpdateTaobaoUserMapRequest) (response *BatchUpdateTaobaoUserMapResponse, err error) {
	response = CreateBatchUpdateTaobaoUserMapResponse()
	err = client.DoAction(request, response)
	return
}

// BatchUpdateTaobaoUserMapWithChan invokes the cloudcallcenter.BatchUpdateTaobaoUserMap API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/batchupdatetaobaousermap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchUpdateTaobaoUserMapWithChan(request *BatchUpdateTaobaoUserMapRequest) (<-chan *BatchUpdateTaobaoUserMapResponse, <-chan error) {
	responseChan := make(chan *BatchUpdateTaobaoUserMapResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchUpdateTaobaoUserMap(request)
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

// BatchUpdateTaobaoUserMapWithCallback invokes the cloudcallcenter.BatchUpdateTaobaoUserMap API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/batchupdatetaobaousermap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchUpdateTaobaoUserMapWithCallback(request *BatchUpdateTaobaoUserMapRequest, callback func(response *BatchUpdateTaobaoUserMapResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchUpdateTaobaoUserMapResponse
		var err error
		defer close(result)
		response, err = client.BatchUpdateTaobaoUserMap(request)
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

// BatchUpdateTaobaoUserMapRequest is the request struct for api BatchUpdateTaobaoUserMap
type BatchUpdateTaobaoUserMapRequest struct {
	*requests.RpcRequest
	TaobaoAliyunAccount *[]BatchUpdateTaobaoUserMapTaobaoAliyunAccount `position:"Query" name:"TaobaoAliyunAccount"  type:"Repeated"`
}

// BatchUpdateTaobaoUserMapTaobaoAliyunAccount is a repeated param struct in BatchUpdateTaobaoUserMapRequest
type BatchUpdateTaobaoUserMapTaobaoAliyunAccount struct {
	AliyunPk   string `name:"aliyunPk"`
	TaobaoNick string `name:"TaobaoNick"`
	TaobaoUid  string `name:"TaobaoUid"`
	Type       string `name:"Type"`
	Status     string `name:"Status"`
}

// BatchUpdateTaobaoUserMapResponse is the response struct for api BatchUpdateTaobaoUserMap
type BatchUpdateTaobaoUserMapResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Count          int    `json:"Count" xml:"Count"`
}

// CreateBatchUpdateTaobaoUserMapRequest creates a request to invoke BatchUpdateTaobaoUserMap API
func CreateBatchUpdateTaobaoUserMapRequest() (request *BatchUpdateTaobaoUserMapRequest) {
	request = &BatchUpdateTaobaoUserMapRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "BatchUpdateTaobaoUserMap", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchUpdateTaobaoUserMapResponse creates a response to parse from BatchUpdateTaobaoUserMap response
func CreateBatchUpdateTaobaoUserMapResponse() (response *BatchUpdateTaobaoUserMapResponse) {
	response = &BatchUpdateTaobaoUserMapResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
