package bssopenapi

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

// QueryInstanceBill invokes the bssopenapi.QueryInstanceBill API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryinstancebill.html
func (client *Client) QueryInstanceBill(request *QueryInstanceBillRequest) (response *QueryInstanceBillResponse, err error) {
	response = CreateQueryInstanceBillResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInstanceBillWithChan invokes the bssopenapi.QueryInstanceBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryinstancebill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryInstanceBillWithChan(request *QueryInstanceBillRequest) (<-chan *QueryInstanceBillResponse, <-chan error) {
	responseChan := make(chan *QueryInstanceBillResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInstanceBill(request)
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

// QueryInstanceBillWithCallback invokes the bssopenapi.QueryInstanceBill API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryinstancebill.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryInstanceBillWithCallback(request *QueryInstanceBillRequest, callback func(response *QueryInstanceBillResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInstanceBillResponse
		var err error
		defer close(result)
		response, err = client.QueryInstanceBill(request)
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

// QueryInstanceBillRequest is the request struct for api QueryInstanceBill
type QueryInstanceBillRequest struct {
	*requests.RpcRequest
	ProductCode      string           `position:"Query" name:"ProductCode"`
	IsHideZeroCharge requests.Boolean `position:"Query" name:"IsHideZeroCharge"`
	SubscriptionType string           `position:"Query" name:"SubscriptionType"`
	BillingCycle     string           `position:"Query" name:"BillingCycle"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	PageNum          requests.Integer `position:"Query" name:"PageNum"`
	BillOwnerId      requests.Integer `position:"Query" name:"BillOwnerId"`
	BillingDate      string           `position:"Query" name:"BillingDate"`
	ProductType      string           `position:"Query" name:"ProductType"`
	IsBillingItem    requests.Boolean `position:"Query" name:"IsBillingItem"`
	Granularity      string           `position:"Query" name:"Granularity"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
}

// QueryInstanceBillResponse is the response struct for api QueryInstanceBill
type QueryInstanceBillResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryInstanceBillRequest creates a request to invoke QueryInstanceBill API
func CreateQueryInstanceBillRequest() (request *QueryInstanceBillRequest) {
	request = &QueryInstanceBillRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryInstanceBill", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryInstanceBillResponse creates a response to parse from QueryInstanceBill response
func CreateQueryInstanceBillResponse() (response *QueryInstanceBillResponse) {
	response = &QueryInstanceBillResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
