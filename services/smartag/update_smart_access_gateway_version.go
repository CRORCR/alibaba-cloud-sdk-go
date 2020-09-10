package smartag

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

// UpdateSmartAccessGatewayVersion invokes the smartag.UpdateSmartAccessGatewayVersion API synchronously
// api document: https://help.aliyun.com/api/smartag/updatesmartaccessgatewayversion.html
func (client *Client) UpdateSmartAccessGatewayVersion(request *UpdateSmartAccessGatewayVersionRequest) (response *UpdateSmartAccessGatewayVersionResponse, err error) {
	response = CreateUpdateSmartAccessGatewayVersionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSmartAccessGatewayVersionWithChan invokes the smartag.UpdateSmartAccessGatewayVersion API asynchronously
// api document: https://help.aliyun.com/api/smartag/updatesmartaccessgatewayversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSmartAccessGatewayVersionWithChan(request *UpdateSmartAccessGatewayVersionRequest) (<-chan *UpdateSmartAccessGatewayVersionResponse, <-chan error) {
	responseChan := make(chan *UpdateSmartAccessGatewayVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSmartAccessGatewayVersion(request)
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

// UpdateSmartAccessGatewayVersionWithCallback invokes the smartag.UpdateSmartAccessGatewayVersion API asynchronously
// api document: https://help.aliyun.com/api/smartag/updatesmartaccessgatewayversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSmartAccessGatewayVersionWithCallback(request *UpdateSmartAccessGatewayVersionRequest, callback func(response *UpdateSmartAccessGatewayVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSmartAccessGatewayVersionResponse
		var err error
		defer close(result)
		response, err = client.UpdateSmartAccessGatewayVersion(request)
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

// UpdateSmartAccessGatewayVersionRequest is the request struct for api UpdateSmartAccessGatewayVersion
type UpdateSmartAccessGatewayVersionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VersionCode          string           `position:"Query" name:"VersionCode"`
	SerialNumber         string           `position:"Query" name:"SerialNumber"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// UpdateSmartAccessGatewayVersionResponse is the response struct for api UpdateSmartAccessGatewayVersion
type UpdateSmartAccessGatewayVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateSmartAccessGatewayVersionRequest creates a request to invoke UpdateSmartAccessGatewayVersion API
func CreateUpdateSmartAccessGatewayVersionRequest() (request *UpdateSmartAccessGatewayVersionRequest) {
	request = &UpdateSmartAccessGatewayVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "UpdateSmartAccessGatewayVersion", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateSmartAccessGatewayVersionResponse creates a response to parse from UpdateSmartAccessGatewayVersion response
func CreateUpdateSmartAccessGatewayVersionResponse() (response *UpdateSmartAccessGatewayVersionResponse) {
	response = &UpdateSmartAccessGatewayVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
