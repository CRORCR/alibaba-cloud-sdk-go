package videosearch

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

// TaskInfo is a nested struct in videosearch response
type TaskInfo struct {
	AnalysisUseTime  int64   `json:"AnalysisUseTime" xml:"AnalysisUseTime"`
	Duration         float64 `json:"Duration" xml:"Duration"`
	ProcessResultOss string  `json:"ProcessResultOss" xml:"ProcessResultOss"`
	Resolution       string  `json:"Resolution" xml:"Resolution"`
	Status           int     `json:"Status" xml:"Status"`
	SubmitTime       int64   `json:"SubmitTime" xml:"SubmitTime"`
	FinishTime       int64   `json:"FinishTime" xml:"FinishTime"`
	TaskId           int64   `json:"TaskId" xml:"TaskId"`
	ErrorInfo        string  `json:"ErrorInfo" xml:"ErrorInfo"`
	StorageInfo      int     `json:"StorageInfo" xml:"StorageInfo"`
	Description      string  `json:"Description" xml:"Description"`
	VideoId          string  `json:"VideoId" xml:"VideoId"`
	VideoTags        string  `json:"VideoTags" xml:"VideoTags"`
	VideoUrl         string  `json:"VideoUrl" xml:"VideoUrl"`
	QueryTags        string  `json:"QueryTags" xml:"QueryTags"`
}
