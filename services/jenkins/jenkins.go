package jenkins

import (
	"github.com/bndr/gojenkins"
)

var jenkins *gojenkins.Jenkins

// var once sync.Once

func init() {
	jenkins = gojenkins.CreateJenkins(nil, "")

}
