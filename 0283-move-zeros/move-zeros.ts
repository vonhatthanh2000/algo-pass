function moveZeroes(nums: number[]): void {
  let lastNonZeroIdx = 0;

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] != 0) {
      nums[lastNonZeroIdx] = nums[i];
      lastNonZeroIdx++;
    }
  }
  for (let i = lastNonZeroIdx; i < nums.length; i++) {
    nums[i] = 0;
  }
}
