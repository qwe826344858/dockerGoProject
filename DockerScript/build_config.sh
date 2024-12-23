#!/bin/bash 

makeGoFile() {
  # 清空文件
  > envGoConfig.go

  ip=$(getCurrentIp)
	{
		echo 'package envcfg'
		echo ''
		echo 'type EnvConfig struct {'
		echo '	RedisConf RedisConfig `json:"redis_conf"`'
		echo '	MysqlConf MysqlConfig `json:"mysql_conf"`'
		echo '	GRpcConf GRpcConfig `json:"grpc_config"`'
		echo '}'
		echo ''
		echo 'type RedisConfig struct {'
		echo '	Host string `json:"host"`'
		echo '}'
		echo ''
		echo 'type MysqlConfig struct {'
		echo '	Host     string `json:"host"`'
		echo '	Username string `json:"username"`'
		echo '	Password string `json:"password"`'
		echo '}'
		echo ''
		echo 'type GRpcConfig struct {'
    echo '	Host     string `json:"host"`'
    echo '}'
    echo ''
		echo 'func GetEnvConfig() *EnvConfig {'
		echo '	return &EnvConfig{'
		echo '		RedisConf: RedisConfig{'
		echo "			Host: \"$ip\","
		echo '		},'
		echo '		MysqlConf: MysqlConfig{'
		echo "			Host: \"$ip\","
		echo '			Username: "root",'
		echo '			Password: "zoneslee",'
		echo '		},'
		echo '		GRpcConf: GRpcConfig{'
    echo "			Host: \"$ip\","
    echo '		},'
		echo '	}'
		echo '}'
	} > envGoConfig.go
  
  echo "生成go配置文件成功 envGoConfig.go"
}

makePythonFile() {
  # 清空文件
  > envPythonConfig.py

  ip=$(getCurrentIp)
  {
    echo 'def getEnvConfig():'
    echo '    return {'
    echo '        "redis_conf": {'
    echo "            \"host\": \"$ip\","  # 使用变量替换
    echo '        },'
    echo '        "mysql_conf": {'
    echo "            \"host\": \"$ip\","  # 使用变量替换
    echo '            "username": "root",'
    echo '            "password": "zoneslee",'
    echo '        },'
    echo '        "grpc_conf": {'
    echo "            \"host\": \"$ip\","  # 使用变量替换
    echo '        },'
    echo '    }'
  } > envPythonConfig.py

  echo "生成python配置文件成功 envPythonConfig.py"
  cp envPythonConfig.py /home/lighthouse/test_py_beta/docker_project/tools/
  echo "配置文件复制成功!"
}

getCurrentIp() {
  # 使用 hostname 命令获取本地 IP 地址
  EXTERNAL_IP=$(hostname -I | awk '{print $1}')
  echo "${EXTERNAL_IP}"
}


language=$1
if [ -z "$language" ]; then
  echo "请输入要生成配置的语言文件 示例: go python"
  exit 1
fi

if [ "$language" = "go" ]; then
   makeGoFile
elif [ "$language" = "python" ]; then
   makePythonFile
else
   echo "未知的语言: $language"
   exit 1
fi



