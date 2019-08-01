# builds our Redis Docker image
# -----
DATA_DIR=$(pwd)/podata
IMG_NAME=postoffice
IMG_VER=1.0

docker build -t $IMG_NAME:$IMG_VER .
docker run --name ${IMG_NAME}Container -p 6379:6379 -v $DATA_DIR:/data -d $IMG_NAME:$IMG_VER redis-server --appendonly yes
sleep 2
go build -o challenge
echo ""
echo "running challenge..."
echo "make a request to /users using the following methods: GET, POST, DELETE"
echo "  (GET|POST|DELETE) /users/:username"
echo "POSTing to /user/ should create an key in redis, read with GET and remove with DELETE"
./challenge
