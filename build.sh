# Windows

echo "Building for Windows..."
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./dist/Codemao-Scrapy-win64.exe

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/Codemao-Scrapy-win32.exe
# Mac OS

echo "Building for Mac OS..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./dist/Codemao-Scrapy-mac

# Linux

echo "Building for Linux..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/Codemao-Scrapy-linux

echo "Done!"
