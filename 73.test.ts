import { describe, expect, it } from "vitest";
import { setZeroes } from "./73.ts";

describe("problem 73", () => {
  it("case 1", () => {
    const matrix = [
      [1, 1, 1],
      [1, 0, 1],
      [1, 1, 1],
    ];
    setZeroes(matrix);
    expect(matrix).toEqual([
      [1, 0, 1],
      [0, 0, 0],
      [1, 0, 1],
    ]);
  });

  it("case 2", () => {
    const matrix = [
      [0, 1, 2, 0],
      [3, 4, 5, 2],
      [1, 3, 1, 5],
    ];
    setZeroes(matrix);
    expect(matrix).toEqual([
      [0, 0, 0, 0],
      [0, 4, 5, 0],
      [0, 3, 1, 0],
    ]);
  });

  it("case 2", () => {
    const matrix = [
      [-4, -2147483648, 6, -7, 0],
      [-8, 6, -8, -6, 0],
      [2147483647, 2, -9, -6, -10],
    ];
    setZeroes(matrix);
    expect(matrix).toEqual([
      [0, 0, 0, 0, 0],
      [0, 0, 0, 0, 0],
      [2147483647, 2, -9, -6, 0],
    ]);
  });
});
