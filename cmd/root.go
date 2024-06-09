/*
 * Minecraft Server Tool(MCST) is a command-line utility making Minecraft server creation quick and easy for beginners.
 * Copyright (c) 2024-2024 Arama.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package cmd

import (
	"github.com/Arama-Vanarana/MCServerTool/pkg/lib"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	goversion "github.com/caarlos0/go-version"
	"github.com/spf13/cobra"
)

func Execute(exit func(int), args []string, version goversion.Info) error {
	cmd := newRootCmd(version)
	cmd.SetArgs(args)
	if err := cmd.Execute(); err != nil {
		log.WithError(err).Error("错误")
		exit(1)
	}
	return nil
}

func newRootCmd(version goversion.Info) *cobra.Command {
	var verbose bool
	cmd := &cobra.Command{
		Use:           "MCST",
		Long:          version.ASCIIName,
		Version:       version.String(),
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(_ *cobra.Command, args []string) error {
			log.SetHandler(cli.Default)
			if verbose {
				log.SetLevel(log.DebugLevel)
				log.WithField("参数", args).Debug("调试模式开启")
			}
			return lib.InitAll(version)
		},
		PersistentPostRun: func(*cobra.Command, []string) {
			log.Info("感谢使用MCST!")
		},
	}
	cmd.SetVersionTemplate("{{.Version}}")
	cmd.PersistentFlags().BoolVar(&verbose, "debug", false, "调试模式(更多的日志)")
	cmd.AddCommand(newCreateCmd(), newDownloadCmd(), newConfigCmd(), newStartCmd(), newListCmd(), newManCmd())
	return cmd
}