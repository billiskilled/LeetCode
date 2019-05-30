/*
 * @lc app=leetcode.cn id=279 lang=javascript
 *
 * [279] 完全平方数
 */
/**
 * @param {number} n
 * @return {number}
 */
var numSquares = function(n) {
    let dp = [Number.MAX_SAFE_INTEGER, 1, 2, 3];
    if (n <= 3) return n;

    for (let i = 3; i <= n; i++) {
        dp[i] = findMin(i, dp);
    }

    return dp[n];
};

function findMin(n, dp) {
    let maxSqrtV = Math.floor(Math.sqrt(n));

    if (maxSqrtV * maxSqrtV === n) return 1;

    let tempArr = [];
    for (let i = 1; i <= maxSqrtV; i++) {
        tempArr.push(dp[n - i * i]);
    }

    return 1 + getMin(tempArr);
}

function getMin(arr) {
    let min = Number.MAX_SAFE_INTEGER
    for (let v of arr) {
        if (v < min) {
            min = v
        }
    }

    return min;
}
