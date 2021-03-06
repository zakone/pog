package popcount

var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func PopCountLoop(x uint64) int {
    var tmp byte
    var i uint8
    for i < 8 {
        tmp += pc[byte(x>>(i*8))]
        i++
    }
    return int(tmp)
}

func PopCount(x uint64) int {
    return int(pc[byte(x>>(0*8))] +
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] +
        pc[byte(x>>(3*8))] +
        pc[byte(x>>(4*8))] +
        pc[byte(x>>(5*8))] +
        pc[byte(x>>(6*8))] +
        pc[byte(x>>(7*8))])
}

func PopCountOnlyLoop(x uint64) int {
    var tmp byte
    for i := uint8(0); i < 64; i++ {
        tmp += byte((x >> i) & 1)
    }
    return int(tmp)

}

func PopCountLastClear(x uint64) int {
    var tmp byte
    for x != 0 {
        x = x & (x - 1)
        tmp += 1
    }
    return int(tmp)
}
