package system

import (
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"syscall"
	"time"

	"github.com/golang/glog"
)

var enablePProf bool

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "/data/logKingtalk")
	flag.Set("v", "3")
	enablePProf = *flag.Bool("pprof_enable", true, "Enable profiling")
	flag.Parse()
}

// ------------------------------------------------------------------------------------------------
type ServicesInterface interface {
	Install() error 
	Start()
	Shutdown(signals chan os.Signal)
}

func RunAppService(instance ServicesInterface) {
	defer glog.Flush()

	// 1. Setup the runtime CPU
	if runtime.NumCPU() > 1 {
		runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	}

	// 2. Seed the time to random
	rand.Seed(time.Now().UnixNano())

	// 3. Install the instance
	if err := instance.Install(); err != nil {
		glog.V(1).Infof("Error installing instance: %v", err)
	}

	// 4. Run the instance's loop
	go instance.Start()
	glog.V(2).Infof("go instance.Start(), successfully started)")

	// 5. Handle signals for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	instance.Shutdown(signals)

	if enablePProf {
		pprof.StopCPUProfile()
	}
}
