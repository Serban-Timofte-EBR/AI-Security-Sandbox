package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	modelPath string
	intensity float64
)

var attackCmd = &cobra.Command{
	Use:   "attack",
	Short: "Run an adversarial attack against a local AI model",
	Long:  `This command runs an actual adversarial attack on a local AI model by applying controlled input perturbations and measuring the impact.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[INFO] Starting adversarial attack...")

		pythonCmd := exec.Command("python3", "python/infer_resnet50.py", "assets/cat.jpg", modelPath, "--perturb", fmt.Sprintf("%f", intensity))
		output, err := pythonCmd.CombinedOutput()
		if err != nil {
			fmt.Println("[ERROR] Failed to run Python inference:", err)
			fmt.Println(string(output))
			return
		}

		fmt.Println("[SUCCESS] Python inference result:")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(attackCmd)

	attackCmd.Flags().StringVarP(&modelPath, "model", "m", "", "Path to ONNX model file (required)")
	attackCmd.Flags().Float64VarP(&intensity, "intensity", "i", 0.5, "Perturbation intensity between 0.0 and 1.0")
	attackCmd.MarkFlagRequired("model")
}
