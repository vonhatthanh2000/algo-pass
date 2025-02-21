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