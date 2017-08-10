// Problem 9
/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    if (x < 0) {
        return false;
    }
    var num = new String(x);
    var len = num.length;
    var i = 0;
    while (i < Math.floor(len / 2)) {
        if (num[i] === num[len - 1 - i++]) {
            continue;
        } else {
            return false;
        }
    }
    return true;
};

console.log(isPalindrome(process.argv.slice(2)));
