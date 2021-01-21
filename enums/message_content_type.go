package enums

type MsgContType string

const (
	TEXT MsgContType = "text"
	//IMAGE       MsgContType = "image"
	//VOICE       MsgContType = "voice"
	//FILE        MsgContType = "file"
	LINK MsgContType = "link"
	//OA          MsgContType = "oa"
	//MARKDOWN    MsgContType = "markdown"
	CARD MsgContType = "action_card"
)

var MsgContTypeList = []MsgContType{
	TEXT, LINK, CARD,
}
var MsgContTypeMap = map[MsgContType]string{
	TEXT: "文本消息",
	//IMAGE:       "图片消息",
	//VOICE:       "语音消息",
	//FILE:        "文件消息",
	LINK: "链接消息",
	//OA:          "OA消息",
	//MARKDOWN:    "markdown消息",
	CARD: "卡片消息",
}

func (static MsgContType) String() string {
	return string(static)
}

func (static MsgContType) Value() string {
	return static.String()
}

func (static MsgContType) Label() string {
	str, _ := MsgContTypeMap[static]
	return str
}

func (static MsgContType) Is() bool {
	for _, v := range MsgContTypeList {
		if v == static {
			return true
		}
	}
	return false
}

func MsgContTypeTypeList() []MsgContType {
	return MsgContTypeList
}

func MsgContTypeTypeToOptions() []map[string]interface{} {
	opts := make([]map[string]interface{}, 0)
	for _, v := range MsgContTypeList {
		opts = append(opts, map[string]interface{}{
			"value": v.Value(),
			"label": v.Label(),
		})
	}
	return opts
}
