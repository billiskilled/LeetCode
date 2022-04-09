import com.sun.xml.internal.ws.api.ha.HaInfo;

import java.util.*;

public class Solution2022040302 {
    public List<List<Integer>> findWinners(int[][] matches) {
        if (matches == null) {
            return null;
        }

        int len = matches.length;
        Map<Integer, Integer> map = new HashMap<>();
        for (int[] cur : matches) {
            if (cur == null) {
                continue;
            }
            map.putIfAbsent(cur[0], 0);
            if (map.get(cur[1]) == null) {
                map.put(cur[1], 1);
            } else {
                map.put(cur[1], map.get(cur[1]) + 1);
            }
        }
        List<Integer> win = new ArrayList<>();
        List<Integer> loser = new ArrayList<>();

        for (Map.Entry<Integer, Integer> e : map.entrySet()) {
            if (e.getValue() == 0) {
                win.add(e.getKey());
            }
            if (e.getValue() == 1) {
                loser.add(e.getKey());
            }
        }

        List<List<Integer>> result = new ArrayList<>();
        Collections.sort(win);
        result.add(win);
        Collections.sort(loser);
        result.add(loser);
        return result;
    }
}
