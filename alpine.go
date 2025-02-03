// Package alpine provides Alpine.js integration for NodX Go
//
//   - https://alpinejs.dev
//   - https://github.com/nodxdev/nodxgo
package alpine

import nodx "github.com/nodxdev/nodxgo"

// Template is a generic function for rendering a <template> element.
//
//	func ExampleTemplate() {
//		node := alpine.Template(
//			nodx.Div(nodx.Text("Div")),
//			nodx.SpanEl(nodx.Text("Span")),
//		)
//		fmt.Println(node)
//		// Output: <template><div>Div</div><span>Span</span></template>
//	}
func Template(children ...nodx.Node) nodx.Node {
	return nodx.El("template", children...)
}

// X is an attribute that renders a x-[key]="[value]" attribute.
//
//	func ExampleX() {
//		node := nodx.Div(
//			alpine.X("if", "true"),
//		)
//		fmt.Println(node)
//		// Output: <div x-if="true"></div>
//	}
func X(key string, value string) nodx.Node {
	return nodx.Attr("x-"+key, value)
}
