# GoBase

修复记录：

1. 将模块路径从 `module GoBase` 修改为 `module github.com/fan-chao-sys/GoBase`，以匹配代码中的导入路径。
2. 将 `main.go` 的包名改为 `package main`，并保留主函数（`main()`）。
3. 将 `test.go` 中的 `main()` 移除并调整为 `package main`（提供辅助函数和导出变量）。
4. 保持 `pkg1` 和 `pkg2` 两个子包不变，使用空导入顺序触发 init。

如何运行：

在项目根目录下运行：

```powershell
cd d:\Go\Project\GoBase
go run .
```

输出预期包含 init 初始化打印顺序，以及主函数输出。