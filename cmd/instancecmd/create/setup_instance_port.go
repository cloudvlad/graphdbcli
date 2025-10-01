package create

import (
	"context"
	"fmt"
	cc "graphdbcli/internal/channels/commons"
	c "graphdbcli/internal/data_objects/graphdb_cluster"
	"graphdbcli/internal/tool_configurations/logging"
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
		logging.LOGGER.Info("Port number was not provided, setting a random one ", zap.Int("httpPort", newPort), zap.Int("grpcPort", newPort+100))
		c.Instance.Port = strconv.Itoa(newPort)
		fmt.Printf("%s Setting port %d for HTTP, and %d for GRPC connectivity\n",
			common_components.PadStatusIndicator(emoji.SeeNoEvilMonkey.String(), 0),
			newPort, newPort+100)
	} else {
		if isPortOpen(c.Instance.Port) {
			fmt.Printf("%s Port %s is open\n", common_components.PadStatusIndicator(emoji.GreenCircle.String(), 0), c.Instance.Port)
			logging.LOGGER.Debug("using provided port", zap.String("port", c.Instance.Port))
		} else {
			cc.HandleEvent(&cc.Failure, p)
			fmt.Printf("%s Port %s is not open\n", common_components.PadStatusIndicator(emoji.RedCircle.String(), 0), c.Instance.Port)
			CleanUp(c.Instance.Name)
			logging.LOGGER.Fatal("provided port is not free", zap.String("port", c.Instance.Port))
		}
	}
	cc.HandleEvent(&cc.Success, p)
}
