# .PHONY=build_bin
build_bin:	
	# For Windows 64-bit
	GOOS=windows GOARCH=amd64 go build -o mycli-windows.exe ./path/to/package

	# For macOS 64-bit
	GOOS=darwin GOARCH=amd64 go build -o mycli-macos ./path/to/package

	# For Linux 64-bit
	GOOS=linux GOARCH=amd64 go build -o mycli-linux ./path/to/package
