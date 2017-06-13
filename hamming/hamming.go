package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (int, error) {
    var dif int

    // Return error if they are different lengths
    if len(a) != len(b) {
        return -1, errors.New("DNAs are not the same length")
    }

    // Check letter by letter in DNAs
    for i := range a {
        if a[i] != b[i] {
            dif++
        }
    }

    return dif, nil
}
