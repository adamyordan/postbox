package postbox

import (
	"github.com/sevlyar/go-daemon"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"syscall"
)

func getContext() *daemon.Context {
	return &daemon.Context{
		PidFileName: path.Join(getPostboxDir(), "postbox.pid"),
		PidFilePerm: 0644,
		LogFileName: path.Join(getPostboxDir(), "postbox.log"),
		LogFilePerm: 0640,
		WorkDir: getPostboxDir(),
		Umask: 027,
	}
}

func IsDaemonRunning(ctx *daemon.Context) (bool, *os.Process, error) {
	d, err := ctx.Search()
	if err != nil {
		return false, d, err
	}
	if err := d.Signal(syscall.Signal(0)); err != nil {
		return false, d, err
	}
	return true, d, nil
}

func StartDaemon(f func() error) {
	ctx := getContext()
	if ok, _, _ := IsDaemonRunning(ctx); ok {
		log.Fatalf("service is already running.")
	}

	log.Info("starting daemon...")
	d, err := ctx.Reborn()
	if err != nil {
		log.Fatalf("unable to run: %v", err)
	}

	if d != nil {
		return
	}

	defer ctx.Release()

	log.Info("- - - - - - - - - - - - - - -")
	log.Info("daemon started")
	log.Fatal(f())
}

func StopDaemon() {
	ctx := getContext()

	if ok, p, _ := IsDaemonRunning(ctx); ok {
		log.Info("stopping daemon")
		err := p.Signal(syscall.Signal(syscall.SIGQUIT))
		if err != nil {
			log.Fatalf("failed to kill daemon %v", err)
		}
		log.Info("stopped daemon")
	} else {
		ctx.Release()
		log.Fatalf("service is not running.")
	}
}
