all: main deploy


saas: api mail meet main deploy

doc:api mail meet


i18n:
	cd hedwi-mail-suite && cd source && sphinx-build -b gettext ./ ../build/gettext && cd .. && sphinx-intl update -p build/gettext -l zh_hans -l en_us
	cd hedwi-meet && cd source && sphinx-build -b gettext ./ ../build/gettext && cd .. && sphinx-intl update -p build/gettext -l zh_hans -l en_us


api:
	cd ./hedwi-api/ && rm -rf build/* && make html && cd ..
	rm -rf static/hedwi-api/*
	cp -r  ./hedwi-api/build/html/* static/hedwi-api/
	date -r ./hedwi-api/build/html/index.html "+%Y-%m-%d %H:%M:%S" > date.modified
	cd ./hedwi-api/ && rm -rf build/* && make -e SPHINXOPTS="-D language='zh_hans'" html && cd ..
	rm -rf document/hedwi-api/zh-hans/*
	cp -r  ./hedwi-api/build/html/* document/hedwi-api/zh-hans/
	cd ./hedwi-api/ && rm -rf build/* && make -e SPHINXOPTS="-D language='en_us'" html && cd ..
	rm -rf document/hedwi-api/en-us/*
	cp -r  ./hedwi-api/build/html/* document/hedwi-api/en-us/
	date -r ./hedwi-api/build/html/index.html "+%Y-%m-%d %H:%M:%S" > mail.modified


mail:
	cd ./hedwi-mail-suite/ && rm -rf build/* && make -e SPHINXOPTS="-D language='zh_hans'" html && cd ..
	rm -rf document/hedwi-mail-suite/zh-hans/*
	cp -r  ./hedwi-mail-suite/build/html/* document/hedwi-mail-suite/zh-hans/
	rm -rf ../hedwi-inbox/document/hedwi-mail-suite/zh-hans/*
	cp -r  ./hedwi-mail-suite/build/html/* ../hedwi-inbox/document/hedwi-mail-suite/zh-hans
	cd ./hedwi-mail-suite/ && rm -rf build/* && make -e SPHINXOPTS="-D language='en_us'" html && cd ..
	rm -rf document/hedwi-mail-suite/en-us/*
	cp -r  ./hedwi-mail-suite/build/html/* document/hedwi-mail-suite/en-us/
	rm -rf ../hedwi-inbox/document/hedwi-mail-suite/en-us/*
	cp -r  ./hedwi-mail-suite/build/html/* ../hedwi-inbox/document/hedwi-mail-suite/en-us
	date -r ./hedwi-mail-suite/build/html/index.html "+%Y-%m-%d %H:%M:%S" > mail.modified

meet:
	cd ./hedwi-meet/ && rm -rf build/* && make -e SPHINXOPTS="-D language='zh_hans'" html && cd ..
	rm -rf document/hedwi-meet/zh-hans/*
	cp -r  ./hedwi-meet/build/html/* document/hedwi-meet/zh-hans/
	rm -rf ../hedwi-inbox/document/hedwi-meet/zh-hans/*
	cp -r  ./hedwi-meet/build/html/* ../hedwi-inbox/document/hedwi-meet/zh-hans
	cd ./hedwi-meet/ && rm -rf build/* && make -e SPHINXOPTS="-D language='en_us'" html && cd ..
	rm -rf document/hedwi-meet/en-us/*
	cp -r  ./hedwi-meet/build/html/* document/hedwi-meet/en-us/
	rm -rf ../hedwi-inbox/document/hedwi-meet/en-us/*
	cp -r  ./hedwi-meet/build/html/* ../hedwi-inbox/document/hedwi-meet/en-us
	date -r ./hedwi-meet/build/html/index.html "+%Y-%m-%d %H:%M:%S" > mail.modified



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
