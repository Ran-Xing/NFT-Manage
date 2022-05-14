## README

<u>由于技术优先，还未开始就放弃了</u>
<u>有幸看到这个仓库欢迎</u> [提交 ISSUES](issues)

```mermaid
flowchart LR
    User-->Chrome{Chrome}
    Chrome{Chrome}--Request-->Web
		Chrome{Chrome}-->PlugIn
		PlugIn<--GetCookies-->Server
    PlugIn<--GetDom  /  SetCookies-->Web
```
