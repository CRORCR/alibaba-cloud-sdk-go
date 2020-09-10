package teslamaxcompute

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

// GetInstancesStatusCount invokes the teslamaxcompute.GetInstancesStatusCount API synchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/getinstancesstatuscount.html
func (client *Client) GetInstancesStatusCount(request *GetInstancesStatusCountRequest) (response *GetInstancesStatusCountResponse, err error) {
	response = CreateGetInstancesStatusCountResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstancesStatusCountWithChan invokes the teslamaxcompute.GetInstancesStatusCount API asynchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/getinstancesstatuscount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstancesStatusCountWithChan(request *GetInstancesStatusCountRequest) (<-chan *GetInstancesStatusCountResponse, <-chan error) {
	responseChan := make(chan *GetInstancesStatusCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstancesStatusCount(request)
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

// GetInstancesStatusCountWithCallback invokes the teslamaxcompute.GetInstancesStatusCount API asynchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/getinstancesstatuscount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetInstancesStatusCountWithCallback(request *GetInstancesStatusCountRequest, callback func(response *GetInstancesStatusCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstancesStatusCountResponse
		var err error
		defer close(result)
		response, err = client.GetInstancesStatusCount(request)
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

// GetInstancesStatusCountRequest is the request struct for api GetInstancesStatusCount
type GetInstancesStatusCountRequest struct {
	*requests.RpcRequest
	Cluster   string `position:"Query" name:"Cluster"`
	QuotaId   string `position:"Query" name:"QuotaId"`
	Region    string `position:"Query" name:"Region"`
	QuotaName string `position:"Query" name:"QuotaName"`
}

// GetInstancesStatusCountResponse is the response struct for api GetInstancesStatusCount
type GetInstancesStatusCountResponse struct {
	*responses.BaseResponse
	Code      int        `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateGetInstancesStatusCountRequest creates a request to invoke GetInstancesStatusCount API
func CreateGetInstancesStatusCountRequest() (request *GetInstancesStatusCountRequest) {
	request = &GetInstancesStatusCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetInstancesStatusCount", "teslamaxcompute", "openAPI")
	return
}

// CreateGetInstancesStatusCountResponse creates a response to parse from GetInstancesStatusCount response
func CreateGetInstancesStatusCountResponse() (response *GetInstancesStatusCountResponse) {
	response = &GetInstancesStatusCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
