package models

import (
	"fmt"
	"github.com/microsoft/onnxruntime-go"
	"github.com/microsoft/onnxruntime-go/onnxruntime"
)

func RunResNetInference(modelPath string, inputData []float32) error {
	session, err := onnxruntime.NewSession(modelPath)
	if err != nil {
		return fmt.Errorf("error loading model: %w", err)
	}
	defer session.Close()

	// Create input tensor
	tensor, err := onnxruntime.NewTensor(onnxruntime.TENSOR_FLOAT, []int64{1, 3, 224, 224}, inputData)
	if err != nil {
		return fmt.Errorf("error creating tensor: %w", err)
	}

	output, err := session.Run(map[string]*onnxruntime.Value{"data": tensor})
	if err != nil {
		return fmt.Errorf("error running inference: %w", err)
	}

	fmt.Println("Output shape:", output[0].Shape())
	fmt.Println("Output values (first 5):")
	for i := 0; i < 5; i++ {
		fmt.Printf("  %.4f\n", output[0].Float32Values()[i])
	}

	return nil
}
