package snapshot

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

const timeFormat = "01-02-2006_15-04-05"

func Snapshot(snapRoot, vol string) error {
	snap, err := initSnapPath(snapRoot, vol)
	if err != nil {
		return err
	}

	cmd := exec.Command("btrfs", "subvolume", "snapshot", vol, snap)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("created new snapshot of %s at %s\n", vol, snap)
	return nil
}

func initSnapPath(snapRoot string, vol string) (string, error) {
	err := exec.Command("mkdir", "-p", snapRoot).Run()
	if err != nil {
		return "", err
	}

	_, volName := path.Split(vol)
	return path.Join(snapRoot, volName+"."+time.Now().Format(timeFormat)), nil
}
