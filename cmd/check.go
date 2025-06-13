package cmd

import (
	"errors"
	"fmt"
	"gowatcher_g3/config"
	"gowatcher_g3/internal/checker"
	"gowatcher_g3/reporter"
	"sync"

	"github.com/spf13/cobra"
)

var (
	inputFilePath  string
	outputFilePath string
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Lance la fonction pour vérifier les URLs",
	Long:  `La commande 'check' parcourt une liste prédéfinie d'URLs et affiche leur statut d'accessibilité en utilisant des goroutines pour la concurrence.`,
	Run: func(cmd *cobra.Command, args []string) {

		if inputFilePath == "" {
			fmt.Println("Erreur sur le chemin du fichier d'entrée (--input)")
			return
		}

		targets, err := config.LoadTargetsFromFile(inputFilePath)
		if err != nil {
			fmt.Printf("Erreur lors du chargement des URLs: %v\n", err)
			return
		}

		if len(targets) == 0 {
			fmt.Println("Aucune URL à vérifier trouvée dans le fichier d'entrée.")
			return
		}

		// On crée un waitgroup qui est un compteur de Goroutines en attente
		var wg sync.WaitGroup
		resultsChan := make(chan checker.CheckResult, len(targets))

		wg.Add(len(targets))
		for _, url := range targets {
			go func(t config.InputTarget) {
				defer wg.Done()
				result := checker.CheckURL(t)
				resultsChan <- result // Envoyer le résultat au channel
			}(url)
		}
		// Cette ligne bloque l'éxecution du main() jusqu'à ce que toutes les goroutines aient appelé wd.Done()
		wg.Wait()
		close(resultsChan) // Fermer le canal après que tous les résultats ont été envoyés

		var finalReport []checker.ReportEntry
		for res := range resultsChan { // Récupérer tous les résultats du channel
			reportEntry := checker.ConvertToReportEntry(res)
			finalReport = append(finalReport, reportEntry)

			// Affichage immédiat comme avant
			if res.Err != nil {
				var unreachable *checker.UnreachableURLError
				if errors.As(res.Err, &unreachable) {
					fmt.Printf("🚫 %s (%s) est inaccessible : %v\n", res.InputTarget.Name, unreachable.URL, unreachable.Err)
				} else {
					fmt.Printf("❌ %s (%s) : erreur - %v\n", res.InputTarget.Name, res.InputTarget.URL, res.Err)
				}
			} else {
				fmt.Printf("✅ %s (%s) : OK - %s\n", res.InputTarget.Name, res.InputTarget.URL, res.Status)
			}

			// Exporter les résultats si outputFilePath est spécifié
			if outputFilePath != "" {
				err := reporter.ExportResultsToJsonfile(outputFilePath, finalReport)
				if err != nil {
					fmt.Printf("Erreur lors de l'exportation des résultats: %v\n", err)
				} else {
					fmt.Printf("Résultats exportés vers %s\n", outputFilePath)
				}
			}
		}
	},
}

func init() {
	// Cette ligne est cruciale : elle "ajoute" la sous-commande `checkCmd` à la commande racine `rootCmd`.
	// C'est ainsi que Cobra sait que 'check' est une commande valide sous 'gowatcher'.
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "Absolute path to JSON file containing targets to check")
	checkCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "Absolute path to JSON file containing targets to export")

	checkCmd.MarkFlagRequired("input")
}
