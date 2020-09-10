package ft

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

// FtGatedLaunchPolicy4 invokes the ft.FtGatedLaunchPolicy4 API synchronously
// api document: https://help.aliyun.com/api/ft/ftgatedlaunchpolicy4.html
func (client *Client) FtGatedLaunchPolicy4(request *FtGatedLaunchPolicy4Request) (response *FtGatedLaunchPolicy4Response, err error) {
	response = CreateFtGatedLaunchPolicy4Response()
	err = client.DoAction(request, response)
	return
}

// FtGatedLaunchPolicy4WithChan invokes the ft.FtGatedLaunchPolicy4 API asynchronously
// api document: https://help.aliyun.com/api/ft/ftgatedlaunchpolicy4.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FtGatedLaunchPolicy4WithChan(request *FtGatedLaunchPolicy4Request) (<-chan *FtGatedLaunchPolicy4Response, <-chan error) {
	responseChan := make(chan *FtGatedLaunchPolicy4Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FtGatedLaunchPolicy4(request)
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

// FtGatedLaunchPolicy4WithCallback invokes the ft.FtGatedLaunchPolicy4 API asynchronously
// api document: https://help.aliyun.com/api/ft/ftgatedlaunchpolicy4.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FtGatedLaunchPolicy4WithCallback(request *FtGatedLaunchPolicy4Request, callback func(response *FtGatedLaunchPolicy4Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FtGatedLaunchPolicy4Response
		var err error
		defer close(result)
		response, err = client.FtGatedLaunchPolicy4(request)
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

// FtGatedLaunchPolicy4Request is the request struct for api FtGatedLaunchPolicy4
type FtGatedLaunchPolicy4Request struct {
	*requests.RpcRequest
	IsGatedLaunch string `position:"Query" name:"IsGatedLaunch"`
}

// FtGatedLaunchPolicy4Response is the response struct for api FtGatedLaunchPolicy4
type FtGatedLaunchPolicy4Response struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	IsGatedLaunch string `json:"IsGatedLaunch" xml:"IsGatedLaunch"`
}

// CreateFtGatedLaunchPolicy4Request creates a request to invoke FtGatedLaunchPolicy4 API
func CreateFtGatedLaunchPolicy4Request() (request *FtGatedLaunchPolicy4Request) {
	request = &FtGatedLaunchPolicy4Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2018-07-13", "FtGatedLaunchPolicy4", "", "")
	return
}

// CreateFtGatedLaunchPolicy4Response creates a response to parse from FtGatedLaunchPolicy4 response
func CreateFtGatedLaunchPolicy4Response() (response *FtGatedLaunchPolicy4Response) {
	response = &FtGatedLaunchPolicy4Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
