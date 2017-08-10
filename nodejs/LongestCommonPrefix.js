// Problem 14
/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    if (strs.length === 0) {
        return '';
    }
    if (strs.length === 1) {
        return strs[0];
    }
    let commonPrefix = '';
    let sentinel = '';
    let strsLength = strs.length;
    let minLength = 0; // 该变量记录字符串数组中最短的字符串长度
    let breakFlag = false; // 这个标志记录终止while循环时机
    let i = 0;
    for (let j = 0; j < strsLength; j++) {
        if (j === 0) {
            minLength = strs[j].length;
        } else {
            minLength = minLength < strs[i].length ? strs[i].length : minLength;
        }
    }
    if (minLength === 0) {
        return '';
    }
    while (true) {
        let j;
        for (j = 0; j < strsLength; j++) {
            if (j === 0) {
                sentinel = strs[0][i];
            } else {
                if (sentinel === strs[j][i]) {
                    continue;
                } else {
                    breakFlag = true;
                    break;
                }
            }
        }
        if (j === strsLength) {
            commonPrefix += sentinel;
        }
        i++;
        if (breakFlag) {
            break;
        }
        if (i >= minLength) {
            break;
        }
    }
    return commonPrefix;
};

let strs = ['', '', '', ''];
console.log(longestCommonPrefix(strs));
