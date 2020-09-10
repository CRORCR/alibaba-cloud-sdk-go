package dms_enterprise

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

// GetDataCorrectOrderDetail invokes the dms_enterprise.GetDataCorrectOrderDetail API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdatacorrectorderdetail.html
func (client *Client) GetDataCorrectOrderDetail(request *GetDataCorrectOrderDetailRequest) (response *GetDataCorrectOrderDetailResponse, err error) {
	response = CreateGetDataCorrectOrderDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataCorrectOrderDetailWithChan invokes the dms_enterprise.GetDataCorrectOrderDetail API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdatacorrectorderdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataCorrectOrderDetailWithChan(request *GetDataCorrectOrderDetailRequest) (<-chan *GetDataCorrectOrderDetailResponse, <-chan error) {
	responseChan := make(chan *GetDataCorrectOrderDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataCorrectOrderDetail(request)
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

// GetDataCorrectOrderDetailWithCallback invokes the dms_enterprise.GetDataCorrectOrderDetail API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdatacorrectorderdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataCorrectOrderDetailWithCallback(request *GetDataCorrectOrderDetailRequest, callback func(response *GetDataCorrectOrderDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataCorrectOrderDetailResponse
		var err error
		defer close(result)
		response, err = client.GetDataCorrectOrderDetail(request)
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

// GetDataCorrectOrderDetailRequest is the request struct for api GetDataCorrectOrderDetail
type GetDataCorrectOrderDetailRequest struct {
	*requests.RpcRequest
	OrderId requests.Integer `position:"Query" name:"OrderId"`
	Tid     requests.Integer `position:"Query" name:"Tid"`
}

// GetDataCorrectOrderDetailResponse is the response struct for api GetDataCorrectOrderDetail
type GetDataCorrectOrderDetailResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	Success                bool                   `json:"Success" xml:"Success"`
	ErrorMessage           string                 `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode              string                 `json:"ErrorCode" xml:"ErrorCode"`
	DataCorrectOrderDetail DataCorrectOrderDetail `json:"DataCorrectOrderDetail" xml:"DataCorrectOrderDetail"`
}

// CreateGetDataCorrectOrderDetailRequest creates a request to invoke GetDataCorrectOrderDetail API
func CreateGetDataCorrectOrderDetailRequest() (request *GetDataCorrectOrderDetailRequest) {
	request = &GetDataCorrectOrderDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetDataCorrectOrderDetail", "dmsenterprise", "openAPI")
	return
}

// CreateGetDataCorrectOrderDetailResponse creates a response to parse from GetDataCorrectOrderDetail response
func CreateGetDataCorrectOrderDetailResponse() (response *GetDataCorrectOrderDetailResponse) {
	response = &GetDataCorrectOrderDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
