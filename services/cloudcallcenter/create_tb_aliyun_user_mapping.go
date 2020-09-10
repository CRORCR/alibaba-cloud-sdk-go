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

// CreateTbAliyunUserMapping invokes the cloudcallcenter.CreateTbAliyunUserMapping API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createtbaliyunusermapping.html
func (client *Client) CreateTbAliyunUserMapping(request *CreateTbAliyunUserMappingRequest) (response *CreateTbAliyunUserMappingResponse, err error) {
	response = CreateCreateTbAliyunUserMappingResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTbAliyunUserMappingWithChan invokes the cloudcallcenter.CreateTbAliyunUserMapping API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createtbaliyunusermapping.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTbAliyunUserMappingWithChan(request *CreateTbAliyunUserMappingRequest) (<-chan *CreateTbAliyunUserMappingResponse, <-chan error) {
	responseChan := make(chan *CreateTbAliyunUserMappingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTbAliyunUserMapping(request)
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

// CreateTbAliyunUserMappingWithCallback invokes the cloudcallcenter.CreateTbAliyunUserMapping API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createtbaliyunusermapping.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateTbAliyunUserMappingWithCallback(request *CreateTbAliyunUserMappingRequest, callback func(response *CreateTbAliyunUserMappingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTbAliyunUserMappingResponse
		var err error
		defer close(result)
		response, err = client.CreateTbAliyunUserMapping(request)
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

// CreateTbAliyunUserMappingRequest is the request struct for api CreateTbAliyunUserMapping
type CreateTbAliyunUserMappingRequest struct {
	*requests.RpcRequest
	AliyunPk         requests.Integer `position:"Query" name:"AliyunPk"`
	TbUserid         requests.Integer `position:"Query" name:"TbUserid"`
	SubAccount       requests.Boolean `position:"Query" name:"SubAccount"`
	PrimaryAccountPk requests.Integer `position:"Query" name:"PrimaryAccountPk"`
	AliyunAkId       string           `position:"Query" name:"AliyunAkId"`
	AliyunAkSecret   string           `position:"Query" name:"AliyunAkSecret"`
	TbNickname       string           `position:"Query" name:"TbNickname"`
}

// CreateTbAliyunUserMappingResponse is the response struct for api CreateTbAliyunUserMapping
type CreateTbAliyunUserMappingResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Success          bool   `json:"Success" xml:"Success"`
	Code             string `json:"Code" xml:"Code"`
	Message          string `json:"Message" xml:"Message"`
	HttpStatusCode   int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Id               int64  `json:"Id" xml:"Id"`
	TbUserid         int64  `json:"TbUserid" xml:"TbUserid"`
	TbNickname       string `json:"TbNickname" xml:"TbNickname"`
	AliyunPk         int64  `json:"AliyunPk" xml:"AliyunPk"`
	SubAccount       bool   `json:"SubAccount" xml:"SubAccount"`
	PrimaryAccountPk int64  `json:"PrimaryAccountPk" xml:"PrimaryAccountPk"`
	AliyunAkId       string `json:"AliyunAkId" xml:"AliyunAkId"`
}

// CreateCreateTbAliyunUserMappingRequest creates a request to invoke CreateTbAliyunUserMapping API
func CreateCreateTbAliyunUserMappingRequest() (request *CreateTbAliyunUserMappingRequest) {
	request = &CreateTbAliyunUserMappingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateTbAliyunUserMapping", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateTbAliyunUserMappingResponse creates a response to parse from CreateTbAliyunUserMapping response
func CreateCreateTbAliyunUserMappingResponse() (response *CreateTbAliyunUserMappingResponse) {
	response = &CreateTbAliyunUserMappingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
