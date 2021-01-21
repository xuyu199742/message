package enums

type MsgType string

const (
	//TASK     MsgType = "task"
	//GROUP    MsgType = "group"
	WORK MsgType = "work"
	//ORDINARY MsgType = "ordinary"
)

var MsgTypeList = []MsgType{
	WORK,
	//TASK, GROUP, , ORDINARY,
}
var MsgTypeMap = map[MsgType]string{
	//TASK:     "任务类通知",
	//ORDINARY: "普通消息",
	//GROUP:    "群消息",
	WORK: "工作通知消息",
}

func (static MsgType) String() string {
	return string(static)
}

func (static MsgType) Value() string {
	return static.String()
}

func (static MsgType) Label() string {
	str, _ := MsgTypeMap[static]
	return str
}

func (static MsgType) Is() bool {
	for _, v := range MsgTypeList {
		if v == static {
			return true
		}
	}
	return false
}

func MsgTypeTypeList() []MsgType {
	return MsgTypeList
}

func MsgTypeToOptions() []map[string]interface{} {
	opts := make([]map[string]interface{}, 0)
	for _, v := range MsgTypeList {
		opts = append(opts, map[string]interface{}{
			"value": v.Value(),
			"label": v.Label(),
		})
	}
	return opts
}
