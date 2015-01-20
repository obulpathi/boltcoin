sudo apt-get update && sudo apt-get -yy upgrade

sudo apt-get install -y libgmp3-dev libreadline6-dev
sudo apt-get install -y mercurial golang

sudo apt-get install -y build-essential g++-4.8 git cmake libboost-all-dev
sudo apt-get install -y automake unzip libgmp-dev libtool libleveldb-dev yasm libminiupnpc-dev libreadline-dev scons
sudo apt-get install -y libncurses5-dev libcurl4-openssl-dev wget
sudo apt-get install -y libjsoncpp-dev libargtable2-dev

go get -u github.com/obulpathi/boltcoin/cmd/ethereum

# Reference: https://github.com/obulpathi/boltcoin/wiki/Instructions-for-getting-the-Go-implementation-of-Ethereum-and-the-Mist-browser-installed-on-Ubuntu-14.04-%28trusty%29
