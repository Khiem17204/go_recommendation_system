#!/bin/bash

set -e  # Exit on error

CONFIG_DIR="./config"

echo "Applying all YAML files in $CONFIG_DIR..."

# Apply all .yaml and .yml files
for file in "$CONFIG_DIR"/*.yaml "$CONFIG_DIR"/*.yml; do
  if [ -f "$file" ]; then
    echo "Applying: $file"
    kubectl apply -f "$file"
  fi
done

echo "✅ All configs applied."
