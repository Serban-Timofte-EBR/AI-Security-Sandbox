package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var injectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Run multiple prompt injection tests against an LLM",
	Long: `This command reads a list of crafted prompts and tests them against a local LLM. 
It generates a detailed report with the model's responses and whether the injection succeeded.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[INFO] Running prompt injection batch...")

		pythonCmd := exec.Command("python3", "python/run_injections.py", modelPath)
		output, err := pythonCmd.CombinedOutput()
		if err != nil {
			fmt.Println("[ERROR] Python injection script failed:", err)
			fmt.Println(string(output))
			return
		}

		fmt.Println("[SUCCESS] Injection test complete:")
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(injectCmd)
	injectCmd.Flags().StringVarP(&modelPath, "model", "m", "", "Path to LLM model (optional if using default)")
	injectCmd.MarkFlagRequired("model")
}
