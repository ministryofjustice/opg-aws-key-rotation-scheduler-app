SHELL := $(shell which bash)
APPNAME := OPGAWSKeyRotation

OS := $(shell uname | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

BUILD_FOLDER = ./builds/
OS_AND_ARCHS_TO_BUILD := darwin_arm64 darwin_amd64
HOST_ARCH := ${OS}_${ARCH}
ICON_FILE := ../../pkg/icons/files/main.png

PLIST := ./'${APPNAME}.app'/Contents/Info.plist
PLIST_TEMP := ./plist.tmp
PWD := $(shell pwd)
USER_PROFILE := ~/.zprofile

.DEFAULT_GOAL: self
.PHONY: self all requirements darwin_arm64 darwin_amd64
.ONESHELL: self all requirements darwin_arm64 darwin_amd64
.EXPORT_ALL_VARIABLES:


self: $(HOST_ARCH)
	
all: $(OS_AND_ARCHS_TO_BUILD)

# this is to handle the intel macs that identify as x86_64
# from uname, but golang treats that as amd64
darwin_x86_64: requirements
	@${MAKE} darwin_amd64

darwin_arm64: requirements
	@mkdir -p $(BUILD_FOLDER)$@/
	env GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -o $(BUILD_FOLDER)$@/main main.go
	cd $(BUILD_FOLDER)$@/ && fyne package --executable ./main --name "${APPNAME}" --icon "${ICON_FILE}"
	@cat $(BUILD_FOLDER)$@/${PLIST} | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > $(BUILD_FOLDER)$@/${PLIST_TEMP}
	@mv $(BUILD_FOLDER)$@/${PLIST_TEMP} $(BUILD_FOLDER)$@/${PLIST}
	@echo Build $@ complete.

darwin_amd64: requirements
	@mkdir -p $(BUILD_FOLDER)$@/
	env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o $(BUILD_FOLDER)$@/main main.go
	cd $(BUILD_FOLDER)$@/ && fyne package --executable ./main --name "${APPNAME}" --icon "${ICON_FILE}"
	@cat $(BUILD_FOLDER)$@/${PLIST} | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > $(BUILD_FOLDER)$@/${PLIST_TEMP}
	@mv $(BUILD_FOLDER)$@/${PLIST_TEMP} $(BUILD_FOLDER)$@/${PLIST}
	@echo Build $@ complete.
	
requirements:
ifeq (, $(shell which go))
	$(error go command not found)
endif
ifndef GOBIN
	$(warning GOBIN is not defined, configuring as ${HOME}/go/bin)
	$(shell mkdir -p ${HOME}/go/bin )
	$(shell echo "" >> ${USER_PROFILE};)
	$(shell echo "# ADDED BY ${PWD}/Makefile" >> ${USER_PROFILE};)
	$(shell echo export GOBIN="\$${HOME}/go/bin" >> ${USER_PROFILE};)
	$(shell echo export PATH="\$${PATH}:\$${GOBIN}" >> ${USER_PROFILE})
endif
ifeq (, $(shell which fyne))
	$(warning fyne command not found, installing)	
	$(shell source ${USER_PROFILE} && go install fyne.io/fyne/v2/cmd/fyne@v2.3.0)
endif
	@echo All requirements checked
	@rm -Rf ${BUILD_FOLDER}
	@test -f ${USER_PROFILE} && source ${USER_PROFILE} || echo ${USER_PROFILE} not found
	
