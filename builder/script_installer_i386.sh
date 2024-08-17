#!/bin/bash

# Define variables
DOWNLOAD_LINK="https://github.com/AlexDeFoc/goRunner/releases/latest/download/linux_i386.tar.xz"
TARBALL="linux_i386.tar.xz"
SCRIPT_DIR=$(pwd)

# 1. Download the file
echo "Downloading file..."
curl -L "$DOWNLOAD_LINK" -o "$TARBALL"

# 2. Uncompress the tarball
echo "Extracting $TARBALL..."
tar -xvJf "$TARBALL"

# 3. Delete the tarball
echo "Removing $TARBALL..."
rm "$TARBALL"

# 4. Add the current path to /etc/paths
echo "Updating /etc/paths..."
if ! grep -q "^$SCRIPT_DIR$" /etc/paths; then
  echo "$SCRIPT_DIR" | sudo tee -a /etc/paths > /dev/null
  if [ $? -ne 0 ]; then
    echo "Error updating /etc/paths."
    exit 1
  fi
else
  echo "Directory already in /etc/paths."
fi

# 5. Print Done
echo "Done"
