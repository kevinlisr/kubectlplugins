package lib

import "fmt"

// jiang map zhuan huan wei string
func Map2String(m map[string]string) (ret string) {
	for k, v := range m{
		ret += fmt.Sprintf("%s=%s\n", k, v)
	}
	return ret
}
