SHELL := $(shell which bash)
APPNAME := OPG AWS Key Rotation

OS := $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

BUILD_FOLDER = ./fyne-cross/

# PLIST := ./'${APPNAME}.app'/Contents/Info.plist
# PLIST_TEMP := ./plist.tmp

.DEFAULT_GOAL: self
.PHONY: self all requirements build plist-fix
.ONESHELL:
.EXPORT_ALL_VARIABLES:


self: requirements
	fyne-cross $(OS) -arch=$(ARCH) --name "${APPNAME}" --icon="./icons/main.png" 
	@cat $(BUILD_FOLDER)/dist/$(OS)-$(ARCH)/'$(APPNAME).app'/Contents/Info.plist | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > $(OS)-$(ARCH).plist
	@mv $(OS)-$(ARCH).plist $(BUILD_FOLDER)/dist/$(OS)-$(ARCH)/'$(APPNAME).app'/Contents/Info.plist

all: requirements 
	fyne-cross darwin -arch=amd64,arm64 --name "${APPNAME}" --icon="./icons/main.png" 
	@cat $(BUILD_FOLDER)/dist/darwin-arm64/'$(APPNAME).app'/Contents/Info.plist | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > darwin-arm64.plist
	@mv darwin-arm64.plist $(BUILD_FOLDER)/dist/darwin-arm64/'$(APPNAME).app'/Contents/Info.plist
	@cat $(BUILD_FOLDER)/dist/darwin-amd64/'$(APPNAME).app'/Contents/Info.plist | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > darwin-amd64.plist
	@mv darwin-amd64.plist $(BUILD_FOLDER)/dist/darwin-amd64/'$(APPNAME).app'/Contents/Info.plist
	
	
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
ifeq (, $(shell which fyne-cross))
	$(error fyne-cross command not found, check https://developer.fyne.io/started/cross-compiling)	
endif
	@rm -Rf ${BUILD_FOLDER}
	@echo All requirements checked
	
