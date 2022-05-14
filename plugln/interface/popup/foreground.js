function getCookies() {
    const url = window.location;
    const cookie = document.cookie;
    if (cookie === "") {
        Notice("cookies 不存在!")
    } else {
        Notice("已登录!")
    }
}
function Notice(noticeInfo) {
    //新建一个div元素节点
    const div = document.createElement("div");
    div.innerText = noticeInfo;
    div.id = "snackbar"
    div.className = "show"
    //把div元素节点添加到body元素节点中成为其子节点，但是放在body的现有子节点的最后
    document.body.appendChild(div);
    //插入到最后面
    document.body.insertBefore(div, document.body.lastElementChild);
    setTimeout(function () {
        document.body.removeChild(div)
    }, 3000);
}

getCookies();
