package dyplsapi

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

// QuerySubsId invokes the dyplsapi.QuerySubsId API synchronously
// api document: https://help.aliyun.com/api/dyplsapi/querysubsid.html
func (client *Client) QuerySubsId(request *QuerySubsIdRequest) (response *QuerySubsIdResponse, err error) {
	response = CreateQuerySubsIdResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySubsIdWithChan invokes the dyplsapi.QuerySubsId API asynchronously
// api document: https://help.aliyun.com/api/dyplsapi/querysubsid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySubsIdWithChan(request *QuerySubsIdRequest) (<-chan *QuerySubsIdResponse, <-chan error) {
	responseChan := make(chan *QuerySubsIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySubsId(request)
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

// QuerySubsIdWithCallback invokes the dyplsapi.QuerySubsId API asynchronously
// api document: https://help.aliyun.com/api/dyplsapi/querysubsid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QuerySubsIdWithCallback(request *QuerySubsIdRequest, callback func(response *QuerySubsIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySubsIdResponse
		var err error
		defer close(result)
		response, err = client.QuerySubsId(request)
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

// QuerySubsIdRequest is the request struct for api QuerySubsId
type QuerySubsIdRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PoolKey              string           `position:"Query" name:"PoolKey"`
	PhoneNoX             string           `position:"Query" name:"PhoneNoX"`
}

// QuerySubsIdResponse is the response struct for api QuerySubsId
type QuerySubsIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	SubsId    string `json:"SubsId" xml:"SubsId"`
}

// CreateQuerySubsIdRequest creates a request to invoke QuerySubsId API
func CreateQuerySubsIdRequest() (request *QuerySubsIdRequest) {
	request = &QuerySubsIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "QuerySubsId", "dypls", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySubsIdResponse creates a response to parse from QuerySubsId response
func CreateQuerySubsIdResponse() (response *QuerySubsIdResponse) {
	response = &QuerySubsIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
