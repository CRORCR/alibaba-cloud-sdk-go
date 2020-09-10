package imm

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

// GetOfficePreviewURL invokes the imm.GetOfficePreviewURL API synchronously
// api document: https://help.aliyun.com/api/imm/getofficepreviewurl.html
func (client *Client) GetOfficePreviewURL(request *GetOfficePreviewURLRequest) (response *GetOfficePreviewURLResponse, err error) {
	response = CreateGetOfficePreviewURLResponse()
	err = client.DoAction(request, response)
	return
}

// GetOfficePreviewURLWithChan invokes the imm.GetOfficePreviewURL API asynchronously
// api document: https://help.aliyun.com/api/imm/getofficepreviewurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOfficePreviewURLWithChan(request *GetOfficePreviewURLRequest) (<-chan *GetOfficePreviewURLResponse, <-chan error) {
	responseChan := make(chan *GetOfficePreviewURLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOfficePreviewURL(request)
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

// GetOfficePreviewURLWithCallback invokes the imm.GetOfficePreviewURL API asynchronously
// api document: https://help.aliyun.com/api/imm/getofficepreviewurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetOfficePreviewURLWithCallback(request *GetOfficePreviewURLRequest, callback func(response *GetOfficePreviewURLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOfficePreviewURLResponse
		var err error
		defer close(result)
		response, err = client.GetOfficePreviewURL(request)
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

// GetOfficePreviewURLRequest is the request struct for api GetOfficePreviewURL
type GetOfficePreviewURLRequest struct {
	*requests.RpcRequest
	SrcType             string           `position:"Query" name:"SrcType"`
	Project             string           `position:"Query" name:"Project"`
	WatermarkVertical   requests.Integer `position:"Query" name:"WatermarkVertical"`
	WatermarkType       requests.Integer `position:"Query" name:"WatermarkType"`
	WatermarkRotate     requests.Float   `position:"Query" name:"WatermarkRotate"`
	WatermarkValue      string           `position:"Query" name:"WatermarkValue"`
	WatermarkFont       string           `position:"Query" name:"WatermarkFont"`
	WatermarkHorizontal requests.Integer `position:"Query" name:"WatermarkHorizontal"`
	SrcUri              string           `position:"Query" name:"SrcUri"`
	WatermarkFillStyle  string           `position:"Query" name:"WatermarkFillStyle"`
}

// GetOfficePreviewURLResponse is the response struct for api GetOfficePreviewURL
type GetOfficePreviewURLResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	PreviewURL              string `json:"PreviewURL" xml:"PreviewURL"`
	AccessToken             string `json:"AccessToken" xml:"AccessToken"`
	RefreshToken            string `json:"RefreshToken" xml:"RefreshToken"`
	AccessTokenExpiredTime  string `json:"AccessTokenExpiredTime" xml:"AccessTokenExpiredTime"`
	RefreshTokenExpiredTime string `json:"RefreshTokenExpiredTime" xml:"RefreshTokenExpiredTime"`
}

// CreateGetOfficePreviewURLRequest creates a request to invoke GetOfficePreviewURL API
func CreateGetOfficePreviewURLRequest() (request *GetOfficePreviewURLRequest) {
	request = &GetOfficePreviewURLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetOfficePreviewURL", "", "")
	request.Method = requests.POST
	return
}

// CreateGetOfficePreviewURLResponse creates a response to parse from GetOfficePreviewURL response
func CreateGetOfficePreviewURLResponse() (response *GetOfficePreviewURLResponse) {
	response = &GetOfficePreviewURLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
