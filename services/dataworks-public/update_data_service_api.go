package dataworks_public

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

// UpdateDataServiceApi invokes the dataworks_public.UpdateDataServiceApi API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/updatedataserviceapi.html
func (client *Client) UpdateDataServiceApi(request *UpdateDataServiceApiRequest) (response *UpdateDataServiceApiResponse, err error) {
	response = CreateUpdateDataServiceApiResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDataServiceApiWithChan invokes the dataworks_public.UpdateDataServiceApi API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/updatedataserviceapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDataServiceApiWithChan(request *UpdateDataServiceApiRequest) (<-chan *UpdateDataServiceApiResponse, <-chan error) {
	responseChan := make(chan *UpdateDataServiceApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDataServiceApi(request)
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

// UpdateDataServiceApiWithCallback invokes the dataworks_public.UpdateDataServiceApi API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/updatedataserviceapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDataServiceApiWithCallback(request *UpdateDataServiceApiRequest, callback func(response *UpdateDataServiceApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDataServiceApiResponse
		var err error
		defer close(result)
		response, err = client.UpdateDataServiceApi(request)
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

// UpdateDataServiceApiRequest is the request struct for api UpdateDataServiceApi
type UpdateDataServiceApiRequest struct {
	*requests.RpcRequest
	ScriptDetails       string           `position:"Body" name:"ScriptDetails"`
	RequestMethod       requests.Integer `position:"Body" name:"RequestMethod"`
	ApiPath             string           `position:"Body" name:"ApiPath"`
	WizardDetails       string           `position:"Body" name:"WizardDetails"`
	VisibleRange        requests.Integer `position:"Body" name:"VisibleRange"`
	ApiDescription      string           `position:"Body" name:"ApiDescription"`
	Timeout             requests.Integer `position:"Body" name:"Timeout"`
	RegistrationDetails string           `position:"Body" name:"RegistrationDetails"`
	TenantId            requests.Integer `position:"Body" name:"TenantId"`
	Protocols           string           `position:"Body" name:"Protocols"`
	ProjectId           requests.Integer `position:"Body" name:"ProjectId"`
	ApiId               requests.Integer `position:"Body" name:"ApiId"`
	ResponseContentType requests.Integer `position:"Body" name:"ResponseContentType"`
}

// UpdateDataServiceApiResponse is the response struct for api UpdateDataServiceApi
type UpdateDataServiceApiResponse struct {
	*responses.BaseResponse
	Data           bool   `json:"Data" xml:"Data"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Success        bool   `json:"Success" xml:"Success"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDataServiceApiRequest creates a request to invoke UpdateDataServiceApi API
func CreateUpdateDataServiceApiRequest() (request *UpdateDataServiceApiRequest) {
	request = &UpdateDataServiceApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateDataServiceApi", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDataServiceApiResponse creates a response to parse from UpdateDataServiceApi response
func CreateUpdateDataServiceApiResponse() (response *UpdateDataServiceApiResponse) {
	response = &UpdateDataServiceApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
