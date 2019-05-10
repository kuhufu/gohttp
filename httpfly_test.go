package flyhttp

import (
	"fmt"
	"testing"
)

func TestClient_Get(t *testing.T) {
	fmt.Println(obj.Get("/idol/idollist").String())
}


