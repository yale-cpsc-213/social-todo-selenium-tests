if [ -z "$1" ]
  then
    echo "No tag supplied"
    exit 1
fi

GOPATH=$HOME/go
NAME="social-todo-selenium-tests"
VERSION="$1"
export CGO_ENABLED=0
GOOS=windows GOARCH=386 go build -o "$NAME-windows-$VERSION".exe
GOOS=linux GOARCH=386 go build -o "$NAME-linux-$VERSION"
go build -o "$NAME-mac-$VERSION"
exit 0
