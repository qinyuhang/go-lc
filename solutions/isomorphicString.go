package solutions

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mapS2T := make(map[byte]byte)
    mapT2S := make(map[byte]byte)
    for i := range len(s) {
        sChar := s[i]
        tChar := t[i]
        if val, ok := mapS2T[sChar]; ok && val != tChar {
            return false
        }
        mapS2T[sChar] = tChar
        if val, ok := mapT2S[tChar]; ok && val != sChar {
            return false
        }
        mapT2S[tChar] = sChar
    }
    return true
}
