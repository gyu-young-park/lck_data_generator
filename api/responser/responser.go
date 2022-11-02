package responser

import (
	"fmt"
	"net/http"
)

func Response(res http.ResponseWriter, stateCode int, msg string, arg ...string) {
	res.WriteHeader(stateCode)
	data := fmt.Sprintf(msg, arg)
	res.Write([]byte(data))
}
