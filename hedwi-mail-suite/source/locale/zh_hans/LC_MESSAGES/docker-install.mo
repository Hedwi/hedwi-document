��                        �     �       &     
   E  0   P     �      �     �  �   �  z   �  c   .     �  +   �     �  "   �  C     �   J  �   �  2   �  i   �  2   [     �     �     �     �     �     �     �  .   �  �       �	     �	  $   �	  	   �	     �	     
     
     9
  �   ;
  ^     U   {     �  )   �       !     =   8  d   v  �   �  .   �  W   �  *   /     Z     g  	   �     �  	   �     �     �     �   0. Requirements 1. configure DNS 2. Install docker, docker-compose, git 3. Install 4. Automatic generation of free SSL certificates 5. Restart the service 6. Use https to access and login A Access https://work.example.com/login?admin=true (work.example.com is the domain name used), use the administrator account (the email and password configured) to log in Access https://work.example.com/hello page, continue to add users, teams, etc. After obtaining the certificate successfully, the certificate will be written to inbox/certs/cert.key inbox/certs/cert.crt After running, http 80 port can be accessed, should be able to access http://configured domain name For example: Minimum 2G memory 2 CPU 64-bit Linux system Start docker containers Stop all running docker containers The hostname can be any secondary domain name prefix, here use work The service has built-in Let’s Encrypt HTTP-01 verification service interface, running this command will request the /acme/create interface to create a certificate The service will automatically create an administrator account (for configuring the system, adding users, adding teams, etc.) and a normal user account admin@your domain name (for using the system service) The web access address is https://work.example.com This command will use the generated free certificate to replace nginx/certs/cert.key nginx/certs/cert.crt `Configuration file example <./env-example.html>`_ docker deploy example.com is your domain name hostname record type record value work x.x.x.x x.x.x.x is the public IP address of the server Project-Id-Version: Hedwi work suite 
Report-Msgid-Bugs-To: 
POT-Creation-Date: 2024-12-01 12:00+0800
PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE
Last-Translator: FULL NAME <EMAIL@ADDRESS>
Language: zh_hans
Language-Team: zh_hans <LL@li.org>
Plural-Forms: nplurals=1; plural=0;
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 8bit
Generated-By: Babel 2.15.0
 0. 环境要求 1. 配置DNS 2. 安装docker, docker-compose, git 3. 安装 4. 自动生成免费SSL证书 5. 重启服务 6. 使用https访问 登录 A 访问 https://work.example.com/login?admin=true (work.example.com为使用的域名), 使用管理员账号登录(配置的邮箱和密码) 访问 https://work.example.com/hello 页面,  继续添加用户、团队等等。 成功获取证书之后，证书将会被写入到 inbox/certs/cert.key inbox/certs/cert.crt  运行之后,80端口可以访问。应该可以通过 http://work.example.com 访问 例如: 最低2G内存 2核CPU  linux 64位系统 启动docker容器 停止所有运行的docker容器 主机名可以是任意二级域名前缀，这里使用work 本服务有内置的let's Encrypt HTTP-012验证服务接口，请求/acme/create可以创建证书 本服务会自动创建管理员账号（根据配置文件中的邮箱和密码）用于配置系统、添加用户、团队等等，还有一个普通用户账户admin@你的域名，用于使用服务 网页访问地址为 https://work.example.com 此条命令将会使用生成的证书替换 nginx/certs/cert.key nginx/certs/cert.crt `配置文件示例 <./env-example.html>`_ docker部署 example.com为你的域名 主机名 记录类型 记录值 work x.x.x.x x.x.x.x 为服务器的IP地址 