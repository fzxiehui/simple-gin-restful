package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/gin-gonic/gin"
	"github.com/fzxiehui/simple-gin-restful/routers"
	// "github.com/fzxiehui/simple-gin-restful/global"
)

var caseCmd = &cobra.Command{
	Use:   "case",
	Short: "case",
	Long:  `case`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("case called")
		// global.InitDB()
		// defer global.CloseDB()
		r := gin.Default() // 创建路由
		routers.InitRouter(r)
		r.Run(":8080") // 监听并在 8080
	},
}

func init() {
	rootCmd.AddCommand(caseCmd)
}

