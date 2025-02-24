#!/bin/bash

echo "🐍🐍 Installing DURL! 🐍🐍"

echo "Thank you for installing durl!"

# Define the source binary and the target directory
BINARY_PATH="./build/durl"
TARGET_DIR="/usr/local/bin"

# Check if the binary exists
if [ ! -f "$BINARY_PATH" ]; then
    echo "Error: $BINARY_PATH does not exist."
    exit 1
fi

# Move the binary to the target directory
sudo cp "$BINARY_PATH" "$TARGET_DIR"

# Set the permissions to make it executable
sudo chmod +x "$TARGET_DIR/durl"

echo "durl has been installed successfully and is now globally executable."
