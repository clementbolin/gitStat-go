if [[ $EUID -ne 0 ]]; then
    echo "You must run this with superuser priviliges.  Try \"sudo ./install\"" 2>&1
    exit 1
else
    echo "Installing gitStat-go..."
fi
make build
sudo cp ./bin/gitStat-go /usr/local/bin
echo "gitStat-go has been successfully installed"