import java.util.Arrays;

/**
 * 6037. 按奇偶性交换后的最大数字
 *
 * 给你一个正整数 num 。你可以交换 num 中 奇偶性 相同的任意两位数字（即，都是奇数或者偶数）。
 *
 * 返回交换 任意 次之后 num 的 最大 可能值。
 */
class Solution2022041001 {
    public static void main(String[] args) {
        System.out.println(largestInteger(65875));
    }

    public static int largestInteger(int num) {
        char[] cArr = String.valueOf(num).toCharArray();

        int len = cArr.length;

        int[] flags = new int[len];
        int jiNum = 0;
        int ouNum = 0;
        for (int i = 0; i < len; i++) {
            if ((cArr[i] - '0') % 2 == 1) {
                jiNum++;
            } else {
                ouNum++;
            }
        }

        Character[] ji = new Character[jiNum];
        Character[] ou = new Character[ouNum];

        int jiIndex = 0;
        int ouIndex = 0;
        for (int i = 0; i < len; i++) {
            flags[i] = (cArr[i] - '0') % 2 == 0 ? 0 : 1;
            if ((cArr[i] - '0') % 2 == 1) {
                ji[jiIndex++] = cArr[i];
            } else {
                ou[ouIndex++] = cArr[i];
            }
        }

        Arrays.sort(ji, (a, b) -> b - a);
        Arrays.sort(ou, (a, b) -> b - a);

        StringBuilder sb = new StringBuilder();
        jiIndex = 0;
        ouIndex = 0;
        for (int i = 0; i < len; i++) {
            if (flags[i] == 1) {
                sb.append(ji[jiIndex++]);
            } else {
                sb.append(ou[ouIndex++]);
            }
        }

        return Integer.parseInt(sb.toString());
    }
}