package devops_rdc

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

// GetPipleineLatestInstanceStatus invokes the devops_rdc.GetPipleineLatestInstanceStatus API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/getpipleinelatestinstancestatus.html
func (client *Client) GetPipleineLatestInstanceStatus(request *GetPipleineLatestInstanceStatusRequest) (response *GetPipleineLatestInstanceStatusResponse, err error) {
	response = CreateGetPipleineLatestInstanceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetPipleineLatestInstanceStatusWithChan invokes the devops_rdc.GetPipleineLatestInstanceStatus API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/getpipleinelatestinstancestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPipleineLatestInstanceStatusWithChan(request *GetPipleineLatestInstanceStatusRequest) (<-chan *GetPipleineLatestInstanceStatusResponse, <-chan error) {
	responseChan := make(chan *GetPipleineLatestInstanceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPipleineLatestInstanceStatus(request)
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

// GetPipleineLatestInstanceStatusWithCallback invokes the devops_rdc.GetPipleineLatestInstanceStatus API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/getpipleinelatestinstancestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetPipleineLatestInstanceStatusWithCallback(request *GetPipleineLatestInstanceStatusRequest, callback func(response *GetPipleineLatestInstanceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPipleineLatestInstanceStatusResponse
		var err error
		defer close(result)
		response, err = client.GetPipleineLatestInstanceStatus(request)
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

// GetPipleineLatestInstanceStatusRequest is the request struct for api GetPipleineLatestInstanceStatus
type GetPipleineLatestInstanceStatusRequest struct {
	*requests.RpcRequest
	UserPk     string           `position:"Body" name:"UserPk"`
	OrgId      string           `position:"Query" name:"OrgId"`
	PipelineId requests.Integer `position:"Query" name:"PipelineId"`
}

// GetPipleineLatestInstanceStatusResponse is the response struct for api GetPipleineLatestInstanceStatus
type GetPipleineLatestInstanceStatusResponse struct {
	*responses.BaseResponse
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Object       Object `json:"Object" xml:"Object"`
}

// CreateGetPipleineLatestInstanceStatusRequest creates a request to invoke GetPipleineLatestInstanceStatus API
func CreateGetPipleineLatestInstanceStatusRequest() (request *GetPipleineLatestInstanceStatusRequest) {
	request = &GetPipleineLatestInstanceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "GetPipleineLatestInstanceStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPipleineLatestInstanceStatusResponse creates a response to parse from GetPipleineLatestInstanceStatus response
func CreateGetPipleineLatestInstanceStatusResponse() (response *GetPipleineLatestInstanceStatusResponse) {
	response = &GetPipleineLatestInstanceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
