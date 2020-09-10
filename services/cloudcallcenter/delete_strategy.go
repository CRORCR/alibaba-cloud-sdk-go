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

// DeleteStrategy invokes the cloudcallcenter.DeleteStrategy API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletestrategy.html
func (client *Client) DeleteStrategy(request *DeleteStrategyRequest) (response *DeleteStrategyResponse, err error) {
	response = CreateDeleteStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteStrategyWithChan invokes the cloudcallcenter.DeleteStrategy API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletestrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteStrategyWithChan(request *DeleteStrategyRequest) (<-chan *DeleteStrategyResponse, <-chan error) {
	responseChan := make(chan *DeleteStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteStrategy(request)
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

// DeleteStrategyWithCallback invokes the cloudcallcenter.DeleteStrategy API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletestrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteStrategyWithCallback(request *DeleteStrategyRequest, callback func(response *DeleteStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteStrategyResponse
		var err error
		defer close(result)
		response, err = client.DeleteStrategy(request)
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

// DeleteStrategyRequest is the request struct for api DeleteStrategy
type DeleteStrategyRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	StrategyId string `position:"Query" name:"StrategyId"`
}

// DeleteStrategyResponse is the response struct for api DeleteStrategy
type DeleteStrategyResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDeleteStrategyRequest creates a request to invoke DeleteStrategy API
func CreateDeleteStrategyRequest() (request *DeleteStrategyRequest) {
	request = &DeleteStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DeleteStrategy", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteStrategyResponse creates a response to parse from DeleteStrategy response
func CreateDeleteStrategyResponse() (response *DeleteStrategyResponse) {
	response = &DeleteStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
