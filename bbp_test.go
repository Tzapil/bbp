package bbp

import "testing"
import "time"

func TestBBP10Length(t *testing.T) {
    cases := []struct {
        in int
        l int
        want string
    }{
        {0, 10, "243F6A8885"},
        {950, 10, "62363F7706"},
        {69950, 10, "C04ED1CF1E"},
        {100000, 10, "35EA16C406"},
        {399950, 10, "5C47DE610F"},
        {1000000, 10, "6C65E52CB4"},
        //{10e7, 10, "5895585A04"},
    }

    for _, c := range cases {
        start := time.Now()
        got := BBP(c.in, c.l)
        f := time.Since(start)
        t.Logf("BBP(%d, %d) = %q in time %v", c.in, c.l, got, f)
        if got != c.want {
            t.Errorf("BBP(%d, %d) == %q, want %q", c.in, c.l, got, c.want)
        }
    }
}

func TestBBPRandomLength(t *testing.T) {
    cases := []struct {
        in int
        l int
        want string
    }{
        {950, 20, "62363F77061BFEDF7242"},
        {49950, 15, "08CE5DB76425C7B"},
        {199950, 14, "3E2ED27C44BC12"},
        {7000, 7, "B155FDF"},
        {399950, 33, "5C47DE610F9004F270F66FE91A3E4148B"},
    }

    for _, c := range cases {
        start := time.Now()
        got := BBP(c.in, c.l)
        f := time.Since(start)
        t.Logf("BBP(%d, %d) = %q in time %v", c.in, c.l, got, f)
        if got != c.want {
            t.Errorf("BBP(%d, %d) == %q, want %q", c.in, c.l, got, c.want)
        }
    }
}