/**
 Do not return anything, modify matrix in-place instead.
 */
export function setZeroes(matrix: number[][]): void {
  let isFirstColZeroed = false;
  for (let row = 0; row < matrix.length; row++) {
    if (matrix[row][0] === 0) isFirstColZeroed = true;
    for (let col = 1; col < matrix[row].length; col++) {
      if (matrix[row][col] === 0) {
        matrix[0][col] = 0;
        matrix[row][0] = 0;
      }
    }
  }

  for (let row = matrix.length - 1; row >= 0; row--) {
    for (let col = matrix[row].length - 1; col >= 1; col--) {
      if (matrix[row][0] === 0 || matrix[0][col] === 0) matrix[row][col] = 0;
    }
    if (isFirstColZeroed) matrix[row][0] = 0;
  }
}
