/**
 * 6038. 向表达式添加括号后的最小结果
 *
 * 给你一个下标从 0 开始的字符串 expression ，格式为 "<num1>+<num2>" ，其中 <num1> 和 <num2> 表示正整数。
 *
 * 请你向 expression 中添加一对括号，使得在添加之后， expression 仍然是一个有效的数学表达式，并且计算后可以得到 最小 可能值。左括号 必须 添加在 '+' 的左侧，而右括号必须添加在 '+' 的右侧。
 *
 * 返回添加一对括号后形成的表达式 expression ，且满足 expression 计算得到 最小 可能值。如果存在多个答案都能产生相同结果，返回任意一个答案。
 *
 * 生成的输入满足：expression 的原始值和添加满足要求的任一对括号之后 expression 的值，都符合 32-bit 带符号整数范围。
 */
class Solution2022041002 {
    public static void main(String[] args) {
        System.out.println(minimizeResult("247+38"));
    }
    public static String minimizeResult(String expression) {
        if (expression == null || expression.length() == 0) {
            return expression;
        }
        String[] sArr = expression.split("\\+");
        String firstNum = sArr[0];
        String secondNum = sArr[1];

        int firstLen = firstNum.length();
        int secondLen = secondNum.length();

        int firstIndex = 1;
        int secondIndex = 1;

        int min = Integer.MAX_VALUE;

        for (int i = firstLen - 2; i >= -1; i--) {
            for (int j = 1; j <= secondLen; j++) {
                StringBuilder sb1 = new StringBuilder();
                StringBuilder sb2 = new StringBuilder();
                StringBuilder sb3 = new StringBuilder();
                StringBuilder sb4 = new StringBuilder();
                for (int k = 0; k <= i; k++) {
                    sb1.append(firstNum.charAt(k));
                }
                for (int k = i + 1; k < firstLen; k++) {
                    sb2.append(firstNum.charAt(k));
                }
                for (int k = 0; k < j; k++) {
                    sb3.append(secondNum.charAt(k));
                }
                for (int k = j; k < secondLen; k++) {
                    sb4.append(secondNum.charAt(k));
                }

                int temp = Integer.parseInt(sb1.toString().equals("") ? "1" : sb1.toString())
                        * (Integer.parseInt(sb2.toString()) + Integer.parseInt(sb3.toString()))
                        * Integer.parseInt(sb4.toString().equals("") ? "1" : sb4.toString());
                if (temp < min) {
                    min = temp;
                    firstIndex = i;
                    secondIndex = j;
                }
            }
        }

        StringBuilder sb = new StringBuilder();
        for (int i = 0; i <= firstIndex; i++) {
            sb.append(firstNum.charAt(i));
        }
        sb.append("(");
        for (int i = firstIndex + 1; i < firstLen; i++) {
            sb.append(firstNum.charAt(i));
        }
        sb.append("+");
        for (int i = 0; i < secondIndex; i++) {
            sb.append(secondNum.charAt(i));
        }
        sb.append(")");
        for (int i = secondIndex; i < secondLen; i++) {
            sb.append(secondNum.charAt(i));
        }

        return sb.toString();
    }
}