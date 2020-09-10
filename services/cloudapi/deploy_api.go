package cloudapi

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

// DeployApi invokes the cloudapi.DeployApi API synchronously
// api document: https://help.aliyun.com/api/cloudapi/deployapi.html
func (client *Client) DeployApi(request *DeployApiRequest) (response *DeployApiResponse, err error) {
	response = CreateDeployApiResponse()
	err = client.DoAction(request, response)
	return
}

// DeployApiWithChan invokes the cloudapi.DeployApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deployapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeployApiWithChan(request *DeployApiRequest) (<-chan *DeployApiResponse, <-chan error) {
	responseChan := make(chan *DeployApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeployApi(request)
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

// DeployApiWithCallback invokes the cloudapi.DeployApi API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/deployapi.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeployApiWithCallback(request *DeployApiRequest, callback func(response *DeployApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeployApiResponse
		var err error
		defer close(result)
		response, err = client.DeployApi(request)
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

// DeployApiRequest is the request struct for api DeployApi
type DeployApiRequest struct {
	*requests.RpcRequest
	StageName          string `position:"Query" name:"StageName"`
	GroupId            string `position:"Query" name:"GroupId"`
	Description        string `position:"Query" name:"Description"`
	ResourceOwnerToken string `position:"Query" name:"ResourceOwnerToken"`
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	ApiId              string `position:"Query" name:"ApiId"`
}

// DeployApiResponse is the response struct for api DeployApi
type DeployApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeployApiRequest creates a request to invoke DeployApi API
func CreateDeployApiRequest() (request *DeployApiRequest) {
	request = &DeployApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeployApi", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeployApiResponse creates a response to parse from DeployApi response
func CreateDeployApiResponse() (response *DeployApiResponse) {
	response = &DeployApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
