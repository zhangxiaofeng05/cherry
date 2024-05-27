package cmd

import (
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"log"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "generate uuid",
	Long:  `A command line tool for generate uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		newUUID, err := uuid.NewUUID()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(newUUID.String())
	},
}
