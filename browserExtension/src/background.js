chrome.contextMenus.create({
  title: 'ContextMenuSample',
  type: 'normal',
  contexts: ['link'],
  onclick: (info, tab) => {
    const url = 'http://localhost:8081/url';
    const request = new XMLHttpRequest();
    let message = ''; 
    request.open('GET', url + '?url=' + info.linkUrl);
    request.onreadystatechange = function() {
      if(request.status !== 200) {
        return 'REQUEST FAILED';
      }
      return request.responseText;
    }

    message = request.send(null);
    // 参照: https://kic-yuuki.hatenablog.com/entry/2019/11/03/143956
    // alert(
    //   JSON.stringify({
    //     // message: 'AlertSample Clicked!!',
    //     info: info,
    //     linkUrl: info.linkUrl,
    //     url: tab.url,
    //     message: message
    //   })
    // );
  }
})