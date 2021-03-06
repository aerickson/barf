package runners

import (
	"barf/internal/cmd"
	"barf/internal/op"
	"barf/internal/rsync"
	"barf/internal/typeconv"
)

type copyRunner struct {
	operation     *op.Operation
	rsync         *rsync.Rsync
	stdoutHandler cmd.LogHandler
	stderrHandler cmd.LogHandler
	statusHandler statushandler
}

func (r *copyRunner) init(operation *op.Operation) {
	r.operation = operation
	r.rsync = rsync.NewRsync()
	r.rsync.OnStdout(r.handleStdout)
	r.rsync.OnStderr(r.handleStderr)
	r.rsync.OnStatus(r.handleStatus)
}

func (r *copyRunner) Start() {
	args := []string{}
	srcArray, _ := typeconv.ToArray(r.operation.Args["src"])
	src := typeconv.ToStringArray(srcArray)
	dst := r.operation.Args["dst"].(string)

	r.rsync.Copy(args, src, dst)
}

func (r *copyRunner) Abort() error {
	return r.rsync.Abort()
}

func (r *copyRunner) OperationID() op.OperationID {
	return r.operation.ID
}

func (r *copyRunner) OnStdout(handler cmd.LogHandler) {
	r.stdoutHandler = handler
}

func (r *copyRunner) OnStderr(handler cmd.LogHandler) {
	r.stderrHandler = handler
}

func (r *copyRunner) OnStatus(handler statushandler) {
	r.statusHandler = handler
}

func (r *copyRunner) handleStdout(line string) {
	r.stdoutHandler(line)
}

func (r *copyRunner) handleStderr(line string) {
	r.stderrHandler(line)
}

func (r *copyRunner) handleStatus(status *op.OperationStatus) {
	r.statusHandler(status)
}
