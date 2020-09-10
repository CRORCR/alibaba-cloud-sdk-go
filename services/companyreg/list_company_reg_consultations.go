package companyreg

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

// ListCompanyRegConsultations invokes the companyreg.ListCompanyRegConsultations API synchronously
// api document: https://help.aliyun.com/api/companyreg/listcompanyregconsultations.html
func (client *Client) ListCompanyRegConsultations(request *ListCompanyRegConsultationsRequest) (response *ListCompanyRegConsultationsResponse, err error) {
	response = CreateListCompanyRegConsultationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCompanyRegConsultationsWithChan invokes the companyreg.ListCompanyRegConsultations API asynchronously
// api document: https://help.aliyun.com/api/companyreg/listcompanyregconsultations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCompanyRegConsultationsWithChan(request *ListCompanyRegConsultationsRequest) (<-chan *ListCompanyRegConsultationsResponse, <-chan error) {
	responseChan := make(chan *ListCompanyRegConsultationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCompanyRegConsultations(request)
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

// ListCompanyRegConsultationsWithCallback invokes the companyreg.ListCompanyRegConsultations API asynchronously
// api document: https://help.aliyun.com/api/companyreg/listcompanyregconsultations.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCompanyRegConsultationsWithCallback(request *ListCompanyRegConsultationsRequest, callback func(response *ListCompanyRegConsultationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCompanyRegConsultationsResponse
		var err error
		defer close(result)
		response, err = client.ListCompanyRegConsultations(request)
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

// ListCompanyRegConsultationsRequest is the request struct for api ListCompanyRegConsultations
type ListCompanyRegConsultationsRequest struct {
	*requests.RpcRequest
	EndGmtCreate   requests.Integer `position:"Query" name:"EndGmtCreate"`
	City           string           `position:"Query" name:"City"`
	PageNum        requests.Integer `position:"Query" name:"PageNum"`
	BizCode        string           `position:"Query" name:"BizCode"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	BizId          string           `position:"Query" name:"BizId"`
	StartGmtCreate requests.Integer `position:"Query" name:"StartGmtCreate"`
}

// ListCompanyRegConsultationsResponse is the response struct for api ListCompanyRegConsultations
type ListCompanyRegConsultationsResponse struct {
	*responses.BaseResponse
	RequestId      string                            `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int                               `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int                               `json:"CurrentPageNum" xml:"CurrentPageNum"`
	PageSize       int                               `json:"PageSize" xml:"PageSize"`
	TotalPageNum   int                               `json:"TotalPageNum" xml:"TotalPageNum"`
	PrePage        bool                              `json:"PrePage" xml:"PrePage"`
	NextPage       bool                              `json:"NextPage" xml:"NextPage"`
	Data           DataInListCompanyRegConsultations `json:"Data" xml:"Data"`
}

// CreateListCompanyRegConsultationsRequest creates a request to invoke ListCompanyRegConsultations API
func CreateListCompanyRegConsultationsRequest() (request *ListCompanyRegConsultationsRequest) {
	request = &ListCompanyRegConsultationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "ListCompanyRegConsultations", "companyreg", "openAPI")
	return
}

// CreateListCompanyRegConsultationsResponse creates a response to parse from ListCompanyRegConsultations response
func CreateListCompanyRegConsultationsResponse() (response *ListCompanyRegConsultationsResponse) {
	response = &ListCompanyRegConsultationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
