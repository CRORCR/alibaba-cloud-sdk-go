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

// ListDashboard invokes the airec.ListDashboard API synchronously
// api document: https://help.aliyun.com/api/airec/listdashboard.html
func (client *Client) ListDashboard(request *ListDashboardRequest) (response *ListDashboardResponse, err error) {
	response = CreateListDashboardResponse()
	err = client.DoAction(request, response)
	return
}

// ListDashboardWithChan invokes the airec.ListDashboard API asynchronously
// api document: https://help.aliyun.com/api/airec/listdashboard.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDashboardWithChan(request *ListDashboardRequest) (<-chan *ListDashboardResponse, <-chan error) {
	responseChan := make(chan *ListDashboardResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDashboard(request)
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

// ListDashboardWithCallback invokes the airec.ListDashboard API asynchronously
// api document: https://help.aliyun.com/api/airec/listdashboard.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDashboardWithCallback(request *ListDashboardRequest, callback func(response *ListDashboardResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDashboardResponse
		var err error
		defer close(result)
		response, err = client.ListDashboard(request)
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

// ListDashboardRequest is the request struct for api ListDashboard
type ListDashboardRequest struct {
	*requests.RoaRequest
	TraceId    string           `position:"Query" name:"TraceId"`
	InstanceId string           `position:"Path" name:"InstanceId"`
	EndDate    requests.Integer `position:"Query" name:"EndDate"`
	Size       requests.Integer `position:"Query" name:"Size"`
	SceneId    string           `position:"Query" name:"SceneId"`
	Page       requests.Integer `position:"Query" name:"Page"`
	StartDate  requests.Integer `position:"Query" name:"StartDate"`
}

// ListDashboardResponse is the response struct for api ListDashboard
type ListDashboardResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateListDashboardRequest creates a request to invoke ListDashboard API
func CreateListDashboardRequest() (request *ListDashboardRequest) {
	request = &ListDashboardRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2018-10-12", "ListDashboard", "/openapi/instances/[InstanceId]/dashboard/statistics", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListDashboardResponse creates a response to parse from ListDashboard response
func CreateListDashboardResponse() (response *ListDashboardResponse) {
	response = &ListDashboardResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
