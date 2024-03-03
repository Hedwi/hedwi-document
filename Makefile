all: main deploy


saas: api docs mail main deploy

api:
	cd ./hedwi-api/ && rm -rf build/* && make html && cd ..
	rm -rf static/hedwi-api/*
	cp -r  ./hedwi-api/build/html/* static/hedwi-api/
	date -r ./hedwi-api/build/html/index.html "+%Y-%m-%d %H:%M:%S" > date.modified

mail:
	cd ./hedwi-mail/ && rm -rf build/* && make html && cd ..
	rm -rf static/hedwi-mail/*
	cp -r  ./hedwi-mail/build/html/* static/hedwi-mail/
	date -r ./hedwi-mail/build/html/index.html "+%Y-%m-%d %H:%M:%S" > mail.modified

docs:
	cd ./hedwi-docs/ rm -rf build/* && make html && cd ..
	rm -rf static/hedwi-docs/*
	cp -r  ./hedwi-docs/build/html/* static/hedwi-docs/
	date -r ./hedwi-docs/build/html/index.html "+%Y-%m-%d %H:%M:%S" > docs.modified

meet:
	cd ./hedwi-meet/ && rm -rf build/* && make html && cd ..
	rm -rf static/hedwi-meet/*
	cp -r  ./hedwi-meet/build/html/* static/hedwi-meet/
	date -r ./hedwi-meet/build/html/index.html "+%Y-%m-%d %H:%M:%S" > meet.modified



main:
	GOOS=linux GOARCH=amd64 go build -o hedwi-document main.go

deploy:
	echo 'hedwi-document'
	#scp -r ./hedwi-api/build server1:/data/www/hedwi-document/
	scp -r hedwi-document server1:/data/www/api/hedwi-document/_tmp
	ssh server1 "cd /data/www/api/hedwi-document/ && mv _tmp hedwi-document && supervisorctl -c /data/www/api/hedwi-api_supervisord.conf restart hedwi-document"
	#scp -r ./hedwi-api/build server2:/data/www/hedwi-document/
	scp -r hedwi-document server2:/data/www/api/hedwi-document/_tmp
	ssh server2 "cd /data/www/api/hedwi-document/ && mv _tmp hedwi-document && supervisorctl -c /data/www/api/hedwi-api_supervisord.conf restart hedwi-document"
	#scp -r ./hedwi-api/build server3:/data/www/hedwi-document/
	scp -r hedwi-document server3:/data/www/api/hedwi-document/_tmp
	ssh server3 "cd /data/www/api/hedwi-document/ && mv _tmp hedwi-document && supervisorctl -c /data/www/api/hedwi-api_supervisord.conf restart hedwi-document"
