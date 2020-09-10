package domain_intl

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

// SaveBatchTaskForUpdatingContactInfoByNewContact invokes the domain_intl.SaveBatchTaskForUpdatingContactInfoByNewContact API synchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskforupdatingcontactinfobynewcontact.html
func (client *Client) SaveBatchTaskForUpdatingContactInfoByNewContact(request *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) (response *SaveBatchTaskForUpdatingContactInfoByNewContactResponse, err error) {
	response = CreateSaveBatchTaskForUpdatingContactInfoByNewContactResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForUpdatingContactInfoByNewContactWithChan invokes the domain_intl.SaveBatchTaskForUpdatingContactInfoByNewContact API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskforupdatingcontactinfobynewcontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveBatchTaskForUpdatingContactInfoByNewContactWithChan(request *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) (<-chan *SaveBatchTaskForUpdatingContactInfoByNewContactResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForUpdatingContactInfoByNewContactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForUpdatingContactInfoByNewContact(request)
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

// SaveBatchTaskForUpdatingContactInfoByNewContactWithCallback invokes the domain_intl.SaveBatchTaskForUpdatingContactInfoByNewContact API asynchronously
// api document: https://help.aliyun.com/api/domain-intl/savebatchtaskforupdatingcontactinfobynewcontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveBatchTaskForUpdatingContactInfoByNewContactWithCallback(request *SaveBatchTaskForUpdatingContactInfoByNewContactRequest, callback func(response *SaveBatchTaskForUpdatingContactInfoByNewContactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForUpdatingContactInfoByNewContactResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForUpdatingContactInfoByNewContact(request)
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

// SaveBatchTaskForUpdatingContactInfoByNewContactRequest is the request struct for api SaveBatchTaskForUpdatingContactInfoByNewContact
type SaveBatchTaskForUpdatingContactInfoByNewContactRequest struct {
	*requests.RpcRequest
	Country                string           `position:"Query" name:"Country"`
	Address                string           `position:"Query" name:"Address"`
	TelArea                string           `position:"Query" name:"TelArea"`
	ContactType            string           `position:"Query" name:"ContactType"`
	City                   string           `position:"Query" name:"City"`
	DomainName             *[]string        `position:"Query" name:"DomainName"  type:"Repeated"`
	Telephone              string           `position:"Query" name:"Telephone"`
	TransferOutProhibited  requests.Boolean `position:"Query" name:"TransferOutProhibited"`
	RegistrantOrganization string           `position:"Query" name:"RegistrantOrganization"`
	TelExt                 string           `position:"Query" name:"TelExt"`
	Province               string           `position:"Query" name:"Province"`
	PostalCode             string           `position:"Query" name:"PostalCode"`
	UserClientIp           string           `position:"Query" name:"UserClientIp"`
	Lang                   string           `position:"Query" name:"Lang"`
	Email                  string           `position:"Query" name:"Email"`
	RegistrantName         string           `position:"Query" name:"RegistrantName"`
}

// SaveBatchTaskForUpdatingContactInfoByNewContactResponse is the response struct for api SaveBatchTaskForUpdatingContactInfoByNewContact
type SaveBatchTaskForUpdatingContactInfoByNewContactResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForUpdatingContactInfoByNewContactRequest creates a request to invoke SaveBatchTaskForUpdatingContactInfoByNewContact API
func CreateSaveBatchTaskForUpdatingContactInfoByNewContactRequest() (request *SaveBatchTaskForUpdatingContactInfoByNewContactRequest) {
	request = &SaveBatchTaskForUpdatingContactInfoByNewContactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveBatchTaskForUpdatingContactInfoByNewContact", "domain", "openAPI")
	return
}

// CreateSaveBatchTaskForUpdatingContactInfoByNewContactResponse creates a response to parse from SaveBatchTaskForUpdatingContactInfoByNewContact response
func CreateSaveBatchTaskForUpdatingContactInfoByNewContactResponse() (response *SaveBatchTaskForUpdatingContactInfoByNewContactResponse) {
	response = &SaveBatchTaskForUpdatingContactInfoByNewContactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
