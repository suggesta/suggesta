SHOSA_PACKAGE_ROOT_PATH="$GOPATH/src/github.com/go-shosa/shosa"
COV_PARTIAL_FILE=profile.cov.out
COV_FILE=profile-all.cov.out
COV_MODE=count
HEADER="mode: $COV_MODE"

cd $SHOSA_PACKAGE_ROOT_PATH; \
echo "mode: count" > $COV_FILE
for dir in $(find . -maxdepth 10 -not -path './.git*' \
    -not -path '*/_*' -not -path './cmd' -not -path './release*' \
    -not -path './vendor*' -type d)
do
if ls $dir/*.go &> /dev/null; then
  go test -v -covermode=count -coverprofile=$dir/$COV_PARTIAL_FILE $dir
  if [ -f $dir/$COV_PARTIAL_FILE ]; then
    cat $dir/$COV_PARTIAL_FILE | tail -n +2 >> $COV_FILE
    rm $dir/$COV_PARTIAL_FILE
  fi
fi
done
