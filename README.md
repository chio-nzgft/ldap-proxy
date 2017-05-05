# ldap-proxy

Note need python 2.7 & golang

must install python library ldap

copy github.com/vjeantet/ldapserver as ldap server in ldapserver directory

1. Edit conf.json:
2. exec --> go run server.go (or build server.go & exec it)
3. run sample file ..... or other you access ldap file
4. stop server with ctrl-c

```
[server] 2017/05/05 09:52:11 Listening on 192.168.0.101:10389
[server] 2017/05/05 09:52:16 Connection client [1] from 192.168.0.101:51066 accepted
[server] 2017/05/05 09:52:16 <<< 1 - BindRequest - hex=&{301802010160130201030404656e6c6980085040737377307264}
2017/05/05 09:52:16 AD server 192.168.0.100:389
2017/05/05 09:52:16 user admin
2017/05/05 09:52:16 pass P@ssw0rd
2017/05/05 09:52:16 &{DN:CN=admin,OU=user,DC=tw,DC=nzgft Attributes:[0xc420068240 0xc4200682c0 0xc420068300 0xc420068340 0xc420068380 0xc420068400 0xc420068440 0xc420068480 0xc4200684c0 0xc420068500 0xc420068580 0xc4200685c0 0xc420068600 0xc420068640 0xc420068680 0xc4200686c0 0xc420068700 0xc420068740 0xc420068780 0xc4200687c0 0xc420068800 0xc420068840 0xc420068880 0xc4200688c0 0xc420068900 0xc420068940 0xc420068980 0xc4200689c0 0xc420068a00 0xc420068a40 0xc420068a80 0xc420068b00 0xc420068b40 0xc420068b80]}
[server] 2017/05/05 09:52:16 >>> 1 - BindResponse - hex=300c02010161070a010004000400
[server] 2017/05/05 09:52:16 <<< 1 - ExtendedRequest - hex=&{301e02010277198017312e332e362e312e342e312e343230332e312e31312e33}
[server] 2017/05/05 09:52:16 >>> 1 - LDAPResult - hex=302f020102302a0a0135040004234f7065726174696f6e206e6f7420696d706c656d656e74656420627920736572766572
[server] 2017/05/05 09:52:16 <<< 1 - UnbindRequest - hex=&{30050201034200}
[server] 2017/05/05 09:52:16 client 1 close()
[server] 2017/05/05 09:52:16 client 1 close() - stop reading from client
[server] 2017/05/05 09:52:16 client 1 close() - Abandon signal sent to processors
[server] 2017/05/05 09:52:16 client [1] request processors ended
[server] 2017/05/05 09:52:16 client [1] connection closed
[server] 2017/05/05 09:54:53 gracefully closing client connections...
[server] 2017/05/05 09:54:53 all clients connection closed

```

用 golang 當 ldap 主機 .. 

想說當不同的 domain 時　go-ldap server 去找不同的 LDAP 或去找 mysql 認證 

先寫個簡單的用 AD 的 LDAP 當參考 .....

所有原始碼 : https://github.com/chio-nzgft/ldap-proxy

先編輯  conf.json  檔案 

改成服務的主機 & AD 主機的 IP

修改  client-ldap-proxy-test.py  內的 

print authenticate("192.168.0.101","admin","P@ssw0rd")

修改 你的 主機及 帳號密碼 資訊

看執行後成果 是否是認證成功  Succesfully authenticated

這樣你就可以修改  server.go  

看 username 就可找不同 AD 或 ldap

例如 :
admin@test1.com 找 ad1
admin@test1.com 找 ad2

PS : 我用 CentOS7 ...

 安裝 python-ldap 用 yum install python-ldap
 
