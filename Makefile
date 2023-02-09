SHELL := $(shell which bash)
APPNAME := OPGAWSKeyRotation

OS := $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

BUILD_FOLDER = ./builds/
OS_AND_ARCHS_TO_BUILD := darwin_arm64 darwin_amd64
HOST_ARCH := ${OS}_${ARCH}

PLIST := ./'${APPNAME}.app'/Contents/Info.plist
PLIST_TEMP := ./plist.tmp

.DEFAULT_GOAL: self
.PHONY: self all requirements darwin_arm64 darwin_amd64
.ONESHELL:
.EXPORT_ALL_VARIABLES:


self: $(HOST_ARCH)
	
all: $(OS_AND_ARCHS_TO_BUILD)


darwin_arm64: requirements
	@mkdir -p $(BUILD_FOLDER)$@/
	env GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o $(BUILD_FOLDER)$@/main main.go
	cd $(BUILD_FOLDER)$@/ && fyne package --executable ./main --name "${APPNAME}" --icon "../../icons/main.png"
	@cat $(BUILD_FOLDER)$@/${PLIST} | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > $(BUILD_FOLDER)$@/${PLIST_TEMP}
	@mv $(BUILD_FOLDER)$@/${PLIST_TEMP} $(BUILD_FOLDER)$@/${PLIST}
	@echo Build $@ complete.

darwin_amd64: requirements
	@mkdir -p $(BUILD_FOLDER)$@/
	env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o $(BUILD_FOLDER)$@/main main.go
	cd $(BUILD_FOLDER)$@/ && fyne package --executable ./main --name "${APPNAME}" --icon "../../icons/main.png"
	@cat $(BUILD_FOLDER)$@/${PLIST} | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > $(BUILD_FOLDER)$@/${PLIST_TEMP}
	@mv $(BUILD_FOLDER)$@/${PLIST_TEMP} $(BUILD_FOLDER)$@/${PLIST}
	@echo Build $@ complete.
	
requirements:
ifeq (, $(shell which go))
	$(error go command not found)
endif
ifndef GOBIN
	$(error GOBIN is not defined)
endif
ifeq (, $(shell which fyne))
	$(error fyne command not found, check https://developer.fyne.io/started/packaging)	
endif
	@echo All requirements checked
	@rm -Rf ${BUILD_FOLDER}
