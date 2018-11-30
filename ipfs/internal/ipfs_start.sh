#!/bin/sh

# Modified IPFS node initialization script.
# Mount to /usr/local/bin/start_ipfs
# Source: https://github.com/ipfs/go-ipfs/blob/master/bin/container_daemon

set -e
user=ipfs
repo="$IPFS_PATH"

if [ "$(id -u)" -eq 0 ]; then
  echo "Changing user to $user"
  # ensure folder is writable
  su-exec "$user" test -w "$repo" || chown -R -- "$user" "$repo"
  # restart script with new privileges
  exec su-exec "$user" "$0" "$@"
fi

# 2nd invocation with regular user
ipfs version

if [ -e "$repo/config" ]; then
  echo "Found IPFS fs-repo at $repo"
else
  case "$IPFS_PROFILE" in
    "") INIT_ARGS="" ;;
    *) INIT_ARGS="--profile=$IPFS_PROFILE" ;;
  esac
  ipfs init "$INIT_ARGS"
  ipfs config Addresses.API /ip4/0.0.0.0/tcp/5001
  ipfs config Addresses.Gateway /ip4/0.0.0.0/tcp/8080
fi

# set datastore quota
ipfs config Datastore.StorageMax %d

# if the first argument is daemon
if [ "$1" = "daemon" ]; then
  # filter the first argument until
  # https://github.com/ipfs/go-ipfs/pull/3573
  # has been resolved
  shift
else
  # print deprecation warning
  # go-ipfs used to hardcode "ipfs daemon" in it's entrypoint
  # this workaround supports the new syntax so people start setting daemon explicitly
  # when overwriting CMD
  echo "DEPRECATED: arguments have been set but the first argument isn't 'daemon'" >&2
fi

exec ipfs daemon "$@"
