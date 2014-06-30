package main

import (
  "testing"
  . "github.com/onsi/gomega"
)

func TestZeroComp(t *testing.T) {
  actual := ZeroComp("3")
  Expect(actual).To(Equal("03"))

  actual = ZeroComp("30")
  Expect(actual).To(Equal("30"))
}
