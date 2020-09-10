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

// DownloadMedia invokes the cloudcallcenter.DownloadMedia API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadmedia.html
func (client *Client) DownloadMedia(request *DownloadMediaRequest) (response *DownloadMediaResponse, err error) {
	response = CreateDownloadMediaResponse()
	err = client.DoAction(request, response)
	return
}

// DownloadMediaWithChan invokes the cloudcallcenter.DownloadMedia API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadmedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadMediaWithChan(request *DownloadMediaRequest) (<-chan *DownloadMediaResponse, <-chan error) {
	responseChan := make(chan *DownloadMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadMedia(request)
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

// DownloadMediaWithCallback invokes the cloudcallcenter.DownloadMedia API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/downloadmedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DownloadMediaWithCallback(request *DownloadMediaRequest, callback func(response *DownloadMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadMediaResponse
		var err error
		defer close(result)
		response, err = client.DownloadMedia(request)
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

// DownloadMediaRequest is the request struct for api DownloadMedia
type DownloadMediaRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	Channel    string `position:"Query" name:"Channel"`
}

// DownloadMediaResponse is the response struct for api DownloadMedia
type DownloadMediaResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	Success            bool               `json:"Success" xml:"Success"`
	Code               string             `json:"Code" xml:"Code"`
	Message            string             `json:"Message" xml:"Message"`
	HttpStatusCode     int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	MediaDownloadParam MediaDownloadParam `json:"MediaDownloadParam" xml:"MediaDownloadParam"`
}

// CreateDownloadMediaRequest creates a request to invoke DownloadMedia API
func CreateDownloadMediaRequest() (request *DownloadMediaRequest) {
	request = &DownloadMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DownloadMedia", "", "")
	request.Method = requests.POST
	return
}

// CreateDownloadMediaResponse creates a response to parse from DownloadMedia response
func CreateDownloadMediaResponse() (response *DownloadMediaResponse) {
	response = &DownloadMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
