package main
/**
 * A command like cat(1)
 *
 *    $ cat foo.txt
 *    $ cat foo.txt bar.txt
 *    $ cat < foo.txt
 */

import (
    "os"
    "io/ioutil"
    "log"
    "io"
)

func readWrite(src io.Reader, dst io.Writer) {
    buf, err := ioutil.ReadAll(src)
    if err != nil {
        log.Fatal(err)
    }
    dst.Write(buf)
}

func main() {
    if len(os.Args) == 1 {
        readWrite(os.Stdin, os.Stdout)
    } else {
        for _,fname := range os.Args[1:] {
            fh, err := os.Open(fname)
            if err != nil {
                log.Fatal(err)
            }

            readWrite(fh, os.Stdout)
        }
    }
}

