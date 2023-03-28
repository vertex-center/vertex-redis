if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "Installing redis..."
    brew install redis
    echo "$(redis-server -v)" installed.
else
    >&2 echo "$OSTYPE" is not supported to install redis automatically.
    exit 1
fi
