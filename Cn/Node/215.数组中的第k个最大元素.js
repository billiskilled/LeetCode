/*
 * @lc app=leetcode.cn id=215 lang=javascript
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (56.55%)
 * Total Accepted:    22.2K
 * Total Submissions: 39.2K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 * 
 * 示例 1:
 * 
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 * 
 * 说明: 
 * 
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 * 
 */
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findKthLargest = function(nums, k) {
    if (!nums) return null;
    if (nums.length < k) return null;

    return findCore(nums, k, 0, nums.length - 1);
};

function findCore(nums, k, start, end) {
    let index = partition(nums, start, end);
    const theIndex = nums.length - k;
    if (index === theIndex) {
        return nums[index];
    } else if (index > theIndex) {
        return findCore(nums, k, start, index - 1);
    } else if (index < theIndex) {
        return findCore(nums, k, index + 1, end);
    }
}

function partition(arr, start, end) {
    const len = end - start + 1;
    const theIndex = Math.floor(Math.random() * len);

    [arr[end], arr[start + theIndex]] = [arr[start + theIndex], arr[end]];

    let theVal = arr[end];

    let cur = start - 1;
    for(let i = 0; i < len; i++) {
        if (arr[start + i] < theVal) {
            cur++;
            [arr[cur], arr[start + i]] = [arr[start + i], arr[cur]];
        }
    }
    
    cur++;
    [arr[cur], arr[end]] = [arr[end], arr[cur]];

    return cur;
}

