function rotate(nums: number[], k: number): void {
  const mod = k % nums.length;
  const dif = nums.length - mod;
  const rotateArr: number[] = [];
  for (let i = 0; i < nums.length; i++) {
    rotateArr[i] = nums[(i + dif) % nums.length];
  }
  for (let i = 0; i < nums.length; i++) {
    nums[i] = rotateArr[i];
  }
}



///// better solution
function rotateArray(nums: number[], k: number): void {
  const n = nums.length;
  k = k % n; // Handle cases where k is greater than n
  if (k === 0) return; // No rotation needed

  // Reverse the entire array
  reverse(nums, 0, n - 1);
  // Reverse the first k elements
  reverse(nums, 0, k - 1);
  // Reverse the remaining n - k elements
  reverse(nums, k, n - 1);
}

function reverse(nums: number[], start: number, end: number): void {
  while (start < end) {
    [nums[start], nums[end]] = [nums[end], nums[start]]; // Swap elements
    start++;
    end--;
  }
}