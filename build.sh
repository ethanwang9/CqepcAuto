# clear cache
rm -rf release
echo "清除缓存成功"

# ==========
# MacOS
# ==========

# mac amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/mac_x64/main ./main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/mac_x64/fix ./fix.go
cp -r view release/mac_x64/view
cp -r static release/mac_x64/static
zip -r "release/mac_x64.zip" release/mac_x64/*
echo "mac x64 编译完成"

# mac arm64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/mac_arm64/main ./main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/mac_arm64/fix ./fix.go
cp -r view release/mac_arm64/view
cp -r static release/mac_arm64/static
zip -r "release/mac_arm64.zip" release/mac_arm64/*
echo "mac arm64 编译完成"

# ==========
# Windows
# ==========

# windows amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/windows_x64/main.exe ./main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/windows_x64/fix.exe ./fix.go
cp -r view release/windows_x64/view
cp -r static release/windows_x64/static
zip -r "release/windows_x64.zip" release/windows_x64/*
echo "windows x64 编译完成"

# windows 386
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/windows_x86/main.exe ./main.go
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/windows_x86/fix.exe ./fix.go
cp -r view release/windows_x86/view
cp -r static release/windows_x86/static
zip -r "release/windows_x86.zip" release/windows_x86/*
echo "windows x86 编译完成"

# ==========
# linux
# ==========

# linux x86
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/linux_x86/main ./main.go
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/linux_x86/fix ./fix.go
cp -r view release/linux_x86/view
cp -r static release/linux_x86/static
zip -r "release/linux_x86.zip" release/linux_x86/*
echo "linux x86 编译完成"

# linux x64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/linux_x64/main ./main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/linux_x64/fix ./fix.go
cp -r view release/linux_x64/view
cp -r static release/linux_x64/static
zip -r "release/linux_x64.zip" release/linux_x64/*
echo "linux x64 编译完成"

# linux arm
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/linux_arm/main ./main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/linux_arm/fix ./fix.go
cp -r view release/linux_arm/view
cp -r static release/linux_arm/static
zip -r "release/linux_arm.zip" release/linux_arm/*
echo "linux arm 编译完成"

# linux arm64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/linux_arm64/main ./main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags -static -extldflags -static" -o release/linux_arm64/fix ./fix.go
cp -r view release/linux_arm64/view
cp -r static release/linux_arm64/static
zip -r "release/linux_arm64.zip" release/linux_arm64/*
echo "linux arm64 编译完成"