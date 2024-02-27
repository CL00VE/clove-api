while getopts ":m:p:" opt; do
case $opt in
m)
    mode="$OPTARG"
    ;;
p)
    port="$OPTARG"
    ;;
\?)
    echo "Invalid option: -$OPTARG" >&2
    exit 1
    ;;
:)
    echo "Option -$OPTARG requires an argument." >&2
    exit 1
    ;;
esac
done

cd ./application/cmd

case $mode in
"reflex")
if [ -n "$port" ]; then
    $GOPATH/bin/reflex -r '\.go$' -s go run cloveApi.go -- -p "$port"
else
    $GOPATH/bin/reflex -r '\.go$' -s go run cloveApi.go
fi
;;
*)
if [ -n "$port" ]; then
    go run cloveApi.go -p "$port"
else
    go run cloveApi.go
fi
;;
esac
