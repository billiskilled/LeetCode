// Problem 136
/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
    let result = 0;
    for (let i = 0; i < nums.length; i++) {
        result = result ^ nums[i];
    }
    return result;
};

let nums = [22, 11, 33, 22, 11];
console.log(singleNumber(nums));
