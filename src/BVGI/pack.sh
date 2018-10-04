# Notice:
# This script does not support Windows. I hate you, Windows!

# Variables
APP_NAME="qna"
PACKAGE="GQQL"
STATIC_DIRS=("data" "ui") # seperate by a white space
OSES=("darwin" "linux") # linux, darwin. Seperate by a white space
ARCHS=("amd64" "386") # amd64, 386, arm


####################################
# Do not change the following lines
####################################

: ${GOPATH?"Need to set GOPATH"} # Check whether GOPATH is set, if not then exist

mkdir -p $GOPATH/dist/$APP_NAME # make temporary directory

echo "attaching static files..."
for STATIC_DIR in "${STATIC_DIRS[@]}"
do
	cp -r $GOPATH/src/$PACKAGE/$STATIC_DIR $GOPATH/dist/$APP_NAME # copy static files
done

for OS in "${OSES[@]}"
do
	for ARCH in "${ARCHS[@]}"
	do
		echo "building for $OS $ARCH..."
		BIN_FILE=$GOPATH/dist/$APP_NAME/$APP_NAME
		GOOS=$OS GOARCH=$ARCH go build -o $BIN_FILE $PACKAGE # build go package

		echo "compressing for $OS $ARCH..."
		DIST_FILE=$GOPATH/dist/$APP_NAME.$OS.$ARCH.tar.gz
		tar czf $DIST_FILE -C $GOPATH/dist/$APP_NAME . # compress directory without obsolute path

		echo "cleaning for $OS $ARCH..."
		rm -rf $BIN_FILE

		echo "done: $DIST_FILE"
	done
done

echo "cleaning for $APP_NAME"
rm -rf $GOPATH/dist/$APP_NAME

echo "done"
ls $GOPATH/dist/