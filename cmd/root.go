package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goWatcher",
	Short: "GoWatcher est un outil pour vérifier l'accessibilité des URLs.",
	Long:  "un outil CLI en Go pour vérifier l'état d'un URL, gérer la conccurence et exporter les résultats.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur : %v\n", err)
		os.Exit(1)
	}
}
