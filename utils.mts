function setIntersections<T>(left: Set<T>, right: Set<T>): T[] {
  return Array.from(new Set([...left].filter((i) => right.has(i))));
}
