package iot

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

// QueryCertUrlByApplyId invokes the iot.QueryCertUrlByApplyId API synchronously
// api document: https://help.aliyun.com/api/iot/querycerturlbyapplyid.html
func (client *Client) QueryCertUrlByApplyId(request *QueryCertUrlByApplyIdRequest) (response *QueryCertUrlByApplyIdResponse, err error) {
	response = CreateQueryCertUrlByApplyIdResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCertUrlByApplyIdWithChan invokes the iot.QueryCertUrlByApplyId API asynchronously
// api document: https://help.aliyun.com/api/iot/querycerturlbyapplyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCertUrlByApplyIdWithChan(request *QueryCertUrlByApplyIdRequest) (<-chan *QueryCertUrlByApplyIdResponse, <-chan error) {
	responseChan := make(chan *QueryCertUrlByApplyIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCertUrlByApplyId(request)
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

// QueryCertUrlByApplyIdWithCallback invokes the iot.QueryCertUrlByApplyId API asynchronously
// api document: https://help.aliyun.com/api/iot/querycerturlbyapplyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCertUrlByApplyIdWithCallback(request *QueryCertUrlByApplyIdRequest, callback func(response *QueryCertUrlByApplyIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCertUrlByApplyIdResponse
		var err error
		defer close(result)
		response, err = client.QueryCertUrlByApplyId(request)
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

// QueryCertUrlByApplyIdRequest is the request struct for api QueryCertUrlByApplyId
type QueryCertUrlByApplyIdRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	ApplyId       requests.Integer `position:"Query" name:"ApplyId"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryCertUrlByApplyIdResponse is the response struct for api QueryCertUrlByApplyId
type QueryCertUrlByApplyIdResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	CertUrl      string `json:"CertUrl" xml:"CertUrl"`
}

// CreateQueryCertUrlByApplyIdRequest creates a request to invoke QueryCertUrlByApplyId API
func CreateQueryCertUrlByApplyIdRequest() (request *QueryCertUrlByApplyIdRequest) {
	request = &QueryCertUrlByApplyIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryCertUrlByApplyId", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryCertUrlByApplyIdResponse creates a response to parse from QueryCertUrlByApplyId response
func CreateQueryCertUrlByApplyIdResponse() (response *QueryCertUrlByApplyIdResponse) {
	response = &QueryCertUrlByApplyIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
