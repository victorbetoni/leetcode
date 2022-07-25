import java.util.Map;
import java.util.Set;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedHashMap;
import java.util.List;

class Solution {
    public int findLeastNumOfUniqueInts(int[] arr, int k) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int j : arr) {
            if (!map.containsKey(j)) {
                map.put(j, 1);
            } else {
                Integer current = map.get(j);
                map.remove(j);
                map.put(j, current + 1);
            }
        }
        List<Map.Entry<Integer, Integer>> list = new ArrayList<>(map.entrySet());
        list.sort(Map.Entry.comparingByValue());
        LinkedHashMap<Integer, Integer> result = new LinkedHashMap<>();
        for (Map.Entry<Integer, Integer> entry : list) {
            result.put(entry.getKey(), entry.getValue());
        }
        Set<Integer> toRemove = new HashSet<>();
        for (Map.Entry<Integer, Integer> entry : result.entrySet()) {
            if(entry.getValue() - k <= 0) {
                k-=entry.getValue();
                toRemove.add(entry.getKey());
            } else {
                k-=entry.getValue();
            }
        }
        toRemove.forEach(result::remove);
        return result.size();
    }
}