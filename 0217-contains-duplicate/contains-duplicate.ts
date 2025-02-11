function containsDuplicate(nums: number[]): boolean {
    const numsObj: Record<number, number> = {}
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] in numsObj) return true
        numsObj[nums[i]] = i
    }

    return false;
};