package command

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"

	"gomander/internal/event"
	"gomander/internal/logger"
	"gomander/internal/platform"
)

type Runner struct {
	runningCommands map[string]*exec.Cmd
	eventEmitter    *event.EventEmitter
	logger          *logger.Logger
}

func NewCommandRunner(logger *logger.Logger, emitter *event.EventEmitter) *Runner {
	return &Runner{
		runningCommands: make(map[string]*exec.Cmd),
		eventEmitter:    emitter,
		logger:          logger,
	}
}

// ExecCommand executes a command by its ID and streams its output.
func (c *Runner) RunCommand(command Command, extraPaths []string) error {
	// Get the command object based on the command string and OS
	cmd := platform.GetCommand(command.Command)

	// Enable color output and set terminal type
	cmd.Env = append(os.Environ(), "FORCE_COLOR=1", "TERM=xterm-256color")
	cmd.Dir = command.WorkingDirectory

	// Set command attributes based on OS
	platform.SetProcAttributes(cmd)
	platform.SetProcEnv(cmd, extraPaths)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		c.sendStreamError(command, err)
		return err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		c.sendStreamError(command, err)
		return err
	}

	if err := cmd.Start(); err != nil {
		c.sendStreamError(command, err)
		return err
	}

	// Save the command in the runningCommands map
	c.runningCommands[command.Id] = cmd

	// Stream stdout
	go c.streamOutput(command.Id, stdout)
	// Stream stderr
	go c.streamOutput(command.Id, stderr)

	// Optional: Wait in background
	go func() {
		err := cmd.Wait()
		if err != nil {
			c.sendStreamError(command, err)
			c.logger.Error("[ERROR - Waiting for command]: " + err.Error())
			return
		}
		c.eventEmitter.EmitEvent(event.ProcessFinished, command.Id)
	}()

	return nil
}

func (c *Runner) StopRunningCommand(id string) error {
	runningCommand, exists := c.runningCommands[id]

	if !exists {
		return errors.New("No running runningCommand for command: " + id)
	}

	return platform.StopProcessGracefully(runningCommand)
}

func (c *Runner) streamOutput(commandId string, pipeReader io.ReadCloser) {
	scanner := bufio.NewScanner(pipeReader)

	for scanner.Scan() {
		line := scanner.Text()
		c.logger.Debug(line)

		c.sendStreamLine(commandId, line)
	}
}

func (c *Runner) sendStreamError(command Command, err error) {
	c.sendStreamLine(command.Id, err.Error())
	c.logger.Error(err.Error())
	c.eventEmitter.EmitEvent(event.ProcessFinished, command.Id)
}

func (c *Runner) sendStreamLine(commandId string, line string) {
	c.eventEmitter.EmitEvent(event.NewLogEntry, map[string]string{
		"id":   commandId,
		"line": line,
	})
}
