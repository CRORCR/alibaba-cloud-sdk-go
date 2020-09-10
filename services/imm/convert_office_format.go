package imm

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

// ConvertOfficeFormat invokes the imm.ConvertOfficeFormat API synchronously
// api document: https://help.aliyun.com/api/imm/convertofficeformat.html
func (client *Client) ConvertOfficeFormat(request *ConvertOfficeFormatRequest) (response *ConvertOfficeFormatResponse, err error) {
	response = CreateConvertOfficeFormatResponse()
	err = client.DoAction(request, response)
	return
}

// ConvertOfficeFormatWithChan invokes the imm.ConvertOfficeFormat API asynchronously
// api document: https://help.aliyun.com/api/imm/convertofficeformat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConvertOfficeFormatWithChan(request *ConvertOfficeFormatRequest) (<-chan *ConvertOfficeFormatResponse, <-chan error) {
	responseChan := make(chan *ConvertOfficeFormatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConvertOfficeFormat(request)
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

// ConvertOfficeFormatWithCallback invokes the imm.ConvertOfficeFormat API asynchronously
// api document: https://help.aliyun.com/api/imm/convertofficeformat.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConvertOfficeFormatWithCallback(request *ConvertOfficeFormatRequest, callback func(response *ConvertOfficeFormatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConvertOfficeFormatResponse
		var err error
		defer close(result)
		response, err = client.ConvertOfficeFormat(request)
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

// ConvertOfficeFormatRequest is the request struct for api ConvertOfficeFormat
type ConvertOfficeFormatRequest struct {
	*requests.RpcRequest
	SrcType        string           `position:"Query" name:"SrcType"`
	Project        string           `position:"Query" name:"Project"`
	PdfVector      requests.Boolean `position:"Query" name:"PdfVector"`
	Password       string           `position:"Query" name:"Password"`
	StartPage      requests.Integer `position:"Query" name:"StartPage"`
	FitToPagesWide requests.Boolean `position:"Query" name:"FitToPagesWide"`
	TgtFilePrefix  string           `position:"Query" name:"TgtFilePrefix"`
	ModelId        string           `position:"Query" name:"ModelId"`
	MaxSheetRow    requests.Integer `position:"Query" name:"MaxSheetRow"`
	MaxSheetCount  requests.Integer `position:"Query" name:"MaxSheetCount"`
	EndPage        requests.Integer `position:"Query" name:"EndPage"`
	TgtFileSuffix  string           `position:"Query" name:"TgtFileSuffix"`
	SheetOnePage   requests.Boolean `position:"Query" name:"SheetOnePage"`
	MaxSheetCol    requests.Integer `position:"Query" name:"MaxSheetCol"`
	TgtType        string           `position:"Query" name:"TgtType"`
	Hidecomments   requests.Boolean `position:"Query" name:"Hidecomments"`
	FitToPagesTall requests.Boolean `position:"Query" name:"FitToPagesTall"`
	SrcUri         string           `position:"Query" name:"SrcUri"`
	TgtFilePages   string           `position:"Query" name:"TgtFilePages"`
	TgtUri         string           `position:"Query" name:"TgtUri"`
}

// ConvertOfficeFormatResponse is the response struct for api ConvertOfficeFormat
type ConvertOfficeFormatResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PageCount int    `json:"PageCount" xml:"PageCount"`
}

// CreateConvertOfficeFormatRequest creates a request to invoke ConvertOfficeFormat API
func CreateConvertOfficeFormatRequest() (request *ConvertOfficeFormatRequest) {
	request = &ConvertOfficeFormatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ConvertOfficeFormat", "", "")
	request.Method = requests.POST
	return
}

// CreateConvertOfficeFormatResponse creates a response to parse from ConvertOfficeFormat response
func CreateConvertOfficeFormatResponse() (response *ConvertOfficeFormatResponse) {
	response = &ConvertOfficeFormatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
