package qualitycheck

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

// SubmitComplaint invokes the qualitycheck.SubmitComplaint API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/submitcomplaint.html
func (client *Client) SubmitComplaint(request *SubmitComplaintRequest) (response *SubmitComplaintResponse, err error) {
	response = CreateSubmitComplaintResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitComplaintWithChan invokes the qualitycheck.SubmitComplaint API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/submitcomplaint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitComplaintWithChan(request *SubmitComplaintRequest) (<-chan *SubmitComplaintResponse, <-chan error) {
	responseChan := make(chan *SubmitComplaintResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitComplaint(request)
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

// SubmitComplaintWithCallback invokes the qualitycheck.SubmitComplaint API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/submitcomplaint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitComplaintWithCallback(request *SubmitComplaintRequest, callback func(response *SubmitComplaintResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitComplaintResponse
		var err error
		defer close(result)
		response, err = client.SubmitComplaint(request)
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

// SubmitComplaintRequest is the request struct for api SubmitComplaint
type SubmitComplaintRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// SubmitComplaintResponse is the response struct for api SubmitComplaint
type SubmitComplaintResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateSubmitComplaintRequest creates a request to invoke SubmitComplaint API
func CreateSubmitComplaintRequest() (request *SubmitComplaintRequest) {
	request = &SubmitComplaintRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "SubmitComplaint", "", "")
	return
}

// CreateSubmitComplaintResponse creates a response to parse from SubmitComplaint response
func CreateSubmitComplaintResponse() (response *SubmitComplaintResponse) {
	response = &SubmitComplaintResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
