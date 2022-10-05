#!/bin/sh

ENV_FILE="${HOME}/.profile"

# Checking the shell used
if [ -n "`${SHELL} -c 'echo $ZSH_VERSION'`" ]; then
   ENV_FILE="${HOME}/.zprofile"
elif [ -n "`${SHELL} -c 'echo $BASH_VERSION'`" ]; then
   ENV_FILE="${HOME}/.profile"
else
  echo "😢 Installation Script only supports Bash and ZSH"
fi

# Getting Release Tag
TAG=$(curl -LIs -o /dev/null -w %{url_effective} https://github.com/lokesh-balla/GoVersionManager/releases/latest | rev | cut -d / -f 1 | rev)

# Detect OS
case "$(command uname)" in
   'Darwin') OS="darwin";;
   'Linux') OS="linux";;
   'FreeBSD') OS='freebsd';;
   *) echo "gvm binaries are not available for $(command uname)" && exit 1
esac

# Detect Arch
case "$(command uname -m)" in
   'arm64') ARCH="arm64";;
   'amd64') ARCH="amd64";;
   *) echo "gvm binaries are not available for $(command uname -m)" && exit 1
esac

BINARY=$(echo "gvm_${OS}_${ARCH}")
URL=$(echo "https://github.com/Lokesh-Balla/GoVersionManager/releases/download/${TAG}/${BINARY}")

echo "Installing ${BINARY} version ${TAG} from ${URL}"

# Downlading the binary
curl --output /tmp/${BINARY} -Ls ${URL}

# Moving the binary to location
mkdir -p ${HOME}/.gvm/bin
mv /tmp/${BINARY} ${HOME}/.gvm/bin/gvm
chmod +x ${HOME}/.gvm/bin/gvm

# Setting the PATH
echo "# Execute the below commands to set the PATH\n"
echo "$ echo 'export PATH=\"\$HOME/.gvm/bin:\$HOME/.gvm/go/bin:\$PATH\"' >> ${ENV_FILE}"
echo "$ source ${ENV_FILE}"
echo "\n Installation Done"
