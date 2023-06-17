package profile

import (
	"fmt"
)

func Profile(idtype string, id string) (target string) {
	switch idtype {
	case "email":
		fmt.Printf("%s is an email.\n", idtype)
		target := "/account/" + id
		return target
	default:
		fmt.Printf("%s isn't a vaild idtype.\n", idtype)
	}
	return target
}
