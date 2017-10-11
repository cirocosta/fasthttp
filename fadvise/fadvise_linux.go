// +build linux
package fadvise

import (
  "golang.org/x/sys/unix"
  "os"
)

func AdviseFile(f *os.File) (err error) {
  err = unix.Fadvise(int(f.Fd()), 0, 0, 4)
  if err != nil {
    return
  }

  return
}
