/**
 * 5219. 每个小孩最多能分到多少糖果
 *
 * 给你一个 下标从 0 开始 的整数数组 candies 。数组中的每个元素表示大小为 candies[i] 的一堆糖果。你可以将每堆糖果分成任意数量的 子堆 ，但 无法 再将两堆合并到一起。
 *
 * 另给你一个整数 k 。你需要将这些糖果分配给 k 个小孩，使每个小孩分到 相同 数量的糖果。每个小孩可以拿走 至多一堆 糖果，有些糖果可能会不被分配。
 *
 * 返回每个小孩可以拿走的 最大糖果数目 。
 */
public class Solution2022040303 {
    public int maximumCandies(int[] candies, long k) {
        if (candies == null) {
            return 0;
        }
        int len = candies.length;
        int sum = 0;
        int min = Integer.MAX_VALUE;
        for (int candy : candies) {
            sum += candy;
            if (candy < min) {
                min = candy;
            }
        }

        if (sum < k) {
            return 0;
        }

        while (min > 0) {
            int num = 0;
            for (int candy : candies) {
                if (candy >= min) {
                    num += candy / min;
                }
                if (num >= k) {
                    break;
                }
            }
            if (num >= k) {
                break;
            }
            min--;
        }


        return min;
    }
}
