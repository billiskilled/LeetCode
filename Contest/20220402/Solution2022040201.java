class Solution2022040201 {
    public int minBitFlips(int start, int goal) {
        int temp = start ^ goal;
        int result = 0;
        while (temp != 0) {
            result++;
            temp = temp & (temp - 1);
        }
        return result;
    }
}