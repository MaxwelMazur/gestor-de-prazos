go build main.go

cd wasm/

GOOS=js GOARCH=wasm go build -o  ../static/json.wasm

cd ..

./main