docker build -f ./docker/init_data/initdata.dockerfile -t init_data . 
docker run --net pms_services --name init_data init_data
docker rm init_data