package enums

type MsgStatus int8

const (
	Sending MsgStatus = iota
	DONE
	FAIL
)

var _MsgStatuslist = []MsgStatus{
	Sending,
}

var _MsgStatusmap = map[MsgStatus]string{
	Sending: "发送中",
	DONE:    "完成",
	FAIL:    "失败",
}

func (static MsgStatus) String() string {
	return string(static)
}

func (static MsgStatus) Value() int8 {
	return int8(static)
}

func (static MsgStatus) Is() bool {
	for _, v := range _MsgStatuslist {
		if v == static {
			return true
		}
	}
	return false
}

func (static MsgStatus) Label() string {
	str, _ := _MsgStatusmap[static]
	return str
}
