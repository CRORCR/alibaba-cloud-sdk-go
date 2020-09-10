package arms

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

// UpdateAlertContact invokes the arms.UpdateAlertContact API synchronously
// api document: https://help.aliyun.com/api/arms/updatealertcontact.html
func (client *Client) UpdateAlertContact(request *UpdateAlertContactRequest) (response *UpdateAlertContactResponse, err error) {
	response = CreateUpdateAlertContactResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAlertContactWithChan invokes the arms.UpdateAlertContact API asynchronously
// api document: https://help.aliyun.com/api/arms/updatealertcontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAlertContactWithChan(request *UpdateAlertContactRequest) (<-chan *UpdateAlertContactResponse, <-chan error) {
	responseChan := make(chan *UpdateAlertContactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAlertContact(request)
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

// UpdateAlertContactWithCallback invokes the arms.UpdateAlertContact API asynchronously
// api document: https://help.aliyun.com/api/arms/updatealertcontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAlertContactWithCallback(request *UpdateAlertContactRequest, callback func(response *UpdateAlertContactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAlertContactResponse
		var err error
		defer close(result)
		response, err = client.UpdateAlertContact(request)
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

// UpdateAlertContactRequest is the request struct for api UpdateAlertContact
type UpdateAlertContactRequest struct {
	*requests.RpcRequest
	ContactId           requests.Integer `position:"Query" name:"ContactId"`
	PhoneNum            string           `position:"Query" name:"PhoneNum"`
	ProxyUserId         string           `position:"Query" name:"ProxyUserId"`
	ContactName         string           `position:"Query" name:"ContactName"`
	DingRobotWebhookUrl string           `position:"Query" name:"DingRobotWebhookUrl"`
	Email               string           `position:"Query" name:"Email"`
	SystemNoc           requests.Boolean `position:"Query" name:"SystemNoc"`
}

// UpdateAlertContactResponse is the response struct for api UpdateAlertContact
type UpdateAlertContactResponse struct {
	*responses.BaseResponse
	RequestId                   string `json:"RequestId" xml:"RequestId"`
	UpdateAlertContactIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
}

// CreateUpdateAlertContactRequest creates a request to invoke UpdateAlertContact API
func CreateUpdateAlertContactRequest() (request *UpdateAlertContactRequest) {
	request = &UpdateAlertContactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdateAlertContact", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateAlertContactResponse creates a response to parse from UpdateAlertContact response
func CreateUpdateAlertContactResponse() (response *UpdateAlertContactResponse) {
	response = &UpdateAlertContactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
