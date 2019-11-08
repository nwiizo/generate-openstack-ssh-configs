# generate-openstack-ssh-configs
[Semi-]Automatically generate ssh config files for your cloud servers

### Explanation
- No set IdentityFile
- Not set automatically if you have some network interfaces or HostName

```
#Connection ID:c5316248-5756-48b6-b4c4-404a49deca9d
Host web001
#Change it
#  HostName 172.16.0.1
#  HostName 10.0.0.1
  user root
  port 22
  IdentityFile ~/.ssh/keys/key

#Connection ID:7951fcfe-1aa7-450d-880a-bf194abb7a74
Host db01
  HostName 10.0.0.2
  user root
  port 22
  IdentityFile ~/.ssh/keys/key
```
