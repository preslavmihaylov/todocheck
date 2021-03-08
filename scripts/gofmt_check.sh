GOFMT_TARGETS=$(gofmt -s -l `find . -name '*.go'`)

if [ -n "$GOFMT_TARGETS" ]; then
    echo "files listed below have not been properly formatted"
    echo $GOFMT_TARGETS
    echo ""
    echo "please make sure you code has been formatted using gofmt"
    exit 1
fi

echo "::set-output name=gofmt-test:: step succeed"