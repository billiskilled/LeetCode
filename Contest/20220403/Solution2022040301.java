class Solution2022040301 {
    public static void main(String[] args) {
        System.out.println(convertTime("02:30", "04:35"));
    }
    public static int convertTime(String current, String correct) {
        if (current == null || correct == null || current.length() != 5 || correct.length() != 5) {
            return 0;
        }
        int[] cur = new int[2];
        int[] cor = new int[2];

        String[] curS = current.split(":");
        String[] corS = correct.split(":");
        cur[0] = Integer.parseInt(curS[0]);
        cur[1] = Integer.parseInt(curS[1]);
        cor[0] = Integer.parseInt(corS[0]);
        cor[1] = Integer.parseInt(corS[1]);

        int diff = 0;
        if (cor[1] < cur[1]) {
            cor[1] += 60;
            cor[0] -= 1;
        }
        if (cor[0] < cur[0]) {
            cor[0] += 24;
        }

        diff = 60 * (cor[0] - cur[0]) + cor[1] - cur[1];

        int result = 0;
        while (diff >= 60) {
            result++;
            diff -= 60;
        }
        while (diff >= 15) {
            result++;
            diff -= 15;
        }
        while (diff >= 5) {
            result++;
            diff -= 5;
        }
        result += diff;

        return result;
    }
}