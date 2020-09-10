package drds

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

// DescribeDrdsDbInstance invokes the drds.DescribeDrdsDbInstance API synchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbinstance.html
func (client *Client) DescribeDrdsDbInstance(request *DescribeDrdsDbInstanceRequest) (response *DescribeDrdsDbInstanceResponse, err error) {
	response = CreateDescribeDrdsDbInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsDbInstanceWithChan invokes the drds.DescribeDrdsDbInstance API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDbInstanceWithChan(request *DescribeDrdsDbInstanceRequest) (<-chan *DescribeDrdsDbInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsDbInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsDbInstance(request)
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

// DescribeDrdsDbInstanceWithCallback invokes the drds.DescribeDrdsDbInstance API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDbInstanceWithCallback(request *DescribeDrdsDbInstanceRequest, callback func(response *DescribeDrdsDbInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsDbInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsDbInstance(request)
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

// DescribeDrdsDbInstanceRequest is the request struct for api DescribeDrdsDbInstance
type DescribeDrdsDbInstanceRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
	DbInstanceId   string `position:"Query" name:"DbInstanceId"`
}

// DescribeDrdsDbInstanceResponse is the response struct for api DescribeDrdsDbInstance
type DescribeDrdsDbInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string                             `json:"RequestId" xml:"RequestId"`
	Success    bool                               `json:"Success" xml:"Success"`
	DbInstance DbInstanceInDescribeDrdsDbInstance `json:"DbInstance" xml:"DbInstance"`
}

// CreateDescribeDrdsDbInstanceRequest creates a request to invoke DescribeDrdsDbInstance API
func CreateDescribeDrdsDbInstanceRequest() (request *DescribeDrdsDbInstanceRequest) {
	request = &DescribeDrdsDbInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsDbInstance", "Drds", "openAPI")
	return
}

// CreateDescribeDrdsDbInstanceResponse creates a response to parse from DescribeDrdsDbInstance response
func CreateDescribeDrdsDbInstanceResponse() (response *DescribeDrdsDbInstanceResponse) {
	response = &DescribeDrdsDbInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
