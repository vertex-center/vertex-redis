if ! command -v redis-server &> /dev/null; then
    echo "redis-server is not installed."
    exit 1
fi
