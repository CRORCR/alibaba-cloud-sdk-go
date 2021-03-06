package airec

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

// QueryDataMessage invokes the airec.QueryDataMessage API synchronously
// api document: https://help.aliyun.com/api/airec/querydatamessage.html
func (client *Client) QueryDataMessage(request *QueryDataMessageRequest) (response *QueryDataMessageResponse, err error) {
	response = CreateQueryDataMessageResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDataMessageWithChan invokes the airec.QueryDataMessage API asynchronously
// api document: https://help.aliyun.com/api/airec/querydatamessage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDataMessageWithChan(request *QueryDataMessageRequest) (<-chan *QueryDataMessageResponse, <-chan error) {
	responseChan := make(chan *QueryDataMessageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDataMessage(request)
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

// QueryDataMessageWithCallback invokes the airec.QueryDataMessage API asynchronously
// api document: https://help.aliyun.com/api/airec/querydatamessage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDataMessageWithCallback(request *QueryDataMessageRequest, callback func(response *QueryDataMessageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDataMessageResponse
		var err error
		defer close(result)
		response, err = client.QueryDataMessage(request)
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

// QueryDataMessageRequest is the request struct for api QueryDataMessage
type QueryDataMessageRequest struct {
	*requests.RoaRequest
	TraceId    string           `position:"Query" name:"TraceId"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	UserType   string           `position:"Query" name:"UserType"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	UserId     string           `position:"Query" name:"UserId"`
	ItemId     string           `position:"Query" name:"ItemId"`
	InstanceId string           `position:"Path" name:"InstanceId"`
	ItemType   string           `position:"Query" name:"ItemType"`
	CmdType    string           `position:"Query" name:"CmdType"`
	Size       requests.Integer `position:"Query" name:"Size"`
	SceneId    string           `position:"Query" name:"SceneId"`
	BhvType    string           `position:"Query" name:"BhvType"`
	Page       requests.Integer `position:"Query" name:"Page"`
	Table      string           `position:"Path" name:"Table"`
}

// QueryDataMessageResponse is the response struct for api QueryDataMessage
type QueryDataMessageResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Code      string                 `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	Result    map[string]interface{} `json:"Result" xml:"Result"`
}

// CreateQueryDataMessageRequest creates a request to invoke QueryDataMessage API
func CreateQueryDataMessageRequest() (request *QueryDataMessageRequest) {
	request = &QueryDataMessageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2018-10-12", "QueryDataMessage", "/openapi/instances/[InstanceId]/tables/[Table]/data-message", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateQueryDataMessageResponse creates a response to parse from QueryDataMessage response
func CreateQueryDataMessageResponse() (response *QueryDataMessageResponse) {
	response = &QueryDataMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
