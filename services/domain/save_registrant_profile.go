package domain

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

// SaveRegistrantProfile invokes the domain.SaveRegistrantProfile API synchronously
// api document: https://help.aliyun.com/api/domain/saveregistrantprofile.html
func (client *Client) SaveRegistrantProfile(request *SaveRegistrantProfileRequest) (response *SaveRegistrantProfileResponse, err error) {
	response = CreateSaveRegistrantProfileResponse()
	err = client.DoAction(request, response)
	return
}

// SaveRegistrantProfileWithChan invokes the domain.SaveRegistrantProfile API asynchronously
// api document: https://help.aliyun.com/api/domain/saveregistrantprofile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveRegistrantProfileWithChan(request *SaveRegistrantProfileRequest) (<-chan *SaveRegistrantProfileResponse, <-chan error) {
	responseChan := make(chan *SaveRegistrantProfileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveRegistrantProfile(request)
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

// SaveRegistrantProfileWithCallback invokes the domain.SaveRegistrantProfile API asynchronously
// api document: https://help.aliyun.com/api/domain/saveregistrantprofile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveRegistrantProfileWithCallback(request *SaveRegistrantProfileRequest, callback func(response *SaveRegistrantProfileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveRegistrantProfileResponse
		var err error
		defer close(result)
		response, err = client.SaveRegistrantProfile(request)
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

// SaveRegistrantProfileRequest is the request struct for api SaveRegistrantProfile
type SaveRegistrantProfileRequest struct {
	*requests.RpcRequest
	Country                  string           `position:"Query" name:"Country"`
	City                     string           `position:"Query" name:"City"`
	RegistrantProfileId      requests.Integer `position:"Query" name:"RegistrantProfileId"`
	ZhCity                   string           `position:"Query" name:"ZhCity"`
	TelExt                   string           `position:"Query" name:"TelExt"`
	Province                 string           `position:"Query" name:"Province"`
	ZhRegistrantName         string           `position:"Query" name:"ZhRegistrantName"`
	PostalCode               string           `position:"Query" name:"PostalCode"`
	Lang                     string           `position:"Query" name:"Lang"`
	Email                    string           `position:"Query" name:"Email"`
	ZhRegistrantOrganization string           `position:"Query" name:"ZhRegistrantOrganization"`
	Address                  string           `position:"Query" name:"Address"`
	TelArea                  string           `position:"Query" name:"TelArea"`
	ZhAddress                string           `position:"Query" name:"ZhAddress"`
	RegistrantType           string           `position:"Query" name:"RegistrantType"`
	RegistrantProfileType    string           `position:"Query" name:"RegistrantProfileType"`
	Telephone                string           `position:"Query" name:"Telephone"`
	DefaultRegistrantProfile requests.Boolean `position:"Query" name:"DefaultRegistrantProfile"`
	ZhProvince               string           `position:"Query" name:"ZhProvince"`
	RegistrantOrganization   string           `position:"Query" name:"RegistrantOrganization"`
	UserClientIp             string           `position:"Query" name:"UserClientIp"`
	RegistrantName           string           `position:"Query" name:"RegistrantName"`
}

// SaveRegistrantProfileResponse is the response struct for api SaveRegistrantProfile
type SaveRegistrantProfileResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	RegistrantProfileId int64  `json:"RegistrantProfileId" xml:"RegistrantProfileId"`
}

// CreateSaveRegistrantProfileRequest creates a request to invoke SaveRegistrantProfile API
func CreateSaveRegistrantProfileRequest() (request *SaveRegistrantProfileRequest) {
	request = &SaveRegistrantProfileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveRegistrantProfile", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveRegistrantProfileResponse creates a response to parse from SaveRegistrantProfile response
func CreateSaveRegistrantProfileResponse() (response *SaveRegistrantProfileResponse) {
	response = &SaveRegistrantProfileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
