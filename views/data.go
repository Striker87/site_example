package views

const (
	AlertLvlError   = "danger"
	AlertLvlWarning = "warning"
	AlertLvlInfo    = "info"
	AlertLvlSuccess = "success"

	AlertMsgGeneric = "Something went wrong"
)

// used to render Bootstrap Alert messages in the templates
type Alert struct {
	Level   string
	Message string
}

// top level structure that views expect data to come in
type Data struct {
	Alert *Alert
	Yield interface{}
}
