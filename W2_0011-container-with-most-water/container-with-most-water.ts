// my solution with BigO (n!) - time limit exceed
function maxArea(height: number[]): number {
    let maxValue = 0;
    for (let i = 0; i < height.length; i++) {
        for (let j = i + 1; j < height.length; j++) {
            const subMax = Math.min(height[i], height[j]) * (j - i);
            if (subMax > maxValue) maxValue = subMax;
        }
    }

    return maxValue;
};

// consult solution
function maxArea2(height: number[]): number {
    let maxAmount = 0;
    let left = 0;
    let right = height.length - 1;

    while (left < right) {
        const leftVal = height[left];
        const rightVal = height[right];
        const width = right - left;
        const currentAmount = Math.min(leftVal, rightVal) * width;
        maxAmount = Math.max(maxAmount, currentAmount);
        if (leftVal < rightVal) {
            left++
        } else {
            right--
        }
    }

    return maxAmount
};