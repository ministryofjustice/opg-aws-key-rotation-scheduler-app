SHELL := '/bin/bash'
APPNAME := OPG AWS Key Rotation
PLIST := ./'${APPNAME}.app'/Contents/Info.plist
PLISTTEMP := ./plist.tmp
BUILD_FOLDER = ./builds/

.PHONY: all requirements build app plist-fix

all:
	@${MAKE} requirements
	@${MAKE} build

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

build:
	@mkdir -p ${BUILD_FOLDER}
	go build -o ${BUILD_FOLDER}main main.go		
	cd ${BUILD_FOLDER} && fyne package --executable ./main --name "${APPNAME}" --icon "../static/icons/main.png"
	@${MAKE} plist-fix

plist-fix:
	@cd ${BUILD_FOLDER} && cat ${PLIST} | sed -e 's#</dict>#\t<key>LSUIElement</key>\n\t<true/>\n</dict>#' > ${PLISTTEMP}
	@cd ${BUILD_FOLDER} && mv ${PLISTTEMP} ${PLIST}

