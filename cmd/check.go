package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gowatcher_g3/config"
	"gowatcher_g3/internal/checker"
	"sync"
)

var (
	inputFilePath string
	//outputFilePath string
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Lance la fonction pour vérifier les URLs",
	Long:  "La commande 'check' parcourt une liste prédéfinie d'URLs et affiche leur statut d'accessibilité",
	Run: func(cmd *cobra.Command, args []string) {

		if inputFilePath == "" {
			fmt.Println("erreur sur le chemin du fichier d'entrée (--input")
			return
		}

		targets, err := config.LoadTargetsFromFile(inputFilePath)

		if err != nil {
			fmt.Printf("erreur lors du chargement des URLs : %v\n", err)
			return
		}

		if len(targets) == 0 {
			fmt.Println("Aucune URL à vérifier trouvée dans le fichier d'entrée")
			return
		}

		// waitgroup compteur de goroutines en attente
		var wg sync.WaitGroup
		resultsChan := make(chan checker.CheckResult, len(targets))

		wg.Add(len(targets))

		for _, url := range targets {

			go func(t config.InputTarget) {
				defer wg.Done()

				result := checker.CheckURL(t)
				resultsChan <- result // Envoyer le result au channel

			}(url)
		}
		wg.Wait()
		close(resultsChan)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
