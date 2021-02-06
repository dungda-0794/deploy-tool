package deploy

import (
	"fmt"
	"github.com/dung13890/deploy-tool/cmd/task"
)

func Fetch(t *task.Task) error {
	path := t.Dir()
	fmt.Printf("%s/release", path)

	// Update code at release_path on host.
	cmd := fmt.Sprintf("git clone git@github.com:dung13890/deploy-tool.git %s/release 2>&1", path)
	err := t.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}
