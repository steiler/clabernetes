package prepare

import (
	"fmt"

	clabernetesmanagertypes "github.com/srl-labs/clabernetes/manager/types"
	clabernetesutil "github.com/srl-labs/clabernetes/util"
)

// Prepare handles preparation tasks that happen before running the clabernets.start method.
func Prepare(c clabernetesmanagertypes.Clabernetes) {
	logger := c.GetBaseLogger()

	logger.Info("preparing certificates...")

	err := certificates(c)
	if err != nil {
		msg := fmt.Sprintf("failed preparing certificates, err: %s", err)

		logger.Critical(msg)

		clabernetesutil.Panic(msg)
	}

	logger.Debug("preparing certificates complete...")
}
