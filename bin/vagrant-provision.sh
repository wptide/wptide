#!/usr/bin/env bash

set -e

# Map binary name to source URL.
REPOS=(
	"docker-compose,https://github.com/docker/compose/releases/download/1.21.2/docker-compose-`uname -s`-`uname -m`"
	"composer,https://getcomposer.org/download/1.6.5/composer.phar"
)

# Add Docker repo.
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# Add Glide for Golang
sudo add-apt-repository ppa:masterminds/glide -y

sudo apt-get update
sudo apt-get install -y \
	docker-ce avahi-daemon git php-cli php-xml zip golang-go glide

# Install various binary scripts.
for REPOSET in "${REPOS[@]}"; do
	while IFS=',' read -ra REPO; do
		BINPATH="/usr/local/bin/${REPO[0]}"

		if [ ! -f "$BINPATH" ]; then
			echo "Installing ${REPO[1]} to $BINPATH"

			sudo curl -sL "${REPO[1]}" -o "$BINPATH"
			sudo chmod +x "$BINPATH"
		fi
	done <<< "$REPOSET"
done

# Add docker user group to avoid using sudo.
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker

# Ensure we have the environment defined.
if [ ! -f /vagrant/.env ]; then
	cp /vagrant/.env.dist /vagrant/.env
fi

# Move to the root of the project.
cd /vagrant

make api.tpl
make api.composer
make api.up
make api.setup