package flyhttp

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println(Get("/idol/idollist").String())
}


