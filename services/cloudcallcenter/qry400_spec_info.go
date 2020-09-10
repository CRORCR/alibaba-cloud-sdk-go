package cloudcallcenter

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

// Qry400SpecInfo invokes the cloudcallcenter.Qry400SpecInfo API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/qry400specinfo.html
func (client *Client) Qry400SpecInfo(request *Qry400SpecInfoRequest) (response *Qry400SpecInfoResponse, err error) {
	response = CreateQry400SpecInfoResponse()
	err = client.DoAction(request, response)
	return
}

// Qry400SpecInfoWithChan invokes the cloudcallcenter.Qry400SpecInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/qry400specinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) Qry400SpecInfoWithChan(request *Qry400SpecInfoRequest) (<-chan *Qry400SpecInfoResponse, <-chan error) {
	responseChan := make(chan *Qry400SpecInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Qry400SpecInfo(request)
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

// Qry400SpecInfoWithCallback invokes the cloudcallcenter.Qry400SpecInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/qry400specinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) Qry400SpecInfoWithCallback(request *Qry400SpecInfoRequest, callback func(response *Qry400SpecInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *Qry400SpecInfoResponse
		var err error
		defer close(result)
		response, err = client.Qry400SpecInfo(request)
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

// Qry400SpecInfoRequest is the request struct for api Qry400SpecInfo
type Qry400SpecInfoRequest struct {
	*requests.RpcRequest
}

// Qry400SpecInfoResponse is the response struct for api Qry400SpecInfo
type Qry400SpecInfoResponse struct {
	*responses.BaseResponse
	RequestId      string               `json:"RequestId" xml:"RequestId"`
	Success        bool                 `json:"Success" xml:"Success"`
	Code           string               `json:"Code" xml:"Code"`
	Message        string               `json:"Message" xml:"Message"`
	HttpStatusCode int                  `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInQry400SpecInfo `json:"Data" xml:"Data"`
}

// CreateQry400SpecInfoRequest creates a request to invoke Qry400SpecInfo API
func CreateQry400SpecInfoRequest() (request *Qry400SpecInfoRequest) {
	request = &Qry400SpecInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "Qry400SpecInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateQry400SpecInfoResponse creates a response to parse from Qry400SpecInfo response
func CreateQry400SpecInfoResponse() (response *Qry400SpecInfoResponse) {
	response = &Qry400SpecInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
