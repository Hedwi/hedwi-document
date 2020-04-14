all:
	cd ./docs && make html && cd ..
	GOOS=linux GOARCH=amd64 go build
	scp -r ./docs/build hedwi-docs:/data/www/hedwi-docs/
	scp -r hedwi-docs hedwi-docs:/data/www/hedwi-docs/_tmp
	ssh hedwi-docs "cd /data/www/hedwi-docs/ && mv _tmp hedwi-docs && supervisorctl -c ./hedwi-docs_supervisord.conf reload"
