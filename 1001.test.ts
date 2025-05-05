import { describe, expect, it } from "vitest";
import { minDominoRotations } from "./1001.ts";

describe("problem 1001", () => {
  it("has answer", () => {
    const top = [2, 1, 2, 4, 2, 2];
    const bottom = [5, 2, 6, 2, 3, 2];
    expect(minDominoRotations(top, bottom)).toBe(2);
  });

  it("has no answer", () => {
    const top = [3, 5, 1, 2, 3];
    const bottom = [3, 6, 3, 3, 4];
    expect(minDominoRotations(top, bottom)).toBe(-1);
  });

  it("has same 2 numbers everywhere", () => {
    const tops = [1, 2, 1, 1, 1, 2, 2, 2];
    const bottoms = [2, 1, 2, 2, 2, 2, 2, 2];
    expect(minDominoRotations(tops, bottoms)).toBe(1);
  });

  it("flipping once", () => {
    const tops = [1, 2, 3, 4, 6];
    const bottoms = [6, 6, 6, 6, 5];
    expect(minDominoRotations(tops, bottoms)).toBe(1);
  });
});
