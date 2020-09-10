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

// GetNumberRegionInfo invokes the cloudcallcenter.GetNumberRegionInfo API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getnumberregioninfo.html
func (client *Client) GetNumberRegionInfo(request *GetNumberRegionInfoRequest) (response *GetNumberRegionInfoResponse, err error) {
	response = CreateGetNumberRegionInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetNumberRegionInfoWithChan invokes the cloudcallcenter.GetNumberRegionInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getnumberregioninfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNumberRegionInfoWithChan(request *GetNumberRegionInfoRequest) (<-chan *GetNumberRegionInfoResponse, <-chan error) {
	responseChan := make(chan *GetNumberRegionInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNumberRegionInfo(request)
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

// GetNumberRegionInfoWithCallback invokes the cloudcallcenter.GetNumberRegionInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getnumberregioninfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNumberRegionInfoWithCallback(request *GetNumberRegionInfoRequest, callback func(response *GetNumberRegionInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNumberRegionInfoResponse
		var err error
		defer close(result)
		response, err = client.GetNumberRegionInfo(request)
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

// GetNumberRegionInfoRequest is the request struct for api GetNumberRegionInfo
type GetNumberRegionInfoRequest struct {
	*requests.RpcRequest
	Number     string `position:"Query" name:"Number"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetNumberRegionInfoResponse is the response struct for api GetNumberRegionInfo
type GetNumberRegionInfoResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Code           string      `json:"Code" xml:"Code"`
	Message        string      `json:"Message" xml:"Message"`
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PhoneNumber    PhoneNumber `json:"PhoneNumber" xml:"PhoneNumber"`
}

// CreateGetNumberRegionInfoRequest creates a request to invoke GetNumberRegionInfo API
func CreateGetNumberRegionInfoRequest() (request *GetNumberRegionInfoRequest) {
	request = &GetNumberRegionInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetNumberRegionInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateGetNumberRegionInfoResponse creates a response to parse from GetNumberRegionInfo response
func CreateGetNumberRegionInfoResponse() (response *GetNumberRegionInfoResponse) {
	response = &GetNumberRegionInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
