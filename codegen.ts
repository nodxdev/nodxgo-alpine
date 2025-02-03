import { capitalize } from "./codegen_helpers.ts";

genXAttributes();

function genXAttributes() {
  const attrs = [
    "data",
    "init",
    "show",
    "text",
    "html",
    "model",
    "modelable",
    "for",
    "effect",
    "ref",
    "teleport",
    "if",
    "id",
  ];

  const codeArr: string[] = [
    "package alpine",
    `import nodx "github.com/nodxdev/nodxgo"`,
  ];
  const examplesArr: string[] = [
    "package alpine_test",
    `import "fmt"`,
    `import nodx "github.com/nodxdev/nodxgo"`,
    `import alpine "github.com/nodxdev/nodxgo-alpine"`,
  ];

  for (const attr of attrs) {
    const { code, example } = genXAttribute(attr);
    codeArr.push(code);
    examplesArr.push(example);
  }

  const allCode = codeArr.join("\n\n");
  const allExamples = examplesArr.join("\n\n");

  try {
    Deno.removeSync("./generated_x_attributes.go");
    Deno.writeTextFileSync("./generated_x_attributes_test.go", "");
  } catch {
    // no-op
  }
  Deno.writeTextFileSync("./generated_x_attributes.go", allCode);
  Deno.writeTextFileSync("./generated_x_attributes_test.go", allExamples);
}

function genXAttribute(xAttribute: string): { example: string; code: string } {
  const funcName = "X" + capitalize(xAttribute);
  const attrName = "x-" + xAttribute;
  const docUrl = `https://alpinejs.dev/directives/${xAttribute}`;

  const exampleArr: string[] = [];
  exampleArr.push(`func Example${funcName}() {`);
  exampleArr.push(`  node := nodx.Div(`);
  exampleArr.push(`    alpine.${funcName}("value"),`);
  exampleArr.push(`  )`);
  exampleArr.push(`  fmt.Println(node.String())`);
  exampleArr.push(`  // Output: <div ${attrName}="value"></div>`);
  exampleArr.push(`}`);
  const example = exampleArr.join("\n");
  const exampleComment = exampleArr.map((line) => "//\t" + line).join("\n");

  const codeArr: string[] = [];
  codeArr.push(
    `// ${funcName} is an attribute that renders an ${attrName}="[value]" attribute.`,
  );
  codeArr.push(`//`);
  codeArr.push(`// ${docUrl}`);
  codeArr.push(`//`);
  codeArr.push(exampleComment);
  codeArr.push(`func ${funcName}(value string) nodx.Node {`);
  codeArr.push(`  return nodx.Attr("${attrName}", value)`);
  codeArr.push(`}`);
  const code = codeArr.join("\n");

  return { example, code };
}
