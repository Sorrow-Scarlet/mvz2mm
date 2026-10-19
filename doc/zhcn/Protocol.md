# 协议

## 通用

> [!IMPORTANT]
> 压缩包内任何BepInEx文件夹以外的文件夹将**原封不动**的安装到MVZ2根目录下！

- 不管你是什么包，.pack .pck .mvz2pack .mvz2 .mvz2mod .zip，都给我老老实实用zip的压缩算法
- 即便本MOD管理器允许嵌套一层文件夹打包，依旧建议你的压缩包不要嵌套文件夹
- 对于单个的普通mod, 应存在manifest.json
- 对于整合包mod,应存在mvz2pack.json

## 理想结构

- manifest.json
- /BepInEx/plugins/你的mod.dll
- /MinecraftVSZombies2_Data/StreamingAssets/Mods/你的mod资源文件夹名

## manifest.json

此json文件用于描述mod的基础信息

### 模板

```json
{
  "id": "com.username.modid",
  "name": "示例Mod",
  "version": "1.3.0",
  "author": "someone",
  "dependencies": ["com.example.deps"]
}
```

### 介绍

- id:string 为纯小写，用于识别唯一mod，你可以放肆地个性化一些，不过最好遵从域名规范
- name:string 为任意UTF-8字符串，用于在管理器中显示你的MOD名
- version:string 为**语义化版本**字符串, 固定格式为X.Y.Z
- author:string 为作者名，你爱叫啥叫啥
- dependencies:[]string 为依赖项字符串数组，允许声明多个mod为依赖项，仅供mod管理器识别依赖项

## mvz2pack.json

> [!NOTE]
> 大理石，大理石，我要DISS你

此json文件用于描述合集mod/mod整合包

### 模板

此模板仅声明必须项，其他项会被丢弃！

```json
{
  "Name": "MVZ2 模组合集 v2026-10-07",
  "Author": "QoZnoS",
  "CreatedUtc": "2026-10-07T08:50:13.7233751Z",
  "Mods": ["DSHCore", "LaboratoryMod", "Loadout", "ModularCurse", "Replay"]
}
```

### 说明

- Name:string 整合包名你爱叫这玩意啥这玩意就叫啥
- Author:string 作者名，你爱叫啥叫啥
- CreatedUtc:string 任何符合`ISO 8601` 或`RFC 3339` 或`GB/T 7408`的时间字符串
- Mods:[]string 这个地方注册你的dll名称，**我要狠狠diss大理石这一点，注册名和产物名不一致的问题。** **必须必须必须必须必须必须必须必须必须一致！**
