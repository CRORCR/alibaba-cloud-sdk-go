package aegis

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

// CreateOrUpdateJoinRule invokes the aegis.CreateOrUpdateJoinRule API synchronously
// api document: https://help.aliyun.com/api/aegis/createorupdatejoinrule.html
func (client *Client) CreateOrUpdateJoinRule(request *CreateOrUpdateJoinRuleRequest) (response *CreateOrUpdateJoinRuleResponse, err error) {
	response = CreateCreateOrUpdateJoinRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateOrUpdateJoinRuleWithChan invokes the aegis.CreateOrUpdateJoinRule API asynchronously
// api document: https://help.aliyun.com/api/aegis/createorupdatejoinrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrUpdateJoinRuleWithChan(request *CreateOrUpdateJoinRuleRequest) (<-chan *CreateOrUpdateJoinRuleResponse, <-chan error) {
	responseChan := make(chan *CreateOrUpdateJoinRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateOrUpdateJoinRule(request)
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

// CreateOrUpdateJoinRuleWithCallback invokes the aegis.CreateOrUpdateJoinRule API asynchronously
// api document: https://help.aliyun.com/api/aegis/createorupdatejoinrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateOrUpdateJoinRuleWithCallback(request *CreateOrUpdateJoinRuleRequest, callback func(response *CreateOrUpdateJoinRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateOrUpdateJoinRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateOrUpdateJoinRule(request)
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

// CreateOrUpdateJoinRuleRequest is the request struct for api CreateOrUpdateJoinRule
type CreateOrUpdateJoinRuleRequest struct {
	*requests.RpcRequest
	WarnLevel       string           `position:"Query" name:"WarnLevel"`
	DataSourceId2   requests.Integer `position:"Query" name:"DataSourceId2"`
	DataSourceId1   requests.Integer `position:"Query" name:"DataSourceId1"`
	TimeWindow      requests.Integer `position:"Query" name:"TimeWindow"`
	Description     string           `position:"Query" name:"Description"`
	RuleName        string           `position:"Query" name:"RuleName"`
	Expression2     string           `position:"Query" name:"Expression2"`
	Expression1     string           `position:"Query" name:"Expression1"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	StatisticsRules string           `position:"Query" name:"StatisticsRules"`
	JoinRelation    string           `position:"Query" name:"JoinRelation"`
	RuleId          requests.Integer `position:"Query" name:"RuleId"`
	Actions         string           `position:"Query" name:"Actions"`
}

// CreateOrUpdateJoinRuleResponse is the response struct for api CreateOrUpdateJoinRule
type CreateOrUpdateJoinRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateOrUpdateJoinRuleRequest creates a request to invoke CreateOrUpdateJoinRule API
func CreateCreateOrUpdateJoinRuleRequest() (request *CreateOrUpdateJoinRuleRequest) {
	request = &CreateOrUpdateJoinRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "CreateOrUpdateJoinRule", "vipaegis", "openAPI")
	return
}

// CreateCreateOrUpdateJoinRuleResponse creates a response to parse from CreateOrUpdateJoinRule response
func CreateCreateOrUpdateJoinRuleResponse() (response *CreateOrUpdateJoinRuleResponse) {
	response = &CreateOrUpdateJoinRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
