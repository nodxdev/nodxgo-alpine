package alpine_test

import (
	"fmt"

	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
)

func ExampleTemplate() {
	node := alpine.Template(
		nodx.Div(nodx.Text("Div")),
		nodx.SpanEl(nodx.Text("Span")),
	)
	fmt.Println(node)
	// Output: <template><div>Div</div><span>Span</span></template>
}

func ExampleX() {
	node := nodx.Div(
		alpine.X("if", "true"),
	)
	fmt.Println(node)
	// Output: <div x-if="true"></div>
}
