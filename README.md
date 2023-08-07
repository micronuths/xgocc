
## 转换模式 
s2t 
- Simplified Chinese to Traditional Chinese
- 简体 -> 繁体
t2s 
- Traditional Chinese to Simplified Chinese
- 繁体 -> 简体
s2tw 
- Simplified Chinese to Traditional Chinese (Taiwan Standard)
- 简体 -> 繁体（台湾）
tw2s 
- Traditional Chinese (Taiwan Standard) to Simplified Chinese
- 繁体（台湾） -> 简体 
s2hk 
- Simplified Chinese to Traditional Chinese (Hong Kong Standard)
- 简体 -> 繁体（香港）
hk2s 
- Traditional Chinese (Hong Kong Standard) to Simplified Chinese
- 繁体（香港）-> 简体
s2twp 
- Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom
- 简体 -> 繁体（标准台湾，带台湾俗语）
tw2sp 
- Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom
- 繁体（台湾）-> 简体（带大陆俗语）
t2tw 
- Traditional Chinese (OpenCC Standard) to Taiwan Standard
- 繁体（OpenCC 标准）-> 繁体（台湾标准）
t2hk 
- Traditional Chinese (OpenCC Standard) to Hong Kong Standard
- 繁体（OpenCC 标准）-> 繁体（香港标准）

## 命令行用法
gocc  --input ..\input.txt --config t2s --output ../output.txt

echo "我们是工农子弟兵" | gocc

## api调用
```

package main

import (
    "fmt"
    "log"
    
    "github.com/micronut/xgocc"
)

func main() {
    s2t, err := xgocc.New("s2t")
    if err != nil {
        log.Fatal(err)
    }
    in := `自然语言处理是人工智能领域中的一个重要方向。`
    out, err := s2t.Convert(in)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n%s\n", in, out)
    //自然语言处理是人工智能领域中的一个重要方向。
    //自然語言處理是人工智能領域中的一個重要方向。
}
```
源项目
https://github.com/liuzl/gocc
