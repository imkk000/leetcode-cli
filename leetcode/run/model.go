package run

const (
	FetchResultStateSuccess = "SUCCESS"
	FetchResultRetryLimit   = 150
)

type RequestSubmissionBody struct {
	QuestionId string `json:"question_id"`
	Lang       string `json:"lang"`
	TypedCode  string `json:"typed_code"`
}

type ResponseSubmissionResultBody struct {
	SubmissionId int64 `json:"submission_id"`
}

type ResponseRunResultBody struct {
	InterpretId string `json:"interpret_id"`
	TestCase    string `json:"test_case"`
}

type ResponseFetchSubmissionBody struct {
	StatusCode        int         `json:"status_code"`
	Lang              string      `json:"lang"`
	RunSuccess        bool        `json:"run_success"`
	StatusRuntime     string      `json:"status_runtime"`
	Memory            int         `json:"memory"`
	QuestionId        string      `json:"question_id"`
	ElapsedTime       int         `json:"elapsed_time"`
	CompareResult     string      `json:"compare_result"`
	CodeOutput        string      `json:"code_output"`
	StdOutput         string      `json:"std_output"`
	LastTestcase      string      `json:"last_testcase"`
	ExpectedOutput    string      `json:"expected_output"`
	TaskFinishTime    int64       `json:"task_finish_time"`
	TotalCorrect      int         `json:"total_correct"`
	TotalTestcases    int         `json:"total_testcases"`
	RuntimePercentile interface{} `json:"runtime_percentile"`
	StatusMemory      string      `json:"status_memory"`
	MemoryPercentile  interface{} `json:"memory_percentile"`
	PrettyLang        string      `json:"pretty_lang"`
	SubmissionId      string      `json:"submission_id"`
	InputFormatted    string      `json:"input_formatted"`
	Input             string      `json:"input"`
	StatusMsg         string      `json:"status_msg"`
	State             string      `json:"state"`
}

type ResponseFetchRunResultBody struct {
	StatusCode             int           `json:"status_code"`
	Lang                   string        `json:"lang"`
	RunSuccess             bool          `json:"run_success"`
	StatusRuntime          string        `json:"status_runtime"`
	Memory                 int           `json:"memory"`
	CodeAnswer             []string      `json:"code_answer"`
	CodeOutput             []interface{} `json:"code_output"`
	StdOutput              []string      `json:"std_output"`
	ElapsedTime            int           `json:"elapsed_time"`
	TaskFinishTime         int64         `json:"task_finish_time"`
	ExpectedStatusCode     int           `json:"expected_status_code"`
	ExpectedLang           string        `json:"expected_lang"`
	ExpectedRunSuccess     bool          `json:"expected_run_success"`
	ExpectedStatusRuntime  string        `json:"expected_status_runtime"`
	ExpectedMemory         int           `json:"expected_memory"`
	ExpectedDisplayRuntime string        `json:"expected_display_runtime"`
	ExpectedCodeAnswer     []string      `json:"expected_code_answer"`
	ExpectedCodeOutput     []interface{} `json:"expected_code_output"`
	ExpectedStdOutput      []string      `json:"expected_std_output"`
	ExpectedElapsedTime    int           `json:"expected_elapsed_time"`
	ExpectedTaskFinishTime int64         `json:"expected_task_finish_time"`
	CorrectAnswer          bool          `json:"correct_answer"`
	CompareResult          string        `json:"compare_result"`
	TotalCorrect           int           `json:"total_correct"`
	TotalTestcases         int           `json:"total_testcases"`
	RuntimePercentile      interface{}   `json:"runtime_percentile"`
	StatusMemory           string        `json:"status_memory"`
	MemoryPercentile       interface{}   `json:"memory_percentile"`
	PrettyLang             string        `json:"pretty_lang"`
	SubmissionId           string        `json:"submission_id"`
	StatusMsg              string        `json:"status_msg"`
	State                  string        `json:"state"`
}
