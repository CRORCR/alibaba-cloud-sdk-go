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

// QueryDeviceFileList invokes the iot.QueryDeviceFileList API synchronously
// api document: https://help.aliyun.com/api/iot/querydevicefilelist.html
func (client *Client) QueryDeviceFileList(request *QueryDeviceFileListRequest) (response *QueryDeviceFileListResponse, err error) {
	response = CreateQueryDeviceFileListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceFileListWithChan invokes the iot.QueryDeviceFileList API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicefilelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceFileListWithChan(request *QueryDeviceFileListRequest) (<-chan *QueryDeviceFileListResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceFileListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceFileList(request)
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

// QueryDeviceFileListWithCallback invokes the iot.QueryDeviceFileList API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicefilelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceFileListWithCallback(request *QueryDeviceFileListRequest, callback func(response *QueryDeviceFileListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceFileListResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceFileList(request)
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

// QueryDeviceFileListRequest is the request struct for api QueryDeviceFileList
type QueryDeviceFileListRequest struct {
	*requests.RpcRequest
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// QueryDeviceFileListResponse is the response struct for api QueryDeviceFileList
type QueryDeviceFileListResponse struct {
	*responses.BaseResponse
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	Success      bool                      `json:"Success" xml:"Success"`
	Code         string                    `json:"Code" xml:"Code"`
	ErrorMessage string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	CurrentPage  int                       `json:"CurrentPage" xml:"CurrentPage"`
	PageCount    int                       `json:"PageCount" xml:"PageCount"`
	PageSize     int                       `json:"PageSize" xml:"PageSize"`
	Total        int                       `json:"Total" xml:"Total"`
	Data         DataInQueryDeviceFileList `json:"Data" xml:"Data"`
}

// CreateQueryDeviceFileListRequest creates a request to invoke QueryDeviceFileList API
func CreateQueryDeviceFileListRequest() (request *QueryDeviceFileListRequest) {
	request = &QueryDeviceFileListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceFileList", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceFileListResponse creates a response to parse from QueryDeviceFileList response
func CreateQueryDeviceFileListResponse() (response *QueryDeviceFileListResponse) {
	response = &QueryDeviceFileListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
