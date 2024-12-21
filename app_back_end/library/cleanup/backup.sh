#!/bin/bash

if command -v timeshift &>/dev/null; then
  sudo timeshift --delete --all
fi

if command -v snapper &>/dev/null; then
  sudo snapper cleanup --dry-run
  sudo snapper cleanup --cleanup-algorithm timeline
fi

if command -v rsnapshot &>/dev/null; then
  sudo rsnapshot -c /etc/rsnapshot.conf cleanup
fi

exit 0
