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

// AddCustomerNumber invokes the cloudcallcenter.AddCustomerNumber API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/addcustomernumber.html
func (client *Client) AddCustomerNumber(request *AddCustomerNumberRequest) (response *AddCustomerNumberResponse, err error) {
	response = CreateAddCustomerNumberResponse()
	err = client.DoAction(request, response)
	return
}

// AddCustomerNumberWithChan invokes the cloudcallcenter.AddCustomerNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/addcustomernumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCustomerNumberWithChan(request *AddCustomerNumberRequest) (<-chan *AddCustomerNumberResponse, <-chan error) {
	responseChan := make(chan *AddCustomerNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCustomerNumber(request)
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

// AddCustomerNumberWithCallback invokes the cloudcallcenter.AddCustomerNumber API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/addcustomernumber.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCustomerNumberWithCallback(request *AddCustomerNumberRequest, callback func(response *AddCustomerNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCustomerNumberResponse
		var err error
		defer close(result)
		response, err = client.AddCustomerNumber(request)
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

// AddCustomerNumberRequest is the request struct for api AddCustomerNumber
type AddCustomerNumberRequest struct {
	*requests.RpcRequest
	Provider    string `position:"Query" name:"Provider"`
	PhoneNumber string `position:"Query" name:"PhoneNumber"`
}

// AddCustomerNumberResponse is the response struct for api AddCustomerNumber
type AddCustomerNumberResponse struct {
	*responses.BaseResponse
	RequestId      string                          `json:"RequestId" xml:"RequestId"`
	Success        bool                            `json:"Success" xml:"Success"`
	Code           string                          `json:"Code" xml:"Code"`
	Message        string                          `json:"Message" xml:"Message"`
	HttpStatusCode int                             `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PhoneNumbers   PhoneNumbersInAddCustomerNumber `json:"PhoneNumbers" xml:"PhoneNumbers"`
}

// CreateAddCustomerNumberRequest creates a request to invoke AddCustomerNumber API
func CreateAddCustomerNumberRequest() (request *AddCustomerNumberRequest) {
	request = &AddCustomerNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "AddCustomerNumber", "", "")
	request.Method = requests.POST
	return
}

// CreateAddCustomerNumberResponse creates a response to parse from AddCustomerNumber response
func CreateAddCustomerNumberResponse() (response *AddCustomerNumberResponse) {
	response = &AddCustomerNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
