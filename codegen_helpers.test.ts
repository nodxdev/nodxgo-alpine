import { assertEquals } from "jsr:@std/assert@1.0.11";
import { capitalize } from "./codegen_helpers.ts";

Deno.test("capitalize", () => {
  // deno-lint-ignore no-explicit-any
  assertEquals(capitalize(1 as any as string), "");
  assertEquals(capitalize(""), "");
  assertEquals(capitalize("hello"), "Hello");
  assertEquals(capitalize("world"), "World");
});
