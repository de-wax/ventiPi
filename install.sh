#Install git
sudo apt install git
#Install most actual go Version
export GOPKG="$(curl -s https://api.github.com/repos/golang/go/git/matching-refs/tags/go | grep ref | grep -v url | grep -v beta | tail -1 | awk -F\/ {' print $3 '} | sed 's/",//')"
wget https://golang.org/dl/$GOPKG.linux-armv6l.tar.gz
sudo tar -C /usr/local -xzf $GOPKG.linux-armv6l.tar.gz
rm $GOPKG.linux-armv6l.tar.gz
echo PATH=$PATH:/usr/local/go/bin >> $HOME/.zshrc
echo GOPATH=$HOME >> $HOME/.zshrc
source $HOME/.zshrc
#Make Folder and Clone Repository
mkdir ventiPi && cd ventiPi
git clone https://github.com/de-wax/ventiPi.git

