import { default as assert } from "assert";

function setIntersections<T>(left: Set<T>, right: Set<T>): T[] {
  return Array.from(new Set([...left].filter((i) => right.has(i))));
}

function findSameValue(tops: number[], bottoms: number[]): number {
  assert(tops.length === bottoms.length, "tops and bottom have unequal size");

  if (tops.length === 0) return -1;
  if (tops.length === 1) return tops[0];

  const sameNumbers = setIntersections(
    new Set([tops[0], bottoms[0]]),
    new Set([tops[1], bottoms[1]])
  );

  if (sameNumbers.length === 0) return -1;

  let finalSameNumber: number | null = null;
  for (
    let sameNumberIdx = 0;
    sameNumberIdx < sameNumbers.length;
    sameNumberIdx++
  ) {
    const sameNumber = sameNumbers[sameNumberIdx];

    try {
      for (let idx = 1; idx < tops.length; idx++) {
        if (!(tops[idx] === sameNumber || bottoms[idx] === sameNumber))
          throw new Error("does not have the number");
      }
    } catch (err) {
      if (sameNumberIdx === sameNumbers.length - 1) return -1;
      continue;
    }

    finalSameNumber = sameNumber;
    break;
  }

  return finalSameNumber!;
}

function calcFlipCount(
  tops: number[],
  bottoms: number[],
  target: number
): number {
  assert(tops.length === bottoms.length, "tops and bottom have unequal size");

  if (tops.length === 0) return 0;
  if (tops.length === 1) return 0;

  let topCount = 0;
  let bottomCount = 0;
  let commonCount = 0;

  for (let idx = 0; idx < tops.length; idx++) {
    if (tops[idx] !== target && bottoms[idx] !== target)
      throw new Error("target number should be a common number among all");

    if (tops[idx] === target) topCount++;
    if (bottoms[idx] === target) bottomCount++;
    if (tops[idx] === target && bottoms[idx] === target) commonCount++;
  }

  return Math.min(topCount, bottomCount) - commonCount;
}

export function minDominoRotations(tops: number[], bottoms: number[]): number {
  const sameNumber = findSameValue(tops, bottoms);
  if (sameNumber === -1) return -1;

  return calcFlipCount(tops, bottoms, sameNumber);
}
