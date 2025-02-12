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
