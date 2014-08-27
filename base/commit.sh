sudo docker build -t="atomos/default:latest" .
sudo docker push atomos/default

#for i in {1..10}; do sudo docker run -t -i atomos/default /usr/local/bin/jeeves -token="test"; done
