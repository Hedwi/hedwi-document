all: build-hedwiapi hedwiapi

build-hedwiapi:
	cd ./hedwi-api/ && make html && cd ..
	GOOS=linux GOARCH=amd64 go build

build-hedwimail:
	cd ./hedwi-mail/ && make html && cd ..
	GOOS=linux GOARCH=amd64 go build

hedwiapi:
	echo 'hedwi-api'
	scp -r ./hedwi-api/build server1:/data/www/hedwi-document/
	scp -r hedwi-document server1:/data/www/hedwi-document/_tmp
	ssh server1 "cd /data/www/hedwi-document/ && mv _tmp hedwi-document && supervisorctl -c /data/www/hedwi-api_supervisord.conf restart hedwi-document"
	scp -r ./hedwi-api/build server2:/data/www/hedwi-document/
	scp -r hedwi-document server2:/data/www/hedwi-document/_tmp
	ssh server2 "cd /data/www/hedwi-document/ && mv _tmp hedwi-document && supervisorctl -c /data/www/hedwi-api_supervisord.conf restart hedwi-document"
	scp -r ./hedwi-api/build server3:/data/www/hedwi-document/
	scp -r hedwi-document server3:/data/www/hedwi-document/_tmp
	ssh server3 "cd /data/www/hedwi-document/ && mv _tmp hedwi-document && supervisorctl -c /data/www/hedwi-api_supervisord.conf restart hedwi-document"
