package ram

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

// UpdateLoginProfile invokes the ram.UpdateLoginProfile API synchronously
// api document: https://help.aliyun.com/api/ram/updateloginprofile.html
func (client *Client) UpdateLoginProfile(request *UpdateLoginProfileRequest) (response *UpdateLoginProfileResponse, err error) {
	response = CreateUpdateLoginProfileResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLoginProfileWithChan invokes the ram.UpdateLoginProfile API asynchronously
// api document: https://help.aliyun.com/api/ram/updateloginprofile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateLoginProfileWithChan(request *UpdateLoginProfileRequest) (<-chan *UpdateLoginProfileResponse, <-chan error) {
	responseChan := make(chan *UpdateLoginProfileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLoginProfile(request)
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

// UpdateLoginProfileWithCallback invokes the ram.UpdateLoginProfile API asynchronously
// api document: https://help.aliyun.com/api/ram/updateloginprofile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateLoginProfileWithCallback(request *UpdateLoginProfileRequest, callback func(response *UpdateLoginProfileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLoginProfileResponse
		var err error
		defer close(result)
		response, err = client.UpdateLoginProfile(request)
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

// UpdateLoginProfileRequest is the request struct for api UpdateLoginProfile
type UpdateLoginProfileRequest struct {
	*requests.RpcRequest
	PasswordResetRequired requests.Boolean `position:"Query" name:"PasswordResetRequired"`
	Password              string           `position:"Query" name:"Password"`
	MFABindRequired       requests.Boolean `position:"Query" name:"MFABindRequired"`
	UserName              string           `position:"Query" name:"UserName"`
}

// UpdateLoginProfileResponse is the response struct for api UpdateLoginProfile
type UpdateLoginProfileResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLoginProfileRequest creates a request to invoke UpdateLoginProfile API
func CreateUpdateLoginProfileRequest() (request *UpdateLoginProfileRequest) {
	request = &UpdateLoginProfileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "UpdateLoginProfile", "Ram", "openAPI")
	return
}

// CreateUpdateLoginProfileResponse creates a response to parse from UpdateLoginProfile response
func CreateUpdateLoginProfileResponse() (response *UpdateLoginProfileResponse) {
	response = &UpdateLoginProfileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
