package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// go run main.go cmd args
// need to run with root user for permition of operat
// rootfs download from https://github.com/MoimHossain/scratch-container
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what??")
	}
}

func run() {
	fmt.Printf("running %v as PID %d\n", os.Args[2], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2], os.Getpid())

	cg()

	_ = syscall.Sethostname([]byte("container-test"))
	_ = syscall.Chroot("/mnt/e/projector-user/rootfs/ubuntu-rootfs")
	_ = syscall.Chdir("/")
	_ = syscall.Mount("proc", "proc", "proc", 0, "")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	must(cmd.Run())

	_ = syscall.Unmount("proc", 0)
}

func cg() {
	cgroup := "/sys/fs/cgroup"
	pids := filepath.Join(cgroup, "pids")
	err := os.Mkdir(filepath.Join(pids, "aseara"), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	// for test proc limit run :/# :() { : | : & }; :
	must(ioutil.WriteFile(filepath.Join(pids, "aseara/pids.max"), []byte("20"), 0700))
	// remove the new cgroup in place after the container exists
	must(ioutil.WriteFile(filepath.Join(pids, "aseara/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "aseara/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
