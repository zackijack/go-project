package cmd

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zackijack/go-project/internal/config"
	"github.com/zackijack/go-project/internal/helpers"
)

const startDesc = `
Start application on the HTTP-server.
By default the HTTP-server will listen to host 0.0.0.0 & port 8080.
However if you want to specify what host & port to use, just set via:

  - argument:

    $ go-project start --host localhost --port 80

  - env variable:

    $ export APP_HOST=localhost
    $ export APP_PORT=80
`

var startCmd = &cobra.Command{
	Use:   "start",
	Aliases: []string{"run", "serve", "server"},
	Short: "Start the application",
	Long: startDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			zlog.Error().Err(err).Msg(helpers.ErrMsg("config"))
		}

		start(cfg)
	},
}

func init() {
	startCmd.Flags().StringP("host", "H", "0.0.0.0", "host address to serve the application on")
	startCmd.Flags().IntP("port", "P", 8080, "port to serve the application on")

	// Bind to config.
	helpers.CheckErr(viper.BindPFlag("APP_HOST", startCmd.Flags().Lookup("host")), helpers.ErrMsg("config flag: host"), false)
	helpers.CheckErr(viper.BindPFlag("APP_PORT", startCmd.Flags().Lookup("port")), helpers.ErrMsg("config flag: port"), false)

	rootCmd.AddCommand(startCmd)
}


func start(cfg *config.Config) {
	addr := fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppPort)

	app := fiber.New()

	// Get /johnny.
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋! from %s with ❤️", c.Params("name"), cfg.AppName)
		return c.SendString(msg) // => Hello, johnny 👋! from go-project with ❤️
	})

	zlog.Fatal().Err(app.Listen(addr)).Msg(helpers.ErrMsg("server"))
}
