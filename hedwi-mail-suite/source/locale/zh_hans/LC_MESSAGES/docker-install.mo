��          �               �     �     �  &   �  
   �  0   �           (     I  �   K     C  +   P  C   |  2   �  <   �     0     >     ^     g     s     �     �  .   �  �  �     D     T  $   a  	   �     �     �     �     �  �   �     �  )   �  =   �  .   .  4   ]     �     �  	   �     �  	   �     �     �     �   0. Requirements 1. configure DNS 2. Install docker, docker-compose, git 3. Install 4. Automatic generation of free SSL certificates 5. Restart the service 6. Use https to access and login A Access https://work.example.com/login?admin=true (work.example.com is the domain name used), use the administrator account (the email and password configured) to log in Access https://work.example.com/hello page, continue to add users, teams, etc. For example: Minimum 2G memory 2 CPU 64-bit Linux system The hostname can be any secondary domain name prefix, here use work The web access address is https://work.example.com `Configuration file example </mail-suite/env-example.html>`_ docker deploy example.com is your domain name hostname record type record value work x.x.x.x x.x.x.x is the public IP address of the server Project-Id-Version: Hedwi work suite 
Report-Msgid-Bugs-To: 
POT-Creation-Date: 2024-11-26 22:33+0800
PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE
Last-Translator: FULL NAME <EMAIL@ADDRESS>
Language: zh_hans
Language-Team: zh_hans <LL@li.org>
Plural-Forms: nplurals=1; plural=0;
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 8bit
Generated-By: Babel 2.15.0
 0. 环境要求 1. 配置DNS 2. 安装docker, docker-compose, git 3. 安装 4. 自动生成免费SSL证书 5. 重启服务 6. 使用https访问 登录 A 访问 https://work.example.com/login?admin=true (work.example.com为使用的域名), 使用管理员账号登录(配置的邮箱和密码) 访问 https://work.example.com/hello 页面,  继续添加用户、团队等等。 例如: 最低2G内存 2核CPU  linux 64位系统 主机名可以是任意二级域名前缀，这里使用work 网页访问地址为 https://work.example.com `配置文件示例 </mail-suite/env-example.html>`_ docker部署 example.com为你的域名 主机名 记录类型 记录值 work x.x.x.x x.x.x.x 为服务器的IP地址 