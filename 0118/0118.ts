export function generate(numRows: number): number[][] {
  const pascal: number[][] = [];
  if (numRows === 0) return [];
  pascal.push([1]);
  if (numRows === 1) return pascal;
  pascal.push([1, 1]);
  if (numRows === 2) return pascal;

  for (let row = 2; row < numRows; row++) {
    const newRow = [1];
    const lastRow = pascal[pascal.length - 1];
    for (let col = 0; col < lastRow.length - 1; col++) {
      newRow.push(lastRow[col] + lastRow[col + 1]);
    }
    newRow.push(1);
    pascal.push(newRow);
  }

  return pascal;
}
