package pour

import (
	"fmt"
	"io"

	"encoding/json"
	"strings"
)

// JSON pours list of props to json
func JSON(s string, wr io.Writer) {
	enc := json.NewEncoder(wr)
	p := map[string]string{}

	w := strings.Fields(s)
	for i := 1; i < len(w); i += 2 {
		p[strings.Trim(w[i-1], ",;:()")] = strings.Trim(w[i], ",;:()")
	}

	fmt.Println("JSON: ", p)
	enc.Encode(p)
}
