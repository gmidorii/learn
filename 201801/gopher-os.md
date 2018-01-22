# Gopher-OS

## 概要
gopher-osを読んでみます。  

## What's Gopher-OS?
- 64bit POSIX準拠のOS
- Ticklessカーネル
  - idle状態のCPUが完全停止(=消費電力減)
- Linux Compatibleなsyscall実装
- Goがlow levelなcodeが書けることをproofするために作成
- ring-0で実行
  - リングプロテクションの考え方
  - ring-0が最も直接ハードウェアとやり取りする

## [Status](https://github.com/achilleasa/gopher-os/blob/master/STATUS.md) (実装されてる機能)
### Core Kernel
- Bootloader-related
  - Computer起動時にOSを起動するプログラム
  - /boot/cmdline.txt?
  - memory
    - virtual memory
      - 物理メモリの何倍も大きい
    - memory mapping
      - File/Imageをaddress spaceにmapping
      - ProcessのVirtual memoryにmapされる
  - Framebuffer
    - abstract layer
    - graphicを担当
- CPU
  - CPU Identification Wrapper
    - Processor Identification
    - 機能
      - 新機能の有無
      - CPUベンダーの判別
      - Processor Serial Numberの返却
- Memory Management
  - Physical Frame Allocate?
  - VMM(Virtual Memory Manager) System
    - Page Table
    - Virtual address space reservation

## SRC
### kmain
- rt0を利用?
  - [rt0](https://github.com/lpsantil/rt0)
  - Minimal C runtime
- g0 structをsetting
- rt0がg0をinvokeする
- not return. If return, halt the CPU.
- Kmain
  - `allocator.Init(kernelStart, kernelEnd uintptr)`
  - Kmain実行時にpointerを渡している
  - `hal.DetectHardware()` でHardwareのInitializeをしている

## Develop
- GOPATHの設定を変更
  - export GOPATH=`pwd`:$GOPATH

## 参考
- [Tickless](http://enakai00.hatenablog.com/entry/20111117/1321508379)
- [リングプロテクションWiki](https://ja.wikipedia.org/wiki/%E3%83%AA%E3%83%B3%E3%82%B0%E3%83%97%E3%83%AD%E3%83%86%E3%82%AF%E3%82%B7%E3%83%A7%E3%83%B3)
- [Linux Kernel](http://archive.linux.or.jp/JF/JFdocs/The-Linux-Kernel-4.html)
- [framebuffer](https://qiita.com/edo_m18/items/95483cabf50494f53bb5)