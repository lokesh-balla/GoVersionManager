#!/bin/sh

ENV_FILE="$HOME/.profile"

# Checking the shell used
if [ -n "`$SHELL -c 'echo $ZSH_VERSION'`" ]; then
   ENV_FILE="$HOME/.zprofile"
elif [ -n "`$SHELL -c 'echo $BASH_VERSION'`" ]; then
   ENV_FILE="$HOME/.profile"
else
  echo "😢 Installation Script only supports Bash and ZSH"
fi

# Creating gvm directory if it does not exist
mkdir -p $HOME/.gvm/bin

# Downloading the binary

# Setting the PATH
echo "# Execute the below commands to set the PATH\n"
echo "$ echo 'export PATH=\"\$HOME/.gvm/bin:\$HOME/.gvm/go/bin:\$PATH\"' >> $ENV_FILE"
echo "$ source $ENV_FILE"
echo "\n Installation DONE"
