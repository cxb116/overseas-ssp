package impl

type DspCompany struct {
	Id      int64  `json:"id"`       // 预算公司Id
	Name    string `json:"name"`     // 公司名称
	DspCode int64  `json:"dsp_code"` // 自定义预算位Code
	Url     string `json:"url"`      // 预算请求url
	Method  string `json:"method"`   // 请求方式 1=post 2=get
	Timeout int64  `json:"timeout"`  // 请求预算响应时间,毫秒时间戳
}

func (DspCompany) TableName() string {
	return "dsp_company"
}
