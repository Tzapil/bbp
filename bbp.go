// Package contains implementation of The BBP Algorithm for Pi
package bbp

import "io"
import "io/ioutil"
import "math"
import "bytes"

//import "time"
//import "fmt"


// Using IEEE 64-bit floating-point arithmetic yields 8 or more correct hex digits
const size = 8
// It takes 2m to compute 8 digits of pi beginning at an position 10^6
const max_value = 10e6

// Exponentiation performed over a modulus
func modulus_pow(base float64, exp int, modulus float64) float64 {
    // border case
    if modulus == 1 {
        return 0
    }

    var result float64 = 1
    var z float64 = math.Mod(base, modulus)

    for exp > 0 {
        var b byte = byte(exp & 1)
        exp = exp >> 1

        if b == 1 {
            result = math.Mod(result * z, modulus)
        }

        z = math.Mod(z * z, modulus)
    }

    return result
}

// Computing part of the sum
func series(n, m int, out chan<- float64) {
    var eps float64 = 1e-17
    var result float64
    
    //start := time.Now()
    for k := 0; k <= n; k++ {
        var base float64 = 8.0 * float64(k) + float64(m)
        var upside float64 = modulus_pow(16.0, n - k, base)
        var sum float64 = upside / base

        result += sum

        //f := time.Since(start)
        //fmt.Println("modulus: ", f)

        // remove integer part
        result = result - float64(int(result))
    }

    

    for k := n + 1; k <= n + 100; k++ {
        var base float64 = 8.0 * float64(k) + float64(m)
        var upside float64 = math.Pow(16.0, float64(n - k))
        var sum float64 = upside / base

        if sum < eps {
            break
        }

        result += sum

        // remove integer part
        result = result - float64(int(result))
    }

    out <- result
}

// Perform multiplying by 16 to "skim off" the hexadecimal digit at this position
func ihex(x float64, nhx int) []byte{
  var y float64 = math.Abs(x)
  var hx string = "0123456789ABCDEF"
  var result []byte = make([]byte, nhx)

  for i := 0; i < nhx; i++ {
    y = 16.0 * (y - float64(int(y)))
    result[i] = hx[int(y)]
  }

  return result
}

// BBP return 8 hexadecimal digits of pi beginning at an arbitrary starting position id
func bbp10(id int) []byte {
    // calculate series
    ch1 := make(chan float64)
    go series(id, 1, ch1)
    ch2 := make(chan float64)
    go series(id, 4, ch2)
    ch3 := make(chan float64)
    go series(id, 5, ch3)
    ch4 := make(chan float64)
    go series(id, 6, ch4)

    s1 := <- ch1
    s2 := <- ch2
    s3 := <- ch3
    s4 := <- ch4

    // calculate result sum
    pid := 4 * s1 - 2 * s2 - s3 - s4

    // remove integer part
    pid = pid - float64(int(pid)) + 1.0

    // multiply by 16 and get 10 pi numbers
    var r []byte = ihex (pid, size)

    return r
}

func BBP(id int, l int) string {
    var count int = l / size
    var r []io.Reader = make([]io.Reader, count, count + 1)
    for i := 0; i < count; i++ {
        r[i] = bytes.NewReader(bbp10(id + i * size))
    }

    var rest int = l - count * size
    if rest != 0 {
        r = append(r, bytes.NewReader(bbp10(id + count * size)[:rest]))
    }

    result, _ := ioutil.ReadAll(io.MultiReader(r...))

    return string(result)
}