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

// QueryMonthlyInstanceConsumption invokes the bssopenapi.QueryMonthlyInstanceConsumption API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/querymonthlyinstanceconsumption.html
func (client *Client) QueryMonthlyInstanceConsumption(request *QueryMonthlyInstanceConsumptionRequest) (response *QueryMonthlyInstanceConsumptionResponse, err error) {
	response = CreateQueryMonthlyInstanceConsumptionResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMonthlyInstanceConsumptionWithChan invokes the bssopenapi.QueryMonthlyInstanceConsumption API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querymonthlyinstanceconsumption.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMonthlyInstanceConsumptionWithChan(request *QueryMonthlyInstanceConsumptionRequest) (<-chan *QueryMonthlyInstanceConsumptionResponse, <-chan error) {
	responseChan := make(chan *QueryMonthlyInstanceConsumptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMonthlyInstanceConsumption(request)
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

// QueryMonthlyInstanceConsumptionWithCallback invokes the bssopenapi.QueryMonthlyInstanceConsumption API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querymonthlyinstanceconsumption.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMonthlyInstanceConsumptionWithCallback(request *QueryMonthlyInstanceConsumptionRequest, callback func(response *QueryMonthlyInstanceConsumptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMonthlyInstanceConsumptionResponse
		var err error
		defer close(result)
		response, err = client.QueryMonthlyInstanceConsumption(request)
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

// QueryMonthlyInstanceConsumptionRequest is the request struct for api QueryMonthlyInstanceConsumption
type QueryMonthlyInstanceConsumptionRequest struct {
	*requests.RpcRequest
	ProductCode      string           `position:"Query" name:"ProductCode"`
	SubscriptionType string           `position:"Query" name:"SubscriptionType"`
	BillingCycle     string           `position:"Query" name:"BillingCycle"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	PageNum          requests.Integer `position:"Query" name:"PageNum"`
	ProductType      string           `position:"Query" name:"ProductType"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
}

// QueryMonthlyInstanceConsumptionResponse is the response struct for api QueryMonthlyInstanceConsumption
type QueryMonthlyInstanceConsumptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryMonthlyInstanceConsumptionRequest creates a request to invoke QueryMonthlyInstanceConsumption API
func CreateQueryMonthlyInstanceConsumptionRequest() (request *QueryMonthlyInstanceConsumptionRequest) {
	request = &QueryMonthlyInstanceConsumptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryMonthlyInstanceConsumption", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryMonthlyInstanceConsumptionResponse creates a response to parse from QueryMonthlyInstanceConsumption response
func CreateQueryMonthlyInstanceConsumptionResponse() (response *QueryMonthlyInstanceConsumptionResponse) {
	response = &QueryMonthlyInstanceConsumptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
