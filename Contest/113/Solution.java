import java.util.*;

class Solution {
    public static String largestTimeFromDigits(int[] A) {
        int[] result = new int[4];

        int max = -1;
        int theIndex = -1;

        // 1
        for (int i = 0; i < 4; i++) {
            if (0 <= A[i] && A[i] <= 2 && A[i] > max) {
                max = A[i];
                theIndex = i;
                result[0] = max;
            }
        }
        if (theIndex == -1) {
            return "";
        } else {
            max = -1;
            A[theIndex] = -1;
            theIndex = -1;
        }
        System.out.println(Arrays.toString(result));
        
        // 2
        for (int i = 0; i < 4; i++) {
            if (result[0] == 2){
                if (0 <= A[i] && A[i] <= 3 && A[i] > max) {
                    max = A[i];
                    theIndex = i;
                    result[1] = max;
                }
            } else if (0 <= A[i] && A[i] <= 9 && A[i] > max) {
                max = A[i];
                theIndex = i;
                result[1] = max;
            }
        }
        if (theIndex == -1) {
            return "";
        } else {
            max = -1;
            A[theIndex] = -1;
            theIndex = -1;
        }
        System.out.println(Arrays.toString(result));

        // 3
        for (int i = 0; i < 4; i++) {
            if (0 <= A[i] && A[i] <= 5 && A[i] > max) {
                max = A[i];
                theIndex = i;
                result[2] = max;
            }
        }
        if (theIndex == -1) {
            return "";
        } else {
            max = -1;
            A[theIndex] = -1;
            theIndex = -1;
        }
        System.out.println(Arrays.toString(result));

        // 4
        for (int i = 0; i < 4; i++) {
            if (0 <= A[i] && A[i] <= 9 && A[i] > max) {
                max = A[i];
                theIndex = i;
                result[3] = max;
            }
        }
        if (theIndex == -1) {
            return "";
        } else {
            max = -1;
            A[theIndex] = -1;
            theIndex = -1;
        }
        System.out.println(Arrays.toString(result));

        String res = "" + result[0] + result[1] + ":" + result[2] + result[3];
        return res;
    }

    public static void main(String[] args) {
        int[] A = new int[]{2, 0, 6, 6};

        String res = largestTimeFromDigits(A);

        System.out.println(res);
    }
}
