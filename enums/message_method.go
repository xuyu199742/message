package enums

type MsgMethod string

const (
	DD MsgMethod = "dd"
	WX MsgMethod = "wx"
)

var MsgMethodList = []MsgMethod{
	DD, WX,
}
var MsgMethodMap = map[MsgMethod]string{
	DD: "钉钉",
	WX: "微信",
}

func (static MsgMethod) String() string {
	return string(static)
}

func (static MsgMethod) Value() string {
	return static.String()
}

func (static MsgMethod) Label() string {
	str, _ := MsgMethodMap[static]
	return str
}

func (static MsgMethod) Is() bool {
	for _, v := range MsgMethodList {
		if v == static {
			return true
		}
	}
	return false
}

func MsgMethodTypeList() []MsgMethod {
	return MsgMethodList
}

func MsgMethodToOptions() []map[string]interface{} {
	opts := make([]map[string]interface{}, 0)
	for _, v := range MsgMethodList {
		opts = append(opts, map[string]interface{}{
			"value": v.Value(),
			"label": v.Label(),
		})
	}
	return opts
}
