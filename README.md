# go-poc
Small PoC for testing web interfaces

docker rmi -f $(docker images -a -q)

docker rmi $(docker images -f "dangling=true" -q)

sudo lsof -nP -i4TCP:$PORT | grep LISTEN
