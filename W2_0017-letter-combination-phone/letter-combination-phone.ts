function letterCombinations(digits: string): string[] {
    let obj: Record<string, string[]> = {};
    obj["2"] = ["a", "b", "c"];
    obj["3"] = ["d", "e", "f"];
    obj["4"] = ["g", "h", "i"];
    obj["5"] = ["j", "k", "l"];
    obj["6"] = ["m", "n", "o"];
    obj["7"] = ["p", "q", "r", "s"];
    obj["8"] = ["t", "u", "v"];
    obj["9"] = ["w", "x", "y", "z"];

    let resArr: string[] = [];


    if (digits.length == 0) return []

    resArr = obj[digits[0]]
    if (digits.length == 1)
        return resArr

    for (let i = 1; i < digits.length; i++) {
        const char: string = digits[i];
        resArr = combine2arr(resArr, obj[char])
    }

    return resArr
};

function combine2arr(arr1: string[], arr2: string[]): string[] {
    const resArr: string[] = [];
    for (let i = 0; i < arr1.length; i++) {
        const fstLetter = arr1[i]
        for (let j = 0; j < arr2.length; j++) {
            resArr.push(fstLetter.concat(arr2[j]))
        }
    }
    return resArr
}