package unimkt

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

// PopUpQuery2 invokes the unimkt.PopUpQuery2 API synchronously
// api document: https://help.aliyun.com/api/unimkt/popupquery2.html
func (client *Client) PopUpQuery2(request *PopUpQuery2Request) (response *PopUpQuery2Response, err error) {
	response = CreatePopUpQuery2Response()
	err = client.DoAction(request, response)
	return
}

// PopUpQuery2WithChan invokes the unimkt.PopUpQuery2 API asynchronously
// api document: https://help.aliyun.com/api/unimkt/popupquery2.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PopUpQuery2WithChan(request *PopUpQuery2Request) (<-chan *PopUpQuery2Response, <-chan error) {
	responseChan := make(chan *PopUpQuery2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PopUpQuery2(request)
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

// PopUpQuery2WithCallback invokes the unimkt.PopUpQuery2 API asynchronously
// api document: https://help.aliyun.com/api/unimkt/popupquery2.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PopUpQuery2WithCallback(request *PopUpQuery2Request, callback func(response *PopUpQuery2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PopUpQuery2Response
		var err error
		defer close(result)
		response, err = client.PopUpQuery2(request)
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

// PopUpQuery2Request is the request struct for api PopUpQuery2
type PopUpQuery2Request struct {
	*requests.RpcRequest
	Extra        string `position:"Body" name:"Extra"`
	UrlId        string `position:"Body" name:"UrlId"`
	AlipayOpenId string `position:"Body" name:"AlipayOpenId"`
	ChannelId    string `position:"Body" name:"ChannelId"`
	OuterCode    string `position:"Body" name:"OuterCode"`
}

// PopUpQuery2Response is the response struct for api PopUpQuery2
type PopUpQuery2Response struct {
	*responses.BaseResponse
	Status    bool   `json:"Status" xml:"Status"`
	Msg       string `json:"Msg" xml:"Msg"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Url       string `json:"Url" xml:"Url"`
}

// CreatePopUpQuery2Request creates a request to invoke PopUpQuery2 API
func CreatePopUpQuery2Request() (request *PopUpQuery2Request) {
	request = &PopUpQuery2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-07", "PopUpQuery2", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePopUpQuery2Response creates a response to parse from PopUpQuery2 response
func CreatePopUpQuery2Response() (response *PopUpQuery2Response) {
	response = &PopUpQuery2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
