package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	url      = "http://localhost:8001"
	admin    = "admin"
	password = "bbe4fffb857b469a9d92e7c45c5576d2"
)

var (
	jenkins *gojenkins.Jenkins
)

var once sync.Once

func init() {
	ctx := context.Background()
	jenkins = gojenkins.CreateJenkins(nil, url, admin, password)
	_, err := jenkins.Init(ctx)
	if err != nil {
		logrus.Error("Something Went Wrong: ", err)
		return
	}

	logrus.Info("[jenkins][init] success")
}

func CreateJob(ctx context.Context, repo string) {
	logrus.Info("[jenkins][CreateJob] version", jenkins.Version)

	repo = "https://github.com/gogolang20/template.git"
	jobID := "job_test"
	config := "./job_temp.xml"
	_, err := jenkins.CreateJob(ctx, config, jobID)
	if err != nil {
		logrus.Error("[jenkins][CreateJob] error: ", err)
		return
	}

	// job, err := jenkins.GetJob(ctx, "#jobname")
	// if err != nil {
	// 	panic(err)
	// }
	// queueid, err := job.InvokeSimple(ctx, params) // or  jenkins.BuildJob(ctx, "#jobname", params)
	// if err != nil {
	// 	panic(err)
	// }
	// build, err := jenkins.GetBuildFromQueueID(ctx, job, queueid)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Wait for build to finish
	// for build.IsRunning(ctx) {
	// 	time.Sleep(5000 * time.Millisecond)
	// 	build.Poll(ctx)
	// }
	//
	// fmt.Printf("build number %d with result: %v\n", build.GetBuildNumber(), build.GetResult())
}
