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

berlioz local push-run --quick --cluster berliozgo --service example --pathoverride .,../support/secret
echo '==============================================================================='
echo '==============================================================================='
echo '==============================================================================='
read -p "Pausing to fetch logs..." -t 2
echo ""

docker $dockerPrefix ps -a | grep berliozgo-example | head -n 1 | awk '{print $1}' | xargs docker $dockerPrefix logs
