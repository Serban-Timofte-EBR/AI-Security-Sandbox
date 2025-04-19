import sys
import numpy as np
from PIL import Image
import onnxruntime as ort
import json
import argparse
import os

def preprocess_image(image_path):
    img = Image.open(image_path).convert("RGB")
    img = img.resize((224, 224))
    img_data = np.asarray(img).astype(np.float32) / 255.0
    img_data = np.transpose(img_data, (2, 0, 1)) 
    img_data = np.expand_dims(img_data, axis=0)
    return img_data

def apply_perturbation(data, intensity=0.1):
    noise = np.random.normal(loc=0.0, scale=intensity, size=data.shape).astype(np.float32)
    perturbed = np.clip(data + noise, 0.0, 1.0)
    return perturbed

def save_perturbed_image(data, path="assets/cat_perturbed.jpg"):
    data = np.squeeze(data, axis=0)
    data = np.transpose(data, (1, 2, 0)) * 255.0 
    img = Image.fromarray(np.uint8(data))
    os.makedirs(os.path.dirname(path), exist_ok=True)
    img.save(path)

def load_labels():
    return {281: "tabby cat", 282: "tiger cat", 283: "Persian cat", 284: "Siamese cat", 285: "Egyptian cat"}


def run_inference(image_data, model_path):
    session = ort.InferenceSession(model_path, providers=["CPUExecutionProvider"])
    input_name = session.get_inputs()[0].name
    outputs = session.run(None, {input_name: image_data})
    scores = outputs[0][0]
    top_idx = int(np.argmax(scores))
    confidence = float(scores[top_idx])
    label = load_labels().get(top_idx, f"class_{top_idx}")
    return top_idx, label, confidence

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("image", help="Path to image")
    parser.add_argument("model", help="Path to model")
    parser.add_argument("--perturb", type=float, help="Apply perturbation with given intensity")
    args = parser.parse_args()

    data = preprocess_image(args.image)
    result_original = run_inference(data, args.model)

    if args.perturb is not None:
        data = apply_perturbation(data, args.perturb)
        save_perturbed_image(data)
        result_perturbed = run_inference(data, args.model)
        out = {
            "before_attack": {
                "label": result_original[1],
                "confidence": round(result_original[2], 4)
            },
            "after_attack": {
                "label": result_perturbed[1],
                "confidence": round(result_perturbed[2], 4)
            },
            "intensity": args.perturb
        }
    else:
        out = {
            "prediction": {
                "label": result_original[1],
                "confidence": round(result_original[2], 4)
            }
        }

    print(json.dumps(out, indent=2))

if __name__ == "__main__":
    main()