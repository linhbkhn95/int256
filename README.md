# int256

Wrap [uint256](https://github.com/holiman/uint256) fixed size 256-bit math library to allow perform with negative number.

# Example usage:

### Expanding a number to 18 decimals

```go
x := int256.NewInt(123)
decimal18 := int256.MustFromDecimal("1000000000000000000")
xExpanded := new(int256.Int).Mul(x, decimal18)

fmt.Printf("%s expanded to 18 decimals is: %s", x.String(), xExpanded.String())
```

### Mostly correct implementation of UniSwap V3 getTickAtSqrtRatio using both uint256 and int256

```go
var (
    ErrInvalidSqrtRatio = errors.New("invalid sqrt ratio")
    MinSqrtRatio        = uint256.MustFromDecimal("4295128739")
    MaxSqrtRatio        = uint256.MustFromDecimal("1461446703485210103287273052203988822378723970342")
    magicSqrt10001      = int256.MustFromDecimal("255738958999603826347141")
    magicTickLow        = int256.MustFromDecimal("3402992956809132418596140100660247210")
    magicTickHigh       = int256.MustFromDecimal("291339464771989622907027621153398088495")
)

func GetTickAtSqrtRatio(sqrtPriceX96 *uint256.Int) (tick int, err error) {
    if sqrtPriceX96.Cmp(MinSqrtRatio) < 0 || sqrtPriceX96.Cmp(MaxSqrtRatio) >= 0 {
        return 0, fmt.Errorf("%w: %v", ErrInvalidSqrtRatio, sqrtPriceX96)
    }
    ratio := new(int256.Int).Lsh(int256.MustFromBig(sqrtPriceX96.ToBig()), 32)

    msb := ratio.MostSignificantBit()
    var r *int256.Int
    if msb >= 128 {
        r = new(int256.Int).Rsh(ratio, uint(msb-127))
    } else {
        r = new(int256.Int).Lsh(ratio, uint(127-msb))
    }
    log2 := new(int256.Int).Lsh(int256.NewInt(int64(int(msb)-128)), 64)

    for i := 0; i < 14; i++ {
        r = new(int256.Int).Rsh(new(int256.Int).Mul(r, r), 127)
        f := new(int256.Int).Rsh(r, 128)
        log2 = new(int256.Int).Or(log2, new(int256.Int).Lsh(f, uint(63-i)))
        r = new(int256.Int).Rsh(r, uint(f.Int64()))
    }

    logSqrt10001 := new(int256.Int).Mul(log2, magicSqrt10001)

    tickLow := new(int256.Int).Rsh(new(int256.Int).Sub(logSqrt10001, magicTickLow), 128).Int64()
    tickHigh := new(int256.Int).Rsh(new(int256.Int).Add(logSqrt10001, magicTickHigh), 128).Int64()

    if tickLow == tickHigh {
        return int(tickLow), nil
    }
	
    // GetSqrtRatioAtTick to be implemented by the user
    sqrtRatio, err := GetSqrtRatioAtTick(int(tickHigh))
    if err != nil {
        return 0, err
    }
    if sqrtRatio.Cmp(sqrtPriceX96) <= 0 {
        return int(tickHigh), nil
    } else {
        return int(tickLow), nil
    }
}
```

### Json Marshaling/UnMarshaling example

```go
type TickInfo struct {
    Initialized  bool
    LiquidityNet *int256.Int
}

type Pool struct {
    PoolAddress *common.Address
    Ticks       map[int]*TickInfo
}

func main() {
    tick1Info := &TickInfo{
        Initialized:  true,
        LiquidityNet: int256.MustFromDecimal("-111000000000000000000000000000000000099990"),
    }

    tick2Info := &TickInfo{
        Initialized:  true,
        LiquidityNet: int256.MustFromDecimal("1110000000000440000000000000000000000099990"),
    }
    ticks := make(map[int]*TickInfo)
    ticks[100] = tick2Info
    ticks[-100] = tick1Info
    poolAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")
    pool := &Pool{
        PoolAddress: &poolAddr,
        Ticks:       ticks,
    }
    data, err := json.Marshal(pool)
    if err != nil {
        fmt.Println(err)
    } else {
        // {"PoolAddress":"0x0000000000000000000000000000000000000000","Ticks":{"-100":{"Initialized":true,"LiquidityNet":"-111000000000000000000000000000000000099990"},"100":{"Initialized":true,"LiquidityNet":"1110000000000440000000000000000000000099990"}}}
        fmt.Println(string(data))
    }
    newPool := &Pool{}
    poolData := "{\"PoolAddress\":\"0x0000000000000000000000000000000000000000\",\"Ticks\":{\"-100\":{\"Initialized\":true,\"LiquidityNet\":\"-111000000000000000000000000000000000099990\"},\"100\":{\"Initialized\":true,\"LiquidityNet\":\"1110000000000440000000000000000000000099990\"}}}"
    err = json.Unmarshal([]byte(poolData), &newPool)
    if err != nil {
        fmt.Println(err)
    } else {
		// -111000000000000000000000000000000000099990
        fmt.Println(newPool.Ticks[-100].LiquidityNet.String())
		// 1110000000000440000000000000000000000099990
        fmt.Println(newPool.Ticks[100].LiquidityNet.String())
    }
}
```