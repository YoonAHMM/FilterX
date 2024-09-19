package pkg

import "sort"

func SortMap(mp map[int32]int) []int32 {
    
    type kv struct {
        Key   int32
        Value int
    }
    var sortedPairs []kv
    for k, v := range mp {
        sortedPairs = append(sortedPairs, kv{k, v})
    }

    
    sort.Slice(sortedPairs, func(i, j int) bool {
        return sortedPairs[i].Value < sortedPairs[j].Value
    })

   
    var sortedKeys []int32
    for _, pair := range sortedPairs {
        sortedKeys = append(sortedKeys, pair.Key)
    }

    return sortedKeys
}