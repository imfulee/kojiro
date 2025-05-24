import { deepStrictEqual } from "assert";

function findWordsContaining(words: string[], x: string): number[] {
  const indices: number[] = [];
  for (let idx = 0; idx < words.length; idx++) {
    if (words[idx].includes(x)) {
      indices.push(idx);
    }
  }
  return indices;
}

deepStrictEqual(findWordsContaining(["leet", "code"], "e"), [0, 1]);
