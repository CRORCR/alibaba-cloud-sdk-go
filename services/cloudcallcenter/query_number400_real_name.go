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

// QueryNumber400RealName invokes the cloudcallcenter.QueryNumber400RealName API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/querynumber400realname.html
func (client *Client) QueryNumber400RealName(request *QueryNumber400RealNameRequest) (response *QueryNumber400RealNameResponse, err error) {
	response = CreateQueryNumber400RealNameResponse()
	err = client.DoAction(request, response)
	return
}

// QueryNumber400RealNameWithChan invokes the cloudcallcenter.QueryNumber400RealName API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/querynumber400realname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryNumber400RealNameWithChan(request *QueryNumber400RealNameRequest) (<-chan *QueryNumber400RealNameResponse, <-chan error) {
	responseChan := make(chan *QueryNumber400RealNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryNumber400RealName(request)
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

// QueryNumber400RealNameWithCallback invokes the cloudcallcenter.QueryNumber400RealName API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/querynumber400realname.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryNumber400RealNameWithCallback(request *QueryNumber400RealNameRequest, callback func(response *QueryNumber400RealNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryNumber400RealNameResponse
		var err error
		defer close(result)
		response, err = client.QueryNumber400RealName(request)
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

// QueryNumber400RealNameRequest is the request struct for api QueryNumber400RealName
type QueryNumber400RealNameRequest struct {
	*requests.RpcRequest
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Status      string           `position:"Query" name:"Status"`
}

// QueryNumber400RealNameResponse is the response struct for api QueryNumber400RealName
type QueryNumber400RealNameResponse struct {
	*responses.BaseResponse
	RequestId      string                       `json:"RequestId" xml:"RequestId"`
	Success        bool                         `json:"Success" xml:"Success"`
	Code           string                       `json:"Code" xml:"Code"`
	Message        string                       `json:"Message" xml:"Message"`
	HttpStatusCode int                          `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           DataInQueryNumber400RealName `json:"Data" xml:"Data"`
}

// CreateQueryNumber400RealNameRequest creates a request to invoke QueryNumber400RealName API
func CreateQueryNumber400RealNameRequest() (request *QueryNumber400RealNameRequest) {
	request = &QueryNumber400RealNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "QueryNumber400RealName", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryNumber400RealNameResponse creates a response to parse from QueryNumber400RealName response
func CreateQueryNumber400RealNameResponse() (response *QueryNumber400RealNameResponse) {
	response = &QueryNumber400RealNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
