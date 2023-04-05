package jenkins

import (
	"context"
	"sync"

	"github.com/bndr/gojenkins"
	"github.com/sirupsen/logrus"
)

var jenkins *gojenkins.Jenkins

var once sync.Once

func init() {
	once.Do(
		func() {
			jenkins = gojenkins.CreateJenkins(nil, "")
		},
	)
}

func CreateJob() (*gojenkins.Job, error) {
	logrus.Info("[jenkins][CreateJob] version", jenkins.Version)

	ctx := context.Background()
	config := ""
	job, err := jenkins.CreateJob(ctx, config)
	if err != nil {
		logrus.Error("[jenkins][CreateJob] error: ", err)
		return nil, err
	}
	return job, nil
}
