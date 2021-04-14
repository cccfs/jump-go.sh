# jump.sh
a simple script for ssh to Aliyun ECS nodes based on Name and Instance ID, with tab auto-completion

```bash
$ jump-go.sh dev-limestone- # pressing TAB
dev-oss-web-wz9d7qqtzw8    dev-oss-web-wz9d7qqtzw9   dev-oss-web-wz9d7qqtzw7 
```

## Setup

```bash
# set env variables
export ALICLOUD_ACCESS_KEY="xxxxxxxxxxxxxx"
export ALICLOUD_SECRET_KEY="xxxxxxxxxxxxx"
export ALICLOUD_REGION="cn-shenzhen"

# clone it
git clone https://github.com/cccfs/jump-go.sh ~/.jumpgohost
echo '. ~/.jumpgohost/auto-completion.sh' >> ~/.bashrc # or .zshrc, depending which shell you use

# and please modify dev.cn-shenzhen.sh to fit your environment
cd ~/.jumphost/generator.d
cp dev.cn-shenzhen.example dev.cn-shenzhen.sh

# default ECS user `root` or add ssh username as ECS Tag `SshUser` to your ECS instances
```

## Access instances with SSH key

There is no explicit way to config ssh key on specific connection.
We can use ssh-add to set up connection with SSH automatically.
