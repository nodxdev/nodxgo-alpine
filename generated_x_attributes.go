package alpine

import nodx "github.com/nodxdev/nodxgo"

// XData is an attribute that renders an x-data="[value]" attribute.
//
// https://alpinejs.dev/directives/data
//
//	func ExampleXData() {
//	  node := nodx.Div(
//	    alpine.XData("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-data="value"></div>
//	}
func XData(value string) nodx.Node {
	return nodx.Attr("x-data", value)
}

// XInit is an attribute that renders an x-init="[value]" attribute.
//
// https://alpinejs.dev/directives/init
//
//	func ExampleXInit() {
//	  node := nodx.Div(
//	    alpine.XInit("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-init="value"></div>
//	}
func XInit(value string) nodx.Node {
	return nodx.Attr("x-init", value)
}

// XShow is an attribute that renders an x-show="[value]" attribute.
//
// https://alpinejs.dev/directives/show
//
//	func ExampleXShow() {
//	  node := nodx.Div(
//	    alpine.XShow("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-show="value"></div>
//	}
func XShow(value string) nodx.Node {
	return nodx.Attr("x-show", value)
}

// XText is an attribute that renders an x-text="[value]" attribute.
//
// https://alpinejs.dev/directives/text
//
//	func ExampleXText() {
//	  node := nodx.Div(
//	    alpine.XText("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-text="value"></div>
//	}
func XText(value string) nodx.Node {
	return nodx.Attr("x-text", value)
}

// XHtml is an attribute that renders an x-html="[value]" attribute.
//
// https://alpinejs.dev/directives/html
//
//	func ExampleXHtml() {
//	  node := nodx.Div(
//	    alpine.XHtml("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-html="value"></div>
//	}
func XHtml(value string) nodx.Node {
	return nodx.Attr("x-html", value)
}

// XModel is an attribute that renders an x-model="[value]" attribute.
//
// https://alpinejs.dev/directives/model
//
//	func ExampleXModel() {
//	  node := nodx.Div(
//	    alpine.XModel("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-model="value"></div>
//	}
func XModel(value string) nodx.Node {
	return nodx.Attr("x-model", value)
}

// XModelable is an attribute that renders an x-modelable="[value]" attribute.
//
// https://alpinejs.dev/directives/modelable
//
//	func ExampleXModelable() {
//	  node := nodx.Div(
//	    alpine.XModelable("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-modelable="value"></div>
//	}
func XModelable(value string) nodx.Node {
	return nodx.Attr("x-modelable", value)
}

// XFor is an attribute that renders an x-for="[value]" attribute.
//
// https://alpinejs.dev/directives/for
//
//	func ExampleXFor() {
//	  node := nodx.Div(
//	    alpine.XFor("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-for="value"></div>
//	}
func XFor(value string) nodx.Node {
	return nodx.Attr("x-for", value)
}

// XEffect is an attribute that renders an x-effect="[value]" attribute.
//
// https://alpinejs.dev/directives/effect
//
//	func ExampleXEffect() {
//	  node := nodx.Div(
//	    alpine.XEffect("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-effect="value"></div>
//	}
func XEffect(value string) nodx.Node {
	return nodx.Attr("x-effect", value)
}

// XRef is an attribute that renders an x-ref="[value]" attribute.
//
// https://alpinejs.dev/directives/ref
//
//	func ExampleXRef() {
//	  node := nodx.Div(
//	    alpine.XRef("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-ref="value"></div>
//	}
func XRef(value string) nodx.Node {
	return nodx.Attr("x-ref", value)
}

// XTeleport is an attribute that renders an x-teleport="[value]" attribute.
//
// https://alpinejs.dev/directives/teleport
//
//	func ExampleXTeleport() {
//	  node := nodx.Div(
//	    alpine.XTeleport("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-teleport="value"></div>
//	}
func XTeleport(value string) nodx.Node {
	return nodx.Attr("x-teleport", value)
}

// XIf is an attribute that renders an x-if="[value]" attribute.
//
// https://alpinejs.dev/directives/if
//
//	func ExampleXIf() {
//	  node := nodx.Div(
//	    alpine.XIf("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-if="value"></div>
//	}
func XIf(value string) nodx.Node {
	return nodx.Attr("x-if", value)
}

// XId is an attribute that renders an x-id="[value]" attribute.
//
// https://alpinejs.dev/directives/id
//
//	func ExampleXId() {
//	  node := nodx.Div(
//	    alpine.XId("value"),
//	  )
//	  fmt.Println(node.String())
//	  // Output: <div x-id="value"></div>
//	}
func XId(value string) nodx.Node {
	return nodx.Attr("x-id", value)
}
