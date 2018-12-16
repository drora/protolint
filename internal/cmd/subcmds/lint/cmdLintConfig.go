package lint

import (
	"github.com/yoheimuta/protolinter/internal/cmd/subcmds"
	"github.com/yoheimuta/protolinter/internal/linter/config"
	"github.com/yoheimuta/protolinter/internal/linter/file"
	"github.com/yoheimuta/protolinter/internal/linter/rule"
)

// CmdLintConfig is a config for lint command.
type CmdLintConfig struct {
	c config.Config
}

// NewCmdLintConfig creates a new CmdLintConfig.
func NewCmdLintConfig(
	c config.Config,
) CmdLintConfig {
	return CmdLintConfig{
		c: c,
	}
}

// GenRules generates rules which are applied to the filename path.
func (c CmdLintConfig) GenRules(
	f file.ProtoFile,
) ([]rule.HasApply, error) {
	var hasApplies []rule.HasApply
	for _, r := range subcmds.NewAllRules() {
		hasApplies = append(hasApplies, r)
	}
	return hasApplies, nil
}
