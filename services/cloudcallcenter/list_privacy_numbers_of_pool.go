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

// ListPrivacyNumbersOfPool invokes the cloudcallcenter.ListPrivacyNumbersOfPool API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listprivacynumbersofpool.html
func (client *Client) ListPrivacyNumbersOfPool(request *ListPrivacyNumbersOfPoolRequest) (response *ListPrivacyNumbersOfPoolResponse, err error) {
	response = CreateListPrivacyNumbersOfPoolResponse()
	err = client.DoAction(request, response)
	return
}

// ListPrivacyNumbersOfPoolWithChan invokes the cloudcallcenter.ListPrivacyNumbersOfPool API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listprivacynumbersofpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPrivacyNumbersOfPoolWithChan(request *ListPrivacyNumbersOfPoolRequest) (<-chan *ListPrivacyNumbersOfPoolResponse, <-chan error) {
	responseChan := make(chan *ListPrivacyNumbersOfPoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPrivacyNumbersOfPool(request)
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

// ListPrivacyNumbersOfPoolWithCallback invokes the cloudcallcenter.ListPrivacyNumbersOfPool API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listprivacynumbersofpool.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListPrivacyNumbersOfPoolWithCallback(request *ListPrivacyNumbersOfPoolRequest, callback func(response *ListPrivacyNumbersOfPoolResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPrivacyNumbersOfPoolResponse
		var err error
		defer close(result)
		response, err = client.ListPrivacyNumbersOfPool(request)
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

// ListPrivacyNumbersOfPoolRequest is the request struct for api ListPrivacyNumbersOfPool
type ListPrivacyNumbersOfPoolRequest struct {
	*requests.RpcRequest
	ProviderId string `position:"Query" name:"ProviderId"`
	BizId      string `position:"Query" name:"BizId"`
	PoolId     string `position:"Query" name:"PoolId"`
	Type       string `position:"Query" name:"Type"`
}

// ListPrivacyNumbersOfPoolResponse is the response struct for api ListPrivacyNumbersOfPool
type ListPrivacyNumbersOfPoolResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	Code           string         `json:"Code" xml:"Code"`
	Message        string         `json:"Message" xml:"Message"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	VirtualPool    VirtualPool    `json:"VirtualPool" xml:"VirtualPool"`
	PrivacyNumbers PrivacyNumbers `json:"PrivacyNumbers" xml:"PrivacyNumbers"`
}

// CreateListPrivacyNumbersOfPoolRequest creates a request to invoke ListPrivacyNumbersOfPool API
func CreateListPrivacyNumbersOfPoolRequest() (request *ListPrivacyNumbersOfPoolRequest) {
	request = &ListPrivacyNumbersOfPoolRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListPrivacyNumbersOfPool", "", "")
	request.Method = requests.POST
	return
}

// CreateListPrivacyNumbersOfPoolResponse creates a response to parse from ListPrivacyNumbersOfPool response
func CreateListPrivacyNumbersOfPoolResponse() (response *ListPrivacyNumbersOfPoolResponse) {
	response = &ListPrivacyNumbersOfPoolResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
