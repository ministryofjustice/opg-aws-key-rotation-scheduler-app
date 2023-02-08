SHELL := $(shell which bash)
APPNAME := OPG AWS Key Rotation

OS := $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH := $(shell arch)

BUILD_FOLDER = ./builds/
OS_AND_ARCHS_TO_BUILD := darwin_arm64 darwin_amd64
HOST_ARCH := ${OS}_${ARCH}

PLIST := ./'${APPNAME}.app'/Contents/Info.plist
PLIST_TEMP := ./plist.tmp

.DEFAULT_GOAL: self
.PHONY: self all requirements build plist-fix
.ONESHELL:
.EXPORT_ALL_VARIABLES:


self: $(HOST_ARCH)
	
all: requirements $(OS_AND_ARCHS_TO_BUILD)
	

darwin_arm64:
	@env GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o $(BUILD_FOLDER)$@/main main.go	
	@cd $(BUILD_FOLDER)$@/ && fyne package --executable ./main --name "${APPNAME}" --icon "../../icons/main.png"
	@cd ${BUILD_FOLDER}$@/ && cat ${PLIST} | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > ${PLIST_TEMP}
	@echo Build $@ complete.

darwin_amd64: 
	@env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o $(BUILD_FOLDER)$@/main main.go 2>/dev/null
	@cd $(BUILD_FOLDER)$@/ && fyne package --executable ./main --name "${APPNAME}" --icon "../../icons/main.png"
	@cd ${BUILD_FOLDER}$@/ && cat ${PLIST} | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > ${PLIST_TEMP}
	@echo Build $@ complete.
	
requirements:
	rm -Rf ${BUILD_FOLDER}
# ifeq (, $(shell which go))
# 	$(error go command not found)
# endif
# ifndef GOBIN
# 	$(error GOBIN is not defined)
# endif
# ifeq (, $(shell which fyne))
# 	$(error fyne command not found, check https://developer.fyne.io/started/packaging)	
# endif
# 	@echo All requirements checked
# 	@rm -Rf ${BUILD_FOLDER}
