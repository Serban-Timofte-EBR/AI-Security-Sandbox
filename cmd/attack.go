package cmd

import (
	"fmt"

	"github.com/Serban-Timofte-EBR/AI-Security-Sandbox/internal/attacks"
	"github.com/Serban-Timofte-EBR/AI-Security-Sandbox/internal/models"
	"github.com/spf13/cobra"
)

var (
	modelPath string
	intensity float64
)

// attack command
var attackCmd = &cobra.Command{
	Use:   "attack",
	Short: "Run an adversarial attack against a local AI model",
	Long:  `This command runs an actual adversarial attack on a local AI model by applying controlled input perturbations and measuring the impact.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[INFO] Starting adversarial attack...")
		imageData, err := models.LoadAndPreprocessImage("assets/cat.jpg")
		if err != nil {
			fmt.Println("[ERROR] Failed to load image:", err)
			return
		}

		err = models.RunResNetInference(modelPath, imageData)
		if err != nil {
			fmt.Println("[ERROR] Inference failed:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(attackCmd)

	attackCmd.Flags().StringVarP(&modelPath, "model", "m", "", "Path to ONNX model file (required)")

	attackCmd.Flags().Float64VarP(&intensity, "intensity", "i", 0.5, "Perturbation intensity between 0.0 and 1.0")

	attackCmd.MarkFlagRequired("model")
}
