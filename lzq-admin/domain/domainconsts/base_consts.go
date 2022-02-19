package domainconsts

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

const (
	NoConst = "NoConst"
)

var BaseMsgFlags = map[string]string{
	NoConst: "NoConst",
}

// GetConstFlag get error information based on Code
func GetConstFlag(code string, constFlags map[string]string) string {
	msg, ok := constFlags[code]
	if ok {
		return msg
	}

	return BaseMsgFlags[NoConst]
}
