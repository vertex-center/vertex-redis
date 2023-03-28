if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "Installing redis..."
    brew install redis
    echo "$(redis-server -v)" installed.
else
    wget https://download.redis.io/redis-stable.tar.gz
    tar -xzvf redis-stable.tar.gz
    cd redis-stable || exit 1
    make
    make install
    cd .. || exit 1
    rm -rf redis-stable
    rm redis-stable.tar.gz
fi
