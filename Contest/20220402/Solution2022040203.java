/**
 * 6035. 选择建筑的方案数
 *
 * 给你一个下标从 0 开始的二进制字符串 s ，它表示一条街沿途的建筑类型，其中：
 *
 * s[i] = '0' 表示第 i 栋建筑是一栋办公楼，
 * s[i] = '1' 表示第 i 栋建筑是一间餐厅。
 * 作为市政厅的官员，你需要随机 选择 3 栋建筑。然而，为了确保多样性，选出来的 3 栋建筑 相邻 的两栋不能是同一类型。
 *
 * 比方说，给你 s = "001101" ，我们不能选择第 1 ，3 和 5 栋建筑，因为得到的子序列是 "011" ，有相邻两栋建筑是同一类型，所以 不合 题意。
 * 请你返回可以选择 3 栋建筑的 有效方案数 。
 */
class Solution2022040203 {
    public static void main(String[] args) {
        String s = "0001100100";
        System.out.println(numberOfWays(s));
    }
    public static long numberOfWays(String s) {
        if (s == null || s.length() == 0) {
            return 0L;
        }
        int len = s.length();
        char[] sCharArr = s.toCharArray();

        int[] statistic = new int[len];
        int sLen = 0;

        char flag = sCharArr[0];
        int tempSum = 0;
        for (int i = 0; i < len; i++) {
            char cur = sCharArr[i];
            if (cur == flag) {
                tempSum++;
            } else {
                flag = flag == '0' ? '1' : '0';
                statistic[sLen++] = tempSum;
                tempSum = 1;
            }
        }
        statistic[sLen++] = tempSum;

        long result = 0L;
        if (sLen < 3) {
            return 0;
        }

        for (int i = 0; i < sLen - 2; i++) {
            long curResult = 0L;

            int first = i + 1;
            int second = first + 1;

            long firstSum = 0L;
            while (first < sLen) {
                long secondSum = 0L;
                while (second < sLen) {
                    secondSum += statistic[second];
                    second += 2;
                }

                firstSum += statistic[first] * secondSum;

                first += 2;
                second = first + 1;
            }

            curResult = (long)statistic[i] * firstSum;

            result += curResult;
        }

        return result;
    }
}