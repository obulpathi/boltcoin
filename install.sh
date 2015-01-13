sudo apt-get update && sudo apt-get upgrade

sudo apt-get install libgmp3-dev libreadline6-dev
sudo apt-get install mercurial
sudo apt-get install golang

sudo apt-get install build-essential g++-4.8 git cmake libboost-all-dev
sudo apt-get install automake unzip libgmp-dev libtool libleveldb-dev yasm libminiupnpc-dev libreadline-dev scons
sudo apt-get install libncurses5-dev libcurl4-openssl-dev wget
sudo apt-get install qtbase5-dev qt5-default qtdeclarative5-dev libqt5webkit5-dev
sudo apt-get install libjsoncpp-dev libargtable2-dev

sudo add-apt-repository ppa:ubuntu-sdk-team/ppa
sudo apt-get update
sudo apt-get install ubuntu-sdk qtbase5-private-dev qtdeclarative5-private-dev libqt5opengl5-dev
# sudo apt-get install aptitude; sudo aptitude install ubuntu-sdk

go get -u github.com/ethereum/go-ethereum/cmd/ethereum
go get -u -a github.com/ethereum/go-ethereum/cmd/mist

# Reference: https://github.com/ethereum/go-ethereum/wiki/Instructions-for-getting-the-Go-implementation-of-Ethereum-and-the-Mist-browser-installed-on-Ubuntu-14.04-%28trusty%29
