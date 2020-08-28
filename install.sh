#Install git
sudo apt install git
#Install go 1.15
export GOPKG=go1.15
wget https://golang.org/dl/$GOPKG.linux-armv6l.tar.gz
sudo tar -C /usr/local -xzf $GOPKG.linux-armv6l.tar.gz
rm $GOPKG.linux-armv6l.tar.gz
echo PATH=$PATH:/usr/local/go/bin && export PATH
echo GOPATH=$HOME && export GOPATH
#Make Folder and Clone Repository
mkdir ventiPi && cd ventiPi
git clone https://github.com/de-wax/ventiPi.git
