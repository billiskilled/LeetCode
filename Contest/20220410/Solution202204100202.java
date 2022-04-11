import java.util.HashMap;
import java.util.Map;

public class Solution202204100202 {
    public int numFlowers(int[][] roads) {
        Map<Integer, Integer> map = new HashMap<>();

        int len = roads.length;
        for (int i = 0; i < len; i++) {
            map.putIfAbsent(roads[i][0], 0);
            map.put(roads[i][0], map.get(roads[i][0]) + 1);

            map.putIfAbsent(roads[i][1], 0);
            map.put(roads[i][1], map.get(roads[i][1]) + 1);
        }

        int max = Integer.MIN_VALUE;
        for (Map.Entry<Integer, Integer> e : map.entrySet()) {
            if (e.getValue() > max) {
                max = e.getValue();
            }
        }

        return max + 1;
    }
}
