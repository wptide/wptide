#!/usr/bin/env bash

cd /vagrant

make api.upd

echo "Done! Visit http://$HOSTNAME.local (please wait 10 seconds before it's up)"
