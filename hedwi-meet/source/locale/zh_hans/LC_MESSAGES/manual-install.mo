��          �               �     �  	   �     �     �  *     	   1     ;     L     P     j     v  	   �     �     �     �     �     �  0   �     �  
   �     	            4     ;   O  �  �       	        '     7     M  	   k     u     �     �     �     �  	   �     �     �  �   �     U     X  (   e     �     �     �     �     �  ,   �  3   �   0. Requirements 0.0.0.0/0 1. Install dependencies 2. Install meet server 3. nginx and ssl certificate configuration 3478-3500 4. Restart nginx 443 5. Firewall configuration 50000-60000 6. Configure DNS 7880-7890 80 Allow Firewall configuration IP Manual deployment Minimum 1G memory 1 core CPU linux 64-bit system Policy Port range Protocol TCP UDP `Configuration file example </meet/meet_conf.html>`_ `Nginx configuration file example </meet/nginx_conf.html>`_ Project-Id-Version: hedwi-meet 
Report-Msgid-Bugs-To: 
POT-Creation-Date: 2024-11-27 11:10+0800
PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE
Last-Translator: FULL NAME <EMAIL@ADDRESS>
Language: zh_hans
Language-Team: zh_hans <LL@li.org>
Plural-Forms: nplurals=1; plural=0;
MIME-Version: 1.0
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 8bit
Generated-By: Babel 2.15.0
 0. 环境要求 0.0.0.0/0 1. 安装依赖 2. 安装 meet server 3. nginx 和 ssl 证书配置 3478-3500 4. 重启 nginx 443 5. 防火墙配置 50000-60000 6. 配置 DNS 7880-7890 80 允许 开放25, 465, 993, 443端口用于外部访问。邮件服务不需要nginx。如果使用nginx，需要修改配置文件中的端口。 IP 手动部署 最低1G内存 1核CPU linux 64位系统 策略 端口范围 协议 TCP UDP `配置文件示例 </meet/meet_conf.html>`_ `Nginx 配置文件示例 </meet/nginx_conf.html>`_ 