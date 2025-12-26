import { expect, it, test } from "vitest";
import { generate } from "./0118.ts";
import { describe } from "node:test";

describe("Question 0118", () => {
  it("Case 1", () => {
    expect(generate(1)).toEqual([[1]]);
  });
  it("Case 5", () => {
    expect(generate(5)).toEqual([
      [1],
      [1, 1],
      [1, 2, 1],
      [1, 3, 3, 1],
      [1, 4, 6, 4, 1],
    ]);
  });
});
