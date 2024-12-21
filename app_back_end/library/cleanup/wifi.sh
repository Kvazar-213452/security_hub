#!/bin/bash

if [ "$(id -u)" -ne 0 ]; then
    echo "Цей скрипт потрібно запускати з правами суперкористувача (root)."
    exit 1
fi

for profile in $(ls /etc/NetworkManager/system-connections/); do
    sudo rm -f "/etc/NetworkManager/system-connections/$profile"
done

sudo systemctl restart NetworkManager
