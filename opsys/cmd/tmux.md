## 基本管理命令
- tmux new -s \<session-name>   //启动新会话session-name
- tmux attach -t 0  // 使用会话编号连接会话
- tmux attach -t \<session-name> //使用会话名称连接
- tmux kill-session //结束会话，用法同attach
- tmux ls  //列出会话
- $ tmux select-pane -t {pane} -T {title}
    - \# Examples:
    - $ tmux select-pane -T title1          # Change title of current pane
    - $ tmux select-pane -t 1 -T title2     # Change title of pane 1 in current window
    - $ tmux select-pane -t 2.1 -T title3   # Change title of pane 1 in window 2
- tmux set pane-border-status bottom # 显示每个窗格的标题
- tmux set pane-border-status off # 关闭状态栏

## 快捷键
- Ctrl+b d ：退出，后台运行
- Ctrl+b s：列出所有会话
- Ctrl+b %：划分左右两个窗格。
- Ctrl+b "：划分上下两个窗格。
- Ctrl+b <arrow key>：光标切换到其他窗格。<arrow key>是指向要切换到的窗格的方向键，比如切换到下方窗格，就按方向键↓。
- Ctrl+b ;：光标切换到上一个窗格。
- Ctrl+b o：光标切换到下一个窗格。
- Ctrl+b {：当前窗格与上一个窗格交换位置。
- Ctrl+b }：当前窗格与下一个窗格交换位置。
- Ctrl+b Ctrl+o：所有窗格向前移动一个位置，第一个窗格变成最后一个窗格。
- Ctrl+b Alt+o：所有窗格向后移动一个位置，最后一个窗格变成第一个窗格。
- **Ctrl+b x：关闭当前窗格面板。**
- Ctrl+b !：将当前窗格拆分为一个独立窗口。
- Ctrl+b z：当前窗格全屏显示，再使用一次会变回原来大小。
- Ctrl+b Ctrl+<arrow key>：按箭头方向调整窗格大小。
- Ctrl+b q：显示窗格编号。
- **Ctrl+b &：删除当前窗口。**

- 按住shift选择，可进行复制
- 按住shift可进行上下滚动
