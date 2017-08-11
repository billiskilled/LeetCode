// Problem 20
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    let container = [];
    let strLength = s.length;
    let i;

    for (i = 0; i < strLength; i++) {
        if (s[i] == '(' || s[i] == '[' || s[i] == '{') {
            container.push(s[i]);
            continue;
        } 
        if (s[i] == ')' || s[i] == ']' || s[i] == '}') {
            switch(s[i]) {
                case ')':
                    if (container[container.length - 1] === '(') {
                        container.pop();
                    } else {
                        return false;
                    }
                    break;
                case ']':
                    if (container[container.length - 1] === '[') {
                        container.pop();
                    } else {
                        return false;
                    }
                    break;
                case '}':
                    if (container[container.length - 1] === '{') {
                        container.pop();
                    } else {
                        return false;
                    }
                    break;
                default:
                    break;
            }
        }
    }
    // 这个地方的代码才是需要注意的地方
    if (container.indexOf('(') === 0 ||
        container.indexOf('[') === 0 ||
        container.indexOf('{') === 0 ) 
    {
        return false;
    }

    if (i === strLength) {
        return true;
    }
};

var s = '([)';
console.log(isValid(s));
