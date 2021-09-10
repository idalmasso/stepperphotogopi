package main

//NOTE: glog levels:
// V=1 for ERROR
// V=2 for WARNING
// V=3 for INFO
// V=4 for DEBUG

import (
	"flag"

	"github.com/golang/glog"
	"github.com/idalmasso/stepperphotogopi/backend/server"
	"github.com/idalmasso/stepperphotogopi/backend/server/hwdummy"
)

func init() {
	flag.Parse()
	if !isFlagPassed("v") {
		flag.Set("v", "3")
	}
	flag.Set("logtostderr", "1")

}

func main() {
	defer glog.Flush()
	if glog.V(3) {
		glog.Infoln("backend start process")
	}
	//piController := hwinterface.NewController()
	dummyControler := hwdummy.NewController()
	server := server.MachineServer{}
	server.Init(&dummyControler)
	server.ListenAndServe()
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
