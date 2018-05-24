#!/usr/bin/env bash

function fatal() {
	echo "$1" >&2
	exit 1
}

[[ $EUID -ne 0 ]] && fatal "You must run this script as root. This is needed because a file must be installed in /usr/local/bin"

command -v curl >/dev/null 2>&1 || fatal "'curl' is not installed. Try: sudo apt-get install curl. Or the equivalent for your system"

os=
arch=

machine_type=`uname -m`
if [ ${machine_type} == 'x86_64' ]; then
  arch="amd64"
else
	arch="386"
fi

if [[ "$OSTYPE" == "linux-gnu" ]]; then
	os="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
	os="darwin"
	if [[ "$arch" == "386" ]]; then
		fatal "x86 architecture isn't supported on macOS"
	fi
else
	fatal "Unsupported OS: ${OSTYPE}"
fi

VERSION=$(curl -s https://raw.githubusercontent.com/panta82/proxyhop/master/VERSION)
if [[ -z ${VERSION} ]] ;then
	fatal "Couldn't obtain version"
fi

url="https://github.com/panta82/proxyhop/releases/download/${VERSION}/proxyhop_${VERSION}_${os}_${arch}"
echo "Downloading proxyhop $VERSION from $url"
curl -sf -L $url --output /tmp/proxyhop || fatal "Failed to download proxyhop to /usr/local/bin/proxyhop"

cp -f /usr/local/bin/proxyhop /tmp/proxyhop.prev
mv /tmp/proxyhop /usr/local/bin/proxyhop
chmod +x /usr/local/bin/proxyhop

echo "Proxyhop $VERSION has been installed to /usr/local/bin/proxyhop"
