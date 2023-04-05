package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	jenkins *gojenkins.Jenkins
)

var once sync.Once

func init() {
	ctx := context.Background()
	jenkins = gojenkins.CreateJenkins(nil, "http://localhost:8001", "admin", "bbe4fffb857b469a9d92e7c45c5576d2")
	_, err := jenkins.Init(ctx)
	if err != nil {
		logrus.Error("Something Went Wrong: ", err)
		return
	}

	logrus.Info("[jenkins][init] success")
}

func CreateJob() {
	logrus.Info("[jenkins][CreateJob] version", jenkins.Version)

	ctx := context.Background()
	config := ""
	_, err := jenkins.CreateJob(ctx, config)
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
