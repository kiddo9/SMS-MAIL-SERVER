package structures

type EmailTemp struct {
	ID              int    `json:"id"`
	TemplateName    string `json:"TemplateName"`
	TemplateContent string `json:"TemplateContent"`
	Date            string `json:"Date"`
}

type SmsTemp struct {
	ID              int    `json:"id"`
	TemplateName    string `json:"name"`
	TemplateContent string `json:"TempContent"`
	Date            string `json:"Date"`
}

type TemplateStruct struct {
	Templates struct {
		EmailsTemp []EmailTemp `json:"emailsTemp"`
		SmsTemp    []SmsTemp   `json:"smsTemp"`
	} `json:"Templates"`
}