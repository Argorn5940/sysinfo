# sysinfo

sysinfoは、システムの様々な情報（CPU、メモリ、ディスク、ネットワーク）を収集して表示するCLIツールです。

## インストール

```sh
go install github.com/Argorn5940/sysinfo@latest
```

## 使用方法

### 全ての情報を表示

```sh
sysinfo all
```

### 特定の情報を表示

```sh
sysinfo cpu     # CPU情報
sysinfo memory  # メモリ情報
sysinfo disk    # ディスク情報
sysinfo network # ネットワーク情報
```

### バージョン情報

```sh
sysinfo version
```

## サブコマンド

| コマンド  | 説明                     |
|-----------|--------------------------|
| `all`     | すべてのシステム情報を表示 |
| `cpu`     | CPU情報を表示             |
| `memory`  | メモリ情報を表示          |
| `disk`    | ディスク情報を表示        |
| `network` | ネットワーク情報を表示    |
| `version` | バージョン情報を表示      |

## ライセンス

このプロジェクトは [MITライセンス](LICENSE) のもとで提供されています。
