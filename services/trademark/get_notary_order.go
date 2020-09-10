package trademark

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

// GetNotaryOrder invokes the trademark.GetNotaryOrder API synchronously
// api document: https://help.aliyun.com/api/trademark/getnotaryorder.html
func (client *Client) GetNotaryOrder(request *GetNotaryOrderRequest) (response *GetNotaryOrderResponse, err error) {
	response = CreateGetNotaryOrderResponse()
	err = client.DoAction(request, response)
	return
}

// GetNotaryOrderWithChan invokes the trademark.GetNotaryOrder API asynchronously
// api document: https://help.aliyun.com/api/trademark/getnotaryorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNotaryOrderWithChan(request *GetNotaryOrderRequest) (<-chan *GetNotaryOrderResponse, <-chan error) {
	responseChan := make(chan *GetNotaryOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNotaryOrder(request)
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

// GetNotaryOrderWithCallback invokes the trademark.GetNotaryOrder API asynchronously
// api document: https://help.aliyun.com/api/trademark/getnotaryorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNotaryOrderWithCallback(request *GetNotaryOrderRequest, callback func(response *GetNotaryOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNotaryOrderResponse
		var err error
		defer close(result)
		response, err = client.GetNotaryOrder(request)
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

// GetNotaryOrderRequest is the request struct for api GetNotaryOrder
type GetNotaryOrderRequest struct {
	*requests.RpcRequest
	NotaryOrderId requests.Integer `position:"Query" name:"NotaryOrderId"`
}

// GetNotaryOrderResponse is the response struct for api GetNotaryOrder
type GetNotaryOrderResponse struct {
	*responses.BaseResponse
	RequestId                   string  `json:"RequestId" xml:"RequestId"`
	NotaryOrderId               int64   `json:"NotaryOrderId" xml:"NotaryOrderId"`
	AliyunOrderId               string  `json:"AliyunOrderId" xml:"AliyunOrderId"`
	TmRegisterNo                string  `json:"TmRegisterNo" xml:"TmRegisterNo"`
	TmName                      string  `json:"TmName" xml:"TmName"`
	TmImage                     string  `json:"TmImage" xml:"TmImage"`
	TmClassification            string  `json:"TmClassification" xml:"TmClassification"`
	OrderPrice                  float64 `json:"OrderPrice" xml:"OrderPrice"`
	NotaryStatus                int     `json:"NotaryStatus" xml:"NotaryStatus"`
	Type                        string  `json:"Type" xml:"Type"`
	Name                        string  `json:"Name" xml:"Name"`
	Phone                       string  `json:"Phone" xml:"Phone"`
	SellerCompanyName           string  `json:"SellerCompanyName" xml:"SellerCompanyName"`
	BusinessLicenseId           string  `json:"BusinessLicenseId" xml:"BusinessLicenseId"`
	LegalPersonName             string  `json:"LegalPersonName" xml:"LegalPersonName"`
	LegalPersonPhone            string  `json:"LegalPersonPhone" xml:"LegalPersonPhone"`
	LegalPersonIdCard           string  `json:"LegalPersonIdCard" xml:"LegalPersonIdCard"`
	CompanyContactName          string  `json:"CompanyContactName" xml:"CompanyContactName"`
	CompanyContactPhone         string  `json:"CompanyContactPhone" xml:"CompanyContactPhone"`
	TmRegisterCertificate       string  `json:"TmRegisterCertificate" xml:"TmRegisterCertificate"`
	BusinessLicense             string  `json:"BusinessLicense" xml:"BusinessLicense"`
	TmAcceptCertificate         string  `json:"TmAcceptCertificate" xml:"TmAcceptCertificate"`
	SellerFrontOfIdCard         string  `json:"SellerFrontOfIdCard" xml:"SellerFrontOfIdCard"`
	SellerBackOfIdCard          string  `json:"SellerBackOfIdCard" xml:"SellerBackOfIdCard"`
	TmRegisterChangeCertificate string  `json:"TmRegisterChangeCertificate" xml:"TmRegisterChangeCertificate"`
	ReceiverName                string  `json:"ReceiverName" xml:"ReceiverName"`
	ReceiverAddress             string  `json:"ReceiverAddress" xml:"ReceiverAddress"`
	ReceiverPhone               string  `json:"ReceiverPhone" xml:"ReceiverPhone"`
	ReceiverPostalCode          string  `json:"ReceiverPostalCode" xml:"ReceiverPostalCode"`
	OrderDate                   int64   `json:"OrderDate" xml:"OrderDate"`
	NotaryAcceptDate            int64   `json:"NotaryAcceptDate" xml:"NotaryAcceptDate"`
	NotarySucceedDate           int64   `json:"NotarySucceedDate" xml:"NotarySucceedDate"`
	NotaryFailedDate            int64   `json:"NotaryFailedDate" xml:"NotaryFailedDate"`
	NotaryFailedReason          string  `json:"NotaryFailedReason" xml:"NotaryFailedReason"`
	NotaryCertificate           string  `json:"NotaryCertificate" xml:"NotaryCertificate"`
	Success                     bool    `json:"Success" xml:"Success"`
	ErrorMsg                    string  `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode                   string  `json:"ErrorCode" xml:"ErrorCode"`
	BizId                       string  `json:"BizId" xml:"BizId"`
	NotaryType                  int     `json:"NotaryType" xml:"NotaryType"`
	NotaryPlatformName          string  `json:"NotaryPlatformName" xml:"NotaryPlatformName"`
	ApplyPostStatus             int     `json:"ApplyPostStatus" xml:"ApplyPostStatus"`
	NotaryPostReceipt           string  `json:"NotaryPostReceipt" xml:"NotaryPostReceipt"`
}

// CreateGetNotaryOrderRequest creates a request to invoke GetNotaryOrder API
func CreateGetNotaryOrderRequest() (request *GetNotaryOrderRequest) {
	request = &GetNotaryOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Trademark", "2018-07-24", "GetNotaryOrder", "trademark", "openAPI")
	return
}

// CreateGetNotaryOrderResponse creates a response to parse from GetNotaryOrder response
func CreateGetNotaryOrderResponse() (response *GetNotaryOrderResponse) {
	response = &GetNotaryOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
