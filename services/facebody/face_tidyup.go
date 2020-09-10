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

// FaceTidyup invokes the facebody.FaceTidyup API synchronously
func (client *Client) FaceTidyup(request *FaceTidyupRequest) (response *FaceTidyupResponse, err error) {
	response = CreateFaceTidyupResponse()
	err = client.DoAction(request, response)
	return
}

// FaceTidyupWithChan invokes the facebody.FaceTidyup API asynchronously
func (client *Client) FaceTidyupWithChan(request *FaceTidyupRequest) (<-chan *FaceTidyupResponse, <-chan error) {
	responseChan := make(chan *FaceTidyupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FaceTidyup(request)
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

// FaceTidyupWithCallback invokes the facebody.FaceTidyup API asynchronously
func (client *Client) FaceTidyupWithCallback(request *FaceTidyupRequest, callback func(response *FaceTidyupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FaceTidyupResponse
		var err error
		defer close(result)
		response, err = client.FaceTidyup(request)
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

// FaceTidyupRequest is the request struct for api FaceTidyup
type FaceTidyupRequest struct {
	*requests.RpcRequest
	ShapeType requests.Integer `position:"Body" name:"ShapeType"`
	Strength  requests.Float   `position:"Body" name:"Strength"`
	ImageURL  string           `position:"Body" name:"ImageURL"`
}

// FaceTidyupResponse is the response struct for api FaceTidyup
type FaceTidyupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFaceTidyupRequest creates a request to invoke FaceTidyup API
func CreateFaceTidyupRequest() (request *FaceTidyupRequest) {
	request = &FaceTidyupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "FaceTidyup", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFaceTidyupResponse creates a response to parse from FaceTidyup response
func CreateFaceTidyupResponse() (response *FaceTidyupResponse) {
	response = &FaceTidyupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
