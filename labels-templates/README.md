# 如何使用 github sync labels

## 安装
使用 make build 安装
```bash
make build
```

设置 Github_TOKEN 环境变量
```bash
export GITHUB_TOKEN=xxxx
```

然后，在项目跟目录，通过 python 脚本获取到组织中所有的仓库地址：
```bash
# python3 repo.py 
OpenIMSDK/Open-IM-Server
OpenIMSDK/Open-IM-SDK-iOS
OpenIMSDK/Open-IM-SDK-Android
OpenIMSDK/Open-IM-SDK-Flutter
OpenIMSDK/openim-sdk-core
OpenIMSDK/Open-IM-Uniapp-Demo
OpenIMSDK/Open-IM-iOS-Demo
OpenIMSDK/Open-IM-Android-Demo
OpenIMSDK/Open-IM-SDK-Uniapp
OpenIMSDK/cpp_go
OpenIMSDK/Open-IM-SDK-Web
OpenIMSDK/Open-IM-Flutter-Demo
OpenIMSDK/Open-IM-SDK-ReactNative
OpenIMSDK/Open-IM-PC-Web-Demo
OpenIMSDK/OpenIM-Docs
OpenIMSDK/openim-sdk-core-ios
OpenIMSDK/sdk_advanced_function
OpenIMSDK/Open-IM-SDK-Dart
OpenIMSDK/Open-IM-SDK-Web-Wasm
OpenIMSDK/open_utils
OpenIMSDK/open_log
OpenIMSDK/getcdv3
OpenIMSDK/Open-IM-SDK-Unity
OpenIMSDK/Open-IM-SDK-Core-IndexDB-Doc
OpenIMSDK/rockscache
OpenIMSDK/Open-IM-Server-k8s-deploy
OpenIMSDK/community
OpenIMSDK/.github
OpenIMSDK/automation
OpenIMSDK/OpenKF
```

然后，修改 actions 文件：
在 sync_labels_openim.yml 文件中:
```yaml
name: Kubecub Sync labels

on:
  push:
    branches:
      - main
    # paths:
    #   - .github/sync_labels.yml

jobs:
    build:
      runs-on: ubuntu-latest
      steps:
        - name: Checkout
          uses: actions/checkout@1.0.0
        - name: Github lables pull and synchronize
          uses: kubecub/github-label-syncer@v2.0.0
          with:
            manifest: labels-templates/openim-yaml.yaml
            token: ${{ secrets.BOT_GITHUB_TOKEN }}
            repository: |
                OpenIMSDK/Open-IM-Server
                OpenIMSDK/Open-IM-SDK-iOS
                OpenIMSDK/Open-IM-SDK-Android
                OpenIMSDK/Open-IM-SDK-Flutter
                OpenIMSDK/openim-sdk-core
                OpenIMSDK/Open-IM-Uniapp-Demo
                OpenIMSDK/Open-IM-iOS-Demo
                OpenIMSDK/Open-IM-Android-Demo
                OpenIMSDK/Open-IM-SDK-Uniapp
                OpenIMSDK/cpp_go
                OpenIMSDK/Open-IM-SDK-Web
                OpenIMSDK/Open-IM-Flutter-Demo
                OpenIMSDK/Open-IM-SDK-ReactNative
                OpenIMSDK/Open-IM-PC-Web-Demo
                OpenIMSDK/OpenIM-Docs
                OpenIMSDK/openim-sdk-core-ios
                OpenIMSDK/sdk_advanced_function
                OpenIMSDK/Open-IM-SDK-Dart
                OpenIMSDK/Open-IM-SDK-Web-Wasm
                OpenIMSDK/open_utils
                OpenIMSDK/open_log
                OpenIMSDK/getcdv3
                OpenIMSDK/Open-IM-SDK-Unity
                OpenIMSDK/Open-IM-SDK-Core-IndexDB-Doc
                OpenIMSDK/rockscache
                OpenIMSDK/Open-IM-Server-k8s-deploy
                OpenIMSDK/community
                OpenIMSDK/.github
                OpenIMSDK/automation
                OpenIMSDK/OpenKF
          env:
            GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```


# `/examples`

Examples for your applications and/or public libraries.

Examples:

* https://github.com/nats-io/nats.go/tree/master/examples
* https://github.com/docker-slim/docker-slim/tree/master/examples
* https://github.com/hashicorp/packer/tree/master/examples