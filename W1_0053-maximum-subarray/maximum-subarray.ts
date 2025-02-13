function maxSubArray(nums: number[]): number {
  let maxSub = nums[0];
  let subSum = nums[0];

  for (let i = 1; i < nums.length; i++) {
    if (subSum < 0) {
      subSum = nums[i];
    } else {
      subSum += nums[i];
    }
    if (subSum > maxSub) maxSub = subSum;
  }
  return maxSub;
}
