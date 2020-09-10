package green

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

// DescribeImageFromLib invokes the green.DescribeImageFromLib API synchronously
// api document: https://help.aliyun.com/api/green/describeimagefromlib.html
func (client *Client) DescribeImageFromLib(request *DescribeImageFromLibRequest) (response *DescribeImageFromLibResponse, err error) {
	response = CreateDescribeImageFromLibResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageFromLibWithChan invokes the green.DescribeImageFromLib API asynchronously
// api document: https://help.aliyun.com/api/green/describeimagefromlib.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeImageFromLibWithChan(request *DescribeImageFromLibRequest) (<-chan *DescribeImageFromLibResponse, <-chan error) {
	responseChan := make(chan *DescribeImageFromLibResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageFromLib(request)
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

// DescribeImageFromLibWithCallback invokes the green.DescribeImageFromLib API asynchronously
// api document: https://help.aliyun.com/api/green/describeimagefromlib.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeImageFromLibWithCallback(request *DescribeImageFromLibRequest, callback func(response *DescribeImageFromLibResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageFromLibResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageFromLib(request)
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

// DescribeImageFromLibRequest is the request struct for api DescribeImageFromLib
type DescribeImageFromLibRequest struct {
	*requests.RpcRequest
	StartDate   string           `position:"Query" name:"StartDate"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	ImageLibId  requests.Integer `position:"Query" name:"ImageLibId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Id          requests.Integer `position:"Query" name:"Id"`
	TotalCount  requests.Integer `position:"Query" name:"TotalCount"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	EndDate     string           `position:"Query" name:"EndDate"`
}

// DescribeImageFromLibResponse is the response struct for api DescribeImageFromLib
type DescribeImageFromLibResponse struct {
	*responses.BaseResponse
	RequestId        string         `json:"RequestId" xml:"RequestId"`
	PageSize         int            `json:"PageSize" xml:"PageSize"`
	CurrentPage      int            `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount       int            `json:"TotalCount" xml:"TotalCount"`
	ImageFromLibList []ImageFromLib `json:"ImageFromLibList" xml:"ImageFromLibList"`
}

// CreateDescribeImageFromLibRequest creates a request to invoke DescribeImageFromLib API
func CreateDescribeImageFromLibRequest() (request *DescribeImageFromLibRequest) {
	request = &DescribeImageFromLibRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeImageFromLib", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeImageFromLibResponse creates a response to parse from DescribeImageFromLib response
func CreateDescribeImageFromLibResponse() (response *DescribeImageFromLibResponse) {
	response = &DescribeImageFromLibResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
