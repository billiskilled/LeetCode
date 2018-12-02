class Solution {
    public String largestTimeFromDigits(int[] A) {
        int[] result = new int[4];

        for (int i = 0; i < 4; i++) {
            if (i == 0) {
                int maxNum = -1;
                boolean isValidNum = false;
                for (int j = 0; j < 4; j++) {
                    if (0 <= A[j] && A[j] <= 2 && a[j] > maxNum){
                        result[i] = maxNum;
                        isValidNum = true;
                    }
                }
                if (!isValidNum) {
                    return "";
                }
            }

            if (i == 1) {
                int maxNum = -1;
                boolean isValidNum = false;
                for (int j = 0; j < 4; j++) {
                    if (0 <= A[j] && A[j] <= 3 && a[j] > maxNum){
                        result[i] = maxNum;
                        isValidNum = true;
                    }
                }
                if (!isValidNum) {
                    return "";
                }
            } else {
                int maxNum = -1;
                boolean isValidNum = false;
                for (int j = 0; j < 4; j++) {
                    if (0 <= A[j] && A[j] <= 9 && a[j] > maxNum){
                        result[i] = maxNum;
                        isValidNum = true;
                    }
                }
                if (!isValidNum) {
                    return "";
                }
            }
        }
    }
}
