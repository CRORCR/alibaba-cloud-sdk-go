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

// GetDataSetOssHeader invokes the qualitycheck.GetDataSetOssHeader API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/getdatasetossheader.html
func (client *Client) GetDataSetOssHeader(request *GetDataSetOssHeaderRequest) (response *GetDataSetOssHeaderResponse, err error) {
	response = CreateGetDataSetOssHeaderResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataSetOssHeaderWithChan invokes the qualitycheck.GetDataSetOssHeader API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getdatasetossheader.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataSetOssHeaderWithChan(request *GetDataSetOssHeaderRequest) (<-chan *GetDataSetOssHeaderResponse, <-chan error) {
	responseChan := make(chan *GetDataSetOssHeaderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataSetOssHeader(request)
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

// GetDataSetOssHeaderWithCallback invokes the qualitycheck.GetDataSetOssHeader API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getdatasetossheader.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataSetOssHeaderWithCallback(request *GetDataSetOssHeaderRequest, callback func(response *GetDataSetOssHeaderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataSetOssHeaderResponse
		var err error
		defer close(result)
		response, err = client.GetDataSetOssHeader(request)
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

// GetDataSetOssHeaderRequest is the request struct for api GetDataSetOssHeader
type GetDataSetOssHeaderRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetDataSetOssHeaderResponse is the response struct for api GetDataSetOssHeader
type GetDataSetOssHeaderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetDataSetOssHeaderRequest creates a request to invoke GetDataSetOssHeader API
func CreateGetDataSetOssHeaderRequest() (request *GetDataSetOssHeaderRequest) {
	request = &GetDataSetOssHeaderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetDataSetOssHeader", "", "")
	return
}

// CreateGetDataSetOssHeaderResponse creates a response to parse from GetDataSetOssHeader response
func CreateGetDataSetOssHeaderResponse() (response *GetDataSetOssHeaderResponse) {
	response = &GetDataSetOssHeaderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
