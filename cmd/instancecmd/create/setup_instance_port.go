package create

import (
	"context"
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	"graphdbcli/internal/tool_configurations/logging"
	tc "graphdbcli/internal/tool_configurations/statics"
	"graphdbcli/internal/tui/common_components"
	sp "graphdbcli/internal/tui/instancetui/create"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/enescakir/emoji"
	"go.uber.org/zap"
)

func setupInstancePort(ctx context.Context, ctxCancel context.CancelFunc) {
	p = tea.NewProgram(common_components.InitialModel(ctx, ctxCancel, sp.SettingUpInstancePort, &cc.Success, &cc.Failure))
	go func() {
		p.Run()
	}()

	if c.Instance.Port == "" {
		newPort, err := getFreePort()
		if err != nil {
			cc.HandleEvent(&cc.Failure, p)
			logging.LOGGER.Fatal("Unable to find a free port")
		}
		logging.LOGGER.Info("Port number was not provided, setting a random port ", zap.Int("httpPort", newPort), zap.Int("grpcPort", newPort+100))
		c.Instance.Port = strconv.Itoa(newPort)
		fmt.Printf("%s Setting a random port %d for HTTP, and %d for GRPC connectivity\n",
			common_components.PadStatusIndicator(emoji.SeeNoEvilMonkey.String(), 0),
			newPort, newPort+100)
	} else {
		if isPortOpen(c.Instance.Port) {
			fmt.Printf("%s Port %s is open\n", common_components.PadStatusIndicator(emoji.GreenCircle.String(), tc.NotTUIStatusIndicatorAdditionalPadding), c.Instance.Port)
			logging.LOGGER.Debug("using provided port", zap.String("port", c.Instance.Port))
		}
	}
	cc.HandleEvent(&cc.Success, p)
}
