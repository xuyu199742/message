package enums

type ReceiverType int8

const (
	USER_RECEIVER ReceiverType = iota + 1
	DEP_RECEIVER
)

var _ReceiverTypelist = []ReceiverType{
	USER_RECEIVER,
	DEP_RECEIVER,
}

var _ReceiverTypemap = map[ReceiverType]string{
	USER_RECEIVER: "员工",
	DEP_RECEIVER:  "部门",
}

func (static ReceiverType) String() string {
	return string(static)
}

func (static ReceiverType) Value() int8 {
	return int8(static)
}

func (static ReceiverType) Is() bool {
	for _, v := range _ReceiverTypelist {
		if v == static {
			return true
		}
	}
	return false
}

func (static ReceiverType) Label() string {
	str, _ := _ReceiverTypemap[static]
	return str
}

func ReceiverTypeToOptions() []map[string]interface{} {
	opts := make([]map[string]interface{}, 0)
	for _, v := range _ReceiverTypelist {
		opts = append(opts, map[string]interface{}{
			"value": v.Value(),
			"label": v.Label(),
		})
	}
	return opts
}
