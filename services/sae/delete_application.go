package sae

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

// DeleteApplication invokes the sae.DeleteApplication API synchronously
// api document: https://help.aliyun.com/api/sae/deleteapplication.html
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
	response = CreateDeleteApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteApplicationWithChan invokes the sae.DeleteApplication API asynchronously
// api document: https://help.aliyun.com/api/sae/deleteapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteApplicationWithChan(request *DeleteApplicationRequest) (<-chan *DeleteApplicationResponse, <-chan error) {
	responseChan := make(chan *DeleteApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteApplication(request)
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

// DeleteApplicationWithCallback invokes the sae.DeleteApplication API asynchronously
// api document: https://help.aliyun.com/api/sae/deleteapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteApplicationWithCallback(request *DeleteApplicationRequest, callback func(response *DeleteApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteApplicationResponse
		var err error
		defer close(result)
		response, err = client.DeleteApplication(request)
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

// DeleteApplicationRequest is the request struct for api DeleteApplication
type DeleteApplicationRequest struct {
	*requests.RoaRequest
	AppId string `position:"Query" name:"AppId"`
}

// DeleteApplicationResponse is the response struct for api DeleteApplication
type DeleteApplicationResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDeleteApplicationRequest creates a request to invoke DeleteApplication API
func CreateDeleteApplicationRequest() (request *DeleteApplicationRequest) {
	request = &DeleteApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "DeleteApplication", "/pop/v1/sam/app/deleteApplication", "serverless", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteApplicationResponse creates a response to parse from DeleteApplication response
func CreateDeleteApplicationResponse() (response *DeleteApplicationResponse) {
	response = &DeleteApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
