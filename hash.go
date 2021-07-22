package skiplist

import (
    "math"
    "time"
)

func hash(k interface{}) uint64 {
    if k == nil {
        return 0
    }
    switch x := k.(type) {
    case string:
        return bytesHash([]byte(x))
    case []byte:
        return bytesHash(x)
    case bool:
        if x {
            return 0
        } else {
            return 1
        }
    case time.Time:
        return uint64(x.UnixNano())
    case int:
        return uint64(x)
    case int8:
        return uint64(x)
    case int16:
        return uint64(x)
    case int32:
        return uint64(x)
    case int64:
        return uint64(x)
    case uint:
        return uint64(x)
    case uint8:
        return uint64(x)
    case uint16:
        return uint64(x)
    case uint32:
        return uint64(x)
    case uint64:
        return x
    case float32:
        return math.Float64bits(float64(x))
    case float64:
        return math.Float64bits(x)
    case uintptr:
        return uint64(x)
    }
    panic("unsupported key type.")
}

func bytesHash(bytes []byte) uint64 {
    hash := uint32(2166136261)
    const prime32 = uint32(16777619)
    keyLength := len(bytes)
    for i := 0; i < keyLength; i++ {
        hash *= prime32
        hash ^= uint32(bytes[i])
    }
    return uint64(hash)
}