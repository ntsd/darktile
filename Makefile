

default: build

build:
	./scripts/build.sh
	
install:
	 apt-get install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev

build-mobile:
	gomobile build ./mobile/.

run-mobile:
	gomobile install ./mobile/.
