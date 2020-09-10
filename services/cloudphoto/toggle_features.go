package cloudphoto

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

// ToggleFeatures invokes the cloudphoto.ToggleFeatures API synchronously
// api document: https://help.aliyun.com/api/cloudphoto/togglefeatures.html
func (client *Client) ToggleFeatures(request *ToggleFeaturesRequest) (response *ToggleFeaturesResponse, err error) {
	response = CreateToggleFeaturesResponse()
	err = client.DoAction(request, response)
	return
}

// ToggleFeaturesWithChan invokes the cloudphoto.ToggleFeatures API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/togglefeatures.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ToggleFeaturesWithChan(request *ToggleFeaturesRequest) (<-chan *ToggleFeaturesResponse, <-chan error) {
	responseChan := make(chan *ToggleFeaturesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ToggleFeatures(request)
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

// ToggleFeaturesWithCallback invokes the cloudphoto.ToggleFeatures API asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/togglefeatures.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ToggleFeaturesWithCallback(request *ToggleFeaturesRequest, callback func(response *ToggleFeaturesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ToggleFeaturesResponse
		var err error
		defer close(result)
		response, err = client.ToggleFeatures(request)
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

// ToggleFeaturesRequest is the request struct for api ToggleFeatures
type ToggleFeaturesRequest struct {
	*requests.RpcRequest
	DisabledFeatures *[]string `position:"Query" name:"DisabledFeatures"  type:"Repeated"`
	StoreName        string    `position:"Query" name:"StoreName"`
	EnabledFeatures  *[]string `position:"Query" name:"EnabledFeatures"  type:"Repeated"`
}

// ToggleFeaturesResponse is the response struct for api ToggleFeatures
type ToggleFeaturesResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
}

// CreateToggleFeaturesRequest creates a request to invoke ToggleFeatures API
func CreateToggleFeaturesRequest() (request *ToggleFeaturesRequest) {
	request = &ToggleFeaturesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "ToggleFeatures", "cloudphoto", "openAPI")
	return
}

// CreateToggleFeaturesResponse creates a response to parse from ToggleFeatures response
func CreateToggleFeaturesResponse() (response *ToggleFeaturesResponse) {
	response = &ToggleFeaturesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
