// Problem 27
/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
    let numsLength = nums.length;
    let newStart = 0; // 如果已经indexOf查找到
    let valCount = 0;
    let endFlag = numsLength - 1;
    let temp;
    while (true) {
        if (nums.indexOf(val, newStart) !== -1) {
            newStart = nums.indexOf(val, newStart);
            temp = nums[newStart];
            nums[newStart] = nums[endFlag];
            nums[endFlag--] = temp;
            nums.pop();
            if (nums[newStart] !== val) {
                newStart += 1; // 从新位置的下一位开始查找新的val
            }
            valCount++;
        } else {
            break;
        }
    }
    
    console.log(nums);
    return numsLength - valCount;
};

let a = [3, 2, 2, 3];
console.log(removeElement(a, 3));
