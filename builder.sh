
if [[ -z $1 ]]; then 
echo "Give a name to the executable ./builder.sh name"
exit 1
fi

rm $1

wget https://github.com/upx/upx/releases/download/v5.0.0/upx-5.0.0-amd64_linux.tar.xz -O upx.tar.gz
tar -xvf upx.tar.gz
GOOS=linux GOARCH=amd64 go build -o $1 -ldflags="-s -w" cmd/main.go 
./upx-5.0.0-amd64_linux/upx --brute $1
rm upx.tar.gz 
rm -rf /upx-5.0.0-amd64_linux