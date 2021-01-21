package enums

type SenderType int8

const (
	AGENT_SENDER SenderType = iota + 1
)

var _SenderTypelist = []SenderType{
	AGENT_SENDER,
}

var _SenderTypemap = map[SenderType]string{
	AGENT_SENDER: "应用平台",
}

func (static SenderType) String() string {
	return string(static)
}

func (static SenderType) Value() int8 {
	return int8(static)
}

func (static SenderType) Is() bool {
	for _, v := range _SenderTypelist {
		if v == static {
			return true
		}
	}
	return false
}

func (static SenderType) Label() string {
	str, _ := _SenderTypemap[static]
	return str
}

func SenderTypeToOptions() []map[string]interface{} {
	opts := make([]map[string]interface{}, 0)
	for _, v := range _SenderTypelist {
		opts = append(opts, map[string]interface{}{
			"value": v.Value(),
			"label": v.Label(),
		})
	}
	return opts
}
