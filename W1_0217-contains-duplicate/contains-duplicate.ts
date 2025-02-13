function containsDuplicate(nums: number[]): boolean {
    const numsObj: Record<number, number> = {}
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] in numsObj) return true
        numsObj[nums[i]] = i
    }

    return false;
};



// reference
// using Set
function containsDuplicate2(nums: number[]): boolean {
    const used = new Set<number>()

    for (const num of nums) {
        if (used.has(num)) return true

        used.add(num)
    }

    return false
};

// better
function containsDuplicate3(nums: number[]): boolean {
    return new Set(nums).size !== nums.length;
};