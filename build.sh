docker stop $(docker ps -a | awk '{ print $1}' | head -n -1 | tail -n +2)
docker rm $(docker ps -a | awk '{ print $1}' | head -n -1 | tail -n +2)
docker rmi $(docker images | awk '{ print $3}' | head -n -5 | tail -n +2)

docker stop $(docker ps -a | awk '{ print $1}' | tail -n +2)
docker rm $(docker ps -a | awk '{ print $1}' | tail -n +2)
docker rmi $(docker images | awk '{ print $3}' | tail -n +2)

cd db
make
make run
echo "-------------db init success----------------"
cd ../login
make
make run
echo "-------------login_app build success----------------"
cd ../register
make
make run
echo "-------------register_app build success----------------"
cd ../dorm
make
make run
echo "-------------dorm_app build success----------------"
cd ../nodejs
make
make run
echo "-----------------nodejs init success--------------------"