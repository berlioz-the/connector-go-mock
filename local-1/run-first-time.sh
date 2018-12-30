MY_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/$(basename "${BASH_SOURCE[0]}")"
MY_DIR="$(dirname $MY_PATH)"
echo "My Dir: $MY_DIR"
cd $MY_DIR

dockerPrefix=""
osNameStr=$(uname -r)
echo "$osNameStr"
if [[ "$osNameStr" = *'Microsoft'* ]]; then
   echo "INSIDE BASH ON WINDOWS"
   dockerPrefix=" -H tcp://0.0.0.0:2375"
fi

rsync -r ../../connector-go.git vendor/

berlioz local build-run --quick --pathoverride .,../support/simple
