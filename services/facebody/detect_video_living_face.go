package facebody

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

// DetectVideoLivingFace invokes the facebody.DetectVideoLivingFace API synchronously
func (client *Client) DetectVideoLivingFace(request *DetectVideoLivingFaceRequest) (response *DetectVideoLivingFaceResponse, err error) {
	response = CreateDetectVideoLivingFaceResponse()
	err = client.DoAction(request, response)
	return
}

// DetectVideoLivingFaceWithChan invokes the facebody.DetectVideoLivingFace API asynchronously
func (client *Client) DetectVideoLivingFaceWithChan(request *DetectVideoLivingFaceRequest) (<-chan *DetectVideoLivingFaceResponse, <-chan error) {
	responseChan := make(chan *DetectVideoLivingFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectVideoLivingFace(request)
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

// DetectVideoLivingFaceWithCallback invokes the facebody.DetectVideoLivingFace API asynchronously
func (client *Client) DetectVideoLivingFaceWithCallback(request *DetectVideoLivingFaceRequest, callback func(response *DetectVideoLivingFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectVideoLivingFaceResponse
		var err error
		defer close(result)
		response, err = client.DetectVideoLivingFace(request)
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

// DetectVideoLivingFaceRequest is the request struct for api DetectVideoLivingFace
type DetectVideoLivingFaceRequest struct {
	*requests.RpcRequest
	VideoUrl string `position:"Body" name:"VideoUrl"`
}

// DetectVideoLivingFaceResponse is the response struct for api DetectVideoLivingFace
type DetectVideoLivingFaceResponse struct {
	*responses.BaseResponse
	RequestId string                      `json:"RequestId" xml:"RequestId"`
	Data      DataInDetectVideoLivingFace `json:"Data" xml:"Data"`
}

// CreateDetectVideoLivingFaceRequest creates a request to invoke DetectVideoLivingFace API
func CreateDetectVideoLivingFaceRequest() (request *DetectVideoLivingFaceRequest) {
	request = &DetectVideoLivingFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "DetectVideoLivingFace", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetectVideoLivingFaceResponse creates a response to parse from DetectVideoLivingFace response
func CreateDetectVideoLivingFaceResponse() (response *DetectVideoLivingFaceResponse) {
	response = &DetectVideoLivingFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
