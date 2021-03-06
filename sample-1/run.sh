MY_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/$(basename "${BASH_SOURCE[0]}")"
MY_DIR="$(dirname $MY_PATH)"
echo "My Dir: $MY_DIR"
cd $MY_DIR

berlioz local build-run --quick --cluster berliozgo --service example --pathoverride .,../support/simple
echo '==============================================================================='
echo '==============================================================================='
echo '==============================================================================='
read -p "Pausing to fetch logs..." -t 2
echo ""
docker ps -a | grep berliozgo-example | head -n 1 | awk '{print $1}' | xargs docker logs
