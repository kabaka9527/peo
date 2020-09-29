## peo

🎮用于建立翼龙面板的自动售卖系统和附加控制系统，自动化你的翼龙面板出售。

目前已在 [demo](https://order.ntmc.tech) 稳定运行半年并跟进开发版本的部署

##### 特性

- [x] 登录、注册、找回密码、修改密码、改绑邮箱
- [x] 主页商品展示、建立订单、支持余额支付或 KEY 支付
- [x] 用户控制台：展示用户服务器信息、跳转控制台、运行时间记录、用同种 KEY 或余额自助续费
- [x] 工单系统
- [x] 用户消息通知
- [x] 用户可分享的公共相册
- [x] 管理员控制台：添加商品、整合包（Nest.Egg）、处理工单、管理相册、批量添加 KEY、导出KEY
- [x] 周期任务：刷新缓存、检测服务器过期、检测 KEY 过期
- [x] 充值系统：支持 KEY 充值或支付宝扫码支付（基于支付宝当面付 API ）
- [x] 服务器到期自动邮件提醒用户、一定时间后在管理员控制台手动确认删除

##### TODO

- [ ] 优化模板复用，提高渲染效率
- [ ] 添加微信支付
- [ ] 详细使用文档
- [ ] 添加微信支付、支付方式集成到订单页
- [ ] 多语言国际化
- [ ] ~~重构早期代码😥~~
- [ ] 修复若干BUG
- [ ] 跟进 Pterodactyl 的新版本 API

如有改进建议或需求欢迎发送 Issue 或 Pull Request