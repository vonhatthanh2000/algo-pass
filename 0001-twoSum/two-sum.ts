// 1st try

function twoSum(nums: number[], target: number): number[] {
    let arrObj: { [key: number]: number } = {};
    for (let i = 0; i < nums.length; i++) {
        const expected = target - nums[i];
        if (arrObj.hasOwnProperty(expected)) {
            return [arrObj[expected], i];
        }
        arrObj[nums[i]] = i;
    }
    return [];
}


// better
function twoSum2(nums: number[], target: number): number[] {
    const arrObj: Record<number, number> = {}
    for (let i = 0; i < nums.length; i++) {
        const expected = target - nums[i];
        if (expected in arrObj) {
            return [arrObj[expected], i]
        }
        arrObj[nums[i]] = i
    }

    return []
}