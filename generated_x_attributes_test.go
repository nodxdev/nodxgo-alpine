package alpine_test

import "fmt"

import nodx "github.com/nodxdev/nodxgo"

import alpine "github.com/nodxdev/nodxgo-alpine"

func ExampleXData() {
	node := nodx.Div(
		alpine.XData("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-data="value"></div>
}

func ExampleXInit() {
	node := nodx.Div(
		alpine.XInit("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-init="value"></div>
}

func ExampleXShow() {
	node := nodx.Div(
		alpine.XShow("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-show="value"></div>
}

func ExampleXText() {
	node := nodx.Div(
		alpine.XText("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-text="value"></div>
}

func ExampleXHtml() {
	node := nodx.Div(
		alpine.XHtml("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-html="value"></div>
}

func ExampleXModel() {
	node := nodx.Div(
		alpine.XModel("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-model="value"></div>
}

func ExampleXModelable() {
	node := nodx.Div(
		alpine.XModelable("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-modelable="value"></div>
}

func ExampleXFor() {
	node := nodx.Div(
		alpine.XFor("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-for="value"></div>
}

func ExampleXEffect() {
	node := nodx.Div(
		alpine.XEffect("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-effect="value"></div>
}

func ExampleXRef() {
	node := nodx.Div(
		alpine.XRef("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-ref="value"></div>
}

func ExampleXTeleport() {
	node := nodx.Div(
		alpine.XTeleport("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-teleport="value"></div>
}

func ExampleXIf() {
	node := nodx.Div(
		alpine.XIf("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-if="value"></div>
}

func ExampleXId() {
	node := nodx.Div(
		alpine.XId("value"),
	)
	fmt.Println(node.String())
	// Output: <div x-id="value"></div>
}
