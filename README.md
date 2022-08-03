# Geacon

**Using Go to implement CobaltStrike's Beacon**

# 原项目
geacon
https://github.com/darkr4y/geacon

# 说明
简单适配了一个profile配置文件,可直接拿来修改使用.

详细文章:https://www.nctry.com/2595.html

# 解密key文件

BeaconTool.jar 用于解密teamserver端的key文件.
然后放入config.go中.

Usage:
[*] parse the .beacon_keys to RSA private key and public key in pem format
	BeaconTool -i .beacon_keys -rsa
[*] compile geacon with the public key from .beacon_keys,which use default c2profile config for communication
	BeaconTool -i .beacon_keys -compile geacon_sourcecode_folder


java -jar BeaconTool.jar -i .cobaltstrike.beacon_keys -rsa

# 流量截图
GET请求
![get](https://github.com/TRYblog/edit-gencon/blob/main/get.png "get")

POST请求
![post](https://github.com/TRYblog/edit-gencon/blob/main/post.png "post")
