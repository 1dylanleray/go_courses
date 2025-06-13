package cmd

import (
	"errors"
	"fmt"
	"gowatcher_g3/config"
	"os"

	"github.com/spf13/cobra"
)

var (
	addFilePath string
	addName     string
	addURL      string
	addOwner    string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajoute une nouvelle URL à un fichier JSON de configuration.",
	Long: `La commande 'add' permet d'ajouter une nouvelle cible (URL, nom, propriétaire)
à un fichier JSON spécifié. Si le fichier n'existe pas, il sera créé.`,
	Run: func(cmd *cobra.Command, args []string) {
		if addFilePath == "" || addName == "" || addURL == "" || addOwner == "" {
			fmt.Println("Erreur: tous les drapeaux (--file, --name, --url, --owner) sont obligatoires.")
			return
		}

		newTarget := config.InputTarget{
			Name:  addName,
			URL:   addURL,
			Owner: addOwner,
		}

		// Tente de charger les cibles existantes. Si le fichier n'existe pas, on commence avec une liste vide.
		existingTargets, err := config.LoadTargetsFromFile(addFilePath)
		if err != nil {
			// Si l'erreur est que le fichier n'existe pas, on ignore et on continue avec une liste vide.
			if errors.Is(err, os.ErrNotExist) {
				existingTargets = []config.InputTarget{}
				fmt.Printf("Le fichier '%s' n'existe pas. Il sera créé.\n", addFilePath)
			} else {
				fmt.Printf("Erreur lors du chargement des URLs existantes: %v\n", err)
				return
			}
		}

		// Vérifier si l'URL existe déjà
		// A faire à part

		existingTargets = append(existingTargets, newTarget)

		err = config.SaveTargetsToFile(addFilePath, existingTargets)
		if err != nil {
			fmt.Printf("Erreur lors de l'enregistrement de la nouvelle URL: %v\n", err)
		} else {
			fmt.Printf("✅ URL '%s' ajoutée avec succès au fichier '%s'.\n", newTarget.URL, addFilePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Définition des drapeaux pour la commande 'add'
	addCmd.Flags().StringVarP(&addFilePath, "file", "f", "", "Chemin vers le fichier JSON où ajouter l'URL")
	addCmd.Flags().StringVarP(&addName, "name", "n", "", "Nom descriptif de l'URL")
	addCmd.Flags().StringVarP(&addURL, "url", "u", "", "L'URL à ajouter")
	addCmd.Flags().StringVarP(&addOwner, "owner", "o", "", "Propriétaire ou contact lié à l'URL")

	// Marquer tous les drapeaux comme obligatoires
	addCmd.MarkFlagRequired("file")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("url")
	addCmd.MarkFlagRequired("owner")
}
