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

// DeleteSipAgentGroup invokes the cloudcallcenter.DeleteSipAgentGroup API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletesipagentgroup.html
func (client *Client) DeleteSipAgentGroup(request *DeleteSipAgentGroupRequest) (response *DeleteSipAgentGroupResponse, err error) {
	response = CreateDeleteSipAgentGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSipAgentGroupWithChan invokes the cloudcallcenter.DeleteSipAgentGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletesipagentgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSipAgentGroupWithChan(request *DeleteSipAgentGroupRequest) (<-chan *DeleteSipAgentGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteSipAgentGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSipAgentGroup(request)
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

// DeleteSipAgentGroupWithCallback invokes the cloudcallcenter.DeleteSipAgentGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletesipagentgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSipAgentGroupWithCallback(request *DeleteSipAgentGroupRequest, callback func(response *DeleteSipAgentGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSipAgentGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteSipAgentGroup(request)
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

// DeleteSipAgentGroupRequest is the request struct for api DeleteSipAgentGroup
type DeleteSipAgentGroupRequest struct {
	*requests.RpcRequest
	Provider string           `position:"Query" name:"Provider"`
	Id       requests.Integer `position:"Query" name:"Id"`
}

// DeleteSipAgentGroupResponse is the response struct for api DeleteSipAgentGroup
type DeleteSipAgentGroupResponse struct {
	*responses.BaseResponse
	RequestId      string     `json:"RequestId" xml:"RequestId"`
	Success        bool       `json:"Success" xml:"Success"`
	Code           string     `json:"Code" xml:"Code"`
	Message        string     `json:"Message" xml:"Message"`
	HttpStatusCode int        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Provider       string     `json:"Provider" xml:"Provider"`
	SipAgents      []SipAgent `json:"SipAgents" xml:"SipAgents"`
}

// CreateDeleteSipAgentGroupRequest creates a request to invoke DeleteSipAgentGroup API
func CreateDeleteSipAgentGroupRequest() (request *DeleteSipAgentGroupRequest) {
	request = &DeleteSipAgentGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DeleteSipAgentGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteSipAgentGroupResponse creates a response to parse from DeleteSipAgentGroup response
func CreateDeleteSipAgentGroupResponse() (response *DeleteSipAgentGroupResponse) {
	response = &DeleteSipAgentGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
