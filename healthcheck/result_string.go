// Code generated by "stringer --type Result healthcheck/result.go"; DO NOT EDIT

package healthcheck

import "fmt"

const _Result_name = "HCUnknownHCErrorHCBadHCGood"

var _Result_index = [...]uint8{0, 9, 16, 21, 27}

func (i Result) String() string {
	if i < 0 || i >= Result(len(_Result_index)-1) {
		return fmt.Sprintf("Result(%d)", i)
	}
	return _Result_name[_Result_index[i]:_Result_index[i+1]]
}
