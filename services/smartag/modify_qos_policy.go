package smartag

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

// ModifyQosPolicy invokes the smartag.ModifyQosPolicy API synchronously
// api document: https://help.aliyun.com/api/smartag/modifyqospolicy.html
func (client *Client) ModifyQosPolicy(request *ModifyQosPolicyRequest) (response *ModifyQosPolicyResponse, err error) {
	response = CreateModifyQosPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyQosPolicyWithChan invokes the smartag.ModifyQosPolicy API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifyqospolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyQosPolicyWithChan(request *ModifyQosPolicyRequest) (<-chan *ModifyQosPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyQosPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyQosPolicy(request)
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

// ModifyQosPolicyWithCallback invokes the smartag.ModifyQosPolicy API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifyqospolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyQosPolicyWithCallback(request *ModifyQosPolicyRequest, callback func(response *ModifyQosPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyQosPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyQosPolicy(request)
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

// ModifyQosPolicyRequest is the request struct for api ModifyQosPolicy
type ModifyQosPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourcePortRange      string           `position:"Query" name:"SourcePortRange"`
	QosPolicyId          string           `position:"Query" name:"QosPolicyId"`
	SourceCidr           string           `position:"Query" name:"SourceCidr"`
	Description          string           `position:"Query" name:"Description"`
	StartTime            string           `position:"Query" name:"StartTime"`
	DestCidr             string           `position:"Query" name:"DestCidr"`
	QosId                string           `position:"Query" name:"QosId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol           string           `position:"Query" name:"IpProtocol"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Priority             requests.Integer `position:"Query" name:"Priority"`
	DestPortRange        string           `position:"Query" name:"DestPortRange"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyQosPolicyResponse is the response struct for api ModifyQosPolicy
type ModifyQosPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyQosPolicyRequest creates a request to invoke ModifyQosPolicy API
func CreateModifyQosPolicyRequest() (request *ModifyQosPolicyRequest) {
	request = &ModifyQosPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifyQosPolicy", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyQosPolicyResponse creates a response to parse from ModifyQosPolicy response
func CreateModifyQosPolicyResponse() (response *ModifyQosPolicyResponse) {
	response = &ModifyQosPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
