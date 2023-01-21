package hi

import (
	"fmt"

	"github.com/julianossilva/goworkspacetest"
)

func Hi() {
	fmt.Println("Hi World!")
}

func Other() {
	goworkspacetest.Message()
}
