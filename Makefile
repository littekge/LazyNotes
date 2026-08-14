# variables
APP := lazynotes
PKG := ./cmd/lazynotes
BIN := bin
DIST := dist

# creates bin dirs
$(BIN):
	mkdir -p $(BIN)
$(DIST):
	mkdir -p $(DIST)

# ---------- DEV ----------
# build native binary
build: | $(BIN)
	go build -o $(BIN)/$(APP) $(PKG)

# run without generating binary
run: 
	go run $(PKG)/main.go
# ---------- END DEV ----------

# ---------- CHORES ----------
# remove all buidl artifacts
clean:
	rm -rf $(DIST)
	rm -rf $(BIN)

# tidy and vendor deps
vendor:
	go mod tidy
	go mod vendor
# ---------- END CHORES ----------

# ---------- RELEASE ----------
# build the full app release in /dist/
release: build-linux-amd64 build-linux-arm64 build-linux-armv7 build-linux-armv6 build-windows-amd64 build-darwin-amd64 build-darwin-arm64
	cd $(DIST) && sha256sum $(APP)_* > sha256sums.txt

# build-linux-amd64: 64-bit x86 Linux (primary desktop/server)
build-linux-amd64: | $(DIST)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(DIST)/$(APP)_linux_amd64 $(PKG)

# build-linux-arm64: 64-bit Raspberry Pi OS (Pi 3/4/5, Zero 2 W on 64-bit)
build-linux-arm64: | $(DIST)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(DIST)/$(APP)_linux_arm64 $(PKG)

# build-linux-armv7: 32-bit Raspberry Pi OS (Pi 2/3/4, Zero 2 W)
build-linux-armv7: | $(DIST)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o $(DIST)/$(APP)_linux_armv7 $(PKG)

# build-linux-armv6: Pi 1 / Pi Zero / Zero W
build-linux-armv6: | $(DIST)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -o $(DIST)/$(APP)_linux_armv6 $(PKG)

# build-windows-amd64: 64-bit Windows (secondary target)
build-windows-amd64: | $(DIST)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(DIST)/$(APP)_windows_amd64.exe $(PKG)

# build-darwin-amd64: Intel macOS
build-darwin-amd64: | $(DIST)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(DIST)/$(APP)_darwin_amd64 $(PKG)

# build-darwin-arm64: Apple Silicon macOS
build-darwin-arm64: | $(DIST)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(DIST)/$(APP)_darwin_arm64 $(PKG)
# ---------- END RELEASE ----------

